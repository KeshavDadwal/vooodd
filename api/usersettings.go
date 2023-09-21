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

type UserSetting struct {
	ID          int32       `json:"id"`
	UserID      int32       `json:"user_id"`
	MallRadius  pgtype.Int4 `json:"mall_radius"`
	StoreRadius pgtype.Int4 `json:"store_radius"`
	CreatedAt   time.Time   `json:"created_at"`
}

func (server *Server) createUserSetting(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	usersetting := UserSetting{}
	err := json.NewDecoder(r.Body).Decode(&usersetting)

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

	arg := db.CreateUserSettingParams{
		UserID: usersetting.UserID,
		MallRadius: usersetting.MallRadius,
		StoreRadius: usersetting.StoreRadius,
	}

	userSettingInfo, err := server.store.CreateUserSetting(ctx,arg)
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
	json.NewEncoder(w).Encode(userSettingInfo)

}

