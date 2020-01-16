package main

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/veluts77/http-go-server/pkg/swagger/server/restapi"
	"github.com/veluts77/http-go-server/pkg/swagger/server/restapi/operations"
	"log"
)

func main() {
	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")

	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	// Implement the CheckHealth handler
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(checkHealthHandler)

	// Implement the GetHelloUser handler
	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(helloHandler)

	// Start listening using having the handlers and port
	// already set up.
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func checkHealthHandler(user operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

func helloHandler(user operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
}
