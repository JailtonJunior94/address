package dtos

import "encoding/xml"

type CorreiosResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                string `xml:",chardata"`
		ConsultaCEPResponse struct {
			Text   string `xml:",chardata"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text         string `xml:",chardata"`
				Bairro       string `xml:"bairro"`
				Cep          string `xml:"cep"`
				Cidade       string `xml:"cidade"`
				Complemento2 string `xml:"complemento2"`
				End          string `xml:"end"`
				Uf           string `xml:"uf"`
			} `xml:"return"`
		} `xml:"consultaCEPResponse"`
	} `xml:"Body"`
}
