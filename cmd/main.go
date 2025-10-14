package main

import (
	"context"
	"fmt"
	"mcp-server/internal/infrastructure/config"
	"mcp-server/internal/interfaces/mcp"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log/slog"

	useCases "mcp-server/internal/application/inventory"
	invRepo "mcp-server/internal/infrastructure/repository/firestore/inventory"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Cargar configuración YAML por defecto
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		fmt.Printf("Error cargando configuración: %v\n", err)
		os.Exit(1)
	}

	slog.Info("Starting MCP Server...")

	// Crear contexto con cancelación
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Configurar manejo de señales para shutdown graceful
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		slog.Info("Received shutdown signal, gracefully shutting down...")
		cancel()
	}()

	// Inicializar dependencias usando Dependency Injection
	app, err := initializeApplication(ctx, cfg)
	if err != nil {
		slog.Error("Failed to initialize application: %v", err.Error(), err)
	}

	// Ejecutar aplicación
	if err := app.Run(ctx); err != nil {
		slog.Error("Application failed: %v", err.Error(), err)
	}

	slog.Info("MCP Server stopped")
}

// Application representa la aplicación principal
type Application struct {
	server     *server.MCPServer
	httpServer *server.StreamableHTTPServer
	addr       string
}

// initializeApplication inicializa todas las dependencias de la aplicación
func initializeApplication(ctx context.Context, cfg *config.Config) (*Application, error) {
	slog.Info("Initializing application dependencies...")
	slog.Info("Server config loaded", "host", cfg.Server.Host, "port", cfg.Server.Port)
	slog.Info("Database config loaded", "project_id", cfg.Database.ProjectID)

	// Inicializar repositorio
	invRepo := invRepo.NewInventoryRepository(&cfg.Database)
	slog.Info("User repository initialized")

	// Inicializar casos de uso
	userUseCase := useCases.NewInventoryCases(invRepo)
	slog.Info("User use cases initialized")

	// Inicializar servidor MCP
	mcpServer := server.NewMCPServer(
		"mcp-server",
		"1.0.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	slog.Info("MCP server initialized")

	// Registrar handlers MCP
	userHandler := mcp.NewMCPHandler(userUseCase)
	if err := userHandler.RegisterTools(mcpServer); err != nil {
		return nil, fmt.Errorf("failed to register MCP tools: %w", err)
	}
	slog.Info("MCP tools registered")

	// Configurar servidor HTTP (streamable-http) en modo stateless
	httpServer := server.NewStreamableHTTPServer(
		mcpServer,
		server.WithStateLess(true), // Modo stateless - no requiere gestión de sesiones
		// server.WithEndpointPath("/mcp"), // opcional, por defecto "/mcp"
	)

	// Crear aplicación
	app := &Application{
		server:     mcpServer,
		httpServer: httpServer,
		addr:       fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
	}

	slog.Info("Application initialization completed")
	return app, nil
}

// Run ejecuta la aplicación
func (app *Application) Run(ctx context.Context) error {
	// Ejecutar servidor MCP usando HTTP (streamable-http)
	errChan := make(chan error, 1)
	go func() {
		slog.Info("Starting HTTP MCP server...", "addr", app.addr)
		if err := app.httpServer.Start(app.addr); err != nil {
			errChan <- fmt.Errorf("HTTP MCP server failed: %w", err)
		}
	}()

	// Esperar contexto o error
	select {
	case <-ctx.Done():
		slog.Info("Shutting down HTTP MCP server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := app.httpServer.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("error en apagado HTTP MCP server: %w", err)
		}
		return nil
	case err := <-errChan:
		return err
	}
}
