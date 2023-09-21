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

type CouponLanguage struct {
	ID          int32       `json:"id"`
	CouponID    pgtype.Int4 `json:"coupon_id"`
	LanguageID  pgtype.Int4 `json:"language_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Disclaimer  string      `json:"disclaimer"`
}
func (server *Server) createCouponLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	couponlanguage := CouponLanguage{}
	err := json.NewDecoder(r.Body).Decode(&couponlanguage)

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

	arg := db.CreateCouponLanguageParams{
		Name: couponlanguage.Name,
		Description: couponlanguage.Description,
		Disclaimer: couponlanguage.Disclaimer,
		CouponID: couponlanguage.CouponID,
		LanguageID: couponlanguage.LanguageID,
	}

	couponLanguageInfo, err := server.store.CreateCouponLanguage(ctx,arg)
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
	json.NewEncoder(w).Encode(couponLanguageInfo)

}

func (server *Server) getAllCouponLanguage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    couponLanguages, err := server.store.SelectAllCouponLanguages(ctx)
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

    if err := json.NewEncoder(w).Encode(couponLanguages); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}
