package entities

import "time"

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
