package schemaless

import (
	"code.jogchat.internal/go-schemaless/core"
	"code.jogchat.internal/jogchat-backend/utils"
	"code.jogchat.internal/go-schemaless/storage/mysql"
	"os"
	"io/ioutil"
	"encoding/json"
	"context"
	"errors"
	"github.com/satori/go.uuid"
	"time"
	"code.jogchat.internal/go-schemaless/models"
	"golang.org/x/crypto/bcrypt"
)

const hashCost = 8
var DataStore *core.KVStore


func newBackend(user, pass, host, port, schemaName string) *mysql.Storage {
	m := mysql.New().WithUser(user).
		WithPass(pass).
		WithHost(host).
		WithPort(port).
		WithDatabase(schemaName)

	err := m.WithZap()
	utils.CheckErr(err)
	err = m.Open()
	utils.CheckErr(err)

	// TODO(rbastic): defer Sync() on all backend storage loggers
	return m
}

func getShards(config map[string][]map[string]string) []core.Shard {
	var shards []core.Shard
	hosts := config["hosts"]

	for _, host := range hosts {
		shard := core.Shard{
			Name: host["database"],
			Backend: newBackend(host["user"], host["password"], host["ip"], host["port"], host["database"])}
		shards = append(shards, shard)
	}

	return shards
}

func newUUID() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}

func InitDB()  {
	jsonFile, err := os.Open("config/config.json")
	utils.CheckErr(err)
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	utils.CheckErr(err)

	var config map[string][]map[string]string
	json.Unmarshal(bytes, &config)

	shards := getShards(config)

	DataStore = core.New(shards)
}

func CloseDB()  {
	DataStore.Destroy(context.TODO())
}

func SignupDB(username string, email string, password string) (successful bool, err error) {
	_, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "users", "email", email)
	if found {
		return false, errors.New("email already registered")
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), hashCost)

	body, _ := json.Marshal(map[string]interface{} {
		"id": newUUID(),
		"username": username,
		"email": email,
		"password": string(hashed),
		"activate": false,
	})

	cell := models.Cell{
		RowKey: newUUID().Bytes(),
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
	info = map[string]string {
		"id": cell["id"].(string),
		"username": cell["username"].(string),
		"email": cell["email"].(string),
	}
	return info, true, nil
}
