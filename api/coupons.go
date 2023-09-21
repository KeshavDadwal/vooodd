package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	//"github.com/jackc/pgx/v5/pgtype"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type Coupon struct {
	ID          int32          `json:"id"`
	StoreID     pgtype.Int4    `json:"store_id"`
	BrandID     pgtype.Int4    `json:"brand_id"`
	ProductID   pgtype.Int4    `json:"product_id"`
	Name        string         `json:"name"`
	Code        string         `json:"code"`
	Description pgtype.Text    `json:"description"`
	Path        pgtype.Text    `json:"path"`
	Isdeleted   pgtype.Bool    `json:"isdeleted"`
	Isblocked   pgtype.Bool    `json:"isblocked"`
	UserID      int32          `json:"user_id"`
	Price       pgtype.Numeric `json:"price"`
	Ispercent   int32          `json:"ispercent"`
	StartDate   pgtype.Date    `json:"start_date"`
	EndDate     pgtype.Date    `json:"end_date"`
	UnitID      pgtype.Int4    `json:"unit_id"`
	Qty         pgtype.Int4    `json:"qty"`
	CreatedAt   time.Time      `json:"created_at"`
}
func (server *Server) createCoupon(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	coupon := Coupon{}
	err := json.NewDecoder(r.Body).Decode(&coupon)

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

	arg := db.CreateCouponParams{
		Name: coupon.Name,
		Code: coupon.Code,
		Description: coupon.Description,
		Path: coupon.Description,
		BrandID: coupon.BrandID,
		StoreID: coupon.StoreID,
		ProductID: coupon.ProductID,
		Isdeleted: coupon.Isdeleted,
		Isblocked: coupon.Isblocked,
		UserID: coupon.UserID,
		Price: coupon.Price,
		Ispercent: coupon.Ispercent,
		StartDate: coupon.StartDate,
		EndDate: coupon.EndDate,
		UnitID: coupon.UnitID,
		Qty: coupon.Qty,
	}

	couponInfo, err := server.store.CreateCoupon(ctx,arg)
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
	json.NewEncoder(w).Encode(couponInfo)

}
func (server *Server) getCouponById(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    vars := mux.Vars(r)
    idParam, ok := vars["id"]
    if !ok {
        errorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
        return
    }

    id, err := strconv.Atoi(idParam)
    if err != nil {
        errorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
        return
    }

    coupon, err := server.store.SelectCoupons(ctx, int32(id))
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

    if err := json.NewEncoder(w).Encode(coupon); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}


func (server *Server) getAllCoupon(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    coupons, err := server.store.SelectAllCoupons(ctx)
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

    if err := json.NewEncoder(w).Encode(coupons); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}

