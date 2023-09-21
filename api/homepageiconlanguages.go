package api

import (
	"encoding/json"
	"net/http"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type HomePageIconLanguage struct {
	ID             int32       `json:"id"`
	HomePageIconID pgtype.Int4 `json:"home_page_icon_id"`
	LanguageID     pgtype.Int4 `json:"language_id"`
	Name           pgtype.Text `json:"name"`
}

func (server *Server) createHomePageIconLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	homepageiconlanguage:= HomePageImageLanguage{}
	err := json.NewDecoder(r.Body).Decode(&homepageiconlanguage)

	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateHomePageIconLanguageParams{
		Name: homepageiconlanguage.Name,
		LanguageID: homepageiconlanguage.LanguageID,
		HomePageIconID: homepageiconlanguage.HomePageImageID,
	}

	homePageIconLanguageInfo, err := server.store.CreateHomePageIconLanguage(ctx,arg)
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
	json.NewEncoder(w).Encode(homePageIconLanguageInfo)

}

func (server *Server) getAllHomePageIconLanguage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    homePageIconLanguage, err := server.store.SelectAllHomePageIconLanguages(ctx)
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

    if err := json.NewEncoder(w).Encode(homePageIconLanguage); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}
