package main

import (
	"net/http"

	"github.com/beharc/go-example-apps/pkg/common/health"
	"github.com/beharc/go-example-apps/pkg/common/logger"
)

func main() {
	log := logger.New()

	log.Info("Starting balance service")
	mux := http.NewServeMux()

	health.AddHealthCheck(mux)

	log.Info("Starting balance service on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
