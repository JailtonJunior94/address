package handlers

import (
	"log"
	"net/http"

	"github.com/jailtonjunior94/address/internal/dtos"
	"github.com/jailtonjunior94/address/internal/interfaces"
	"github.com/jailtonjunior94/address/pkg/responses"

	"github.com/go-chi/chi/v5"
)

type AddressHandler struct {
	CorreiosService interfaces.CorreiosService
	ViaCepService   interfaces.ViaCepService
}

func NewAdressHandler(c interfaces.CorreiosService, v interfaces.ViaCepService) *AddressHandler {
	return &AddressHandler{CorreiosService: c, ViaCepService: v}
}

// Get Address By CEP godoc
// @Summary     Get Address By CEP
// @Description Get Address By CEP
// @Tags        Address
// @Accept      json
// @Produce     json
// @Param       cep path     string true "CEP"
// @Success     200 {object} dtos.AddressResponse
// @Failure     404 {object} dtos.Error
// @Failure     500 {object} dtos.Error
// @Router      /address/{cep} [get]
func (h *AddressHandler) Address(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		responses.Error(w, http.StatusUnprocessableEntity, "CEP ausente ou mal formatado")
		return
	}

	errCh := make(chan error)
	addressCorreiosCh := make(chan *dtos.AddressResponse)
	addressViaCepCh := make(chan *dtos.AddressResponse)

	go func() {
		address, err := h.CorreiosService.AddressByCEP(cep)
		if err != nil {
			errCh <- err
		}
		addressCorreiosCh <- address
	}()

	go func() {
		address, err := h.ViaCepService.AddressByCEP(cep)
		if err != nil {
			errCh <- err
		}
		addressViaCepCh <- address
	}()

	var address *dtos.AddressResponse

	select {
	case msg := <-addressCorreiosCh:
		address = msg
	case msg := <-addressViaCepCh:
		address = msg
	case errMsg := <-errCh:
		log.Fatalln(errMsg)
	}

	responses.JSON(w, http.StatusOK, address)
}
