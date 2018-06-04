package postgres

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/satori/go.uuid"
)

type Credentials struct {
	Username string
	Password string
}


func SignupDB(w http.ResponseWriter, params map[string]string) {
	username, email, password := params["Username"], params["Password"], params["Email"]

	if username == "" || email == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	user_id := uuid.Must(uuid.NewV4())

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)

	// Next, insert the username, along with the hashed password into the database
	if _, err = db.Exec("INSERT INTO users VALUES ($1, $2, $3, $4, false)", user_id, username, email, string(hashed)); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
}

func LoginDB(w http.ResponseWriter, params map[string]string){
	// Parse and decode the request body into a new `Credentials` instance	
	// Get the existing entry present in the database for the given username
	username, password := params["Username"], params["Password"]

	result := db.QueryRow("SELECT password FROM users WHERE username=$1", username)

	var stored_pswd string
	// Store the obtained password in `storedCreds`
	err := result.Scan(&stored_pswd)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(stored_pswd), []byte(password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
	}

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
}