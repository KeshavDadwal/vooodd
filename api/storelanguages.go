package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)
type StoreLanguage struct {
	ID          int32  `json:"id"`
	LanguageID  int32  `json:"language_id"`
	StoreID     int32  `json:"store_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
func (server *Server) createStoreLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	storelangauge := StoreLanguage{}
	err := json.NewDecoder(r.Body).Decode(&storelangauge)

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

	arg := db.CreateStoreLanguageParams{
		Name: storelangauge.Name,
		Description: storelangauge.Description,
		StoreID: storelangauge.StoreID,
		LanguageID: storelangauge.LanguageID,
	}

	storeLangaugeInfo, err := server.store.CreateStoreLanguage(ctx,arg)
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
	json.NewEncoder(w).Encode(storeLangaugeInfo)

}
