package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/schattenbrot/explerror"
	"github.com/schattenbrot/go-pizza-api/pkgs/responder"
)

// statusResponse represents the server status.
// swagger:model
type statusResponse struct {
	Status      string    `json:"status"`
	StartUptime time.Time `json:"startUptime"`
	Uptime      time.Time `json:"uptime"`
	Environment string    `json:"environment"`
	Version     string    `json:"version"`
}

// @Summary Get server status
// @Description Retrieves the running server status
// @Tags app
// @Produce json
// @Success 200 {object} app.statusResponse
// @Router /app/status [get]
// @Router /app [get]
func status(w http.ResponseWriter, r *http.Request) {
	explerror.NotImplemented(w, errors.New("status handler not implemented yet"))
}

// @Summary Ping the server
// @Description Responds with status 204 No Content
// @Tags app
// @Produce json
// @Success 204
// @Router /app/ping [get]
func ping(w http.ResponseWriter, r *http.Request) {
	responder.Send(w, http.StatusNoContent)
}
