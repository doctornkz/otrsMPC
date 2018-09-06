package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type userPreferences struct {
	user_id           int32
	preferences_key   string
	preferences_value []byte
}

func check(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}
}

type dbConf struct {
	db     string
	dbuser string
	dbpwd  string
	dbhost string
	dbport string
}

var (
	conf dbConf
	DSN  string
)

func init() {

	// db, dbuser, dbpwd, dbhost, dbport,
	db := flag.String("db", "default", "Database to read")
	dbuser := flag.String("dbuser", "root", "Mysql user")
	dbpwd := flag.String("dbpwd", "change_me", "DBuser password")
	dbhost := flag.String("dbhost", "localhost", "Mysql host")
	dbport := flag.String("dbport", "3306", "Mysql port")

	flag.Parse()

	conf.db = *db
	conf.dbhost = *dbhost
	conf.dbport = *dbport
	conf.dbpwd = *dbpwd
	conf.dbuser = *dbuser
	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *dbuser, *dbpwd, *dbhost, *dbport, *db)
	log.Printf(DSN)
}

func main() {
	db, err := sql.Open("mysql", DSN)
	check(err)
	defer db.Close()

	rows, err := db.Query("select * from user_preferences")
	check(err)

	ups := make([]*userPreferences, 0)

	for rows.Next() {
		log.Printf("!")
		up := new(userPreferences)
		err := rows.Scan(&up.user_id, &up.preferences_key, &up.preferences_value)
		if err != nil {
			log.Fatal(err)
		}
		ups = append(ups, up)
	}

	log.Printf("Result: ,%v", ups)

}
