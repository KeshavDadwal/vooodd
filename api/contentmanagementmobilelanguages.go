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

type ContentManagementMobileLanguage struct {
	ID                        int32  `json:"id"`
	ContentManagementMobileID int32  `json:"content_management_mobile_id"`
	LanguageID                int32  `json:"language_id"`
	Title                     string `json:"title"`
	Link                      string `json:"link"`
	LinkText                  string `json:"link_text"`
	Content                   string `json:"content"`
}

func (server *Server) createContentManagementMobileLanguage(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed,"Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	contentmanagementmobilelanguage := ContentManagementMobileLanguage{}
	err := json.NewDecoder(r.Body).Decode(&contentmanagementmobilelanguage)

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

	arg := db.CreateContentManagementMobileLanguageParams{
		Title: contentmanagementmobilelanguage.Title,
		Link: contentmanagementmobilelanguage.Link,
		LinkText: contentmanagementmobilelanguage.LinkText,
		Content: contentmanagementmobilelanguage.Content,
		LanguageID: contentmanagementmobilelanguage.LanguageID,
		ContentManagementMobileID: contentmanagementmobilelanguage.ContentManagementMobileID,
	}

	contentManagementMobileLanguageInfo, err := server.store.CreateContentManagementMobileLanguage(ctx,arg)
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
	json.NewEncoder(w).Encode(contentManagementMobileLanguageInfo)

}

func (server *Server) getContentManagementMobileLanguage(w http.ResponseWriter, r *http.Request) {
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

    contentManagementMobileLanguage, err := server.store.SelectContentManagementMobileLanguages(ctx, int32(id))
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

    if err := json.NewEncoder(w).Encode(contentManagementMobileLanguage); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}


func (server *Server) getAllContentManagementMobileLanguage(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    contentManagementMobileLanguages, err := server.store.SelectAllContentManagementMobileLanguages(ctx)
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

    if err := json.NewEncoder(w).Encode(contentManagementMobileLanguages); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}
