package api

import (
	"encoding/json"
	"net/http"
	"time"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type DisLikeCount struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	OfferID   int32     `json:"offer_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) createDislikeCount(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	dislikecount := DisLikeCount{}
	err := json.NewDecoder(r.Body).Decode(&dislikecount)

	if err != nil {

		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateDislikeCountParams{
		UserID: dislikecount.UserID,
		OfferID: dislikecount.OfferID,
	}

	dislikeCountInfo, err := server.store.CreateDislikeCount(ctx,arg)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dislikeCountInfo)

}

func (server *Server) getAllDislikeCount(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    dislikeCounts, err := server.store.SelectAllDislikeCounts(ctx)
    if err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to fetch brand languages",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(dislikeCounts); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}
