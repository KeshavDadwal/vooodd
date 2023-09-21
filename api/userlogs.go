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

type UserLog struct {
	ID        int32       `json:"id"`
	UserID    int32       `json:"user_id"`
	MallID    int32       `json:"mall_id"`
	StoreID   int32       `json:"store_id"`
	OfferID   int32       `json:"offer_id"`
	Name      string      `json:"name"`
	Text      string      `json:"text"`
	Isblocked pgtype.Bool `json:"isblocked"`
	Isdeleted pgtype.Bool `json:"isdeleted"`
	CreatedAt time.Time   `json:"created_at"`
}


func (server *Server) createUserLog(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	userlog := UserLog{}
	err := json.NewDecoder(r.Body).Decode(&userlog)

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

	arg := db.CreateUserLogParams{
		Name: userlog.Name,
		MallID: userlog.MallID,
		UserID: userlog.UserID,
		StoreID: userlog.StoreID,
		OfferID: userlog.OfferID,
		Text: userlog.Text,
	}

	userLogInfo, err := server.store.CreateUserLog(ctx,arg)
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
	json.NewEncoder(w).Encode(userLogInfo)

}
