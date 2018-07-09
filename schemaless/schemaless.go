package schemaless

import (
	"code.jogchat.internal/go-schemaless/core"
	"code.jogchat.internal/golang_backend/utils"
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

func SignupDB(category string, email string, token string) (successful bool, err error) {
	duplicate, _ := DataStore.CheckValueExist(context.TODO(), "users", "email", email)
	if duplicate {
		return false, errors.New("email already registered")
	}
	token_hash, _ := bcrypt.GenerateFromPassword([]byte(token), hashCost)

	body := map[string]interface{} {
		"email": email,
		"token": string(token_hash),
		category: true,
		"activate": false,
	}
	_, cell, err := constructCell("users", body)
	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "token")
		utils.CheckErr(err)
	}()
	return true, nil
}

func ReverifyEmail(email string, token string) (successful bool, err error) {
	cell, body, found, err := getUserByUniqueField("email", email)
	if !found {
		return false, errors.New("email not registered")
	}
	if body["activate"].(bool) {
		return false, errors.New("email already activated")
	}
	body["token"] = token
	cell = mutateCell(cell, body)
	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password", "token")
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
	if body["activate"].(bool) {
		return nil, false, errors.New("email already activated")
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
	_, body, found, _ := getUserByUniqueField("email", email)
	if !found {
		return nil, false, errors.New("email not registered")
	}
	if !body["activate"].(bool) {
		return nil, false, errors.New("please verify your email")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(body["password"].(string)), []byte(password)); err != nil {
		return nil, false, errors.New("invalid password")
	}
	info = map[string]string {
		"id": body["id"].(string),
		"username": body["username"].(string),
		"email": body["email"].(string),
	}
	return info, true, nil
}

func ResetRequest(email string, token string) (found bool, err error) {
	cell, body, found, err := getUserByUniqueField("email", email)
	if !found {
		return false, errors.New("email not registered")
	}
	token_hash, _ := bcrypt.GenerateFromPassword([]byte(token), hashCost)
	body["token"] = string(token_hash)
	cell = mutateCell(cell, body)
	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password", "token", "resume")
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
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password", "token", "resume")
		utils.CheckErr(err)
	}()
	info = map[string]string {
		"id": body["id"].(string),
		"username": body["username"].(string),
		"email": body["email"].(string),
	}
	return info, true, nil
}

func UploadResume(username string, filename string) (successful bool, err error) {
	cell, body, found, err := getUserByUniqueField("username", username)
	if !found {
		return false, errors.New("username does not exist")
	}
	body["resume"] = filename
	cell = mutateCell(cell, body)
	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "password", "token", "resume")
		utils.CheckErr(err)
	}()

	return true, nil
}

func GetResume(username string) (filename string, found bool, err error) {
	_, body, found, err := getUserByUniqueField("username", username)
	if !found {
		return filename, false, errors.New("username does not exist")
	}
	resume, ok := body["resume"]
	if !ok {
		return filename, false, errors.New("resume not uploaded")
	}
	return resume.(string), true, nil
}

func AddCompanySchool(category, name string, domain string, filename string) (successful bool, err error) {
	duplicate, _ := DataStore.CheckValueExist(context.TODO(), category, "name", name)
	if duplicate {
		return false, errors.New(category + " name already exist")
	}
	duplicate, _ = DataStore.CheckValueExist(context.TODO(), category, "domain", domain)
	if duplicate {
		return false, errors.New(category + " domain already exist")
	}
	body := map[string]interface{} {
		"name": name,
		"domain": domain,
		"icon": filename,
	}
	_, cell, err := constructCell(category, body)
	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "icon")
		utils.CheckErr(err)
	}()

	return true, nil
}

func PostPosition(username string, company string, position string, description string) (info map[string]interface{}, successful bool, err error) {
	found, _ := DataStore.CheckValueExist(context.TODO(), "users", "username", username)
	if !found {
		return nil, false, errors.New("username does not exist")
	}
	body := map[string]interface{} {
		"postedBy": username,
		"position": position,
		"description": description,
	}
	id, cell, err := constructCell("position", body)
	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "description")
		utils.CheckErr(err)
	}()

	info = map[string]interface{} {
		"id": id,
	}
	return info, true, nil
}

func CommentOn(username string, positionId string, parentId string, parentType string, content string) (info map[string]interface{}, successful bool, err error) {
	found, _ := DataStore.CheckValueExist(context.TODO(), "users", "username", username)
	if !found {
		return nil, false, errors.New("username does not exist")
	}
	found, _ = DataStore.CheckValueExist(context.TODO(), "position", "id", positionId)
	if !found {
		return nil, false, errors.New("invalid position id")
	}
	found, _ = DataStore.CheckValueExist(context.TODO(), parentType, "id", parentId)
	if !found {
		return nil, false, errors.New("invalid parent id")
	}
	body := map[string]interface{} {
		"username": username,
		"positionId": positionId,
		"parentId": parentId,
		"content": content,
	}
	id, cell, err := constructCell("comment", body)
	go func() {
		err = DataStore.PutCell(context.TODO(), cell.RowKey, cell.ColumnName, cell.RefKey, cell, "content")
		utils.CheckErr(err)
	}()
	info = map[string]interface{} {
		"id": id,
	}
	return info, true, nil
}
