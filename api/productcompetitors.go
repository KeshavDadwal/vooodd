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

type ProductCompetitor struct {
	ID              int32          `json:"id"`
	ProductID       int32          `json:"product_id"`
	CompetitorID    int32          `json:"competitor_id"`
	StoreLocationID int32          `json:"store_location_id"`
	Price           pgtype.Numeric `json:"price"`
	CreatedAt       time.Time      `json:"created_at"`
}
func (server *Server) createProductCompetitor(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	productcompetitor := ProductCompetitor{}
	err := json.NewDecoder(r.Body).Decode(&productcompetitor)

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

	arg := db.CreateProductCompetitorParams{
		ProductID: productcompetitor.ProductID,
		CompetitorID: productcompetitor.CompetitorID,
		StoreLocationID: productcompetitor.StoreLocationID,
		Price: productcompetitor.Price,
	}

	productCompetitorInfo, err := server.store.CreateProductCompetitor(ctx,arg)
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
	json.NewEncoder(w).Encode(productCompetitorInfo)

}
