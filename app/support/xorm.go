package support

import (
	"fmt"
	"github.com/alecthomas/log4go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/revel/config"
	"github.com/revel/revel"
)

var Xorm *xorm.Engine

func InitXorm() {

	file := (revel.BasePath + "/conf/speedy.conf")
	data, _ := config.ReadDefault(file)

	driver, _ := data.String("database", "database.driver")
	dbname, _ := data.String("database", "database.dbname")
	user, _ := data.String("database", "database.user")
	passwd, _ := data.String("database", "database.password")
	host, _ := data.String("database", "database.read.host")
	prefix, _ := data.String("database", "database.prefix")

	var err error
	Xorm, err = xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, passwd, host, dbname))
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
	Xorm.SetTableMapper(tbMapper)

	if err != nil {
		log4go.Error(err)
	} else {
		Xorm.Ping()
	}
}