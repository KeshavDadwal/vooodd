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
type UserMallPermission struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	MallID    int32     `json:"mall_id"`
	CreatedAt time.Time `json:"created_at"`
}
func (server *Server) createUserMallPermission(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	usermallpermission := UserMallPermission{}
	err := json.NewDecoder(r.Body).Decode(&usermallpermission)

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

	arg := db.CreateUserMallPermissionParams{
		UserID: usermallpermission.UserID,
		MallID: usermallpermission.MallID,
	}

	userMallPermissionInfo, err := server.store.CreateUserMallPermission(ctx,arg)
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
	json.NewEncoder(w).Encode(userMallPermissionInfo)

}
