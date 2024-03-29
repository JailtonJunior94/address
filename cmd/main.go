package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/jailtonjunior94/address/configs"
	_ "github.com/jailtonjunior94/address/docs"
	"github.com/jailtonjunior94/address/internal/handlers"
	"github.com/jailtonjunior94/address/internal/interfaces"
	"github.com/jailtonjunior94/address/internal/services"
	"github.com/jailtonjunior94/address/pkg/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
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

// @BasePath	/
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(config)
	defer logger.Sync()

	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/health"))

	httpClient := interfaces.NewHttpClient(config)
	viaCepService := services.NewViaCepService(config, logger, httpClient)
	correiosService := services.NewCorreiosService(config, logger, httpClient)
	addressHandler := handlers.NewAdressHandler(correiosService, viaCepService)

	router.Get("/address/{cep}", addressHandler.Address)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", config.HttpServerPort)),
	))

	server := http.Server{
		ReadTimeout:       time.Duration(10) * time.Second,
		ReadHeaderTimeout: time.Duration(10) * time.Second,
		Handler:           router,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.HttpServerPort))
	if err != nil {
		panic(err)
	}
	server.Serve(listener)
}
