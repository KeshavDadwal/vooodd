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

type Store struct {
	ID          int32       `json:"id"`
	Name        string      `json:"name"`
	StoreNo     pgtype.Text `json:"store_no"`
	Description string      `json:"description"`
	Logo        pgtype.Text `json:"logo"`
	PhoneNo     pgtype.Text `json:"phone_no"`
	Isdeleted   pgtype.Bool `json:"isdeleted"`
	Isblocked   pgtype.Bool `json:"isblocked"`
	UserID      pgtype.Int4 `json:"user_id"`
	Icon        pgtype.Text `json:"icon"`
}
func (server *Server) createStore(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	store := Store{}
	err := json.NewDecoder(r.Body).Decode(&store)

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

	arg := db.CreateStoreParams{
		Name: store.Name,
		StoreNo: store.StoreNo,
		Description: store.Description,
		Logo: store.Logo,
		PhoneNo: store.PhoneNo,
		UserID: store.UserID,
		Icon: store.Icon,
	}

	storeInfo, err := server.store.CreateStore(ctx,arg)
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
	json.NewEncoder(w).Encode(storeInfo)

}
