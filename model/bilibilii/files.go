//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/15 12:26 下午

package bilibilii

import (
	"bilibili-spirder/global"
	"bilibili-spirder/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type CSVData struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	AuthorId    string `json:"author_id"`
	Video       string `json:"video"`
	HttpVideo   string `json:"http_video"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

type CSV struct {
	mx  sync.RWMutex
	w   *bufio.Writer
	c   chan bool
	len int
}

func NewCSV() (*CSV, error) {
	w, e := newBuf()
	if e != nil {
		return nil, e
	}
	csv := &CSV{
		w:   w,
		c:   make(chan bool),
		len: 0,
	}
	csv.AutoFlush()
	defer csv.Flush()
	return csv, nil
}

func newBuf() (*bufio.Writer, error) {
	fileName := fmt.Sprintf("%s/%s-%d%d%d.csv", utils.Config.Path.Csv, "csv", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		global.Logger.Error("打开文件失败，err:", err.Error())
		return nil, err
	}
	return bufio.NewWriter(f), nil
}

func (c *CSV) Write(csvData *CSVData) error {
	data := csvData.Title + "," +
		csvData.Author + "," +
		csvData.AuthorId + "," +
		csvData.Video + "," +
		csvData.HttpVideo + "," +
		csvData.Tag + "," +
		csvData.Description
	c.mx.Lock()
	n, err := c.w.WriteString(data)
	c.mx.Unlock()
	if err != nil {
		global.Logger.Error("写入失败,err:", err.Error())
		return err
	}
	c.len += n
	global.Logger.Info("写入成功,n:", strconv.FormatInt(int64(c.len), 10))
	if c.len >= 10*1024*4 {
		c.c <- true
	}
	return nil
}

func (c *CSV) Flush() error {
	c.mx.RLock()
	defer c.mx.RUnlock()
	if err := c.w.Flush(); err != nil {
		global.Logger.Error("flush error:", err.Error())
		return err
	} else {
		global.Logger.Info("flush success")
		c.len = 0
		return nil
	}
}

func (c *CSV) AutoFlush() {
	go func(c *CSV) {
		for {
			select {
			case <-c.c:
				global.Logger.Info("flush start...")
				c.Flush()
			}
		}
	}(c)

	go func(c *CSV) {
		//每30秒落地一次
		t := time.NewTicker(10 * time.Second)
		defer t.Stop()
		for {
			<-t.C
			global.Logger.Info("30秒开始flush")
			if c.len > 0 {
				c.c <- true
			}
		}
	}(c)
}
