package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/italolelis/epioca/service/pkg/domain/bid"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// Bidding holds all handlers for an bid
type Bidding struct {
	repo bid.Repository
}

// NewBidding creates a new instance of Bidding
func NewBidding(repo bid.Repository) *Bidding {
	return &Bidding{repo}
}

// ShowByAuction handler to list biddings for an auction
func (h *Bidding) ShowByAuction(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "auctionId"))
	if err != nil {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusBadRequest, "The provided auction ID is invalid")
		return
	}

	bids, err := h.repo.FindByAuction(id)
	if err != nil {
		JSON(w, http.StatusInternalServerError, "Failed during searching bids")
		return
	}

	JSON(w, http.StatusOK, bids)
}

// ShowByLowest handler to show the lowest bidding for an auction
func (h *Bidding) ShowByLowest(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "auctionId"))
	if err != nil {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusBadRequest, "The provided auction ID is invalid")
		return
	}

	bids, err := h.repo.FindLowest(id)
	if err != nil {
		JSON(w, http.StatusInternalServerError, "Failed during searching bids")
		return
	}

	JSON(w, http.StatusOK, bids)
}

// Create handler to create an auction
func (h *Bidding) Create(w http.ResponseWriter, r *http.Request) {
	auctionID, err := uuid.FromString(chi.URLParam(r, "auctionId"))
	if err != nil {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusBadRequest, "The provided auction ID is invalid")
		return
	}

	bid := bid.NewBid(uuid.NewV4())
	bid.AuctionID = auctionID
	if err := json.NewDecoder(r.Body).Decode(bid); err != nil {
		log.WithError(err).Warn("Failed to decode body request during creating a bid")
		JSON(w, http.StatusBadRequest, "Failed to decode body request")
		return
	}

	if err := h.repo.Add(bid); err != nil {
		log.WithError(err).Error("Failed to create a bid")
		JSON(w, http.StatusInternalServerError, "FFailed to create a bid")
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/auctions/%s/bids", auctionID))
	w.WriteHeader(http.StatusCreated)
}