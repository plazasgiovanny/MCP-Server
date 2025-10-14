package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	getProductBySku            = "get_product_by_sku"
	getProductBySkuDescription = "Herramienta que permite obtener un producto de la base de datos por el sku"
)

// GetProductBySku crea y configura una herramienta MCP para obtener un producto específico de la base de datos.
// La herramienta requiere únicamente el SKU del producto que se desea consultar.
// Utiliza el SKU como identificador único para localizar y retornar la información del producto específico.
// Retorna una herramienta MCP configurada con el parámetro SKU requerido.
func (m *ManagmentTools) GetProductBySku() mcp.Tool {
	tool := mcp.NewTool(getProductBySku,
		mcp.WithDescription(getProductBySkuDescription),
		mcp.WithString(sku, mcp.Description(skuDescription)),
	)
	return tool
}

// HandleGetProductBySku procesa la solicitud para obtener un producto específico de la base de datos.
// Extrae el SKU del producto de la solicitud MCP y ejecuta el caso de uso correspondiente
// para buscar y retornar la información del producto identificado por ese SKU.
//
// Parámetros:
//   - ctx: Contexto de la operación
//   - request: Solicitud MCP que contiene el SKU del producto a consultar
//
// Retorna:
//   - *mcp.CallToolResult: Resultado de la operación en formato JSON con la información completa del producto
//   - error: Error si ocurre algún problema durante la ejecución del caso de uso o formateo de la respuesta
func (m *ManagmentTools) HandleGetProductBySku(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	sku := request.GetString(sku, "")

	// Ejecutar caso de uso
	resp, err := m.inventoryUseCase.GetProductBySku(ctx, sku)
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
