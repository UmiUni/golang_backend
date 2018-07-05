package schemaless

import (
	"code.jogchat.internal/go-schemaless/core"
	"code.jogchat.internal/golang_backend/utils"
	"encoding/json"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"code.jogchat.internal/go-schemaless"
)

const hashCost = 8
var DataStore *core.KVStore


func InitDB()  {
	DataStore = schemaless.InitDataStore()
}

func CloseDB()  {
	DataStore.Destroy(context.TODO())
}

func SignupDB(email string, token string) (successful bool, err error) {
	duplicate, _ := DataStore.CheckValueExist(context.TODO(), "users", "email", email)
	if duplicate {
		return false, errors.New("email already registered")
	}
	token_hash, _ := bcrypt.GenerateFromPassword([]byte(token), hashCost)

	body := map[string]interface{} {
		"email": email,
		"token": string(token_hash),
		"activate": false,
	}
	_, cell, err := constructCell("users", body)

	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "token")
		utils.CheckErr(err)
	}()

	return true, nil
}

func ActivateEmail(email string, username string, password string, token string) (info map[string]string, successful bool, err error) {
	_, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "users", "username", username)
	if found {
		return nil, false, errors.New("username already in use")
	}
	cell, body, successful, err := verifyEmailToken(email, token)
	if !successful {
		return nil, false, err
	}

	password_hash, _ := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	body["username"] = string(username)
	body["password"] = string(password_hash)
	body["activate"] = true
	cell = mutateCell(cell, body)
	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password", "token")
		utils.CheckErr(err)
	}()

	info = map[string]string {
		"id": body["id"].(string),
		"username": body["username"].(string),
		"email": body["email"].(string),
	}
	return info, true, nil
}

func SigninDB(email string, password string) (info map[string]string, successful bool, err error) {
	cells, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "users", "email", email)
	if !found {
		return nil, false, errors.New("unregistered email")
	}
	var cell map[string]interface{}
	json.Unmarshal(cells[0].Body, &cell)

	if !cell["activate"].(bool) {
		return nil, false, errors.New("please verify your email")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(cell["password"].(string)), []byte(password)); err != nil {
		return nil, false, errors.New("invalid password")
	}
	info = map[string]string {
		"id": cell["id"].(string),
		"username": cell["username"].(string),
		"email": cell["email"].(string),
	}
	return info, true, nil
}

func ResetRequest(email string, token string) (found bool, err error) {
	cells, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "users", "email", email)
	if !found {
		return false, errors.New("unregistered email")
	}
	cell := cells[0]

	var body map[string]interface{}
	json.Unmarshal(cell.Body, &body)
	token_hash, _ := bcrypt.GenerateFromPassword([]byte(token), hashCost)
	body["token"] = string(token_hash)
	cell = mutateCell(cell, body)

	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password", "token")
		utils.CheckErr(err)
	}()

	return true, nil
}

func ResetPassword(email string, password string, token string) (info map[string]string, successful bool, err error) {
	cell, body, successful, err := verifyEmailToken(email, token)
	if !successful {
		return nil, false, err
	}
	if !body["activate"].(bool) {
		return nil, false, errors.New("please verify your email")
	}

	password_hash, _ := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	body["password"] = string(password_hash)
	cell = mutateCell(cell, body)
	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password", "token")
		utils.CheckErr(err)
	}()

	info = map[string]string {
		"id": body["id"].(string),
		"username": body["username"].(string),
		"email": body["email"].(string),
	}
	return info, true, nil
}
