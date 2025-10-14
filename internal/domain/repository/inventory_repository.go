package repository

import (
	"context"
	"mcp-server/internal/domain/entities"
)

type ProductRepository interface {
	// Crea o actualiza un registro
	Upsert(ctx context.Context, product *entities.Product) error

	// GetBySku obtiene un producto por su SKU
	GetBySku(ctx context.Context, sku string) (*entities.Product, error)

	// Delete elimina un producto por su SKU
	Delete(ctx context.Context, sku string) error

	// List obtiene una lista de productos con paginación
	List(ctx context.Context, limit, offset int) ([]*entities.Product, error)

	// Count retorna el número total de productos
	Count(ctx context.Context) (int64, error)
}
