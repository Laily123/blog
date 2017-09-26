package models

import (
	"blog/app/support"
	"testing"
)

func init() {
	initDB()
}

func initDB() {
	host := "laily.net"
	port := "3306"
	user := "Laily"
	password := "CLARK0618"
	dbname := "test"
	prefix := "t_"
	driver := "mysql"
	support.TestXorm(driver, user, password, host, port, dbname, prefix)
}

func TestSeriesModel_Create(t *testing.T) {
	var s SeriesModel
	s.Create("first", "1")
}

func TestSeriesBlogModel_AddBlogForSeries(t *testing.T) {
	var s SeriesBlogModel
	s.AddBlogForSeries(1, []int64{23, 45})
}
