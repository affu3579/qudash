package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id        int    `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Mobile    int    `json:"Mobile"`
	UserName  string `json:"UserName"`
	Orders    int    `json:"Orders"`
	Password  string `json:"Password"`
	UserRole  uint8  `json:"UserRole"`
	IsActive  uint8  `json:"IsActive"`
}

func CreateUser(u User) (int64, error) {
	log.Print("going to insert record in database for name : ", u.UserName)
	stmt, err := db.Prepare("INSERT INTO user (id, firstname, lastname, username, mobile, password, orders, user_role, is_active)VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Print("error preparing query :: ", err)
		return 0, err
	}
	result, err := stmt.Exec(u.Id, u.FirstName, u.LastName, u.UserName, u.Mobile, u.Password, u.Orders, u.UserRole, u.IsActive)
	if err != nil {
		log.Print("error executing query :: ", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	return id, err

}

func GetUsers() ([]User, error) {
	log.Print("reading records from database")
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Print("error executing select query :: ", err)
		return nil, err
	}
	var users []User
	var id int
	var username string
	var firstname string
	var lastname string
	var mobile int
	var password string
	var orders int
	var userRole uint8
	var isActive uint8
	for rows.Next() {
		err = rows.Scan(&id, &firstname, &lastname, &username, &mobile, &password, &orders, &userRole, &isActive)
		user := User{Id: id,  FirstName: firstname, LastName: lastname, UserName: username, Mobile: mobile,
			Password: password, Orders: orders, UserRole: userRole, IsActive: isActive}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return users, err
}

func GetUser(uid int) *User {
	log.Print("reading records from database")
	rows, err := db.Query("SELECT * FROM user WHERE id=?", uid)
	if err != nil {
		log.Print("error executing select query :: ", err)
		return nil
	}
	var id int
	var username string
	var firstname string
	var lastname string
	var mobile int
	var password string
	var orders int
	var userRole uint8
	var isActive uint8

	var user User
	for rows.Next() {
		err = rows.Scan(&id, &username, &firstname, &lastname, &mobile, &password, &orders, &userRole, &isActive)
		user = User{Id: id, UserName: username, FirstName: firstname, LastName: lastname, Mobile: mobile,
			Password: password, Orders: orders, UserRole: userRole, IsActive: isActive}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return &user
}

func UpdateUser(id int, u User) (int, error) {

	log.Print("going to update record in database for user : ", id)
	sqlStatement := `UPDATE user SET firstname=?, lastname=?, username=?, mobile=?, password=?, orders=?, user_role=?, is_active=? where id=?;`
	stmt, err := db.Prepare(sqlStatement)
	stmt.Exec(u.FirstName, u.LastName, u.UserName, u.Mobile, u.Password, u.Orders, u.UserRole, u.IsActive, id)
	//sqlStatement := `UPDATE user SET firstname=$2, lastname=$3, username=$3, mobile=$4, password=$5, orders=$6, user_role=$7, is_active=$8 where id=5;`
	//_, err := db.Exec(sqlStatement, u.FirstName, u.LastName, u.UserName, u.Mobile, u.Password, u.Orders, u.UserRole, u.IsActive, id)
	if err != nil {
		log.Print("error executing query :: ", err)
		return 0, err
	}
	return u.Id, err
}

func DeleteUser(id int) (int, error) {
	log.Print("going to delete record in database for userid : ", id)
	stmt, err := db.Prepare("DELETE from user where id=?")
	if err != nil {
		log.Print("error preparing query :: ", err)
		return 0, err
	}
	if _, err := stmt.Exec(id); err != nil {
		log.Print("error executing query :: ", err)
		return 0, err
	}
	return id, err
}
