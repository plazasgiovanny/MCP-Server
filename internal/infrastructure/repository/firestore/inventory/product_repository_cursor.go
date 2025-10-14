package inventory

import (
	"context"
	"mcp-server/internal/domain/entities"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// ListWithCursor implementa paginación eficiente usando cursores de Firestore
// Esta es la implementación recomendada para producción
func (r *inventoryRepository) ListWithCursor(ctx context.Context, limit int, lastDoc *entities.Product) ([]*entities.Product, error) {
	// Validar parámetros
	if limit <= 0 {
		limit = 10 // Valor por defecto
	}
	if limit > 100 {
		limit = 100 // Límite máximo para evitar consultas muy grandes
	}

	// Crear query base ordenada por stock descendente
	query := r.client.Collection(productCollection).OrderBy(stockOrder, firestore.Desc)

	// Si tenemos un documento de referencia (cursor), usar startAfter
	if lastDoc != nil {
		// Crear un documento de referencia para el cursor
		docRef := r.client.Collection(productCollection).Doc(lastDoc.Sku)
		query = query.StartAfter(docRef)
	}

	// Aplicar límite
	query = query.Limit(limit)

	items := query.Documents(ctx)
	products := []*entities.Product{}

	defer items.Stop()
	for {
		doc, err := items.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var product entities.Product
		doc.DataTo(&product)
		products = append(products, &product)
	}

	return products, nil
}

// ListWithOffset implementa paginación tradicional con offset
// Menos eficiente pero más simple de entender
func (r *inventoryRepository) ListWithOffset(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	// Validar parámetros
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	if limit > 100 {
		limit = 100
	}

	// Para offset > 0, necesitamos obtener más documentos y aplicar el offset manualmente
	// Esto es menos eficiente pero funcional
	totalToFetch := offset + limit

	query := r.client.Collection(productCollection).OrderBy(stockOrder, firestore.Desc).Limit(totalToFetch)
	items := query.Documents(ctx)
	allProducts := []*entities.Product{}

	defer items.Stop()
	for {
		doc, err := items.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var product entities.Product
		doc.DataTo(&product)
		allProducts = append(allProducts, &product)
	}

	// Aplicar offset
	if offset >= len(allProducts) {
		return []*entities.Product{}, nil
	}

	end := offset + limit
	if end > len(allProducts) {
		end = len(allProducts)
	}

	return allProducts[offset:end], nil
}
