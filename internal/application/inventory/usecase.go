package usecase

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"math"
	"mcp-server/internal/application/dto"
	"mcp-server/internal/domain/entities"
	"mcp-server/internal/domain/repository"
	"os"
	"strconv"
	"time"
)

const (
	defaultPath  string = "configs/smartphones_catalog.csv"
	requestEmpty string = "la peticion esta vacia"
	TypeSum             = "Sum"
	TypeRest            = "Rest"
)

type InventoryCases struct {
	ProductRepository repository.ProductRepository
}

func NewInventoryCases(productRepository repository.ProductRepository) *InventoryCases {
	return &InventoryCases{
		ProductRepository: productRepository,
	}
}

func (i *InventoryCases) CreateProduct(ctx context.Context, request *dto.CreateProductRequest) (response *dto.SuccessResponse, err error) {
	if request == nil {
		return nil, errors.New(requestEmpty)
	}

	product := &entities.Product{
		Sku:          request.Sku,
		Active:       request.Active,
		Category:     request.Category,
		ImageUrl:     request.ImageUrl,
		Name:         request.Name,
		NameProvider: request.NameProvider,
		Price:        request.Price,
		Stock:        0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = i.ProductRepository.Upsert(ctx, product)
	if err != nil {
		return nil, err
	}

	return &dto.SuccessResponse{
		Message: fmt.Sprintf("Producto creado correctamente: %v", product.Sku),
		Data:    product,
	}, nil

}

func (i *InventoryCases) InitalProductData(ctx context.Context, path string) (*dto.SuccessResponse, error) {
	// Si no se proporciona path, usar el default
	if path == "" {
		path = defaultPath
	}

	// Abrir el archivo CSV
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo CSV: %v", err)
	}
	defer file.Close()

	// Crear el lector CSV
	reader := csv.NewReader(file)

	// Leer todas las líneas del CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo CSV: %v", err)
	}

	// Verificar que hay al menos una línea (header)
	if len(records) < 2 {
		return nil, fmt.Errorf("archivo CSV vacío o sin datos")
	}

	// Saltar la primera línea (header) y procesar los datos
	products := []entities.Product{}
	for j, record := range records[1:] {
		product, err := mapCSVRecordToProduct(record)
		if err != nil {
			return nil, fmt.Errorf("error procesando línea %d: %v", j+2, err) // +2 porque empezamos desde línea 2
		}
		err = i.ProductRepository.Upsert(ctx, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return &dto.SuccessResponse{
		Message: fmt.Sprintf("Se insertaron %d registros en la base de datos", len(products)),
	}, nil
}

func (i *InventoryCases) ChangeStockOfProduct(ctx context.Context, request *dto.ChangeStockProductRequest) (response *dto.SuccessResponse, err error) {
	if request == nil {
		return nil, errors.New(requestEmpty)
	}

	product, err := i.ProductRepository.GetBySku(ctx, request.Sku)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("el producto %v no existe", request.Sku)
	}
	valueStock := int(math.Abs(float64(request.Stock)))
	if TypeRest == request.Type {
		valueStock = valueStock * -1
	}
	product.Stock += valueStock

	err = i.ProductRepository.Upsert(ctx, product)
	if err != nil {
		return nil, err
	}

	return &dto.SuccessResponse{
		Message: fmt.Sprintf("Producto creado correctamente: %v, nuevo stock: %d", product.Sku, product.Stock),
		Data:    product,
	}, nil

}

func (i *InventoryCases) Delete(ctx context.Context, sku string) (response *dto.SuccessResponse, err error) {
	if sku == "" {
		return nil, errors.New(requestEmpty)
	}

	err = i.ProductRepository.Delete(ctx, sku)
	if err != nil {
		return nil, err
	}

	return &dto.SuccessResponse{
		Message: fmt.Sprintf("Producto Eliminado correctamente: %v", sku),
		Data:    nil,
	}, nil

}

