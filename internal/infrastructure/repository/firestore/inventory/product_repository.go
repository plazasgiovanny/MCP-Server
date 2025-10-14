package inventory

import (
	"context"
	"errors"
	"log/slog"
	"mcp-server/internal/domain/entities"
	"mcp-server/internal/domain/repository"

	"mcp-server/internal/infrastructure/config"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/firestore/apiv1/firestorepb"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

const (
	productCollection = "productos"
	stockOrder        = "stock"
)

type inventoryRepository struct {
	client *firestore.Client
	config *config.FireStoreConfig
}

func NewInventoryRepository(config *config.FireStoreConfig) repository.ProductRepository {
	conf := &firebase.Config{ProjectID: config.ProjectID}
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		slog.Error("Error creating Firebase app", "error", err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		slog.Error("Error creating Firestore client", "error", err)
	}
	return &inventoryRepository{client: client, config: config}
}

func (r *inventoryRepository) Upsert(ctx context.Context, product *entities.Product) error {
	_, err := r.client.Collection(productCollection).Doc(product.Sku).Set(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (r *inventoryRepository) GetBySku(ctx context.Context, sku string) (*entities.Product, error) {
	docsnap, err := r.client.Collection(productCollection).Doc(sku).Get(ctx)
	if err != nil {
		return nil, err
	}
	var product entities.Product
	if err := docsnap.DataTo(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *inventoryRepository) Delete(ctx context.Context, sku string) error {
	_, err := r.client.Collection(productCollection).Doc(sku).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *inventoryRepository) List(ctx context.Context, limit, offset int) ([]*entities.Product, error) {
	// Validar parámetros
	if limit <= 0 {
		limit = 10 // Valor por defecto
	}
	if offset < 0 {
		offset = 0
	}

	// Crear query base ordenada por stock descendente
	query := r.client.Collection(productCollection).OrderBy(stockOrder, firestore.Desc)

	// Si hay offset, necesitamos implementar paginación con cursores
	// Para simplificar, usamos Limit con offset (menos eficiente pero funcional)
	// En producción, se recomienda usar cursores con startAfter()
	if offset > 0 {
		// Obtener todos los documentos hasta el offset + limit
		allQuery := query.Limit(offset + limit)
		allItems := allQuery.Documents(ctx)

		var allProducts []*entities.Product
		defer allItems.Stop()

		for {
			doc, err := allItems.Next()
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

		// Aplicar offset y limit manualmente
		start := offset
		end := offset + limit
		if start >= len(allProducts) {
			return []*entities.Product{}, nil
		}
		if end > len(allProducts) {
			end = len(allProducts)
		}

		return allProducts[start:end], nil
	}

	// Sin offset, usar limit directamente (más eficiente)
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

func (r *inventoryRepository) Count(ctx context.Context) (int64, error) {
	collection := r.client.Collection(productCollection)

	aggregationQuery := collection.Query.NewAggregationQuery().WithCount("all")
	results, err := aggregationQuery.Get(ctx)
	if err != nil {
		return 0, err
	}

	count, ok := results["all"]
	if !ok {
		return 0, errors.New("firestore: couldn't get alias for COUNT from results")
	}

	countValue := count.(*firestorepb.Value)
	return countValue.GetIntegerValue(), nil
}
