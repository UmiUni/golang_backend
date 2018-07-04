package schemaless

import (
	"code.jogchat.internal/go-schemaless/core"
	"code.jogchat.internal/golang_backend/utils"
	"encoding/json"
	"context"
	"errors"
	"time"
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

func SignupDB(username string, email string, password string, category, token string) (info map[string]string, successful bool, err error) {
	duplicate, _ := DataStore.CheckValueExist(context.TODO(), "users", "email", email)
	if duplicate {
		return nil, false, errors.New("email already registered")
	}
	duplicate, _ = DataStore.CheckValueExist(context.TODO(), "users", "username", username)
	if duplicate {
		return nil, false, errors.New("username already exist")
	}
	password_hash, _ := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	token_hash, _ := bcrypt.GenerateFromPassword([]byte(token), hashCost)

	body := map[string]interface{}{
		"username": username,
		"email":    email,
		"password": string(password_hash),
		"category": category,
		"token":    string(token_hash),
		"activate": false,
	}
	id, cell, err := constructCell("users", body)

	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password", "token")
		utils.CheckErr(err)
	}()

	info = map[string]string {
		"id": id.String(),
		"username": username,
		"email": email,
	}

	return info, true, nil
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
	body := map[string]interface{} {
		"domain": domain,
		"timestamp": timestamp,
		"author": author,
		"summary": summary,
		"title": title,
		"text": text,
		"url": url,
	}
	_, cell, err := constructCell("news", body)
	if err != nil {
		return false, err
	}

	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "summary", "text", "url")
		utils.CheckErr(err)
	}()

	return true, nil
}

func GetNewsByField(field string, value interface{}) (news []map[string]interface{}, found bool, err error) {
	cells, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "news", field, value)
	if !found {
		return nil, false, errors.New("no news found")
	}

	news, err = cellsToMaps(cells)
	return news, true, nil
}

func CommentOn(user_id string, news_id string, parent_id string, comment_on string, content string, timestamp int64) (comment_id string, successful bool, err error) {
	if exist, _ := DataStore.CheckValueExist(context.TODO(), "users", "id", user_id); !exist {
		return comment_id, false, errors.New("invalid user_id")
	}
	if exist, _ := DataStore.CheckValueExist(context.TODO(), "news", "id", news_id); !exist {
		return comment_id, false, errors.New("invalid news_id")
	}
	if exist, _ := DataStore.CheckValueExist(context.TODO(), comment_on, "id", parent_id); !exist {
		return comment_id, false, errors.New("invalid parent_id")
	}

	body := map[string]interface{} {
		"userId": user_id,
		"newsId": news_id,
		"parentId": parent_id,
		"content": content,
		"timestamp": timestamp,
	}
	id, cell, err := constructCell("comment", body)
	if err != nil {
		return comment_id, false, err
	}

	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "content")
		utils.CheckErr(err)
	}()

	return id.String(), true, nil
}

func GetComment(parent_id string) (comments []map[string]interface{}, found bool, err error) {
	cells, found, err := DataStore.GetCellsByFieldLatest(context.TODO(), "comment", "parentId", parent_id)
	if !found {
		return nil, false, errors.New("no comments found")
	}

	comments, err = cellsToMaps(cells)
	return comments, true, nil
}
