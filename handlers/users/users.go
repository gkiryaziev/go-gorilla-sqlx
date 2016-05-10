package users

import (
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/models"
	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/utils"
)

// UserHandler struct
type UserHandler struct {
	db  *sqlx.DB
	lck sync.RWMutex
}

// NewUserHandler return new UserHandler object
func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{db: db}
}

// errorMessage return error message as json string
func errorMessage(status int, msg string) string {
	msgFinal := &models.ResponseMessage{Status: status, Message: msg, Info: "/docs/api/errors"}
	result, _ := utils.NewResultTransformer(msgFinal).ToJSON()
	return result
}
