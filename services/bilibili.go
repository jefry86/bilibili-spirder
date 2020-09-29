//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/17 3:54 下午

package services

import (
	"bilibili-spirder/global"
	"bilibili-spirder/model/bilibilii"
	"bilibili-spirder/tools"
)

func SpiderVideo() {
	b := &biliBiliSpider{
		ra: tools.NewRate(10),
		ch: make(chan *[]tools.SearchData, 1),
	}
	go b.getList()
	go b.getVideo()
	go b.execVideo()
}

type biliBiliSpider struct {
	ra *tools.Rate
	ch chan *[]tools.SearchData
}

func (b *biliBiliSpider) getList() {
	keyword := "FPV穿越机"
	pageNo := 1
	pageSize := 20
	num, err := b.search(keyword, pageNo, pageSize)
	if err != nil {
		global.Logger.Error("获取列表报错，err:", err.Error())
	}
	for {
		select {
		case <-b.ra.Get():
			global.Logger.Info("get list获取速度，开始执行...")
			pageNo++
			if num < int64(pageNo) {
				return
			}
			_, err := b.search(keyword, pageNo, pageSize)
			if err != nil {
				global.Logger.Error("获取列表报错，err:", err.Error())
			}
		}
	}
}

func (b *biliBiliSpider) search(keyword string, pageNo, pageSize int) (int64, error) {
	list, num, err := tools.Search(keyword, pageNo, pageSize)
	if err != nil {
		global.Logger.Error("获取列表报错，err:", err.Error())
		return 0, err
	}
	b.ch <- &list
	return num, nil
}

func (b *biliBiliSpider) getVideo() {
	for {
		select {
		case list := <-b.ch:
			global.Logger.Info("get video获取速度，开始执行...")
			for _, l := range *list {

				bili := bilibilii.Bili{
					Title:       l.Title,
					Author:      l.Author,
					AuthorId:    l.Author,
					VideoPage:   l.Url,
					VideoUrl:    "",
					Tags:        l.Tag,
					Description: l.Description,
					Pic:         l.Pic,
				}
				if err := bili.Save(); err != nil {
					global.Logger.Error(err.Error())
				}

			}
		}
	}
}

func (b *biliBiliSpider) execVideo() {
	bili := new(bilibilii.Bili)
	num := 5
	ch := make(chan bool, num)

	for i := 0; i < num; i++ {
		ch <- true
	}

	for {
		list, err := bili.List(num)
		if err != nil {
			global.Logger.Error(err.Error())
			continue
		}
		for _, v := range *list {
			<-ch
			go func(info bilibilii.Bili) {
				videoData, err := tools.ParseVideoData(info.VideoPage)
				if err != nil {
					global.Logger.Error(err.Error())
					return
				}
				url, err := tools.VideoDo(videoData, info.VideoPage)
				if err != nil {
					global.Logger.Error(err.Error())
					return
				}
				upB := bilibilii.Bili{
					VideoUrl: url,
				}
				err = upB.Update(info.Id)
				if err != nil {
					global.Logger.Error(err.Error())
				}
				ch <- true
			}(v)
		}

	}

}
