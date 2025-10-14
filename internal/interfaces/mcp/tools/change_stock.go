package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"mcp-server/internal/application/dto"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ChangeStockOfProduct          = "change_stock_product"
	stockDescription              = "Herramienta que permite cambiar el stock de un producto, sumando o restando, segun el type_operation"
	TypeOperation                 = "type_operation"
	paramTypeOperationDescription = "Tipo de operacion a realizar, para suma (Sum) y para resta (Rest)"
	sku                           = "sku"
	skuDescription                = "Stock Keeping Unit: Codigo unico que permite identificar un producto ejemplo > TEL-APP-I13P-256"
	valueStock                    = "value_stock"
	valueStockDescription         = "valor del stock (numeros enteros positivos) a sumar o restar dependiendo del tipo de operacion"
)

// ChangeStockOfProduct crea y configura una herramienta MCP para modificar el stock de un producto existente.
// La herramienta permite sumar o restar stock de un producto específico identificado por su SKU.
// Requiere tres parámetros: SKU del producto, tipo de operación (Sum/Rest) y el valor del stock a modificar.
// Retorna una herramienta MCP configurada con los parámetros necesarios para la operación de cambio de stock.
func (m *ManagmentTools) ChangeStockOfProduct() mcp.Tool {
	tool := mcp.NewTool(ChangeStockOfProduct,
		mcp.WithDescription(stockDescription),
		mcp.WithString(sku, mcp.Description(skuDescription)),
		mcp.WithString(TypeOperation, mcp.Description(paramTypeOperationDescription)),
		mcp.WithNumber(valueStock, mcp.Description(valueStockDescription)),
	)
	return tool
}

// HandleChangeStockOfProduct procesa la solicitud para modificar el stock de un producto específico.
// Extrae los parámetros SKU, tipo de operación y valor del stock de la solicitud MCP, construye el DTO
// correspondiente y ejecuta el caso de uso para realizar la operación de cambio de stock.
//
// Parámetros:
//   - ctx: Contexto de la operación
//   - request: Solicitud MCP que contiene el SKU del producto, tipo de operación y valor del stock
//
// Retorna:
//   - *mcp.CallToolResult: Resultado de la operación en formato JSON con la información actualizada del producto
//   - error: Error si ocurre algún problema durante la ejecución del caso de uso o formateo de la respuesta
func (m *ManagmentTools) HandleChangeStockOfProduct(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	sku := request.GetString(sku, "")
	typeOp := request.GetString(TypeOperation, "Sum")
	valueStock := request.GetInt(valueStock, 0)

	req := &dto.ChangeStockProductRequest{
		Sku:   sku,
		Type:  typeOp,
		Stock: valueStock,
	}
	// Ejecutar caso de uso
	resp, err := m.inventoryUseCase.ChangeStockOfProduct(ctx, req)
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
