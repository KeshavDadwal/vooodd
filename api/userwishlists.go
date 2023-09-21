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

type UserWishlist struct {
	ID              int32         `json:"id"`
	UserID          int32         `json:"user_id"`
	ProductID       int32         `json:"product_id"`
	StoreLocationID pgtype.Int4   `json:"store_location_id"`
	WishPrice       float64       `json:"wish_price"`
	MaxPrice        float64       `json:"max_price"`
	Immediately     pgtype.Bool   `json:"immediately"`
	DateTill        pgtype.Date   `json:"date_till"`
	CreatedAt       time.Time     `json:"created_at"`
}

func (server *Server) createUserWishlistItem(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	userwishlist := UserWishlist{}
	err := json.NewDecoder(r.Body).Decode(&userwishlist)

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

	arg := db.CreateUserWishlistItemParams{
		UserID:  userwishlist.UserID,
		ProductID: userwishlist.ProductID,
		StoreLocationID:  userwishlist.StoreLocationID,
		WishPrice: userwishlist.WishPrice,
		MaxPrice: userwishlist.MaxPrice,
		Immediately: userwishlist.Immediately,
		DateTill: userwishlist.DateTill,
	}

	userWishlistInfo, err := server.store.CreateUserWishlistItem(ctx,arg)
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
	json.NewEncoder(w).Encode(userWishlistInfo)

}
