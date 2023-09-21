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

type Product struct {
	ID                int32       `json:"id"`
	BrandID           pgtype.Int4 `json:"brand_id"`
	ProductCategoryID pgtype.Int4 `json:"product_category_id"`
	Name              string      `json:"name"`
	Sku               string      `json:"sku"`
	Description       pgtype.Text `json:"description"`
	ShortDescription  pgtype.Text `json:"short_description"`
	NewFrom           pgtype.Date `json:"new_from"`
	NewEnd            pgtype.Date `json:"new_end"`
	Isdeleted         pgtype.Bool `json:"isdeleted"`
	Isblocked         pgtype.Bool `json:"isblocked"`
	Isbestseller      pgtype.Bool `json:"isbestseller"`
	Isfeatured        pgtype.Bool `json:"isfeatured"`
	MetaTitle         string      `json:"meta_title"`
	MetaKeywords      string      `json:"meta_keywords"`
	MetaDescription   string      `json:"meta_description"`
	Price             int32       `json:"price"`
	UnitID            pgtype.Int4 `json:"unit_id"`
}
func (server *Server) createProduct(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	product := Product{}
	err := json.NewDecoder(r.Body).Decode(&product)

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

	arg := db.CreateProductParams{
		Name: product.Name,
		Sku: product.Sku,
		ShortDescription: product.ShortDescription,
		Description: product.Description,
		BrandID: product.BrandID,
		ProductCategoryID: product.ProductCategoryID,
		NewFrom: product.NewFrom,
		NewEnd: product.NewEnd,
		Isdeleted: product.Isdeleted,
		Isblocked: product.Isblocked,
		Isbestseller: product.Isbestseller,
		Isfeatured: product.Isfeatured,
		MetaTitle: product.MetaTitle,
		MetaKeywords: product.MetaKeywords,
		MetaDescription: product.MetaDescription,
		Price: product.Price,
		UnitID: product.UnitID,
	}

	productInfo, err := server.store.CreateProduct(ctx,arg)
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
	json.NewEncoder(w).Encode(productInfo)

}
