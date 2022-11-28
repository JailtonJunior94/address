package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jailtonjunior94/address/internal/dtos"
)

type ViaCepService struct {
	HTTPClient *http.Client
}

func NewViaCepService() *ViaCepService {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	return &ViaCepService{HTTPClient: client}
}

func (s *ViaCepService) AddressByCEP(cep string) (*dtos.AddressResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")

	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result *dtos.ViaCepResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	response := dtos.NewAddressResponse(result.Cep,
		result.Logradouro,
		result.Bairro,
		result.Localidade,
		result.Uf)
	return response, nil
}
