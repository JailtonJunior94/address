package services

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/jailtonjunior94/address/configs"
	"github.com/jailtonjunior94/address/internal/dtos"
	serviceMocks "github.com/jailtonjunior94/address/internal/services/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CorreiosServicesTestSuite struct {
	suite.Suite

	config *configs.Config
}

func TestCorreiosServicesTestSuite(t *testing.T) {
	suite.Run(t, new(CorreiosServicesTestSuite))
}

func (s *CorreiosServicesTestSuite) SetupTest() {
	s.config = &configs.Config{
		CorreiosBaseURL: "http://localhost:3000",
	}
}

func (s *CorreiosServicesTestSuite) TestAddressByCEP() {
	type (
		args struct {
			cep string
		}
		fields struct {
			httpClient *serviceMocks.HttpClient
		}
	)

	responseValid := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(strings.NewReader(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		<soap:Body>
			<ns2:consultaCEPResponse xmlns:ns2="http://cliente.bean.master.sigep.bsb.correios.com.br/">
				<return>
					<bairro>Parque Fernão Dias</bairro>
					<cep>06503015</cep>
					<cidade>Santana de Parnaíba</cidade>
					<complemento2></complemento2>
					<end>Rua José Pontes Zé Buraco</end>
					<uf>SP</uf>
				</return>
			</ns2:consultaCEPResponse>
		</soap:Body>
	</soap:Envelope>`)),
	}

	respondeInvalid := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body: io.NopCloser(strings.NewReader(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		<soap:Body>
			<soap:Fault>
				<faultcode>soap:Server</faultcode>
				<faultstring>CEP INVÁLIDO</faultstring>
				<detail>
					<ns2:SigepClienteException xmlns:ns2="http://cliente.bean.master.sigep.bsb.correios.com.br/">CEP INVÁLIDO</ns2:SigepClienteException>
				</detail>
			</soap:Fault>
		</soap:Body>
	</soap:Envelope>`)),
	}

	scenarios := []struct {
		name     string
		args     args
		fields   fields
		expected func(res *dtos.AddressResponse, err error)
	}{
		{
			name: "must return an address given a zip code",
			args: args{cep: "06503015"},
			fields: fields{
				httpClient: func() *serviceMocks.HttpClient {
					httpClient := &serviceMocks.HttpClient{}
					httpClient.On("Do", mock.Anything).Return(responseValid, nil)
					return httpClient
				}(),
			},
			expected: func(res *dtos.AddressResponse, err error) {
				s.NoError(err)
				s.NotNil(res)
				s.Equal("Santana de Parnaíba", res.City)
			},
		},
		{
			name: "should return error when unable to return address",
			args: args{cep: "06503015"},
			fields: fields{
				httpClient: func() *serviceMocks.HttpClient {
					httpClient := &serviceMocks.HttpClient{}
					httpClient.On("Do", mock.Anything).Return(respondeInvalid, nil)
					return httpClient
				}(),
			},
			expected: func(res *dtos.AddressResponse, err error) {
				s.NotNil(err)
				s.Nil(res)
			},
		},
	}

	for _, scenario := range scenarios {
		s.T().Run(scenario.name, func(t *testing.T) {
			service := NewCorreiosService(s.config, scenario.fields.httpClient)
			address, err := service.AddressByCEP(scenario.args.cep)
			scenario.expected(address, err)
		})
	}
}
