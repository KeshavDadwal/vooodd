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


type StoreNameCategory struct {
	ID       int32       `json:"id"`
	ParentID pgtype.Int4 `json:"parent_id"`
	Lft      pgtype.Int4 `json:"lft"`
	Rght     pgtype.Int4 `json:"rght"`
	Name     pgtype.Text `json:"name"`
}

func (server *Server) CreateStoreNameCategory(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	storenamecategory := StoreNameCategory{}
	err := json.NewDecoder(r.Body).Decode(&storenamecategory)

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

	arg := db.CreateStoreNameCategoryParams{
		Name: storenamecategory.Name,
		ParentID: storenamecategory.ParentID,
		Lft: storenamecategory.Lft,
		Rght: storenamecategory.Rght,
	}

	storeNameCategoryInfo, err := server.store.CreateStoreNameCategory(ctx,arg)
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
	json.NewEncoder(w).Encode(storeNameCategoryInfo)

}
