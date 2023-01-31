package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (u Users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (u Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = "%" + nameOrNick + "%"
	rows, err := u.db.Query("select id, name, nick, email, created_at from users where name like ? or nick like ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u Users) GetByID(userID uint64) (models.User, error) {
	row, err := u.db.Query("select id, name, nick, email, created_at from users where id = ?", userID)

	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}

func (u Users) Update(userID uint64, user models.User) error {
	statement, err := u.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil
}

func (u Users) Delete(userID uint64) error {
	statement, err := u.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(userID); err != nil {
		return err
	}

	return nil
}
