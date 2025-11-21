package http

import (
	"log"
	stdhttp "net/http"
	"payment-ecosystem-clean/api/internal/metrics"
)

func Start(address string, handler *Handler) error {
	mux := stdhttp.NewServeMux()
	mux.HandleFunc("/payment", handler.Payment)
	mux.Handle("/metrics", metrics.Handler())

	log.Printf("api listening on %s", address)
	return stdhttp.ListenAndServe(address, mux)
}
