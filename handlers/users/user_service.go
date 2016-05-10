package users

import (
	"errors"

	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/models"
	"github.com/gkiryaziev/go-gorilla-mysql-sqlx-example/utils"
)

// getUsers return all users from db
func (us *UserHandler) getUsers() *utils.ResultTransformer {

	// concurrency safe
	us.lck.RLock()
	defer us.lck.RUnlock()

	users := []User{}

	err := us.db.Select(&users, "select * from tbl_users order by id")
	if err != nil {
		panic(err)
	}

	header := models.Header{Status: "ok", Count: len(users), Data: users}
	result := utils.NewResultTransformer(header)

	return result
}

// getUser return user by id from db
func (us *UserHandler) getUser(id int64) (*utils.ResultTransformer, error) {

	// concurrency safe
	us.lck.RLock()
	defer us.lck.RUnlock()

	user := User{}

	err := us.db.Get(&user, "select * from tbl_users where id = ?", id)
	if err != nil {
		return nil, err
	}

	header := models.Header{Status: "ok", Count: 1, Data: user}
	result := utils.NewResultTransformer(header)

	return result, nil
}

// updateUser update user and get rows affected in db
func (us *UserHandler) updateUser(user User) error {

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

// deleteUserByID delete user by id and get rows affected in db
func (us *UserHandler) deleteUserByID(id int64) error {

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

// deleteUser delete user and get rows affected in db
func (us *UserHandler) deleteUser(user User) error {

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

// insertUser insert new user and get last id from db
func (us *UserHandler) insertUser(user User) (int64, error) {

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
