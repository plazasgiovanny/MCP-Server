package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	loadInitalData           = "load_initial_data"
	loadDescription          = "Carga la data inicial del catalogo de smartphones en la base de datos"
	pathSmartPhoneCatalog    = "configs/%v"
	fileName                 = "file_name"
	paramFileNameDescription = "Nombre del catalogo a cargar, definido por el usuario"
)

// LoadInitialSmartphoneCatalog crea y configura una herramienta MCP para cargar datos iniciales del catálogo de smartphones.
// La herramienta permite cargar un archivo CSV específico con información de productos smartphones en la base de datos.
// Requiere el nombre del archivo de catálogo que se encuentra en la carpeta configs/.
// Retorna una herramienta MCP configurada con el parámetro de nombre de archivo requerido.
func (m *ManagmentTools) LoadInitialSmartphoneCatalog() mcp.Tool {
	loadInitialData := mcp.NewTool(loadInitalData,
		mcp.WithDescription(loadDescription),
		mcp.WithString(fileName, mcp.Description(paramFileNameDescription)),
	)
	return loadInitialData
}

// HandleLoadInitialSmartphoneCatalog procesa la solicitud para cargar datos iniciales del catálogo de smartphones.
// Extrae el nombre del archivo de la solicitud MCP, construye la ruta completa del archivo CSV y ejecuta
// el caso de uso correspondiente para cargar los productos smartphones en la base de datos.
//
// Parámetros:
//   - ctx: Contexto de la operación
//   - request: Solicitud MCP que contiene el nombre del archivo de catálogo a cargar
//
// Retorna:
//   - *mcp.CallToolResult: Resultado de la operación en formato JSON con información sobre los productos cargados
//   - error: Error si ocurre algún problema durante la carga del archivo o formateo de la respuesta
func (m *ManagmentTools) HandleLoadInitialSmartphoneCatalog(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileN := request.GetString(fileName, "")
	// Ejecutar caso de uso
	resp, err := m.inventoryUseCase.InitalProductData(ctx, fmt.Sprintf(pathSmartPhoneCatalog, fileN))
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	// Responder
	payload, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("error formateando respuesta: %v", err)), nil
	}
	return mcp.NewToolResultText(string(payload)), nil
}
