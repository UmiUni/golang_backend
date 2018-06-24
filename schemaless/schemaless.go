package schemaless

import (
	"code.jogchat.internal/go-schemaless/core"
	"code.jogchat.internal/golang_backend/utils"
	"encoding/json"
	"context"
	"errors"
	"time"
	"code.jogchat.internal/go-schemaless/models"
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

func SignupDB(username string, email string, password string, token string) (successful bool, err error) {
	_, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "users", "email", email)
	if found {
		return false, errors.New("email already registered")
	}
	password_hash, _ := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	token_hash, _ := bcrypt.GenerateFromPassword([]byte(token), hashCost)

	body, _ := json.Marshal(map[string]interface{} {
		"id": utils.NewUUID(),
		"username": username,
		"email": email,
		"password": string(password_hash),
		"token": string(token_hash),
		"activate": false,
	})

	cell := models.Cell{
		RowKey: utils.NewUUID().Bytes(),
		ColumnName: "users",
		RefKey: time.Now().UnixNano(),
		Body: body,
	}
	err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell)
	utils.CheckErr(err)

	return true, nil
}

func SigninDB(email string, password string) (info map[string]string, successful bool, err error) {
	cells, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "users", "email", email)
	if !found {
		return nil, false, errors.New("unregistered email")
	}
	if len(cells) != 1 {
		panic("error: duplicate email address")
	}
	var cell map[string]interface{}
	err = json.Unmarshal(cells[0].Body, &cell)
	utils.CheckErr(err)

	if err = bcrypt.CompareHashAndPassword([]byte(cell["password"].(string)), []byte(password)); err != nil {
		return nil, false, errors.New("invalid password")
	}
	if !cell["activate"].(bool) {
		return nil, false, errors.New("please verify your email")
	}
	info = map[string]string {
		"id": cell["id"].(string),
		"username": cell["username"].(string),
		"email": cell["email"].(string),
	}
	return info, true, nil
}

func VerifyEmail(email string, token string) (rowKey []byte, successful bool, err error) {
	cells, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "users", "email", email)
	if !found {
		return nil, false, errors.New("unregistered email")
	}
	if len(cells) != 1 {
		panic("error: duplicate email address")
	}
	var body map[string]interface{}
	err = json.Unmarshal(cells[0].Body, &body)
	utils.CheckErr(err)

	if err = bcrypt.CompareHashAndPassword([]byte(body["token"].(string)), []byte(token)); err != nil {
		return nil, false, errors.New("invalid password")
	}

	return cells[0].RowKey, true, nil
}

func ActivateEmail(rowKey []byte) {
	cell, _, _ := DataStore.GetCellLatest(context.TODO(), rowKey, "users")
	var body map[string]interface{}
	err := json.Unmarshal(cell.Body, &body)
	utils.CheckErr(err)

	body["activate"] = true
	cell.Body, err = json.Marshal(body)
	cell.RefKey = time.Now().UnixNano()
	err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell)
	utils.CheckErr(err)
}
