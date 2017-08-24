package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/italolelis/epioca/service/pkg/domain/auction"
	"github.com/italolelis/epioca/service/pkg/repo"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// Auction holds all handlers for an auction
type Auction struct {
	repo auction.Repository
}

// NewAuction creates a new instance of Auction
func NewAuction(repo auction.Repository) *Auction {
	return &Auction{repo}
}

// Index handler to list auctions
func (h *Auction) Index(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("s")

	auctions, err := h.repo.Find(status)
	if err == repo.ErrInvalidStatusProvided {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if err != nil {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusInternalServerError, "Failed during searching auctions")
		return
	}

	JSON(w, http.StatusOK, auctions)
}

// Show handler to show an auction
func (h *Auction) Show(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusBadRequest, "The provided id is invalid")
		return
	}

	auctions, err := h.repo.FindByID(id)
	if err != nil {
		log.WithError(err).Error("Failed to create an auction")
		JSON(w, http.StatusNotFound, "Failed during searching auctions")
		return
	}

	JSON(w, http.StatusOK, auctions)
}

// Create handler to create an auction
func (h *Auction) Create(w http.ResponseWriter, r *http.Request) {
	act := auction.NewAuction(uuid.NewV4())
	if err := json.NewDecoder(r.Body).Decode(act); err != nil {
		log.WithError(err).Warn("Failed to decode body request during creating an auction")
		JSON(w, http.StatusBadRequest, "Failed to decode body request")
		return
	}

	if err := h.repo.Add(act); err != nil {
		log.WithError(err).Error("Failed to create an auction")
		JSON(w, http.StatusInternalServerError, "FFailed to create an auction")
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/auctions/%s", act.ID))
	w.WriteHeader(http.StatusCreated)
}

// Update handler to create an auction
func (h *Auction) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusBadRequest, "The provided id is invalid")
		return
	}

	auction, err := h.repo.FindByID(id)
	if err != nil {
		log.WithError(err).WithField("id", id).Warn("Could not find an auction")
		JSON(w, http.StatusNotFound, "Could not find an auction")
		return
	}

	if !auction.IsScheduled() {
		JSON(w, http.StatusBadRequest, "You cannot change an auction that is running or completed")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(auction); err != nil {
		log.WithError(err).Warn("Failed to decode body request during creating an auction")
		JSON(w, http.StatusBadRequest, "Failed to decode body request")
		return
	}

	if err := h.repo.Add(auction); err != nil {
		log.WithError(err).Error("Failed to update an auction")
		JSON(w, http.StatusInternalServerError, "Failed to update an auction")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Remove handler to remove an auction
func (h *Auction) Remove(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusBadRequest, "The provided id is invalid")
		return
	}

	if err := h.repo.Remove(id); err != nil {
		log.WithError(err).Error("Failed to remove an auction")
		JSON(w, http.StatusInternalServerError, "Failed to remove an auction")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
