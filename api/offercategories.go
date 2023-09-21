package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type OffersCategory struct {
	ID      int32       `json:"id"`
	DealID  pgtype.Int4 `json:"deal_id"`
	OfferID pgtype.Int4 `json:"offer_id"`
}
func (server *Server) createOfferCategory (w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	offercategory := OffersCategory{}
	err := json.NewDecoder(r.Body).Decode(&offercategory)

	if err != nil {
		fmt.Println("------error1------", err)

		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateOfferCategoryParams{
		DealID: offercategory.DealID,
		OfferID: offercategory.OfferID,
	}

	offerCategoryInfo, err := server.store.CreateOfferCategory(ctx,arg)
	if err != nil {
		fmt.Println("------error2------", err)
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offerCategoryInfo)

}
