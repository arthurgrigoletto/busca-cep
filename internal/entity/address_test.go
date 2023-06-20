package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	a := NewAddress("Avenida Paulista", "Bela Vista", "São Paulo", "SP", "12345012")

	assert.NotNil(t, a)
	assert.Equal(t, "Avenida Paulista", a.Street)
	assert.Equal(t, "Bela Vista", a.Neighborhood)
	assert.Equal(t, "São Paulo", a.City)
	assert.Equal(t, "SP", a.State)
	assert.Equal(t, "12345012", a.Zip)
}
