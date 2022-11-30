package services

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jailtonjunior94/address/internal/dtos"
)

type CorreiosService struct {
	HTTPClient *http.Client
}

func NewCorreiosService() *CorreiosService {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	return &CorreiosService{HTTPClient: client}
}

func (s *CorreiosService) AddressByCEP(cep string) (*dtos.AddressResponse, error) {
	url := "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente?wsdl"
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
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/soap+xml;charset=utf-8")
	req.Header.Set("cache-control", "no-cache")

	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		b, _ := ioutil.ReadAll(res.Body)
		return nil, errors.New(fmt.Sprintf("[ERROR] [StatusCode] [%d] [Detail] [%s]", res.StatusCode, string(b)))
	}

	var result *dtos.CorreiosResponse
	err = xml.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	provider := dtos.NewProviderResponse("Correios")
	response := dtos.NewAddressResponse(result.Body.ConsultaCEPResponse.Return.Cep,
		result.Body.ConsultaCEPResponse.Return.End,
		result.Body.ConsultaCEPResponse.Return.Bairro,
		result.Body.ConsultaCEPResponse.Return.Cidade,
		result.Body.ConsultaCEPResponse.Return.Uf,
		provider)
	return response, nil
}
