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
type DevicetypeEnum string

const (
	DevicetypeEnumANDROID DevicetypeEnum = "ANDROID"
	DevicetypeEnumIOS     DevicetypeEnum = "IOS"
)
func (e *DevicetypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DevicetypeEnum(s)
	case string:
		*e = DevicetypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for DevicetypeEnum: %T", src)
	}
	return nil
}

type NullDevicetypeEnum struct {
	DevicetypeEnum DevicetypeEnum `json:"devicetype_enum"`
	Valid          bool           `json:"valid"` // Valid is true if DevicetypeEnum is not NULL
}

type UserDevice struct {
	ID          int32          `json:"id"`
	UserID      int32          `json:"user_id"`
	DeviceID    string         `json:"device_id"`
	DeviceType  DevicetypeEnum `json:"device_type"`
	Wuid        pgtype.Text    `json:"wuid"`
	DeviceToken pgtype.Text    `json:"device_token"`
}
func (server *Server) createUserDevice(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	userdevice := UserDevice{}
	err := json.NewDecoder(r.Body).Decode(&userdevice)

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

	arg := db.CreateUserDeviceParams{
		UserID: userdevice.UserID,
		DeviceID: userdevice.DeviceID,
		Wuid: userdevice.Wuid,
		DeviceToken: userdevice.DeviceToken,
		Column3: db.DevicetypeEnum(userdevice.DeviceType),
	}

	userDeviceInfo, err := server.store.CreateUserDevice(ctx,arg)
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
	json.NewEncoder(w).Encode(userDeviceInfo)

}
