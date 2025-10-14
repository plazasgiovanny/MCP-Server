# FireStoreAPI - MCP Server para Gestión de Inventario

## 📋 Descripción

FireStoreAPI es un servidor **Model Context Protocol (MCP)** desarrollado en Go que proporciona herramientas para la gestión completa de inventario de productos tecnológicos utilizando Google Firestore como base de datos. Este proyecto está diseñado específicamente para startups de comercio electrónico que necesitan un sistema robusto y escalable para administrar su catálogo de productos.

## 🎯 Propósito

El proyecto tiene como objetivo principal proporcionar una API RESTful a través del protocolo MCP que permita:

- **Gestión completa de productos**: Crear, leer, actualizar y eliminar productos del inventario
- **Control de stock**: Modificar cantidades de inventario con operaciones de suma y resta
- **Carga masiva de datos**: Importar catálogos completos desde archivos CSV
- **Consultas paginadas**: Obtener listados de productos con paginación eficiente
- **Integración con IA**: Exponer funcionalidades a través de herramientas MCP para uso con modelos de IA

## 🏗️ Arquitectura

El proyecto implementa una **arquitectura hexagonal (Clean Architecture)** con separación clara de responsabilidades:

```
┌─────────────────────────────────────────────────────────────┐
│                    PRESENTATION LAYER                       │
│  ┌─────────────────┐  ┌─────────────────┐  ┌──────────────┐  │
│  │   MCP Tools    │  │   HTTP Server   │  │   Handlers   │  │
│  │   (create_*)   │  │   (Streamable)  │  │              │  │
│  └─────────────────┘  └─────────────────┘  └──────────────┘  │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                   APPLICATION LAYER                         │
│  ┌─────────────────┐  ┌─────────────────┐  ┌──────────────┐  │
│  │   Use Cases     │  │      DTOs       │  │   Services    │  │
│  │   (Inventory)   │  │   (Requests)    │  │              │  │
│  └─────────────────┘  └─────────────────┘  └──────────────┘  │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                     DOMAIN LAYER                            │
│  ┌─────────────────┐  ┌─────────────────┐  ┌──────────────┐  │
│  │   Entities      │  │   Repositories  │  │   Business   │  │
│  │   (Product)     │  │   (Interfaces)  │  │   Logic      │  │
│  └─────────────────┘  └─────────────────┘  └──────────────┘  │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                 INFRASTRUCTURE LAYER                        │
│  ┌─────────────────┐  ┌─────────────────┐  ┌──────────────┐  │
│  │   Firestore     │  │   Configuration │  │   External   │  │
│  │   Repository    │  │   (YAML)       │  │   Services   │  │
│  └─────────────────┘  └─────────────────┘  └──────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

### Componentes Principales

#### 1. **Presentation Layer** (`internal/interfaces/`)
- **MCP Tools**: Herramientas expuestas a través del protocolo MCP
- **HTTP Server**: Servidor HTTP con soporte para streamable-http
- **Handlers**: Manejo de solicitudes y respuestas

#### 2. **Application Layer** (`internal/application/`)
- **Use Cases**: Lógica de negocio para gestión de inventario
- **DTOs**: Objetos de transferencia de datos para requests/responses
- **Services**: Servicios de aplicación

#### 3. **Domain Layer** (`internal/domain/`)
- **Entities**: Modelos de dominio (Product)
- **Repository Interfaces**: Contratos para acceso a datos
- **Business Logic**: Reglas de negocio puras

#### 4. **Infrastructure Layer** (`internal/infrastructure/`)
- **Firestore Repository**: Implementación concreta del repositorio
- **Configuration**: Gestión de configuración con YAML
- **External Services**: Integraciones con servicios externos

## 🛠️ Tecnologías Implementadas

### Backend
- **Go 1.24.0**: Lenguaje de programación principal
- **MCP-Go**: Framework para implementar servidores MCP
- **Google Firestore**: Base de datos NoSQL en la nube
- **Firebase SDK**: SDK oficial de Google para Firebase
- **YAML**: Configuración de la aplicación

### Herramientas y Librerías
- **Structured Logging**: Logging estructurado con slog
- **Dependency Injection**: Inyección de dependencias manual
- **Graceful Shutdown**: Manejo elegante de cierre de servidor
- **Context Management**: Gestión de contexto para operaciones asíncronas

### Protocolos y Estándares
- **Model Context Protocol (MCP)**: Protocolo para integración con IA
- **HTTP/REST**: API RESTful para comunicación
- **JSON**: Formato de intercambio de datos
- **CSV**: Importación masiva de datos

## 🚀 Funcionalidades

### Gestión de Productos
- ✅ **Crear producto**: Registro de nuevos productos en el inventario
- ✅ **Consultar producto**: Búsqueda por SKU único
- ✅ **Actualizar producto**: Modificación de datos del producto
- ✅ **Eliminar producto**: Remoción del inventario
- ✅ **Listar productos**: Consulta paginada del catálogo

### Control de Inventario
- ✅ **Modificar stock**: Operaciones de suma y resta de inventario
- ✅ **Validación de stock**: Control de cantidades negativas
- ✅ **Trazabilidad**: Registro de fechas de creación y actualización

### Carga Masiva
- ✅ **Importación CSV**: Carga masiva desde archivos CSV
- ✅ **Catálogo inicial**: 100 smartphones comerciales pre-configurados
- ✅ **Validación de datos**: Verificación de integridad durante la importación

### Integración MCP
- ✅ **Herramientas MCP**: Exposición de funcionalidades para IA
- ✅ **Esquemas de entrada**: Validación automática de datos
- ✅ **Respuestas estructuradas**: Formato JSON consistente

## 📁 Estructura del Proyecto

```
FireStoreAPI/
├── cmd/
│   └── main.go                    # Punto de entrada de la aplicación
├── configs/
│   ├── config.yaml               # Configuración de la aplicación
│   └── smartphones_catalog.csv   # Catálogo inicial de productos
├── internal/
│   ├── application/
│   │   ├── dto/                  # Data Transfer Objects
│   │   └── inventory/           # Casos de uso de inventario
│   ├── domain/
│   │   ├── entities/            # Entidades de dominio
│   │   └── repository/         # Interfaces de repositorio
│   ├── infrastructure/
│   │   ├── config/             # Gestión de configuración
│   │   └── repository/         # Implementaciones de repositorio
│   └── interfaces/
│       ├── mcp/                # Herramientas y handlers MCP
│       └── middleware/         # Middleware de la aplicación
├── pkg/
│   └── utils/                  # Utilidades compartidas
├── go.mod                      # Dependencias de Go
├── go.sum                      # Checksums de dependencias
├── Dockerfile                  # Configuración de Docker
└── README.md                   # Documentación del proyecto
```

## ⚙️ Configuración

### Archivo de Configuración (`configs/config.yaml`)

```yaml
server:
  host: "localhost"
  port: 8080

