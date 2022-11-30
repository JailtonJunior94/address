package dtos

type ProviderResponse struct {
	Name string `json:"name"`
}

func NewProviderResponse(name string) *ProviderResponse {
	return &ProviderResponse{
		Name: name,
	}
}