func (i *InventoryCases) GetAllProductoByPagination(ctx context.Context, limit, offset int) (response *dto.SuccessResponse, err error) {
	if limit < 0 {
		return nil, errors.New("el limite debe ser mayor a 0")
	}
	if offset < 0 {
		return nil, errors.New("el offset debe ser mayor a 0")
	}
	count, _ := i.ProductRepository.Count(ctx)
	if count == 0 {
		return nil, errors.New("no existen productos para consultar")
	}

	list, err := i.ProductRepository.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.SuccessResponse{
		Message: fmt.Sprintf("Se consultaron %d productos de %d", len(list), count),
		Data:    list,
	}, nil

}

func (i *InventoryCases) UpdateProduct(ctx context.Context, request *dto.UpdateProductRequest) (response *dto.SuccessResponse, err error) {
	if request == nil {
		return nil, errors.New(requestEmpty)
	}
	if request.Sku == "" {
		return nil, errors.New("el sku no puede venir vacio")
	}

	product, err := i.ProductRepository.GetBySku(ctx, request.Sku)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("el producto %v no existe", request.Sku)
	}

	product.ImageUrl = request.ImageUrl
	product.Name = request.Name
	product.Active = request.Active
	product.Price = request.Price

	err = i.ProductRepository.Upsert(ctx, product)
	if err != nil {
		return nil, err
	}

	return &dto.SuccessResponse{
		Message: fmt.Sprintf("Producto Actualizado correctamente: %v", product.Sku),
		Data:    product,
	}, nil

}

func (i *InventoryCases) GetProductBySku(ctx context.Context, sku string) (response *dto.SuccessResponse, err error) {
	if sku == "" {
		return nil, errors.New(requestEmpty)
	}

	product, err := i.ProductRepository.GetBySku(ctx, sku)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, fmt.Errorf("el producto %v no existe", sku)
	}

	return &dto.SuccessResponse{
		Message: fmt.Sprintf("Producto consultado correctamente: %v", sku),
		Data:    product,
	}, nil

}

// Función auxiliar para mapear una fila del CSV a un Product
func mapCSVRecordToProduct(record []string) (entities.Product, error) {
	// Verificar que tenemos todos los campos necesarios
	if len(record) != 10 {
		return entities.Product{}, fmt.Errorf("número incorrecto de campos: esperado 10, obtenido %d", len(record))
	}

	// Parsear los campos que necesitan conversión
	active, err := strconv.ParseBool(record[1])
	if err != nil {
		return entities.Product{}, fmt.Errorf("error parseando campo 'activo': %v", err)
	}

	price, err := strconv.ParseFloat(record[8], 64)
	if err != nil {
		return entities.Product{}, fmt.Errorf("error parseando campo 'precio': %v", err)
	}

	stock, err := strconv.Atoi(record[9])
	if err != nil {
		return entities.Product{}, fmt.Errorf("error parseando campo 'stock': %v", err)
	}

	// Parsear las fechas
	createdAt, err := time.Parse(time.RFC3339, record[4])
	if err != nil {
		return entities.Product{}, fmt.Errorf("error parseando campo 'fechaCreacion': %v", err)
	}

	updatedAt, err := time.Parse(time.RFC3339, record[3])
	if err != nil {
		return entities.Product{}, fmt.Errorf("error parseando campo 'fechaActualizacion': %v", err)
	}

	// Crear y retornar el producto
	return entities.Product{
		Sku:          record[0], // sku
		Active:       active,    // activo
		Category:     record[2], // categoria
		UpdatedAt:    updatedAt, // fechaActualizacion
		CreatedAt:    createdAt, // fechaCreacion
		ImageUrl:     record[5], // imagenUrl
		Name:         record[6], // nombre
		NameProvider: record[7], // nombreProveedor
		Price:        price,     // precio
		Stock:        stock,     // stock
	}, nil
}
