/*
Конфигурация проекта через .env
*/

package _config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

// configEnv - Конфигурация вашего проекта
//
type configEnv struct {
	DB_PSQL_URL string `env:"DB_PSQL_URL,required"`
}

// Сингл-тон для конфига
//
var (
	instance *configEnv
	once     sync.Once
)

// GetConfigEnv - Получить конфигурацию переменных среды
// Дабы не бегать за своим хвостом - выполнить обязательно в init()
//
func GetConfigEnv() *configEnv {
	once.Do(
		func() {

			instance = &configEnv{}

			//Грузим переменные среды из файла
			err := godotenv.Load()
			if err != nil {
				//Если файла нет, то мы нестроены через Кубер/Докер и прочую нечисть
				log.Println("SYSTEM SETTINGS ANOTHER")
			}

			//Заполняем ими конфигурацию
			if err := env.Parse(instance); err != nil {
				log.Fatal("CONFIG ERROR", err)
			}

			log.Println("SYSTEM SUCCESS CONFIGURE")
		},
	)

	return instance
}
