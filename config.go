package main

import (
	"fmt"
	"os"
	"strings"
)

const MysqlHost = "MYSQL_HOST"
const MysqlPort = "MYSQL_PORT"
const MysqlUser = "MYSQL_USER"
const MysqlPassword = "MYSQL_PASSWORD"
const MysqlDB = "MYSQL_DB"

var Config = map[string]string{
	MysqlHost:     "localhost",
	MysqlPort:     "3306",
	MysqlUser:     "demo",
	MysqlPassword: "demo123",
	MysqlDB:       "demo",
	"SERVER_PORT": "8080",
}

func LoadEnv() {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
		_, ok := Config[pair[0]]
		if ok {
			Config[pair[0]] = pair[1]
		}
	}
}

func GetMysqlConnectionURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Config[MysqlUser],
		Config[MysqlPassword], Config[MysqlHost], Config[MysqlPort], Config[MysqlDB])
}
