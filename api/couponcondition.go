package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	//"github.com/jackc/pgx/v5/pgtype"
	"github.com/gorilla/mux"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
)

type TypeEnum string

const (
	TypeEnumDepartment      TypeEnum = "Department"
	TypeEnumProductCategory TypeEnum = "ProductCategory"
	TypeEnumProduct         TypeEnum = "Product"
	TypeEnumBrand           TypeEnum = "Brand"
	TypeEnumFilter          TypeEnum = "Filter"
	TypeEnumMoreFIlter      TypeEnum = "More_FIlter"
	TypeEnumPrivate         TypeEnum = "Private"
	TypeEnumSku             TypeEnum = "Sku"
)

func (e *TypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TypeEnum(s)
	case string:
		*e = TypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for TypeEnum: %T", src)
	}
	return nil
}

type NullTypeEnum struct {
	TypeEnum TypeEnum `json:"type_enum"`
	Valid    bool     `json:"valid"` // Valid is true if TypeEnum is not NULL
}

type CouponCondition struct {
	ID        int32        `json:"id"`
	CouponID  int32        `json:"coupon_id"`
	Idx       int32        `json:"idx"`
	Type      NullTypeEnum `json:"type"`
	TypeID    string       `json:"type_id"`
	CreatedAt time.Time    `json:"created_at"`
}

func (server *Server) createCouponCondition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	couponcondition := CouponCondition{}
	err := json.NewDecoder(r.Body).Decode(&couponcondition)

	if err != nil {

		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	arg := db.CreateCouponConditionParams{
		CouponID: couponcondition.CouponID,
		Idx:      couponcondition.Idx,
		Type:     db.NullTypeEnum{},
		TypeID:   couponcondition.TypeID,
	}

	couponConditionInfo, err := server.store.CreateCouponCondition(ctx, arg)
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
	json.NewEncoder(w).Encode(couponConditionInfo)

}
func (server *Server) getCouponConditionById(w http.ResponseWriter, r *http.Request) {
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

    couponCondition, err := server.store.SelectCouponConditions(ctx, int32(id))
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

    if err := json.NewEncoder(w).Encode(couponCondition); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}


func (server *Server) getAllCouponCondition(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        errorResponse(w, http.StatusMethodNotAllowed, "Only GET requests are allowed")
        return
    }
    ctx := r.Context()

    couponConditions, err := server.store.SelectAllCouponConditions(ctx)
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

    if err := json.NewEncoder(w).Encode(couponConditions); err != nil {
        jsonResponse := JsonResponse{
            Status:     false,
            Message:    "Failed to encode response",
            StatusCode: http.StatusInternalServerError,
        }
        util.WriteJSONResponse(w, http.StatusInternalServerError, jsonResponse)
        return
    }
}


