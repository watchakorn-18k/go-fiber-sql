package repositories

import (
	"fmt"
	"go-fiber-sql/src/domain/datasources"
	"go-fiber-sql/src/domain/entities"
)

type userRepo struct {
	DB *datasources.SqlDb
}

type IUsersRepository interface {
	InsertNewUser(data entities.UserDataFormat) error
	FindAll() (*[]entities.UserDataFormat, error)
	GetUser(userID string) (*entities.UserDataFormat, error)
	UpdateUser(userID string, data *entities.NewUserBody) error
	DeleteUser(userID string) error
}

func NewUsersRepository(db *datasources.SqlDb) IUsersRepository {
	return &userRepo{db}
}

func (r *userRepo) InsertNewUser(data entities.UserDataFormat) error {
	_, err := r.DB.Connect.Exec("INSERT INTO users (user_id, username, email) VALUES (?, ?, ?)", data.UserID, data.Username, data.Email)
	if err != nil {
		fmt.Println("error insert : ", err)
		return err
	}

	return nil
}

func (r *userRepo) FindAll() (*[]entities.UserDataFormat, error) {
	var userData []entities.UserDataFormat
	cursor, err := r.DB.Connect.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for cursor.Next() {
		var data entities.UserDataFormat
		cursor.Scan(&data.UserID, &data.Username, &data.Email)
		userData = append(userData, data)
	}

	return &userData, nil
}

func (r *userRepo) GetUser(userID string) (*entities.UserDataFormat, error) {
	var userData entities.UserDataFormat
	data := r.DB.Connect.QueryRow("SELECT * FROM users WHERE user_id = ?", userID).Scan(&userData.UserID, &userData.Username, &userData.Email)
	if data != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &userData, nil
}

func (r *userRepo) UpdateUser(userID string, data *entities.NewUserBody) error {
	_, err := r.DB.Connect.Exec("UPDATE users SET username = ?, email = ? WHERE user_id = ?", data.Username, data.Email, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) DeleteUser(userID string) error {
	_, err := r.DB.Connect.Exec("DELETE FROM users WHERE user_id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}
