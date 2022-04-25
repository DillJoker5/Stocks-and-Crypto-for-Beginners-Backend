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
	"github.com/google/uuid"
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

	// Handle api endpoints
	router.HandleFunc("/login", Login).Methods(http.MethodPost)
	router.HandleFunc("/register", Register).Methods(http.MethodPost)
	router.HandleFunc("/forgotPassword", ForgotPassword).Methods(http.MethodPost)

	router.HandleFunc("/readThread", ReadThreadTable).Methods(http.MethodPost)
	router.HandleFunc("readApiFavorites", ReadApiFavoritesTable).Methods(http.MethodPost)
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
	
	// Verify database is running
	err := db.PingContext(ctx)
	if err != nil {
		return false
	}

	var s model.Session
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		return false
	}

	if s.UserGuid == "" {
		return false
	}

	tsqlQuery := fmt.Sprintf("SELECT SessionId FROM Session WHERE UserGuid='%s' AND IsActive=1;", s.UserGuid)

	row := db.QueryRowContext(ctx, tsqlQuery)

	var sid int32
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
}

func ReadThreadTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReadApiFavoritesTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReadResponseTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReadThreadFavoritesTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateNewThread(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
