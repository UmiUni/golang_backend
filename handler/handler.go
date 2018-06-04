package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"log"
	"errors"
	"strconv"
	"math/rand"
	"code.jogchat.internal/jogchat_golang_backend/postgres"
)

// Creds holds the credentials we send back
type Creds struct {
	Status      string
	APIKey      string
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


func myLookupKey(key string) []byte {
	return []byte(key)
}



func hasValidToken(jwtToken, key string) bool {
	ret := false
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return myLookupKey(key), nil
	})

	if err == nil && token.Valid {
		ret = true
	}
	return ret
}


// FakeData provides some fake data for testing...
func FakeData(env *Env, w http.ResponseWriter, r *http.Request) error {

	// If this is an OPTION method, then we don't do anything since it's just
	// validating the preflight info
	if r.Method == "OPTIONS" {
		return nil
	}

	// Validate the API call
	secret := r.Header.Get("x-authentication")
	isValid := hasValidToken(secret, env.Secret)
	if !isValid {
		return StatusError{401, errors.New("Invalid authorization token")}
	}

	// If we get here, then we've got a valid call.  So, go ahead and
	// process the request.  In this instance, all we want to do
	// is pass back a list of integers that is N items long
	N, err := strconv.Atoi(r.FormValue("N"))
	if err != nil {
		N = 10
	}

	// Create N random integers
	rand.Seed(time.Now().UTC().UnixNano())
	var data []int
	for i := 0; i < N; i++ {
		data = append(data, rand.Intn(100))
	}

	time.Sleep(1 * time.Second) // just for fun, pause a bit

	// return the results
	out, _ := json.MarshalIndent(&data, "", "  ")
	fmt.Fprintf(w, string(out))
	return nil
}


// GetCredentials determines if the username and password is valid
// This is where logic would go to validate and return account info
func GetCredentials(env *Env, username, password string) Creds {
	credentials := Creds{
		Status:      "UNAUTHORIZED",
		APIKey:      "",
		AccountType: "",
		Email:       "",
		AuthToken:   "",
		IsLoggedIn:  false,
	}
	if (username == "admin") && (password == "admin") {
		credentials.Status = "OK"
		credentials.APIKey = "12345"
		credentials.AccountType = "admin"
		credentials.Email = "admin@example.com"
		credentials.IsLoggedIn = true
		// Now create a JWT for user
		// Create the token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set some claims
		claims := token.Claims.(jwt.MapClaims)
		claims["sub"] = username
		claims["iss"] = "example.com"
		claims["exp"] = time.Now().Add(time.Hour *72).Unix()
		var err error
		credentials.AuthToken, err = token.SignedString([]byte(env.Secret))
		if err != nil {
			log.Println(err)
		}
	}
	return credentials
}

func addCredentials(env *Env, w http.ResponseWriter, params map[string]string) {
	credentials := GetCredentials(env, params["Username"], params["Password"])

	out, _ := json.MarshalIndent(&credentials, "", "  ")
	fmt.Fprintf(w, string(out))
}


// Login captures the data posted to the /login route
func Login(env *Env, w http.ResponseWriter, r *http.Request) error {
	data, _ := ioutil.ReadAll(r.Body) // Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(data, &params)

	postgres.LoginDB(w, params)
	addCredentials(env, w, params)
	return nil
}

func Signup(env *Env, w http.ResponseWriter, r *http.Request) error {
	data, _ := ioutil.ReadAll(r.Body)

	var params map[string]string
	json.Unmarshal(data, &params)

	postgres.SignupDB(w, params)
	addCredentials(env, w, params)
	return nil
}
