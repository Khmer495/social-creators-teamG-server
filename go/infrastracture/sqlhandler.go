package infrastracture

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
	"github.com/Khmer495/social-creators-teamG-server/go/interface/database"
)

var count = 0

func DBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")"
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	conn, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Println("DB not ready. Retconnecting...")
		time.Sleep(1 * time.Second)
		count++
		log.Println("Try count: ", count)
		if count > 30 {
			log.Println("Cannot connect to DB.")
			panic(err.Error())
		}
		return DBConnect()
	}

	conn.LogMode(true)

	return conn
}

func DBMigrate(conn *gorm.DB) {
	log.Println("Start auto migration")
	conn.AutoMigrate(
		&model.Prefecture{},
		&model.City{},
		&model.User{},
		&model.Restaurant{},
		&model.Comment{},
		&model.Good{},
		&model.Favorite{},
	)
	log.Println("Finish auto migration")
}

func InitDBServer() {
	conn := DBConnect()
	DBMigrate(conn)
}

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = DBConnect()
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}

func (handler *SqlHandler) Query() *gorm.DB {
	return handler.Conn
}
