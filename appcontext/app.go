package appcontext

import (
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type appContext struct {
	DB *gorm.DB
}

var context *appContext

var (
	dbPath = "./toko-ijah.db"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbPath)
	checkError(err)

	return db, nil
}

//GetDB will return DB connection from appContext
func GetDB() *gorm.DB {
	return context.DB
}

/*InitContext will inject all dependency to the appContext struct,
so it will be available across application
*/
func InitContext() {
	db, _ := initDB()

	context = &appContext{
		DB: db,
	}
}
