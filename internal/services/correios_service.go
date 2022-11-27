package services

import (
	"bytes"
	"encoding/xml"
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
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/soap+xml;charset=utf-8")
	req.Header.Set("cache-control", "no-cache")

	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result *dtos.Envelope
	err = xml.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	response := dtos.NewAddressResponse(result.Body.ConsultaCEPResponse.Return.Cep,
		result.Body.ConsultaCEPResponse.Return.End,
		result.Body.ConsultaCEPResponse.Return.Bairro,
		result.Body.ConsultaCEPResponse.Return.Cidade,
		result.Body.ConsultaCEPResponse.Return.Uf)
	return response, nil
}
