package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)


type StoreLocationOffer struct {
	ID              int32          `json:"id"`
	StoreLocationID int32          `json:"store_location_id"`
	ProductID       int32          `json:"product_id"`
	OfferID         int32          `json:"offer_id"`
	Price           pgtype.Numeric `json:"price"`
	OfferedPrice    pgtype.Numeric `json:"offered_price"`
	IsPercent       pgtype.Bool    `json:"is_percent"`
	StartDate       pgtype.Date    `json:"start_date"`
	EndDate         pgtype.Date    `json:"end_date"`
	Sequence        pgtype.Int4    `json:"sequence"`
	Isblocked       pgtype.Bool    `json:"isblocked"`
	Isdeleted       pgtype.Bool    `json:"isdeleted"`
	Priority        pgtype.Bool    `json:"priority"`
	Weight          pgtype.Text    `json:"weight"`
	UnitID          int32          `json:"unit_id"`
	CreatedAt       time.Time      `json:"created_at"`
}

func (server *Server) createStoreLocationOffers(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	storelocationoffer := StoreLocationOffer{}
	err := json.NewDecoder(r.Body).Decode(&storelocationoffer)

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

	arg := db.CreateStoreLocationOfferParams{
		StoreLocationID: storelocationoffer.StoreLocationID,
		ProductID: storelocationoffer.ProductID,
		OfferID: storelocationoffer.OfferID,
		Price: storelocationoffer.Price,
		OfferedPrice: storelocationoffer.OfferedPrice,
		IsPercent: storelocationoffer.IsPercent,
		StartDate: storelocationoffer.StartDate,
		EndDate: storelocationoffer.EndDate,
		Sequence: storelocationoffer.Sequence,
		Isblocked: storelocationoffer.Isblocked,
		Isdeleted: storelocationoffer.Isdeleted,
		Priority: storelocationoffer.Priority,
		Weight: storelocationoffer.Weight,
		UnitID: storelocationoffer.UnitID,
	}

	storeLocationOfferInfo, err := server.store.CreateStoreLocationOffer(ctx,arg)
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
	json.NewEncoder(w).Encode(storeLocationOfferInfo)

}
