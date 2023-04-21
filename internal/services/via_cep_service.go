package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jailtonjunior94/address/configs"
	"github.com/jailtonjunior94/address/internal/dtos"
	"github.com/jailtonjunior94/address/internal/interfaces"
	"go.uber.org/zap"
)

type viaCepService struct {
	config     *configs.Config
	logger     *zap.SugaredLogger
	httpClient interfaces.HttpClient
}

func NewViaCepService(config *configs.Config, logger *zap.SugaredLogger, httpClient interfaces.HttpClient) *viaCepService {
	return &viaCepService{config: config, logger: logger, httpClient: httpClient}
}

func (s *viaCepService) AddressByCEP(cep string) (*dtos.AddressResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(s.config.ViaCepBaseURL, cep), nil)
	if err != nil {
		s.logger.Errorw("could not make to request", zap.Error(err))
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")

	res, err := s.httpClient.Do(req)
	if err != nil {
		s.logger.Errorw("could not make to request", zap.Error(err))
		return nil, err
	}

	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		b, _ := io.ReadAll(res.Body)
		s.logger.Errorw("could not make to request", zap.Error(err))
		return nil, fmt.Errorf("[ERROR] [StatusCode] [%d] [Detail] [%s]", res.StatusCode, string(b))
	}

	var result *dtos.ViaCepResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		s.logger.Errorw("could not decoder json", zap.Error(err))
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
