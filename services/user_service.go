package services

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"../models"
	"../utils"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{db: db}
}

// get all users
func (this *UserService) GetUsers() *utils.ResultTransformer {

	users := []models.User{}

	err := this.db.Select(&users, "select * from tbl_users order by id")
	if err != nil {
		panic(err)
	}

	header := models.Header{"ok", len(users), users}
	result := utils.NewResultTransformer(header)

	return result
}

// get user by id
func (this *UserService) GetUser(id int64) (*utils.ResultTransformer, error) {

	user := models.User{}

	err := this.db.Get(&user, "select * from tbl_users where id = ?", id)
	if err != nil {
		return nil, err
	}

	header := models.Header{"ok", 1, user}
	result := utils.NewResultTransformer(header)

	return result, nil
}

// update user and get rows affected
func (this *UserService) UpdateUser(user models.User) error {

	result, err := this.db.NamedExec("update tbl_users set "+
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

// delete user by id and get rows affected
func (this *UserService) DeleteUserById(id int64) error {

	result, err := this.db.NamedExec("delete from tbl_users where id = :id", map[string]interface{}{"id": id})
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

// delete user and get rows affected
func (this *UserService) DeleteUser(user models.User) error {

	result, err := this.db.NamedExec("delete from tbl_users where id = :id", user)
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

// insert new user and get last id
func (this *UserService) InsertUser(user models.User) (int64, error) {

	result, err := this.db.NamedExec("insert into tbl_users ("+
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
