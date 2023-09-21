package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type OfferView struct {
	ID              int32       `json:"id"`
	ProductID       int32       `json:"product_id"`
	OfferID         int32       `json:"offer_id"`
	MallID          int32       `json:"mall_id"`
	StoreID         int32       `json:"store_id"`
	StoreLocationID int32       `json:"store_location_id"`
	UserID          int32       `json:"user_id"`
	Viewed          pgtype.Int4 `json:"viewed"`
	ViewDate        pgtype.Date   `json:"view_date"`
	IsClicked       pgtype.Bool `json:"is_clicked"`
	CreatedAt       time.Time   `json:"created_at"`
}

func (server *Server) createOfferView(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	offerview := OfferView{}
	err := json.NewDecoder(r.Body).Decode(&offerview)

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

	arg := db.CreateOfferViewParams{
		ProductID: offerview.ProductID,
		OfferID: offerview.OfferID,
		MallID: offerview.MallID,
		StoreID: offerview.StoreID,
		StoreLocationID: offerview.StoreLocationID,
		UserID: offerview.StoreLocationID,
		Viewed: offerview.Viewed,
		ViewDate: offerview.ViewDate,
		IsClicked: offerview.IsClicked,

	}

	offerViewInfo, err := server.store.CreateOfferView(ctx,arg)
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
	json.NewEncoder(w).Encode(offerViewInfo)

}

