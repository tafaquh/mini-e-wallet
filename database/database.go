package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/tafaquh/mini-e-wallet/seed"
)

var (
	DBConn *gorm.DB
	User = os.Getenv("EWALLET_DB_USER")
	Password = os.Getenv("EWALLET_DB_PASSWORD")
	DB = os.Getenv("EWALLET_DB_NAME")
    dsn = fmt.Sprintf("%v:%v@/%v?parseTime=true", User, Password, DB)
)

func ConnectMySQL() (err error) {
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, //Cache to speed up query process
		SkipDefaultTransaction: true, //Disable default transaction for auto store created_at, updated_at, deleted_at
	})
	
	seed.Load(DBConn)
	
    if err != nil {
        return err
    }
 
    return nil
}