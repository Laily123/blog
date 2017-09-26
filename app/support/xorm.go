package support

import (
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/revel/config"
	"github.com/revel/revel"
)

var Xorm *xorm.Engine
var Isinstalled bool

// 初始化 Xorm
// 根据配置文件里的数据库类型调用相应的方法
func InitXorm(appConfig *config.Config) error {
	AppConfig = appConfig
	dbdriver, _ := AppConfig.String("database", "database.driver")
	switch dbdriver {
	case "mysql":
		return initMySQL()
	}
	return errors.New("no db driver")
}

// 初始化 mysql 数据库的方法
func initMySQL() error {
	dbname, _ := AppConfig.String("database", "database.dbname")
	user, _ := AppConfig.String("database", "database.user")
	passwd, _ := AppConfig.String("database", "database.password")
	host, _ := AppConfig.String("database", "database.host")
	port, _ := AppConfig.String("database", "database.port")
	prefix, _ := AppConfig.String("database", "database.prefix")

	return TestXorm("mysql", user, passwd, host, port, dbname, prefix)
}

// 测试数据库是否初始化成功
func TestXorm(driver, user, pass, host, port, dbname string, prefix string) error {
	var err error
	Xorm, err = xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, dbname))
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
	Xorm.SetTableMapper(tbMapper)
	Xorm.ShowSQL(true)
	if err != nil {
		return err
	}
	return Xorm.Ping()
}

func AddDB(dbhost, dbport, dbuser, dbpass, dbname, dbprefix, dbtype string) error {
	AppConfig.AddOption("database", "database.host", dbhost)
	AppConfig.AddOption("database", "database.port", dbport)
	AppConfig.AddOption("database", "database.user", dbuser)
	AppConfig.AddOption("database", "database.password", dbpass)
	AppConfig.AddOption("database", "database.dbname", dbname)
	AppConfig.AddOption("database", "database.prefix", dbprefix)
	AppConfig.AddOption("database", "database.driver", dbtype)
	return nil
}

func writeConfig() error {
	filepath := revel.BasePath + "/conf/speedy.conf"
	_, err := os.Open(filepath)
	if err != nil {
		os.Create(filepath)
	}

	err = AppConfig.WriteFile(filepath, 0775, "default config")
	if err != nil {
		return err
	}
	return nil
}

func FinishInstall() {
	err := writeConfig()
	if err != nil {
		revel.ERROR.Println("write config error: ", err)
	}
}
