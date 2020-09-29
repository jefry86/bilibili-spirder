//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/28 5:43 下午

package model

import (
	"bilibili-spirder/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func NewDb() *xorm.Engine {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", utils.Config.Db.Username, utils.Config.Db.Password, utils.Config.Db.Host, utils.Config.Db.Port, utils.Config.Db.Name)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	return engine
}
