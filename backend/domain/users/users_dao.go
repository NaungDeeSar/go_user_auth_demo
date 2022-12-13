package users

import (
	"user_auth_golang/backend/datasource/mysql/users_db"
	"user_auth_golang/backend/utils/errors"
)

var (
	queryInsertUser    = "INSERT INTO users_db_01.user_db(first_name,last_name,email,password)VALUES(?,?,?,?);"
	querySelectByEmail = "SELECT id first_name,last_name,email,password FROM users_db_01.user_db WHERE email = ?;"
    queryGetByID = "SELECT id,first_name,last_name,email FROM users_db_01.user_db WHERE id=?;"
)

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServerError("Database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("Database error")
	}
	user.ID = userID
	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(querySelectByEmail)
	if err != nil {
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password); getErr != nil {
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) GetByID() *errors.RestErr{
	stmt, err :=users_db.Client.Prepare(queryGetByID)

	if err != nil {
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return errors.NewInternalServerError("Database error")
	}
	return nil
}