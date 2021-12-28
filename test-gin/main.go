package main

import (
	"flag"
	"fmt"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"log"
	"main/api"
	"main/controller"
	"net/http"
	"os"
)

func NewGinPetServer(getPing controller.GetPingRequest, port int) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic chi router
	r := gin.Default()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	r = api.RegisterHandlers(r, &getPing)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}
	return s
}

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()
	// Create an instance of our handler which satisfies the generated interface

	getPing := controller.NewPingRequest()
	s := NewGinPetServer(getPing, *port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
