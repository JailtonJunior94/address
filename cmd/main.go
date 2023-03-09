package main

import (
	"fmt"
	"net/http"

	"github.com/jailtonjunior94/address/internal/handlers"
	"github.com/jailtonjunior94/address/internal/interfaces"
	"github.com/jailtonjunior94/address/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jailtonjunior94/address/docs"
	swagger "github.com/swaggo/http-swagger"
)

//	@title			Address API
//	@version		1.0
//	@description	Address API
//	@termsOfService	http://swagger.io/terms

//	@contact.name	Jailton Junior
//	@contact.url	http://jailton.junior.net
//	@contact.email	jailton.junior94@outlook.com

//	@license.name	Jailton Junior License
//	@license.url	http://jailton.junior.net

//	@BasePath	/
func main() {
	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/health"))

	httpClient := interfaces.NewHttpClient()
	viaCepService := services.NewViaCepService()
	correiosService := services.NewCorreiosService(httpClient)
	addressHandler := handlers.NewAdressHandler(correiosService, viaCepService)

	router.Get("/address/{cep}", addressHandler.Address)

	router.Get("/docs/*", swagger.Handler(swagger.URL("http://localhost:8000/docs/doc.json")))
	fmt.Printf("🚀 API is running on http://localhost:%v", "8000")
	http.ListenAndServe(":8000", router)
}
