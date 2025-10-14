package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	DeleteProduct            = "delete_product"
	DeleteProductDescription = "Herramienta que permite eliminar un producto de la base de datos"
)

// DeleteProduct crea y configura una herramienta MCP para eliminar un producto de la base de datos.
// La herramienta requiere únicamente el SKU del producto que se desea eliminar.
// Utiliza el SKU como identificador único para localizar y eliminar el producto específico.
// Retorna una herramienta MCP configurada con el parámetro SKU requerido.
func (m *ManagmentTools) DeleteProduct() mcp.Tool {
	tool := mcp.NewTool(DeleteProduct,
		mcp.WithDescription(DeleteProductDescription),
		mcp.WithString(sku, mcp.Description(skuDescription)),
	)
	return tool
}

// HandleDeleteProduct procesa la solicitud para eliminar un producto de la base de datos.
// Extrae el SKU del producto de la solicitud MCP y ejecuta el caso de uso correspondiente
// para eliminar el producto identificado por ese SKU.
//
// Parámetros:
//   - ctx: Contexto de la operación
//   - request: Solicitud MCP que contiene el SKU del producto a eliminar
//
// Retorna:
//   - *mcp.CallToolResult: Resultado de la operación en formato JSON confirmando la eliminación
//   - error: Error si ocurre algún problema durante la ejecución del caso de uso o formateo de la respuesta
func (m *ManagmentTools) HandleDeleteProduct(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	sku := request.GetString(sku, "")

	// Ejecutar caso de uso
	resp, err := m.inventoryUseCase.Delete(ctx, sku)
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
