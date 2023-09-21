package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"github.com/jackc/pgx/v5/pgtype"
	"github.com/gorilla/mux"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type ContentManagementBannerLanguage struct {
	ID                  int32  `json:"id"`
	ContentManagementID int32  `json:"content_management_id"`
	LanguageID          int32  `json:"language_id"`
	Name                string `json:"name"`
	Title               string `json:"title"`
	Content             string `json:"content"`
}

func (server *Server) createContentManagementBannerLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	contentmanagemnetbannerlanguages := ContentManagementBannerLanguage{}
	err := json.NewDecoder(r.Body).Decode(&contentmanagemnetbannerlanguages)

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

	arg := db.CreateContentManagementBannerLanguageParams{
		Name: contentmanagemnetbannerlanguages.Name,
		Title: contentmanagemnetbannerlanguages.Title,
		Content: contentmanagemnetbannerlanguages.Content,
		LanguageID: contentmanagemnetbannerlanguages.LanguageID,
		ContentManagementID: contentmanagemnetbannerlanguages.ContentManagementID,
	}

	contentManagementBannerLanguageInfo, err := server.store.CreateContentManagementBannerLanguage(ctx,arg)
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
	json.NewEncoder(w).Encode(contentManagementBannerLanguageInfo)

}


func (server *Server) getContentManagementBannerLanguageById(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    vars := mux.Vars(r)
    idParam, ok := vars["id"]
    if !ok {
        errorResponse(w, http.StatusBadRequest, "Missing 'id' URL parameter")
        return
    }

    id, err := strconv.Atoi(idParam)
    if err != nil {
        errorResponse(w, http.StatusBadRequest, "Invalid 'id' URL parameter")
        return
    }

    contentManagementBannerLanguage, err := server.store.SelectContentManagementBannerLanguages(ctx, int32(id))
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

    if err := json.NewEncoder(w).Encode(contentManagementBannerLanguage); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}


func (server *Server) getAllContentManagementBannerLanguage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    contentManagementBannerLanguages, err := server.store.SelectAllContentManagementBannerLanguages(ctx)
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

    if err := json.NewEncoder(w).Encode(contentManagementBannerLanguages); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}
