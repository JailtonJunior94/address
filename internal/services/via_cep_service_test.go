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

type ViaCepServicesTestSuite struct {
	suite.Suite

	config *configs.Config
}

func TestViaCepServicesTestSuite(t *testing.T) {
	suite.Run(t, new(ViaCepServicesTestSuite))
}

func (s *ViaCepServicesTestSuite) SetupTest() {
	s.config = &configs.Config{
		ViaCepBaseURL: "http://localhost:3000/address/%s",
	}
}

func (s *ViaCepServicesTestSuite) TestAddressByCEP() {
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
		Body: io.NopCloser(strings.NewReader(`{
		  "cep": "06503-015",
		  "logradouro": "Rua José Pontes Zé Buraco",
		  "complemento": "",
		  "bairro": "Parque Fernão Dias",
		  "localidade": "Santana de Parnaíba",
		  "uf": "SP",
		  "ibge": "3547304",
		  "gia": "6233",
		  "ddd": "11",
		  "siafi": "7047"
		}`)),
	}

	respondeInvalid := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body: io.NopCloser(strings.NewReader(`{
			"errorMessage": "error",
		  }`)),
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
			service := NewViaCepService(s.config, scenario.fields.httpClient)
			address, err := service.AddressByCEP(scenario.args.cep)
			scenario.expected(address, err)
		})
	}
}
