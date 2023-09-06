package DataBases

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/Golang")
	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		fmt.Println("Error closing the database connection:", err)
	}
}

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/Golang")
	if err != nil {
		panic(err)
	}

	return db
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {

	}
}

type UserData struct {
	ID          int64     `json:"IdUser"`
	DateOfBirth time.Time `json:"DateOfBirth"`
	ZodiacSign  string    `json:"ZodiacSign"`
	Position    int       `json:"Position"`
}

func GetUserData(Id int64) *UserData {

	var user UserData

	res, err := db.Query(fmt.Sprintf("SELECT `IdUser`,`Position` FROM `UsersTelega` WHERE `IdUser` = '%d'", Id))

	if err != nil {
		panic(err)
	}
	for res.Next() {

		err = res.Scan(&user.ID, &user.Position)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("User : %s with age %d", user.Name, user.Age))
	}
	return &user

}

func GetUserZodiac(Id int64) string {

	var user UserData

	res, err := db.Query(fmt.Sprintf("SELECT `IdUser`,`ZodiacSign` FROM `UsersTelega` WHERE `IdUser` = '%d'", Id))

	if err != nil {
		panic(err)
	}
	for res.Next() {

		err = res.Scan(&user.ID, &user.ZodiacSign)
		if err != nil {
			//panic(err)
			return ""
		}
		//fmt.Println(fmt.Sprintf("User : %s with age %d", user.Name, user.Age))
	}
	return user.ZodiacSign

}

func AddUser(Id int64, date time.Time, zodiac int) {

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `UsersTelega`(`IdUser`,`DateOfBirth`,`ZodiacSign`) "+
		"VALUES ('%d','%d','%d')", Id, date, zodiac))
	if err != nil {
		panic(err)
	}
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {

		}
	}(insert)

}

func AddUserId(Id int64) {

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `UsersTelega`(`IdUser`,`Position`) "+
		"VALUES ('%d','%d')", Id, 0))
	if err != nil {
		panic(err)
	}
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {

		}
	}(insert)

}

func UpdateZodiacSign(Id int64, ZodiacSign string) {

	update, err := db.Query(fmt.Sprintf("UPDATE `UsersTelega` SET `ZodiacSign` = '%s' WHERE `IdUser` = '%d' ", ZodiacSign, Id))
	if err != nil {
		panic(err)
	}
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {

		}
	}(update)

}

func UpdatePosition(Id int64, position int) {

	update, err := db.Query(fmt.Sprintf("UPDATE `UsersTelega` SET `Position` = '%d' WHERE `IdUser` = '%d' ", position, Id))
	if err != nil {
		panic(err)
	}
	defer func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {

		}
	}(update)

}

func Gethoroscope(ZodiacSign string) string {

	Horoscope := ""

	res, err := db.Query(fmt.Sprintf("SELECT `Horoscope` FROM `Horoscope` WHERE `ZodiacSign` = '%s'", ZodiacSign))

	if err != nil {
		panic(err)
	}
	for res.Next() {

		err = res.Scan(&Horoscope)
		if err != nil {
			//panic(err)
			return ""
		}

	}
	return Horoscope

}
