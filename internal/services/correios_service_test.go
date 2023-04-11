package services

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (m *MockHTTP) Do(_ *http.Request) (*http.Response, error) {
	return m.response, m.err
}

type MockHTTP struct {
	response *http.Response
	err      error
}

func TestCorreiosAddressByCEP(t *testing.T) {
	res := http.Response{
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

	mock := &MockHTTP{
		response: &res,
		err:      nil,
	}

	service := NewCorreiosService(mock)

	response, err := service.AddressByCEP("06503015")
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "06503015", response.CEP)
}

func TestCorreiosAddressByCEPWithInvalidCEP(t *testing.T) {
	res := http.Response{
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

	mock := &MockHTTP{
		response: &res,
		err:      nil,
	}

	service := NewCorreiosService(mock)

	response, err := service.AddressByCEP("0657878703015")
	assert.NotNil(t, err)
	assert.Nil(t, response)
}
