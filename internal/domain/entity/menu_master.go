package entity

import "time"

type MenuMaster struct {
	ID        int
	BankID    string
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type ListMenuMasters []MenuMaster

func NewRegisterMenuMaster(
	bankID string,
	name string,
	createdBy string,
	updatedBy string,
) *MenuMaster {
	return &MenuMaster{
		BankID:    bankID,
		Name:      name,
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}
}