database:
  project_id: "technology-cloud-75d1f"
```

### Variables de Entorno Requeridas

```bash
# Firebase/Firestore
GOOGLE_APPLICATION_CREDENTIALS="path/to/service-account.json"
FIREBASE_PROJECT_ID="your-project-id"

# Servidor
SERVER_HOST="localhost"
SERVER_PORT="8080"
```

## 🚀 Instalación y Uso

### Prerrequisitos

1. **Go 1.24.0+** instalado
2. **Cuenta de Google Cloud** con Firestore habilitado
3. **Archivo de credenciales** de servicio de Firebase

### Instalación

1. **Clonar el repositorio**:
```bash
git clone <repository-url>
cd FireStoreAPI
```

2. **Instalar dependencias**:
```bash
go mod download
```

3. **Configurar Firebase**:
```bash
# Descargar credenciales de Firebase Console
export GOOGLE_APPLICATION_CREDENTIALS="path/to/service-account.json"
```

4. **Ejecutar la aplicación**:
```bash
go run cmd/main.go
```

### Uso con Docker

```bash
# Construir imagen
docker build -t firestore-api .

# Ejecutar contenedor
docker run -p 8080:8080 \
  -e GOOGLE_APPLICATION_CREDENTIALS=/app/credentials.json \
  -v /path/to/credentials.json:/app/credentials.json \
  firestore-api
