package handlers

import (
	"net/http"

	"github.com/jailtonjunior94/address/internal/interfaces"
	"github.com/jailtonjunior94/address/pkg/responses"

	"github.com/go-chi/chi/v5"
)

type AddressHandler struct {
	CorreiosService interfaces.CorreiosService
}

func NewAdressHandler(c interfaces.CorreiosService) *AddressHandler {
	return &AddressHandler{CorreiosService: c}
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

	address, err := h.CorreiosService.AddressByCEP(cep)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Não foi possível obter endereço através do CEP informado")
		return
	}

	responses.JSON(w, http.StatusOK, address)
}
