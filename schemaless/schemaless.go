package schemaless

import (
	"code.jogchat.internal/go-schemaless/core"
	"code.jogchat.internal/go-schemaless/utils"
	"code.jogchat.internal/go-schemaless/storage/mysql"
	"os"
	"io/ioutil"
	"encoding/json"
	"context"
)

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

func SignupDB()  {

}

func SigninDB()  {

}
