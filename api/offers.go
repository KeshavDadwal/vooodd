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

type Offer struct {
	ID          int32       `json:"id"`
	DealNo      pgtype.Text `json:"deal_no"`
	StoreID     pgtype.Int4 `json:"store_id"`
	ProductID   pgtype.Int4 `json:"product_id"`
	Name        string      `json:"name"`
	Path        pgtype.Text `json:"path"`
	Description pgtype.Text `json:"description"`
	Isblocked   pgtype.Bool `json:"isblocked"`
	Isdeleted   pgtype.Bool `json:"isdeleted"`
	UserID      pgtype.Int4 `json:"user_id"`
}

func (server *Server) createOffer (w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	offer := Offer{}
	err := json.NewDecoder(r.Body).Decode(&offer)

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

	arg := db.CreateOfferParams{
		Name: offer.Name,
		Path: offer.Path,
		DealNo: offer.DealNo,
		StoreID: offer.StoreID,
		ProductID: offer.ProductID,
		Description: offer.Description,
		Isblocked: offer.Isblocked,
		Isdeleted: offer.Isdeleted,
		UserID: offer.UserID,
	}

	offerInfo, err := server.store.CreateOffer(ctx,arg)
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
	json.NewEncoder(w).Encode(offerInfo)

}