```

## 🔧 Herramientas MCP Disponibles

### `create_product`
Crea un nuevo producto en el inventario.

**Parámetros**:
```json
{
  "sku": "TEL-APP-I15P-256",
  "active": true,
  "category": "Celulares",
  "image_url": "https://example.com/image.jpg",
  "name": "iPhone 15 Pro 256GB",
  "name_provider": "Apple",
  "price": 4500000
}
```

### `change_stock_product`
Modifica el stock de un producto existente.

**Parámetros**:
```json
{
  "sku": "TEL-APP-I15P-256",
  "stock": 10,
  "type": "Sum"  // o "Rest"
}
```

### `update_product`
Actualiza información de un producto existente.

**Parámetros**:
```json
{
  "sku": "TEL-APP-I15P-256",
  "name": "iPhone 15 Pro 256GB - Actualizado",
  "price": 4600000,
  "active": true,
  "image_url": "https://example.com/new-image.jpg"
}
```

### `delete_product`
Elimina un producto del inventario.

**Parámetros**:
```json
{
  "sku": "TEL-APP-I15P-256"
}
```

### `get_product_by_sku`
Consulta un producto específico por su SKU.

**Parámetros**:
```json
{
  "sku": "TEL-APP-I15P-256"
}
```

### `get_all_products_pagination`
Obtiene una lista paginada de productos.

**Parámetros**:
```json
{
  "limit": 10,
  "offset": 0
}
```

### `initial_product_data`
Carga datos iniciales desde un archivo CSV.

**Parámetros**:
```json
{
  "path": "configs/smartphones_catalog.csv"  // opcional
}
```

## 📊 Modelo de Datos

### Entidad Product

```go
type Product struct {
    Sku          string    `firestore:"sku"`
    Active       bool      `firestore:"activo"`
    Category     string    `firestore:"categoria"`
    UpdatedAt    time.Time `firestore:"fechaActualizacion"`
    CreatedAt    time.Time `firestore:"fechaCreacion"`
    ImageUrl     string    `firestore:"imagenUrl"`
    Name         string    `firestore:"nombre"`
    NameProvider string    `firestore:"nombreProveedor"`
    Price        float64   `firestore:"precio"`
    Stock        int       `firestore:"stock"`
}
```

## 🔒 Seguridad

- **Autenticación**: Integración con Firebase Auth
- **Autorización**: Control de acceso basado en roles
- **Validación**: Validación de entrada en todos los endpoints
- **Logging**: Registro de todas las operaciones críticas

## 📈 Monitoreo y Logging

- **Structured Logging**: Logs estructurados con contexto
- **Métricas**: Monitoreo de operaciones de base de datos
- **Trazabilidad**: Seguimiento de requests y responses
- **Health Checks**: Endpoints de salud del servicio

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

## 👥 Autores

- **Desarrollador Principal**: [Tu Nombre]
- **Arquitecto de Software**: [Tu Nombre]

## 📞 Soporte

Para soporte técnico o consultas sobre el proyecto:

- **Email**: [tu-email@ejemplo.com]
- **Issues**: [GitHub Issues](https://github.com/tu-usuario/FireStoreAPI/issues)
- **Documentación**: [Wiki del Proyecto](https://github.com/tu-usuario/FireStoreAPI/wiki)

---

**FireStoreAPI** - Potenciando el comercio electrónico con tecnología de vanguardia 🚀
