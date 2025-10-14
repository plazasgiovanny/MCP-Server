package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// GenerateID genera un ID único usando timestamp y bytes aleatorios
func GenerateID(prefix string) string {
	// Generar bytes aleatorios
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)

	// Crear ID con timestamp y bytes aleatorios
	timestamp := time.Now().UnixNano()
	id := fmt.Sprintf("%s_%d_%s", prefix, timestamp, hex.EncodeToString(randomBytes))

	return id
}

// ValidateEmail valida el formato de un email
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}

	// Validación básica
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	localPart := parts[0]
	domainPart := parts[1]

	// Validar partes
	if len(localPart) == 0 || len(domainPart) == 0 {
		return false
	}

	// Verificar que el dominio tenga al menos un punto
	if !strings.Contains(domainPart, ".") {
		return false
	}

	return true
}

// SanitizeString limpia y sanitiza una cadena de texto
func SanitizeString(input string) string {
	// Remover espacios en blanco al inicio y final
	cleaned := strings.TrimSpace(input)

	// Remover caracteres de control
	cleaned = strings.Map(func(r rune) rune {
		if r < 32 || r == 127 {
			return -1 // Remover caracter
		}
		return r
	}, cleaned)

	return cleaned
}

// Contains verifica si un slice contiene un elemento específico
func Contains[T comparable](slice []T, item T) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

// Max retorna el máximo entre dos enteros
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min retorna el mínimo entre dos enteros
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Clamp limita un valor entre un mínimo y máximo
func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
