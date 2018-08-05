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

func constructCell(columnKey string, body map[string]interface{}, uniqueId bool) (id uuid.UUID, cell models.Cell) {
	if uniqueId {
		id = uniqueUUID(columnKey)
	} else {
		id = utils.NewUUID()
	}
	body["id"] = id.String()
	json_body, err := json.Marshal(body)
	utils.CheckErr(err)

	cell = models.Cell{
		RowKey: utils.NewUUID().Bytes(),
		ColumnName: columnKey,
		RefKey: time.Now().UnixNano(),
		Body: json_body,
	}
	return id, cell
}

func mutateCell(cell models.Cell, body map[string]interface{}) models.Cell {
	cell.Body, _ = json.Marshal(body)
	cell.RefKey = time.Now().UnixNano()
	return cell
}

func verifyEmailToken(email string, token string) (cell models.Cell, body map[string]interface{}, successful bool, err error) {
	cell, body, found, _ := getUserByUniqueField("email", email)
	if !found {
		return cell, nil, false, errors.New("invalid email")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(body["token"].(string)), []byte(token)); err != nil {
		return cell, nil, false, errors.New("invalid token")
	}
	return cell, body, true, nil
}

func getUserByUniqueField(field string, value string) (cell models.Cell, body map[string]interface{}, found bool, err error) {
	return getEntityByUniqueField("users", field, value)
}

func getEntityByUniqueField(category string, field string, value interface{}) (cell models.Cell, body map[string]interface{}, found bool, err error) {
	cell, found, err = DataStore.GetCellByUniqueFieldLatest(context.TODO(), category, field, value)
	if !found {
		return cell, nil, false, err
	}
	json.Unmarshal(cell.Body, &body)
	return cell, body, true, nil
}
