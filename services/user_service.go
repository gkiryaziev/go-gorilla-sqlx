package services

import (
	"errors"
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/models"
	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/utils"
)

// UserService struct
type UserService struct {
	db  *sqlx.DB
	lck sync.RWMutex
}

// NewUserService return new UserService object
func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{db: db}
}

// GetUsers return all users
func (us *UserService) GetUsers() *utils.ResultTransformer {

	// concurrency safe
	us.lck.RLock()
	defer us.lck.RUnlock()

	users := []models.User{}

	err := us.db.Select(&users, "select * from tbl_users order by id")
	if err != nil {
		panic(err)
	}

	header := models.Header{Status: "ok", Count: len(users), Data: users}
	result := utils.NewResultTransformer(header)

	return result
}

// GetUser return user by id
func (us *UserService) GetUser(id int64) (*utils.ResultTransformer, error) {

	// concurrency safe
	us.lck.RLock()
	defer us.lck.RUnlock()

	user := models.User{}

	err := us.db.Get(&user, "select * from tbl_users where id = ?", id)
	if err != nil {
		return nil, err
	}

	header := models.Header{Status: "ok", Count: 1, Data: user}
	result := utils.NewResultTransformer(header)

	return result, nil
}

// UpdateUser update user and get rows affected
func (us *UserService) UpdateUser(user models.User) error {

	// concurrency safe
	us.lck.Lock()
	defer us.lck.Unlock()

	result, err := us.db.NamedExec("update tbl_users set "+
		"first_name=:first_name, last_name=:last_name, middle_name=:middle_name, "+
		"dob=:dob, address=:address, phone=:phone, login=:login, password=:password "+
		"where id=:id", user)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows <= 0 {
		return errors.New("0 Rows Affected")
	}

	return nil
}

// DeleteUserByID delete user by id and get rows affected
func (us *UserService) DeleteUserByID(id int64) error {

	// concurrency safe
	us.lck.Lock()
	defer us.lck.Unlock()

	result, err := us.db.NamedExec("delete from tbl_users where id = :id", map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows <= 0 {
		return errors.New("0 Rows Affected")
	}

	return nil
}

// DeleteUser delete user and get rows affected
func (us *UserService) DeleteUser(user models.User) error {

	// concurrency safe
	us.lck.Lock()
	defer us.lck.Unlock()

	result, err := us.db.NamedExec("delete from tbl_users where id = :id", user)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows <= 0 {
		return errors.New("0 Rows Affected")
	}

	return nil
}

// InsertUser insert new user and get last id
func (us *UserService) InsertUser(user models.User) (int64, error) {

	// concurrency safe
	us.lck.Lock()
	defer us.lck.Unlock()

	result, err := us.db.NamedExec("insert into tbl_users ("+
		"first_name, last_name, middle_name, dob, address, phone, login, password) values ("+
		":first_name, :last_name, :middle_name, :dob, :address, :phone, :login, :password)", user)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
