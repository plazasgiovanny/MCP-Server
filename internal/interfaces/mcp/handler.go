package mcp

import (
	usecase "mcp-server/internal/application/inventory"

	"mcp-server/internal/interfaces/mcp/tools"

	"github.com/mark3labs/mcp-go/server"
)

type MCPHandler struct {
	inventoryUseCase *usecase.InventoryCases
	managementTools  *tools.ManagmentTools
}

func NewMCPHandler(inventory *usecase.InventoryCases) *MCPHandler {
	return &MCPHandler{
		inventoryUseCase: inventory,
		managementTools:  tools.NewManagmentTools(inventory),
	}
}

// RegisterTools registra las herramientas MCP para usuarios
func (h *MCPHandler) RegisterTools(mcpServer *server.MCPServer) error {
	// Registrar herramientas
	mcpServer.AddTool(h.managementTools.CreateNewProductTool(), h.managementTools.HandleCreateProduct)
	mcpServer.AddTool(h.managementTools.LoadInitialSmartphoneCatalog(), h.managementTools.HandleLoadInitialSmartphoneCatalog)
	mcpServer.AddTool(h.managementTools.ChangeStockOfProduct(), h.managementTools.HandleChangeStockOfProduct)
	mcpServer.AddTool(h.managementTools.DeleteProduct(), h.managementTools.HandleDeleteProduct)
	mcpServer.AddTool(h.managementTools.GetListProduct(), h.managementTools.HandleGetListProduct)
	mcpServer.AddTool(h.managementTools.UpdateProduct(), h.managementTools.HandleUpdateProduct)
	mcpServer.AddTool(h.managementTools.GetProductBySku(), h.managementTools.HandleGetProductBySku)
	return nil
}
