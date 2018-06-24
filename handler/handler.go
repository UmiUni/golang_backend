package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"code.jogchat.internal/jwt-go"
	"time"
	"log"
	"errors"
	"code.jogchat.internal/golang_backend/schemaless"
)

// Creds holds the credentials we send back
type Creds struct {
	Status      string
	AccountType string
	Email       string
	AuthToken   string
	IsLoggedIn  bool
}


// Env holds application-wide configuration.
type Env struct {
	Secret string
}

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
// See:
//    https://annevankesteren.nl/2015/02/same-origin-policy
//    http://stackoverflow.com/questions/22972066/how-to-handle-preflight-cors-requests-on-a-go-server
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "x-authentication")

	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}



// GetCredentials determines if the username and password is valid
// This is where logic would go to validate and return account info
func GetCredentials(env *Env, username string, email string) Creds {
	credentials := Creds{
		Status:      "OK",
		AccountType: "user",
		Email:       email,
		IsLoggedIn:  true,
	}
	// Now create a JWT for user
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = username
	claims["iss"] = "jogchat.com"
	claims["exp"] = time.Now().Add(time.Hour *72).Unix()
	var err error
	credentials.AuthToken, err = token.SignedString([]byte(env.Secret))
	if err != nil {
		log.Println(err)
	}
	return credentials
}

func addCredentials(env *Env, w http.ResponseWriter, username string, email string) {
	credentials := GetCredentials(env, username, email)
	out, _ := json.MarshalIndent(&credentials, "", "  ")
	fmt.Fprintf(w, string(out))
}

// Login captures the data posted to the /login route
func Login(env *Env, w http.ResponseWriter, r *http.Request) error {
	data, _ := ioutil.ReadAll(r.Body) // Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(data, &params)

	info, successful, err := schemaless.SigninDB(params["Email"], params["Password"])
	if !successful {
		handleFailure(err, w)
	} else {
		addCredentials(env, w, info["username"], info["email"])
	}
	return nil
}

func Signup(env *Env, w http.ResponseWriter, r *http.Request) error {
	data, _ := ioutil.ReadAll(r.Body)

	var params map[string]string
	json.Unmarshal(data, &params)

	if params["Username"] == "" || params["Email"] == "" || params["Password"] == "" {
		handleFailure(errors.New("username, email and password cannot be empty"), w)
	} else {
		successful, err := schemaless.SignupDB(params["Username"], params["Email"], params["Password"])
		if !successful {
			handleFailure(err, w)
		} else {
			addCredentials(env, w, params["Username"], params["Email"])
		}
	}
	return nil
}


func handleFailure(err error, w http.ResponseWriter) {
	res, _ := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	fmt.Fprintf(w, string(res))
}
