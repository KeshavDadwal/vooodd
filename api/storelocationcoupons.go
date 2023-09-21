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

type StoreLocationCoupon struct {
	ID              int32          `json:"id"`
	StoreLocationID int32          `json:"store_location_id"`
	ProductID       int32          `json:"product_id"`
	CouponID        int32          `json:"coupon_id"`
	Price           pgtype.Numeric `json:"price"`
	CouponPrice     pgtype.Numeric `json:"coupon_price"`
	IsPercent       pgtype.Bool    `json:"is_percent"`
	StartDate       pgtype.Date    `json:"start_date"`
	EndDate         pgtype.Date    `json:"end_date"`
	Sequence        pgtype.Int4    `json:"sequence"`
	Isblocked       pgtype.Bool    `json:"isblocked"`
	Isdeleted       pgtype.Bool    `json:"isdeleted"`
	CreatedAt       time.Time      `json:"created_at"`
}

func (server *Server) createStoreLocationCoupon(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	storelocationcoupon := StoreLocationCoupon{}
	err := json.NewDecoder(r.Body).Decode(&storelocationcoupon)

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

	arg := db.CreateStoreLocationCouponParams{
		StoreLocationID: storelocationcoupon.StoreLocationID,
		ProductID: storelocationcoupon.ProductID,
		CouponID: storelocationcoupon.CouponID,
		Price: storelocationcoupon.Price,
		CouponPrice: storelocationcoupon.CouponPrice,
		IsPercent: storelocationcoupon.IsPercent,
		StartDate: storelocationcoupon.StartDate,
		EndDate: storelocationcoupon.EndDate,
	}

	storelocationcouponInfo, err := server.store.CreateStoreLocationCoupon(ctx,arg)
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
	json.NewEncoder(w).Encode(storelocationcouponInfo)

}
