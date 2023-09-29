package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kokopelli-inc/echo-ddd-demo/internal/domain/entity"
)

func TestNewRegisterMenuMaster(t *testing.T) {
	menuMaster := entity.NewRegisterMenuMaster(
		"bankID",
		"Name",
		"CreatedBy",
		"UpdatedBy",
	)

	assert.Equal(t, "bankID", menuMaster.BankID)
	assert.Equal(t, "Name", menuMaster.Name)
	assert.Equal(t, "CreatedBy", menuMaster.CreatedBy)
	assert.Equal(t, "UpdatedBy", menuMaster.UpdatedBy)
}
