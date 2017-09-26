package models

import (
	"blog/app/support"
	log "github.com/sirupsen/logrus"
)

// 系列表
// 将一批关联的博文作为一个系列合集
type SeriesModel struct {
	Id    int64  `xorm:"not null pk autoincr INT(11)"`
	Name  string `xorm:"VARCHAR(250)"`
	Alias string `xorm:"VARCHAR(250)"`
	Count int    `xorm:"INT(11)"`
}

func (s *SeriesModel) TableName() string {
	return "t_series"
}

// Create 创建一个系列
func (s *SeriesModel) Create(name string, alias string) {
	ss := &SeriesModel{}
	ss.Name = name
	ss.Alias = alias
	_, err := support.Xorm.Insert(ss)
	if err != nil {
		log.Errorf("insert series error:%v, series: %+v \n", err, ss)
	}
}

// 专题和博客映射表
type SeriesBlogModel struct {
	Id       int64 `xorm:"not null pk autoincr INT(11)"`
	SeriesID int64 `xorm:"series_id INT(11)"`
	BlogID   int64 `xorm:"blog_id INT(11)"`
}

func (s *SeriesBlogModel) TableName() string {
	return "t_series_blog"
}

// 为系列添加新的博文
func (s *SeriesBlogModel) AddBlogForSeries(seriesID int64, blogIDs []int64) {
	sbs := make([]*SeriesBlogModel, 0)
	for _, id := range blogIDs {
		ss := &SeriesBlogModel{SeriesID: seriesID, BlogID: id}
		sbs = append(sbs, ss)
	}
	support.Xorm.Insert(&sbs)
}
