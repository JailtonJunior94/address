package services

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/jailtonjunior94/address/configs"
	"github.com/jailtonjunior94/address/internal/dtos"
	"github.com/jailtonjunior94/address/internal/interfaces"
	"go.uber.org/zap"
)

type correiosService struct {
	config     *configs.Config
	logger     *zap.SugaredLogger
	httpClient interfaces.HttpClient
}

func NewCorreiosService(config *configs.Config, logger *zap.SugaredLogger, httpClient interfaces.HttpClient) *correiosService {
	return &correiosService{config: config, logger: logger, httpClient: httpClient}
}

func (s *correiosService) AddressByCEP(cep string) (*dtos.AddressResponse, error) {
	payload := `
			<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cli="http://cliente.bean.master.sigep.bsb.correios.com.br/">
				<soapenv:Header/>
				<soapenv:Body>
					<cli:consultaCEP>
						<cep>` + cep + `s</cep>
					</cli:consultaCEP>
				</soapenv:Body>
			</soapenv:Envelope>
		`
	req, err := http.NewRequest(http.MethodPost, s.config.CorreiosBaseURL, bytes.NewBufferString(payload))
	if err != nil {
		s.logger.Errorw("could not make to request", zap.Error(err))
		return nil, err
	}

	req.Header.Set("content-type", "application/soap+xml;charset=utf-8")
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

	var result *dtos.CorreiosResponse
	err = xml.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		s.logger.Errorw("could not decoder xml", zap.Error(err))
		return nil, err
	}

	provider := dtos.NewProviderResponse("Correios")
	response := dtos.NewAddressResponse(
		result.Body.ConsultaCEPResponse.Return.Cep,
		result.Body.ConsultaCEPResponse.Return.End,
		result.Body.ConsultaCEPResponse.Return.Bairro,
		result.Body.ConsultaCEPResponse.Return.Cidade,
		result.Body.ConsultaCEPResponse.Return.Uf,
		provider,
	)
	return response, nil
}
