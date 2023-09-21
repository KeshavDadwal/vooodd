package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type UserMall struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	MallID    int32     `json:"mall_id"`
	DeviceID  string    `json:"device_id"`
	CreatedAt time.Time `json:"created_at"`
}
func (server *Server) createUserMall(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	usermall := UserMall{}
	err := json.NewDecoder(r.Body).Decode(&usermall)

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

	arg := db.CreateUserMallMappingParams{
		UserID: usermall.UserID,
		MallID: usermall.MallID,
		DeviceID: usermall.DeviceID,
	}

	userMallInfo, err := server.store.CreateUserMallMapping(ctx,arg)
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
	json.NewEncoder(w).Encode(userMallInfo)

}
