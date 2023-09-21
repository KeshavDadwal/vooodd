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

type ProductImage struct {
	ID        int32       `json:"id"`
	ProductID int32       `json:"product_id"`
	Image     string      `json:"image"`
	Isdeleted pgtype.Bool `json:"isdeleted"`
	Isblocked pgtype.Bool `json:"isblocked"`
	Isdefault pgtype.Bool `json:"isdefault"`
}

func (server *Server) createProductImage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	productimage := ProductImage{}
	err := json.NewDecoder(r.Body).Decode(&productimage)

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

	arg := db.CreateProductImageParams{
		Image: productimage.Image,
		ProductID: productimage.ProductID,
		Isdeleted: productimage.Isdeleted,
		Isblocked: productimage.Isblocked,
		Isdefault: productimage.Isdefault,
	}

	productImageInfo, err := server.store.CreateProductImage(ctx,arg)
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
	json.NewEncoder(w).Encode(productImageInfo)

}
