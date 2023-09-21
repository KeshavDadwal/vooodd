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

type UserEvent struct {
	ID        int32       `json:"id"`
	UserID    int32       `json:"user_id"`
	Name      string      `json:"name"`
	Date      pgtype.Date `json:"date"`
	CreatedAt time.Time   `json:"created_at"`
}

func (server *Server) createUserEvent(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	userevent := UserEvent{}
	err := json.NewDecoder(r.Body).Decode(&userevent)

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

	arg := db.CreateUserEventParams{
		Name: userevent.Name,
		UserID: userevent.UserID,
		Date: userevent.Date,
	}

	userEventInfo, err := server.store.CreateUserEvent(ctx,arg)
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
	json.NewEncoder(w).Encode(userEventInfo)

}
