package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/italolelis/epioca/service/pkg/domain/auction"
	"github.com/italolelis/epioca/service/pkg/domain/bid"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

// Bidding holds all handlers for an bid
type Bidding struct {
	repo        bid.Repository
	auctionRepo auction.Repository
}

// NewBidding creates a new instance of Bidding
func NewBidding(repo bid.Repository, auctionRepo auction.Repository) *Bidding {
	return &Bidding{repo, auctionRepo}
}

// ShowByAuction handler to list biddings for an auction
func (h *Bidding) ShowByAuction(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "auctionId"))
	if err != nil {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusBadRequest, "The provided auction ID is invalid")
		return
	}

	threshold := r.URL.Query().Get("threshold")

	var bids bid.Bids

	if threshold == "" {
		bids, err = h.repo.FindByAuction(id)
	} else {
		bids, err = h.repo.FindByAuctionAndQuery(id, map[string]interface{}{"threshold": threshold})
	}

	if err != nil {
		JSON(w, http.StatusInternalServerError, "Failed during searching latest bids")
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
	bid.Created = time.Now()
	if err := json.NewDecoder(r.Body).Decode(bid); err != nil {
		log.WithError(err).Warn("Failed to decode body request during creating a bid")
		JSON(w, http.StatusBadRequest, "Failed to decode body request")
		return
	}

	auction, err := h.auctionRepo.FindByID(bid.AuctionID)
	if err != nil {
		log.WithError(err).WithField("auction ID", bid.AuctionID).Warn("Could not find an auction")
		JSON(w, http.StatusNotFound, "Could not find an auction")
		return
	}

	if auction.IsCompleted() || auction.IsScheduled() {
		JSON(w, http.StatusBadRequest, "You cannot bid for already completed or scheduled auction")
		return
	}

	latestBid, err := h.repo.FindLatestByAuctionUser(bid.AuctionID, bid.User.ID)
	if err != nil && errors.Cause(err) != sql.ErrNoRows {
		log.WithError(err).Error("")
		JSON(w, http.StatusInternalServerError, "Failed during searching bids")
		return
	}
	if errors.Cause(err) != sql.ErrNoRows {
		if latestBid.Value <= bid.Value {
			JSON(w, http.StatusInternalServerError, "Your bid value should be less than previous one")
			return
		}

		if auction.MaxPrice < bid.Value {
			JSON(w, http.StatusInternalServerError, "Your bid value cannot be higher than maximum value defined per auction")
			return
		}
	}
	if err := h.repo.Add(bid); err != nil {
		log.WithError(err).Error("Failed to create a bid")
		JSON(w, http.StatusInternalServerError, "Failed to create a bid")
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/auctions/%s/bids", auctionID))
	w.WriteHeader(http.StatusCreated)
}

// ShowByAuction handler to list biddings for an auction
func (h *Bidding) Winners(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "auctionId"))
	if err != nil {
		log.WithError(err).Warn("Failed to parse the UUID")
		JSON(w, http.StatusBadRequest, "The provided auction ID is invalid")
		return
	}

	bids, err := h.repo.FindByAuction(id)
	if err != nil {
		JSON(w, http.StatusInternalServerError, "Failed during searching latest bids")
		return
	}

	winningBids, err := auction.FindWinningBids(bids, 100)
	if err != nil {
		JSON(w, http.StatusInternalServerError, "Failed to get the winning bids")
		return
	}

	JSON(w, http.StatusOK, winningBids)
}
