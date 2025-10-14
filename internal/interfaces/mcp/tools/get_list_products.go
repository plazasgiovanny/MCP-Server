package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	getListProduct            = "get_list_products"
	getListProductDescription = "Permite Obtener un listado de productos de forma paginada, basada en el limite y el offset"
	limitParam                = "limit"
	offseParam                = "offset"
	limitParamDescription     = "Número máximo de productos a retornar en la consulta"
	offsetParamDescription    = "Número de productos a omitir desde el inicio de la consulta (para paginación)"
)

// GetListProduct crea y configura una herramienta MCP para obtener un listado paginado de productos.
// La herramienta permite especificar un límite de productos a retornar y un offset para la paginación.
// Retorna una herramienta MCP configurada con los parámetros necesarios para la consulta paginada.
func (m *ManagmentTools) GetListProduct() mcp.Tool {
	tool := mcp.NewTool(getListProduct,
		mcp.WithDescription(getListProductDescription),
		mcp.WithNumber(limitParam, mcp.Description(limitParamDescription)),
		mcp.WithNumber(offseParam, mcp.Description(offsetParamDescription)),
	)
	return tool
}

// HandleGetListProduct procesa la solicitud para obtener un listado paginado de productos.
// Extrae los parámetros de limit y offset de la solicitud MCP y ejecuta el caso de uso correspondiente.
//
// Parámetros:
//   - ctx: Contexto de la operación
//   - request: Solicitud MCP que contiene los parámetros limit y offset
//
// Retorna:
//   - *mcp.CallToolResult: Resultado de la operación en formato JSON
//   - error: Error si ocurre algún problema durante la ejecución
func (m *ManagmentTools) HandleGetListProduct(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	limit := request.GetInt(limitParam, 0)
	offset := request.GetInt(offseParam, 0)

	// Ejecutar caso de uso
	resp, err := m.inventoryUseCase.GetAllProductoByPagination(ctx, limit, offset)
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
