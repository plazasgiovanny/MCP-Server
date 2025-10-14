package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"mcp-server/internal/application/dto"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateProduct = "create_product"
)

// CreateNewProductTool crea y configura una herramienta MCP para crear un nuevo producto en la base de datos.
// La herramienta utiliza un esquema de entrada basado en el DTO CreateProductRequest para validar
// los datos de entrada del nuevo producto antes de su creación.
// Retorna una herramienta MCP configurada con el esquema de entrada correspondiente.
func (m *ManagmentTools) CreateNewProductTool() mcp.Tool {
	CreateProduct := mcp.NewTool(CreateProduct,
		mcp.WithDescription("Crea un Nuevo producto en la base de datos"),
		mcp.WithInputSchema[dto.CreateProductRequest](),
	)
	return CreateProduct
}

// HandleCreateProduct procesa la solicitud para crear un nuevo producto en la base de datos.
// Extrae los argumentos de la solicitud MCP, los deserializa al DTO CreateProductRequest y ejecuta
// el caso de uso correspondiente para crear el producto en la base de datos.
//
// Parámetros:
//   - ctx: Contexto de la operación
//   - request: Solicitud MCP que contiene los datos del nuevo producto a crear
//
// Retorna:
//   - *mcp.CallToolResult: Resultado de la operación en formato JSON con la información del producto creado
//   - error: Error si ocurre algún problema durante la serialización, validación o ejecución
func (m *ManagmentTools) HandleCreateProduct(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// Obtener argumentos
	raw := request.GetRawArguments()

	// Parsear a DTO
	bytesArgs, err := json.Marshal(raw)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("no se pudieron serializar los argumentos: %v", err)), nil
	}

	var req dto.CreateProductRequest
	if err := json.Unmarshal(bytesArgs, &req); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("argumentos inválidos para CreateProduct: %v", err)), nil
	}

	// Ejecutar caso de uso
	resp, err := m.inventoryUseCase.CreateProduct(ctx, &req)
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
