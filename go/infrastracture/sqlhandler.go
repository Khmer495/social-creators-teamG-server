package infrastracture

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"

	"github.com/Khmer495/social-creators-teamG-server/go/entity/model"
	"github.com/Khmer495/social-creators-teamG-server/go/interface/database"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = dbConnect()
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

func (handler *SqlHandler) Set(name string, value interface{}) *gorm.DB {
	return handler.Conn.Set(name, value)
}

func (handler *SqlHandler) Query() *gorm.DB {
	return handler.Conn
}

func dbRetryConnect(count int) *gorm.DB {
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
		return dbRetryConnect(count)
	}
	return conn
}

func dbConnect() *gorm.DB {
	conn := dbRetryConnect(0)

	conn.LogMode(true)

	return conn
}

func dbMigrate(conn *gorm.DB) {
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

	log.Println("Start inserting data")
	for _, Prefecture := range Prefectures {
		conn.Create(&Prefecture)
	}
	for _, City := range Cities {
		conn.Create(&City)
	}
	for _, User := range Users {
		conn.Create(&User)
	}
	for _, Restaurant := range Restaurants {
		conn.Create(&Restaurant)
	}
	for _, Comment := range Comments {
		conn.Create(&Comment)
	}
	for _, Good := range Goods {
		conn.Create(&Good)
	}
	for _, Favorite := range Favorites {
		conn.Create(&Favorite)
	}
	log.Println("Finish inserting data")
}

func InitDBServer() {
	conn := dbConnect()
	if !conn.HasTable(&model.User{}) {
		dbMigrate(conn)
	}
}

func dbDrop(conn *gorm.DB) {
	log.Println("Start drop")
	conn.DropTable(&model.Prefecture{})
	conn.DropTable(&model.City{})
	conn.DropTable(&model.User{})
	conn.DropTable(&model.Restaurant{})
	conn.DropTable(&model.Comment{})
	conn.DropTable(&model.Good{})
	conn.DropTable(&model.Favorite{})
	log.Println("Finish drop")
}

func ResetDB(c echo.Context) error {
	conn := dbConnect()
	dbDrop(conn)
	dbMigrate(conn)
	return c.String(http.StatusOK, "Reset Complete!")
}
