package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	model "stock-and-crypto-api/models"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	gomail "gopkg.in/gomail.v2"
)

var db *sql.DB
var server = "localhost"
var port = 59236
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
	router.HandleFunc("/readApiFavorites", mwCheck(ReadApiFavoritesTable)).Methods(http.MethodPost)
	router.HandleFunc("/readThreadFavorites", mwCheck(ReadThreadFavoritesTable)).Methods(http.MethodPost)
	router.HandleFunc("/readResponse", ReadResponseTable).Methods(http.MethodPost)

	router.HandleFunc("/newThread", mwCheck(CreateNewThread)).Methods(http.MethodPost)
	router.HandleFunc("/newThreadResponse", mwCheck(CreateResponse)).Methods(http.MethodPost)
	router.HandleFunc("/newApiFavorite", mwCheck(CreateNewApiFavorite)).Methods(http.MethodPost)
	router.HandleFunc("/newThreadFavorite", mwCheck(CreateThreadFavorite)).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    ":8000",
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.UserJsonResponse{Message: "", Type: "Success", Data: users}
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.ThreadJsonResponse{Message: "", Type: "Success", Data: threads}
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.ApiFavoritesJsonResponse{Message: "", Type: "Success", Data: apiFavorites}
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.ResponsesJsonResponse{Message: "", Type: "Success", Data: responses}
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.ThreadFavoritesJsonResponse{Message: "", Type: "Success", Data: threadFavorites}
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

	tsqlQuery = fmt.Sprintf("INSERT INTO Threads VALUES(%d, '%s', '%s', '%s')", nThread.UserId, nThread.Name, nThread.Description, nThread.DateCreated)

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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.GenericJsonResponse{Message: "Successfully created the Thread!", Type: "Success"}
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

	var nApiFavorite model.ApiFavorite

	err = json.NewDecoder(r.Body).Decode(&nApiFavorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tsqlQuery := "SELECT Api_Favorites_Id, User_Id, Stock_Id, Api_Url FROM Api_Favorites;"

	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var apiFavorite model.ApiFavoritesTable
		rows.Scan(&apiFavorite.ApiFavoritesId, &apiFavorite.UserId, &apiFavorite.StockId, &apiFavorite.ApiUrl)
		if apiFavorite.UserId == nApiFavorite.UserId && apiFavorite.StockId == nApiFavorite.StockId && apiFavorite.ApiUrl == nApiFavorite.ApiUrl {
			http.Error(w, "The given Api Favorite has already been created!", http.StatusBadRequest)
			return
		}
	}

	tsqlQuery = fmt.Sprintf("INSERT INTO Api_Favorites VALUES(%d, '%s', '%s');", nApiFavorite.UserId, nApiFavorite.StockId, nApiFavorite.ApiUrl)

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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.GenericJsonResponse{Message: "Successfully created the Api Favorite!", Type: "Success"}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateResponse(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var nResponse model.Response

	err = json.NewDecoder(r.Body).Decode(&nResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tsqlQuery := "SELECT Response_Id, User_Id, Thread_Id, Description, Date_Created FROM Responses;"

	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var response model.ResponseTable
		rows.Scan(&response.ResponseId, &response.UserId, &response.ThreadId, &response.Description, &response.DateCreated)
		if response.UserId == nResponse.UserId && response.ThreadId == nResponse.ThreadId && response.Description == nResponse.Description {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	tsqlQuery = fmt.Sprintf("INSERT INTO Responses VALUES(%d, %d, '%s', '%s')", nResponse.UserId, nResponse.ThreadId, nResponse.Description, nResponse.DateCreated)

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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.GenericJsonResponse{Message: "Successfully create Response", Type: "Success"}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateThreadFavorite(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var nThread model.ThreadFavorites
	err = json.NewDecoder(r.Body).Decode(&nThread)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tsqlQuery := "SELECT Thread_Favorites_Id, User_Id FROM Thread_Favorites;"

	rows, err := db.QueryContext(ctx, tsqlQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var Thread model.ThreadFavoritesTable
		rows.Scan(&Thread.ThreadFavoritesId, &Thread.UserId)
		if Thread.UserId == nThread.UserId {
			http.Error(w, "The given thread favorite has already been created", http.StatusInternalServerError)
			return
		}
	}

	tsqlQuery = fmt.Sprintf("INSERT INTO Thread_Favorites VALUES(%d);", nThread.UserId)

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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.GenericJsonResponse{Message: "Successfully created Thread Favorite", Type: "Success"}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var u model.UserTable
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Check for valid login
	tsql := fmt.Sprintf("SELECT User_Id FROM Users WHERE Email='%s' AND Password='%s';", u.Email, u.Password)
	row := db.QueryRowContext(ctx, tsql)
	var uid int
	if err = row.Scan(&uid); err != nil {
		http.Error(w, "No login found", http.StatusUnauthorized)
		return
	}

	tsql = fmt.Sprintf("SELECT SessionId FROM Sessions WHERE User_Id=%d AND IsActive=1;", uid)
	row = db.QueryRowContext(ctx, tsql)

	var sid int
	err = row.Scan(&sid)
	fmt.Println(sid)
	if sid > 0 {
		http.Error(w, "You are already logged in!", http.StatusForbidden)
		return
	}
	//Log user in by creating a session
	guid := uuid.New()
	tsql = fmt.Sprintf("INSERT INTO Sessions VALUES(%d, '%s', 1)", uid, guid)
	res, err := db.ExecContext(ctx, tsql)
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.LoginJsonResponse{Message: "Logged In", Type: "Success", UserGuid: guid.String(), UserId: uId }
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var u model.User
	err = json.NewDecoder(r.Body).Decode((&u))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Check to see if email or username is already in use
	tsql := "SELECT Username, Email FROM Users;"
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var nu model.UserTable
		rows.Scan(&nu.Username, &nu.Email)
		if nu.Username == u.Username || nu.Email == u.Email {
			http.Error(w, "The given username or email has already been used!", http.StatusBadRequest)
			return
		}
	}

	//Add username
	tsql = fmt.Sprintf("INSERT INTO Users VALUES ('%s','%s','%s');", u.Email, u.Username, u.Password)
	res, err := db.ExecContext(ctx, tsql)
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	w.WriteHeader(http.StatusOK)
	resp := model.LoginJsonResponse{Message: "Added User ", Type: "Success"}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Verify Database is alive
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var u model.User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", "cooltestemail23@gmail.com")
	msg.SetHeader("To", u.Email)
	msg.SetHeader("Subject", "Password Reset")
	msg.SetBody("text/html", "<b>Your password has been changed to a temporary Password of P@$$word!</b>")

	n := gomail.NewDialer("smtp.gmail.com", 587, "cooltestemail23@gmail.com", "PassTest1!")

	// Send the email
	err = n.DialAndSend(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tsql := fmt.Sprintf("UPDATE Users SET Password = 'P@$$word!' WHERE Email = '%s'", u.Email)

	res, err := db.ExecContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count, err := res.RowsAffected()
	if err != nil || count != 1 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.GenericJsonResponse{Message: "Message Sent", Type: "Successessfully reset Password"}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
