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
type ProductStoreLocation struct {
	ID              int32       `json:"id"`
	ProductID       pgtype.Int4 `json:"product_id"`
	StoreLocationID pgtype.Int4 `json:"store_location_id"`
}

func (server *Server) createProductStoreLocation(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	productstorelocation := ProductStoreLocation{}
	err := json.NewDecoder(r.Body).Decode(&productstorelocation)

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

	arg := db.CreateProductStoreLocationParams{
		ProductID: productstorelocation.ProductID,
		StoreLocationID: productstorelocation.StoreLocationID,
	}

	productStoreLocationInfo, err := server.store.CreateProductStoreLocation(ctx,arg)
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
	json.NewEncoder(w).Encode(productStoreLocationInfo)

}
