package dtos

type ViaCepResponse struct {
	Bairro      string `json:"bairro"`
	Cep         string `json:"cep"`
	Complemento string `json:"complemento"`
	Ddd         string `json:"ddd"`
	Gia         string `json:"gia"`
	Ibge        string `json:"ibge"`
	Localidade  string `json:"localidade"`
	Logradouro  string `json:"logradouro"`
	Uf          string `json:"uf"`
}
