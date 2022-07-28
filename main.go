/*
Название программы
*/
package main

import (
	"fmt"
	"project_1/_config"
	"project_1/db"
)

func init() {
	// Гарантирует правильную инициализацию переменных среды
	_config.GetConfigEnv()
}

func main() {
	cfg := _config.GetConfigEnv()

	db.SomeDBTest(cfg.DB_PSQL_URL)
	fmt.Println(cfg.DB_PSQL_URL)
}
