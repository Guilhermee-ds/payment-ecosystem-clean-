package http

import (
	"encoding/json"
	"io"
	"log/slog"
	stdhttp "net/http"
	"payment-ecosystem-clean/api/internal/domain"
	"payment-ecosystem-clean/api/internal/usecase"
)

type Handler struct {
	ingest *usecase.IngestService
	logger *slog.Logger
}

func NewHandler(ing *usecase.IngestService) *Handler { return &Handler{ingest: ing} }

func (h *Handler) Payment(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	var req domain.Payment

	body, err := io.ReadAll(r.Body)
	if err != nil {
		stdhttp.Error(w, "invalid body", stdhttp.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		stdhttp.Error(w, "invalid json", stdhttp.StatusBadRequest)
		return
	}

	id, err := h.ingest.Ingest(r.Context(), &req)
	if err != nil {
		stdhttp.Error(w, "enqueue error", stdhttp.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(stdhttp.StatusAccepted)

	resp := map[string]string{
		"id":     id,
		"status": "accepted",
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error("failed to write response", "error", err)
	}
}
