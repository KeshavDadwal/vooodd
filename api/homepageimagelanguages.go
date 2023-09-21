package api

import (
	"encoding/json"
	"net/http"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type HomePageImageLanguage struct {
	ID              int32       `json:"id"`
	HomePageImageID pgtype.Int4 `json:"home_page_image_id"`
	LanguageID      pgtype.Int4 `json:"language_id"`
	Name            pgtype.Text `json:"name"`
}

func (server *Server) createHomePageImageLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	homepageimagelanguage := HomePageImageLanguage{}
	err := json.NewDecoder(r.Body).Decode(&homepageimagelanguage)

	if err != nil {

		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateHomePageImageLanguageParams{
		Name: homepageimagelanguage.Name,
		LanguageID: homepageimagelanguage.LanguageID,
		HomePageImageID: homepageimagelanguage.HomePageImageID,
	}

	homePageImageLanguageInfo, err := server.store.CreateHomePageImageLanguage(ctx,arg)
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
	json.NewEncoder(w).Encode(homePageImageLanguageInfo)

}

func (server *Server) getAllHomePageImageLanguage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    homePageImageLanguage, err := server.store.SelectHomePageImageLanguages(ctx)
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

    if err := json.NewEncoder(w).Encode(homePageImageLanguage); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}