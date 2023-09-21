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

type ProductLanguage struct {
	ID               int32          `json:"id"`
	ProductID        pgtype.Int4    `json:"product_id"`
	LanguageID       pgtype.Int4    `json:"language_id"`
	Name             string         `json:"name"`
	Sku              string         `json:"sku"`
	Description      pgtype.Text    `json:"description"`
	ShortDescription pgtype.Text    `json:"short_description"`
	Weight           pgtype.Numeric `json:"weight"`
	MetaTitle        string         `json:"meta_title"`
	MetaKeywords     string         `json:"meta_keywords"`
	MetaDescription  string         `json:"meta_description"`
	CreatedAt        time.Time      `json:"created_at"`
}
func (server *Server) createProductLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	productlanguage := ProductLanguage{}
	err := json.NewDecoder(r.Body).Decode(&productlanguage)

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

	arg := db.CreateProductLanguageParams{
		Name: productlanguage.Name,
		Sku: productlanguage.Sku,
		ShortDescription: productlanguage.ShortDescription,
		Description: productlanguage.Description,
		ProductID: productlanguage.ProductID,
		LanguageID: productlanguage.LanguageID,
		Weight: productlanguage.Weight,
		MetaTitle: productlanguage.MetaTitle,
		MetaKeywords: productlanguage.MetaKeywords,
		MetaDescription: productlanguage.MetaDescription,

	}

	productLanguageInfo, err := server.store.CreateProductLanguage(ctx,arg)
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
	json.NewEncoder(w).Encode(productLanguageInfo)

}
