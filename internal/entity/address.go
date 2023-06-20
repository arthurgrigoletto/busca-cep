package entity

import (
	"github.com/arthurgrigoletto/busca-cep/pkg/entity"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type APICep struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type Address struct {
	id           entity.ID `json:"-"`
	Zip          string    `json:"zip"`
	Street       string    `json:"street"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	State        string    `json:"state"`
}

func NewAddress(street string, neighborhood string, city string, state string, zip string) *Address {
	address := &Address{
		id:           entity.NewID(),
		Zip:          zip,
		Street:       street,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
	}

	return address
}
