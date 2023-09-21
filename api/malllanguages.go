package api

import (
	"encoding/json"
	"net/http"
	"time"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type MallLanguage struct {
	ID         int32     `json:"id"`
	MallID     int32     `json:"mall_id"`
	LanguageID int32     `json:"language_id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Zip        string    `json:"zip"`
	PhoneNo    string    `json:"phone_no"`
	CreatedAt  time.Time `json:"created_at"`
}

func (server *Server) createMallLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	malllanguage := MallLanguage{}
	err := json.NewDecoder(r.Body).Decode(&malllanguage)

	if err != nil {

		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateMallLanguageParams{
		Name: malllanguage.Name,
		Address: malllanguage.Address,
		City: malllanguage.City,
		State: malllanguage.State,
		Zip: malllanguage.Zip,
		MallID: malllanguage.MallID,
		LanguageID: malllanguage.LanguageID,
	}

	mallLanguageInfo, err := server.store.CreateMallLanguage(ctx,arg)
	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mallLanguageInfo)

}
func (server *Server) getAllMallLanguage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    homeBanners, err := server.store.SelectMallLanguages(ctx)
    if err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to fetch brand languages",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(homeBanners); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}
