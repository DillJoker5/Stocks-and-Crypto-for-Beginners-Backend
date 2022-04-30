package main

import (
	"encoding/json"
	"log"
	"net/http"
	model "stock-and-crypto-api/models"
	"github.com/gorilla/mux"
	"database/sql"
	"context"
	_"github.com/denisenkom/go-mssqldb"
	"fmt"
	//"time"
	//"github.com/google/uuid"
)

var db *sql.DB
var server = "DESKTOP-K7IIMGF"
var port = 1433
var user = "finalwebprojectuser"
var password = "finalwebproject2022!"
var database = "Final_Web_Project"

func main() {
	// Build connection string
	connectionString := fmt.Sprintf("server%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	var err error

	// Create connection
	db, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error connecting:", err.Error())
	}

	router := mux.NewRouter()

	// Handle api endpoints
	router.HandleFunc("/login", Login).Methods(http.MethodPost)
	router.HandleFunc("/register", Register).Methods(http.MethodPost)
	router.HandleFunc("/forgotPassword", ForgotPassword).Methods(http.MethodPost)

	router.HandleFunc("/readUsers", ReadUserTable).Methods(http.MethodPost)
	router.HandleFunc("/readThread", ReadThreadTable).Methods(http.MethodPost)
	router.HandleFunc("/readApiFavorites", ReadApiFavoritesTable).Methods(http.MethodPost)
	router.HandleFunc("/readThreadFavorites", ReadThreadFavoritesTable).Methods(http.MethodPost)
	router.HandleFunc("/readResponse", ReadResponseTable).Methods(http.MethodPost)

	router.HandleFunc("/newThread", mwCheck(CreateNewThread)).Methods(http.MethodPost)
	router.HandleFunc("/newThreadResponse", mwCheck(CreateResponse)).Methods(http.MethodPost)
	router.HandleFunc("/newApiFavorite", mwCheck(CreateNewApiFavorite)).Methods(http.MethodPost)
	router.HandleFunc("/newThreadFavorite", mwCheck(CreateThreadFavorite)).Methods(http.MethodPost)

	srv := &http.Server {
		Addr: ":8000",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func mwCheck(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validateUser(r) {
			http.Error(w, "Unauthorized user", http.StatusInternalServerError)
		} else {
			f(w, r)
		}
	}
}

func validateUser(r *http.Request) bool {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		return false
	}
	ug := r.Header.Get("userguid")
	if ug == "" {
		return false
	}

	tsql := fmt.Sprintf("SELECT SessionId FROM Sessions WHERE UserGuid='%s' AND IsActive=1;", ug)
	row := db.QueryRowContext(ctx, tsql)
	if err != nil {
		return false
	}
	var sid int
	if err = row.Scan(&sid); err != nil {
		return false
	}
	return true
}

func ReadUserTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := "SELECT User_Id, Email, Username, Password FROM Users;"

	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var users []model.UserTable
	for rows.Next() {
		var user model.UserTable
		rows.Scan(&user.UserId, &user.Email, &user.Username, &user.Password)
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.LoginJsonResponse{ Message: "", Type: "Success", Data: users }
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadThreadTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := "SELECT Thread_Id, User_Id, Name, Description, Date_Created FROM Threads"

	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var threads []model.ThreadTable
	for rows.Next() {
		var thread model.ThreadTable
		rows.Scan(&thread.ThreadId, &thread.UserId, &thread.Name, &thread.Description, &thread.DateCreated)
		threads = append(threads, thread)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.ThreadJsonResponse{ Message: "", Type: "Success", Data: threads }
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadApiFavoritesTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := "SELECT Api_Favorites_Id, User_Id, Stock_Id, Api_Url FROM Api_Favorites;"

	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var apiFavorites []model.ApiFavoritesTable
	for rows.Next() {
		var apiFavorite model.ApiFavoritesTable
		rows.Scan(&apiFavorite.ApiFavoritesId, &apiFavorite.UserId, &apiFavorite.StockId, &apiFavorite.ApiUrl)
		apiFavorites = append(apiFavorites, apiFavorite)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.ApiFavoritesJsonResponse{ Message: "", Type: "Success", Data: apiFavorites }
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadResponseTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := "SELECT Response_Id, User_Id, Thread_Id, Description, Date_Created FROM Responses;"

	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var responses []model.ResponseTable
	for rows.Next() {
		var response model.ResponseTable
		rows.Scan(&response.ResponseId, &response.UserId, &response.ThreadId, &response.Description, &response.DateCreated)
		responses = append(responses, response)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.ResponsesJsonResponse{ Message: "", Type: "Success", Data: responses }
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadThreadFavoritesTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := "SELECT Thread_Favorites_Id, User_Id FROM Thread_Favorites;"

	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var threadFavorites []model.ThreadFavoritesTable
	for rows.Next() {
		var threadFavorite model.ThreadFavoritesTable
		rows.Scan(&threadFavorite.ThreadFavoritesId, &threadFavorite.UserId)
		threadFavorites = append(threadFavorites, threadFavorite)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.ThreadFavoritesJsonResponse{ Message: "", Type: "Success", Data: threadFavorites }
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateNewThread(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var nThread model.Thread

	err = json.NewDecoder(r.Body).Decode(&nThread)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tsqlQuery := "SELECT Thread_Id, User_Id, Name, Description, Date_Created FROM Threads"

	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var thread model.ThreadTable
		rows.Scan(&thread.ThreadId, &thread.UserId, &thread.Name, &thread.Description, &thread.DateCreated)
		if thread.UserId == nThread.UserId && thread.Name == nThread.Name && thread.Description == nThread.Description {
			http.Error(w, "The given thread has already been created!", http.StatusBadRequest)
			return
		}
	}

	tsqlQuery = fmt.Sprintf("INSERT INTO Threads VALUES(%d, %d, '%s', '%s', '%s')", nThread.UserId, nThread.ResponseId, nThread.Name, nThread.Description, nThread.DateCreated)

	res, err := db.ExecContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, err := res.RowsAffected()
	if err != nil || count != 1 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.GenericJsonResponse{ Message: "Successfully created the Thread!", Type: "Success" }
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateNewApiFavorite(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateResponse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateThreadFavorite(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
