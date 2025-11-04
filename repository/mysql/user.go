package mysql

import (
	"Game/entity"
	"database/sql"
	"fmt"
	"time"
)


func (d *MySQLDB) IsPhoneNumberUnique(phonenumber string) (bool, error) {
	user := entity.User{}
	var createdAt time.Time

	row := d.db.QueryRow(`select * from users where phone_number = ?`, phonenumber)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("cannot scan query row: %w", err)
	}
	return false, nil
}



func (d *MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number) values(?, ?)`, u.Name, u.PhoneNumber)
	if err != nil {
		return entity.User{}, fmt.Errorf("cannnot execute command: %w", err)
	}
	id, _ := res.LastInsertId()
    
    u.ID = int(id)
	return u, nil
}
