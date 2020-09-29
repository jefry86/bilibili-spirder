//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/17 5:52 下午

package bilibilii

import (
	"testing"
	"time"
)

func TestNewCSV(t *testing.T) {
	csv, err := NewCSV()
	if err != nil {
		t.Error(err)
	}
	for i := 0; i <= 1000; i++ {
		data := &CSVData{
			Title:       "test",
			AuthorId:    "1",
			Author:      "au",
			Video:       "./temp/video/11111.mp4",
			HttpVideo:   "http://upos-sz-mirrorcos.bilivideo.com/upgcxcode/53/31/233373153/233373153-1-30080.m4s",
			Tag:         "野生技术协会,野生技术协会,106.1万投稿,574精选视频,进入频道,订阅 15.6万,知识野生技术协会,第一视角,学习,科技,设备,穿戴相机,FPV,穿戴设备,运动相机,insta360",
			Description: "www.youtube.com\n【福利】Driftty-FPV为大家争取到insta 360官方优惠卷，私信我即可领取，享最低价购买Insta 360 Go相机\nInsta360 Go全新固件新增FPV穿越机模式，本视频为具体设置教程，使用Go拇指相机安装在类似75mm室内穿越机或者3寸牙签机、涵道机上，在确保通过性的同时获取最高画质，的确是FPV拍摄的理想选择。",
		}
		if err = csv.Write(data); err != nil {
			t.Error(err)
		}
		time.Sleep(1 * time.Second)
	}
	if err = csv.Flush(); err != nil {
		t.Error(err)
	}
	t.Log("success")
}
