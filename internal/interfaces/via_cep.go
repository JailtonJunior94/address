package interfaces

import "github.com/jailtonjunior94/address/internal/dtos"

type ViaCepService interface {
	AddressByCEP(cep string) (*dtos.AddressResponse, error)
}
