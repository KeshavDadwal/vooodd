package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type ProductCategoryMap struct {
	ID                int32 `json:"id"`
	ProductID         int32 `json:"product_id"`
	ProductCategoryID int32 `json:"product_category_id"`
}

func (server *Server) createProductCategoryMap(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	productcategorymap := ProductCategoryMap{}
	err := json.NewDecoder(r.Body).Decode(&productcategorymap)

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

	arg := db.CreateProductCategoryMapParams{
		ProductID: productcategorymap.ProductID,
		ProductCategoryID: productcategorymap.ProductCategoryID,

	}

	productCategoryMapInfo, err := server.store.CreateProductCategoryMap(ctx,arg)
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
	json.NewEncoder(w).Encode(productCategoryMapInfo)

}
