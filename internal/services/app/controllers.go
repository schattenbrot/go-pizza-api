package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/schattenbrot/explerror"
	"github.com/schattenbrot/go-pizza-api/pkgs/responder"
)

// StatusResponse represents the server status.
// swagger:model
type StatusResponse struct {
	// Status represents the server's current status.
	// Example: available
	Status string `json:"status"`

	// StartupTime represents the server's startup time as iso string.
	// Example: 2017-04-20T11:32:00.000Z
	StartupTime time.Time `json:"startUptime"`

	// Uptime represents the server's current uptime in days:HH:mm.
	// Example: 365:16:20
	Uptime time.Time `json:"uptime"`

	// Environment represents the server's current env.
	// Example: dev
	// Example: prod
	// Example: test
	Environment string `json:"environment"`

	// Version represents the server's current version.
	// Example: 1.0.0
	Version string `json:"version"`
}

// @Summary Get server status
// @Description Retrieves the running server status
// @Tags App
// @Produce json
// @Success 200 {object} app.StatusResponse
// @Router /status [get]
// @Router / [get]
func status(w http.ResponseWriter, r *http.Request) {
	explerror.NotImplemented(w, errors.New("status handler not implemented yet"))
}

// @Summary Ping the server
// @Description Responds with status 204 No Content
// @Tags App
// @Produce json
// @Success 204
// @Router /app/ping [get]
func ping(w http.ResponseWriter, r *http.Request) {
	responder.Send(w, http.StatusNoContent)
}
