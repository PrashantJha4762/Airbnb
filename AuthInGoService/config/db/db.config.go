package db

import (
	config "AuthInGoService/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)


//Here we basically made a fn to connect to the db.We made an object using mysql.NewConfig()
//using that we fetched all env variables
//Then used sql.open fn where we passed the name of the db and cfg.formatdsn()
func SetUpDB() {
	cfg := mysql.NewConfig()
	cfg.User=config.GetString("DB_USER","root");
	cfg.Passwd=config.GetString("DB_PWD","");
	cfg.Addr=config.GetString("DB_ADDr","127.0.0.1:3003")
	cfg.DBName=config.GetString("DB_NAME","AUTH_DEV")
	cfg.Net="tcp"
	fmt.Println("Connecting to the database name",cfg.DBName)
	db,err:=sql.Open("mysql",cfg.FormatDSN()) // this formatdsn takes all the above thing and convert it into
									//dsn string. read about it more
	if(err!=nil){
		fmt.Println("Some error occured");
	}
	fmt.Println("Connecting to the database")
	pingerr:=db.Ping()
	if(pingerr!=nil){
		fmt.Println("Not able to ping the db");
	}
	fmt.Println("Connected to the database sucessfully",cfg.DBName);
}