# Wrapper postgres

## Описание
Данный модуль предназначается для взаимодействия с базой данных postgresql.  
Модуль является оберткой над [jackc/pgx](github.com/jackc/pgx)
## Пример использования
Использование через конструктор.
``` go
package main

import (
	"log"

	"github.com/requiemofthesouls/postgres"
)

func main() {
	var (
		db postgres.Wrapper
		err     error
	)
	if db, err = postgres.New(context.Background(), postgres.Config{
		Host:               "localhost",
		Port:               5432,
		Username:           "test",
		Password:           "test",
		Database:           "test",
	}); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
```
Использование через definitions.
``` go
package config

import (
	"log"

	"github.com/requiemofthesouls/container"
	"github.com/requiemofthesouls/postgres"
	pgCont "github.com/requiemofthesouls/postgres/container"
)

func main() {
	var db postgres.Wrapper
	if err := container.Container.Fill(pgCont.DIWrapper, &db); err != nil {
		log.Fatal(err)
	}
	
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
```
## Пример конфигурации
``` yaml
postgres:
  # Host postgres сервера
  host: 127.0.0.1
  # Port postgres сервера
  port: 5432
  # Имя пользователя postgres сервера
  username: myuser
  # Пароль postgres сервера
  password: mypass
  # Имя базы данных postgres сервера
  database: mydb
  # Максимальное кол-во открытых соединений (по умолчанию 4)
  maxConns: 3
  # Максимальное время жизни подключения (по умолчанию 90 секунд)
  maxConnLifetimeSec: 30
  # Максимальное время жизни неиспользуемых соединений (по умолчанию 10 секунд)
  maxConnIdleTimeSec: 1
  
```
## Зависимости от модулей
- [config](https://github.com/requiemofthesouls/config/-/blob/main/README.md)
- [container](https://github.com/requiemofthesouls/container/-/blob/main/README.md)
