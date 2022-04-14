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
	"time"
)

var db *sql.DB
var server = ""
var user = ""
var password = ""
var database = ""

func main() {
	// Build connection string
	connectionString := fmt.Sprintf("server%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password. port, database)
	var err error

	// Create connection
	db, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error connecting:", err.Error())
	}

	router := mux.NewRouter()

	router.HandleFunc("/login", mwCheck(readUserTable)).Methods(http.MethodPost)
	router.HandleFunc("/viewThreads", mwCheck(readThreadTable)).Methods(http.MethodPost)
	router.HandleFunc("viewStockCryptoTable", mwCheck(readApiFavoritesTable)).Methods(http.MethodPost)
	router.HandleFunc("/viewThreads", mwCheck(readThreadFavoritesTable)).Methods(http.MethodPost)
	router.HandleFunc("/viewThread", mwCheck(readResponseTable)).Methods(http.MethodPost)

	router.HandleFunc("/register", mwCheck(createNewUser)).Methods(http.MethodPost)
	router.HandleFunc("/createThread", mwCheck(createNewThread)).Methods(http.MethodPost)
	router.HandleFunc("/newApiFavorite", mwCheck(createNewApiFavorite)).Methods(http.MethodPost)
	router.HandleFunc("/newResponse", mwCheck(createResponse)).Methods(http.MethodPost)
	router.HandleFunc("/newThreadFavorite", mwCheck(createThreadFavorite)).Methods(http.MethodPost)

	srv := &http.Server {
		Addr: ":8000",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func mwCheck(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) { // Handle authentication
		// if valid auth f(w, r)
		// else send_error(w, r)
		f(w, r)
	}
}

func readUserTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := fmt.Sprintf("SELECT UserId, Email, Username, Password, ThreadId, ResponseId FROM Users")

	// Execute query
	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	defer rows.Close()

	var users []model.UserTable
	for rows.Next() {
		var user model.UserTable
		rows.Scan(&user.UserId, &user.Email, &user.Username, &user.Password, &user.ThreadId, &user.ResponseId)
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = model.UserJsonResponse{ Type: "Success", Data: users }
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func readThreadTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := fmt.Sprintf("SELECT ThreadId, UserId, ResponseId, Name, Description, DateCreated FROM Thread")

	// Execute query
	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	defer rows.Close()

	var threads []model.ThreadTable
	for rows.Next() {
		var thread model.ThreadTable
		rows.Scan(&thread.ThreadId, &thread.UserId, &thread.ResponseId, &thread.Name, &thread.Description, &thread.DateCreated)
		threads = append(threads, thread)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = model.ThreadJsonResponse{ Type: "Success", Data: threads }
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func readApiFavoritesTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := fmt.Sprintf("SELECT ApiFavoritesId, UserId, StockId, ApiUrl FROM ApiFavorites")

	// Execute query
	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	var response = model.ApiFavoritesJsonResponse{ Type: "Success", Data: apiFavorites }
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func readResponseTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := fmt.Sprintf("SELECT ResponseId, UserId, ThreadId, Reply FROM Response")

	// Execute query
	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	defer rows.Close()

	var responses []model.ResponseTable
	for rows.Next() {
		var response model.ResponseTable
		rows.Scan(&response.ResponseId, &response.UserId, &response.ThreadId, &response.Reply)
		responses = append(responses, response)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var resp = model.ResponsesJsonResponse{ Type: "Success", Data: responses }
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func readThreadFavoritesTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlQuery := fmt.Sprintf("SELECT ThreadFavoritesId, UserId FROM ThreadFavorites")

	// Execute query
	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	var response = model.ThreadFavoritesJsonResponse{ Type: "Success", Data: threadFavorites }
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlMutation := fmt.Sprintf("INSERT INTO Users VALUES(%s, %s, %s)") // finish mutation

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = model.UserJsonResponse{ Type: "Success", Data: }
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func createNewThread(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlMutation := fmt.Sprintf("INSERT INTO Thread VALUES(%d, %s, %s, %s)") // finish mutation

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = model.ThreadJsonResponse{ Type: "Success", Data:  }
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func createNewApiFavorite(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlMutation := fmt.Sprintf("INSERT INTO ApiFavorites VALUES(%d, %s, %s)") // finish mutation

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = model.ApiFavoritesJsonResponse{ Type: "Success", Data:  }
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func createResponse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlMutation := fmt.Sprintf("INSERT INTO Response VALUES(%d, %d, %s)") // finish mutation

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var resp = model.ResponsesJsonResponse{ Type: "Success", Data:  }
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func createThreadFavorite(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tsqlMutation := fmt.Sprintf("INSERT INTO ThreadFavorite VALUES(%d)") // finish mutation

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response = model.ThreadFavorite{ Type: "Success", Data:  }
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
