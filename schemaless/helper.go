package schemaless

import (
	"github.com/satori/go.uuid"
	"code.jogchat.internal/golang_backend/utils"
	"context"
	"code.jogchat.internal/go-schemaless/models"
	"time"
	"encoding/json"
)

// Generate a unique UUID for entry
// columnKey represents the table name
func uniqueUUID(columnKey string) uuid.UUID {
	id := utils.NewUUID()
	for {
		// UUIDs are stored as bytes
		duplicate, _ := DataStore.CheckValueExist(context.TODO(), columnKey, "id", id.Bytes())
		if !duplicate {
			break
		} else {
			id = utils.NewUUID()
		}
	}
	return id
}

func constructCell(columnKey string, body map[string]interface{}) (id uuid.UUID, cell models.Cell, err error) {
	id = uniqueUUID(columnKey)
	body["id"] = id.String()
	json_body, err := json.Marshal(body)
	if err != nil {
		return id, cell, err
	}

	cell = models.Cell{
		RowKey: utils.NewUUID().Bytes(),
		ColumnName: columnKey,
		RefKey: time.Now().UnixNano(),
		Body: json_body,
	}
	return id, cell, nil
}

func cellsToMaps(cells []models.Cell) (results []map[string]interface{}, err error) {
	for _, cell := range cells {
		var body map[string]interface{}
		err = json.Unmarshal(cell.Body, &body)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		results = append(results, body)
	}
	return results, nil
}
