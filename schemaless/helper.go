package schemaless

import (
	"github.com/satori/go.uuid"
	"code.jogchat.internal/golang_backend/utils"
	"context"
)

// Generate a unique UUID for entry
// columnKey represents the table name
func uniqueUUID(columnKey string) uuid.UUID {
	id := utils.NewUUID()
	for {
		duplicate, _ := DataStore.CheckValueExist(context.TODO(), columnKey, "id", id)
		if !duplicate {
			break
		} else {
			id = utils.NewUUID()
		}
	}
	return id
}
