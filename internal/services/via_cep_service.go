package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jailtonjunior94/address/internal/dtos"
)

type viaCepService struct {
	HTTPClient *http.Client
}

func NewViaCepService() *viaCepService {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	return &viaCepService{HTTPClient: client}
}

func (s *viaCepService) AddressByCEP(cep string) (*dtos.AddressResponse, error) {
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

	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("[ERROR] [StatusCode] [%d] [Detail] [%s]", res.StatusCode, string(b))
	}

	var result *dtos.ViaCepResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	provider := dtos.NewProviderResponse("ViaCEP")
	response := dtos.NewAddressResponse(result.Cep,
		result.Logradouro,
		result.Bairro,
		result.Localidade,
		result.Uf,
		provider)
	return response, nil
}
