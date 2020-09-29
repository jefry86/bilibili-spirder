//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/28 5:59 下午

package bilibilii

import (
	"bilibili-spirder/global"
	"time"
)

type Bili struct {
	Id          int64
	Title       string    `xorm:"varchar(255)"`
	Author      string    `xorm:"varchar(255)"`
	AuthorId    string    `xorm:"varchar(255)"`
	VideoPage   string    `xorm:"varchar(255)"`
	VideoUrl    string    `xorm:"varchar(255)"`
	Pic         string    `xorm:"varchar(255)"`
	Tags        string    `xorm:"varchar(500)"`
	Description string    `xorm:"varchar(1024)"`
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
}

func (b *Bili) Save() error {
	_, err := global.Db.Insert(b)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bili) List(limit int) (*[]Bili, error) {
	rs := make([]Bili, 0)
	if err := global.Db.
		Select("id,video_page").
		Where("video_url=''").
		Limit(limit, 0).
		Asc("id").
		Find(&rs); err != nil {
		return nil, err
	}
	return &rs, nil

}

func (b *Bili) Update(id int64) error {
	_, err := global.Db.Where("id=?", id).Update(b)
	return err
}
