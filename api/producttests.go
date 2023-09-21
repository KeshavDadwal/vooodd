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


type ProductTest struct {
	ID                int32          `json:"id"`
	BrandID           pgtype.Int4    `json:"brand_id"`
	ProductCategoryID pgtype.Int4    `json:"product_category_id"`
	Name              string         `json:"name"`
	Sku               string         `json:"sku"`
	Description       pgtype.Text    `json:"description"`
	ShortDescription  pgtype.Text    `json:"short_description"`
	Weight            pgtype.Numeric `json:"weight"`
	NewFrom           pgtype.Date    `json:"new_from"`
	NewEnd            pgtype.Date    `json:"new_end"`
	Isdeleted         pgtype.Bool    `json:"isdeleted"`
	Isblocked         pgtype.Bool    `json:"isblocked"`
	Isbestseller      pgtype.Bool    `json:"isbestseller"`
	Isfeatured        pgtype.Bool    `json:"isfeatured"`
	MetaTitle         string         `json:"meta_title"`
	MetaKeywords      string         `json:"meta_keywords"`
	MetaDescription   string         `json:"meta_description"`
	Price             pgtype.Numeric `json:"price"`
	UnitID            pgtype.Int4    `json:"unit_id"`
	CreatedAt         time.Time      `json:"created_at"`
}
func (server *Server) createProductTest(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	producttest := ProductTest{}
	err := json.NewDecoder(r.Body).Decode(&producttest)

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

	arg := db.CreateProductTestParams{
		Name: producttest.Name,
		Sku: producttest.Sku,
		ShortDescription: producttest.ShortDescription,
		Description: producttest.Description,
		BrandID: producttest.BrandID,
		ProductCategoryID: producttest.ProductCategoryID,
		Weight: producttest.Weight,
		NewFrom: producttest.NewFrom,
		NewEnd: producttest.NewEnd,
		Isdeleted: producttest.Isdeleted,
		Isblocked: producttest.Isblocked,
		Isbestseller: producttest.Isbestseller,
		Isfeatured: producttest.Isfeatured,
		MetaTitle: producttest.MetaTitle,
		MetaKeywords: producttest.MetaKeywords,
		MetaDescription: producttest.MetaDescription,
		Price: producttest.Price,
		UnitID: producttest.UnitID,

	}

	productTestInfo, err := server.store.CreateProductTest(ctx,arg)
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
	json.NewEncoder(w).Encode(productTestInfo)

}
