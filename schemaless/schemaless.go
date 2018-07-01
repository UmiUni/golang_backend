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

	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password")
		utils.CheckErr(err)
	}()

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
	if err != nil {
		return nil,false, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(cell["password"].(string)), []byte(password)); err != nil {
		return nil, false, errors.New("invalid password")
	}
	if !cell["activate"].(bool) {
		return nil, false, errors.New("please verify your email")
	}
	info = map[string]string {
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

	if err != nil {
		return nil, false, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(body["token"].(string)), []byte(token)); err != nil {
		return nil, false, errors.New("invalid token")
	}

	return cells[0].RowKey, true, nil
}

func ActivateEmail(rowKey []byte) (err error) {
	cell, _, _ := DataStore.GetCellLatest(context.TODO(), rowKey, "users")
	var body map[string]interface{}
	err = json.Unmarshal(cell.Body, &body)

	if err != nil {
		return err
	}

	if body["activate"].(bool) {
		return errors.New("email already activated")
	}
	body["activate"] = true
	cell.Body, err = json.Marshal(body)
	cell.RefKey = time.Now().UnixNano()

	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password")
		utils.CheckErr(err)
	}()

	return nil
}

func InsertNews(domain string, timestamp int64, author string, summary string, title string, text string, url string) (successful bool, err error) {
	body, err := json.Marshal(map[string]interface{} {
		"id": utils.NewUUID(),
		"domain": domain,
		"timestamp": timestamp,
		"author": author,
		"summary": summary,
		"title": title,
		"text": text,
		"url": url,
	})
	if err != nil {
		return false, err
	}

	cell := models.Cell{
		RowKey: utils.NewUUID().Bytes(),
		ColumnName: "news",
		RefKey: time.Now().UnixNano(),
		Body: body,
	}

	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "summary", "text", "url")
		utils.CheckErr(err)
	}()

	return true, nil
}

func GetNews(domain string) (news []map[string]interface{}, found bool, err error) {
	cells, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "news", "domain", domain)
	if !found {
		return nil, false, errors.New("no news found")
	}

	for _, cell := range cells {
		var body map[string]interface{}
		err = json.Unmarshal(cell.Body, &body)
		if err != nil {
			return nil, false, err
		}
		news = append(news, body)
	}
	return news, true, nil
}
