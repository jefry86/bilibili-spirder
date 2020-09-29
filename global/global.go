//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/16 8:37 下午

package global

import (
	"bilibili-spirder/model"
	"bilibili-spirder/utils"
	"xorm.io/xorm"
)

var Logger utils.Logger
var Db *xorm.Engine

func NewGlobal() {
	Logger = utils.NewLogger()
	Db = model.NewDb()
}
