package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"mcp-server/internal/application/dto"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	updateProduct            = "update_product"
	updateProductDescription = "Herramienta que permite Actualizar Datos de un producto en la base de datos, como ImageUrl, nombre, precio, estado Activo"
)

// UpdateProduct crea y configura una herramienta MCP para actualizar los datos de un producto existente.
// La herramienta permite modificar campos como ImageUrl, nombre, precio y estado activo del producto.
// Utiliza un esquema de entrada basado en el DTO UpdateProductRequest para validar los datos de entrada.
// Retorna una herramienta MCP configurada con el esquema de entrada correspondiente.
func (m *ManagmentTools) UpdateProduct() mcp.Tool {
	tool := mcp.NewTool(updateProduct,
		mcp.WithDescription(updateProductDescription),
		mcp.WithInputSchema[dto.UpdateProductRequest](),
	)
	return tool
}

// HandleUpdateProduct procesa la solicitud para actualizar los datos de un producto existente.
// Extrae los argumentos de la solicitud MCP, los deserializa al DTO UpdateProductRequest y ejecuta
// el caso de uso correspondiente para actualizar el producto en la base de datos.
//
// Parámetros:
//   - ctx: Contexto de la operación
//   - request: Solicitud MCP que contiene los datos del producto a actualizar
//
// Retorna:
//   - *mcp.CallToolResult: Resultado de la operación en formato JSON con los datos actualizados
//   - error: Error si ocurre algún problema durante la serialización, validación o ejecución
func (m *ManagmentTools) HandleUpdateProduct(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Obtener argumentos
	raw := request.GetRawArguments()

	// Parsear a DTO
	bytesArgs, err := json.Marshal(raw)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("no se pudieron serializar los argumentos: %v", err)), nil
	}

	var req dto.UpdateProductRequest
	if err := json.Unmarshal(bytesArgs, &req); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("argumentos inválidos para CreateProduct: %v", err)), nil
	}

	// Ejecutar caso de uso
	resp, err := m.inventoryUseCase.UpdateProduct(ctx, &req)
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
