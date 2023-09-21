package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"strconv"
	//"time"

	//"aidanwoods.dev/go-paseto"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/vod/db/sqlc"
	util "github.com/vod/utils"
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	Username string `json:"username"`
// 	FullName string `json:"full_name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

type GenderEnum string

const (
	GenderEnumMale   GenderEnum = "Male"
	GenderEnumFemale GenderEnum = "Female"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func errorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errResponse := ErrorResponse{
		Status:  statusCode,
		Message: message,
	}
	json.NewEncoder(w).Encode(errResponse)
}

func (e *GenderEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = GenderEnum(s)
	case string:
		*e = GenderEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for GenderEnum: %T", src)
	}
	return nil
}

type NullGenderEnum struct {
	GenderEnum GenderEnum `json:"gender_enum"`
	Valid      bool       `json:"valid"` // Valid is true if GenderEnum is not NULL
}

type User struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Gender    GenderEnum  `json:"gender"`
	Dob       pgtype.Date `json:"dob"`
	Address   string      `json:"address"`
	City      string      `json:"city"`
	State     string      `json:"state"`
	CountryID pgtype.Int4 `json:"country_id"`
	Zip       string      `json:"zip"`
	MobileNo  string      `json:"mobile_no"`
	Email     string      `json:"email"`
	Username  string      `json:"username"`
	Password  string      `json:"password"`
}

type JsonResponse struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"status_code"`
}

type SuceessResponse struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	UserToken interface{} `json:"userToken"`
}

func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}
	ctx := r.Context()

	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		jsonResponse := JsonResponse{
			Status:     false,
			Message:    "invalid JSON request",
			StatusCode: http.StatusTeapot,
		}
		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
		return
	}

	//hashedPassword, err := util.HashPassword(user.Password)
	// if err != nil {

	// 	jsonResponse := JsonResponse{
	// 		Status:     false,
	// 		Message:    "invalid JSON request",
	// 		StatusCode: http.StatusTeapot,
	// 	}
	// 	util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
	// 	return
	// }

	arg := db.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Column3:   db.GenderEnum(user.Gender),
		Dob:       user.Dob,
		Address:   user.Address,
		City:      user.City,
		State:     user.State,
		CountryID: user.CountryID,
		Zip:       user.Zip,
		MobileNo:  user.MobileNo,
		Email:     user.Email,
		Username:  user.Username,
		Password:  string(hash),
	}

	userInfo, err := server.store.CreateUser(ctx, arg)
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
	json.NewEncoder(w).Encode(userInfo)
}

type Login struct {
	Email    string
	Password string
}

// func (server *Server) loginUser(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
// 		return
// 	}

// 	ctx := r.Context()

// 	login := Login{}

// 	err := json.NewDecoder(r.Body).Decode(&login)
// 	if err != nil {
// 		fmt.Println("------error1------", err)

// 		jsonResponse := JsonResponse{
// 			Status:     false,
// 			Message:    "invalid JSON request",
// 			StatusCode: http.StatusTeapot,
// 		}
// 		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
// 		return
// 	}
// 	//var user1 db.User

// 	//query := `SELECT id, password FROM users WHERE email = $1`
// 	//row := db.QueryRow(context.Background(), query, login.Email)
	
// 	userInfo, err := server.store.GetUserByEmail(ctx, login.Email)
// 	// query := `SELECT password FROM users WHERE email = $1`
// 	// row := db.QueryRow(context.Background(), query, login.Email)

// 	//err = row.Scan(&user1.ID, &user1.Password)
// 	if err != nil {
// 		fmt.Println("------error1------", err)

// 		jsonResponse := JsonResponse{
// 			Status:     false,
// 			Message:    "invalid username or password ",
// 			StatusCode: http.StatusTeapot,
// 		}
// 		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
// 		return
// 	}

// 	if userInfo.ID == 0 {
// 		jsonResponse := JsonResponse{
// 			Status:     false,
// 			Message:    "invalid username or password ",
// 			StatusCode: http.StatusTeapot,
// 		}
// 		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
// 		return
// 	}
// 	userID := strconv.Itoa(int(userInfo.ID))

// 	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(login.Password))
// 	//err = bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(body.Password))

// 	if err != nil {

// 		jsonResponse := JsonResponse{
// 			Status:     false,
// 			Message:    "Password Mismatch",
// 			StatusCode: http.StatusTeapot,
// 		}
// 		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
// 		return
// 	}

// 	//symmetricKey := os.Getenv("SECRET")
// 	// now := time.Now()
// 	// exp := now.Add(24 * time.Hour)
// 	// nbt := now

// 	jsonToken := paseto.NewToken()

// 	jsonToken.SetIssuedAt(time.Now())
// 	jsonToken.SetNotBefore(time.Now())
// 	jsonToken.SetExpiration(time.Now().Add(2 * time.Hour))

// 	jsonToken.SetString("user-id", userID)

// 	// jsonToken := paseto.Token{
// 	// 	Jti:        "123",
// 	// 	Subject:    userID,
// 	// 	IssuedAt:   now,
// 	// 	Expiration: exp,
// 	// 	NotBefore:  nbt,
// 	// }

// 	// Encrypt data
// 	jsonToken.Set("data", "this is a signed message")
// 	jsonToken.Set("email", userInfo.Email)

// 	//footer := "some footer"
// 	key := paseto.NewV4SymmetricKey()
// 	tokenString := jsonToken.V4Encrypt(key, nil)

// 	//tokenString, err := paseto.NewV2().Encrypt([]byte(symmetricKey), jsonToken, footer)
// 	if err != nil {

// 		jsonResponse := JsonResponse{
// 			Status:     false,
// 			Message:    "Failed to create a token",
// 			StatusCode: http.StatusTeapot,
// 		}
// 		util.WriteJSONResponse(w, http.StatusTeapot, jsonResponse)
// 		return
// 	}

// 	fmt.Println("Login")
// 	w.Header().Set("Content-Type", "application/json")
// 	jsonResponse := JsonResponse{
// 		Status:     true,
// 		Message:    tokenString,
// 		StatusCode: http.StatusAccepted,
// 	}
// 	util.WriteJSONResponse(w, http.StatusAccepted, jsonResponse)
// 	return

// }

func (server *Server) refreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, http.StatusMethodNotAllowed, "Only POST requests are allowed")
		return
	}

	// dummy JSON response
	userInfo := struct {
		Message string `json:"message"`
		UserID  int    `json:"user_id"`
	}{
		Message: "refresh token",
		UserID:  123, // Replace with an actual user ID if needed
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userInfo)
}
