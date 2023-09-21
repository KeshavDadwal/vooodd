package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)
type StoreCategory struct {
	ID      int32 `json:"id"`
	DealID  int32 `json:"deal_id"`
	StoreID int32 `json:"store_id"`
}
func (server *Server) createStoreCategory(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	storecategory := StoreCategory{}
	err := json.NewDecoder(r.Body).Decode(&storecategory)

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

	arg := db.CreateStoreCategoryParams{
		DealID: storecategory.DealID,
		StoreID: storecategory.StoreID,
	}

	storeCategoryInfo, err := server.store.CreateStoreCategory(ctx,arg)
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
	json.NewEncoder(w).Encode(storeCategoryInfo)

}
