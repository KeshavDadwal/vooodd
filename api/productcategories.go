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
type ProductCategory struct {
	ID       int32       `json:"id"`
	ParentID pgtype.Int4 `json:"parent_id"`
	Lft      pgtype.Int4 `json:"lft"`
	Rght     pgtype.Int4 `json:"rght"`
	Name     pgtype.Text `json:"name"`
	Logo     pgtype.Text `json:"logo"`
	StoreID  pgtype.Int4 `json:"store_id"`
}
func (server *Server) createProductCategory(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	productcategory:= ProductCategory{}
	err := json.NewDecoder(r.Body).Decode(&productcategory)

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

	arg := db.CreateProductCategoryParams{
		Name: productcategory.Name,
		ParentID: productcategory.ParentID,
		Lft: productcategory.Lft,
		Rght: productcategory.Rght,
		Logo: productcategory.Logo,
		StoreID: productcategory.StoreID,
	}

	productcategoryInfo, err := server.store.CreateProductCategory(ctx,arg)
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
	json.NewEncoder(w).Encode(productcategoryInfo)

}
