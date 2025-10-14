package tools

import (
	usecase "mcp-server/internal/application/inventory"
)

type ManagmentTools struct {
	inventoryUseCase *usecase.InventoryCases
}

func NewManagmentTools(inventory *usecase.InventoryCases) *ManagmentTools {
	return &ManagmentTools{
		inventoryUseCase: inventory,
	}
}
