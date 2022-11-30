package dtos

type AddressResponse struct {
	CEP          string            `json:"cep"`
	Street       string            `json:"street"`
	Neighborhood string            `json:"neighborhood"`
	City         string            `json:"city"`
	State        string            `json:"state"`
	Provider     *ProviderResponse `json:"provider"`
}

func NewAddressResponse(cep, street, neighborhood, city, state string, provider *ProviderResponse) *AddressResponse {
	return &AddressResponse{
		CEP:          cep,
		Street:       street,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
		Provider:     provider,
	}
}
