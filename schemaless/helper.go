package schemaless

import (
	"github.com/satori/go.uuid"
	"code.jogchat.internal/golang_backend/utils"
	"context"
	"code.jogchat.internal/go-schemaless/models"
	"time"
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Generate a unique UUID for entry
// columnKey represents the table name
func uniqueUUID(columnKey string) uuid.UUID {
	id := utils.NewUUID()
	for {
		// UUIDs are stored as string
		duplicate, _ := DataStore.CheckValueExist(context.TODO(), columnKey, "id", id)
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

func mutateCell(cell models.Cell, body map[string]interface{}) models.Cell {
	cell.Body, _ = json.Marshal(body)
	cell.RefKey = time.Now().UnixNano()
	return cell
}

func verifyEmailToken(email string, token string) (cell models.Cell, body map[string]interface{}, successful bool, err error) {
	cells, found, _ := DataStore.GetCellsByFieldLatest(context.TODO(), "users", "email", email)
	if !found {
		return cell, nil, false, errors.New("unregistered email")
	}
	cell = cells[0]
	json.Unmarshal(cell.Body, &body)
	if err = bcrypt.CompareHashAndPassword([]byte(body["token"].(string)), []byte(token)); err != nil {
		return cell, nil, false, errors.New("invalid token")
	}
	return cell, body, true, nil
}
