package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/arthurgrigoletto/busca-cep/internal/entity"
	"github.com/go-chi/chi/v5"
)

type CepHandler struct{}

func NewCepHandler() *CepHandler {
	return &CepHandler{}
}

// Get CEP info godoc
// @Summary      Get CEP
// @Description  Get CEP
// @Tags         cep
// @Accept       json
// @Produce      json
// @Param        cep  	path      string  true  "CEP"
// @Success      200		{object} entity.Address
// @Router       /search/{cep}			[get]
func (h *CepHandler) GetCepInfo(w http.ResponseWriter, r *http.Request) {
	viaCep := make(chan entity.Address)
	apiCep := make(chan entity.Address)
	cepParam := chi.URLParam(r, "cep")
	cep := cepParam

	if !strings.Contains(cepParam, "-") {
		cep = cepParam[0:5] + "-" + cepParam[5:8]
	}

	go func() {
		cep, _ := ViaCep(cep)
		address := entity.NewAddress(cep.Logradouro, cep.Bairro, cep.Localidade, cep.Uf, cep.Cep)
		viaCep <- *address
	}()

	go func() {
		cep, _ := ApiCep(cep)
		address := entity.NewAddress(cep.Address, cep.District, cep.City, cep.State, cep.Code)
		apiCep <- *address
	}()

	select {
	case address := <-viaCep:
		fmt.Printf("Provider: ViaCEP - Result: %s, %s - %s/%s - %s", address.Street, address.Neighborhood, address.City, address.State, address.Zip)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(address)
	case address := <-apiCep:
		fmt.Printf("Provider: API Cep - Result: %s, %s - %s/%s - %s", address.Street, address.Neighborhood, address.City, address.State, address.Zip)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(address)
	case <-time.After(time.Second):
		w.WriteHeader(http.StatusRequestTimeout)
	}
}

func ViaCep(cep string) (*entity.ViaCep, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c entity.ViaCep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func ApiCep(cep string) (*entity.APICep, error) {
	resp, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c entity.APICep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	if c.Status != 200 {
		return nil, errors.New("address not found")
	}

	return &c, nil
}
