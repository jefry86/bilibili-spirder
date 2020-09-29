//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/16 5:39 下午

package tools

import (
	"bilibili-spirder/utils"
	"testing"
)

func TestDownloadVideo(t *testing.T) {
	utils.Config.Path.Video = "../temp/video"
	utils.Config.Path.Logs = "../temp/logs"
	videoData := VideoData{
		Video: "http://upos-sz-mirrorhw.bilivideo.com/upgcxcode/85/68/65666885/65666885-1-30064.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&uipk=5&nbs=1&deadline=1600435516&gen=playurl&os=hwbv&oi=1034817494&trid=32853c85fdf14868b9ef0b7c85928e11u&platform=pc&upsig=91ee52a8c78ad58e267facd5a9110e33&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,platform&mid=0&orderid=0,3&agrr=0&logo=80000000",
		Audio: "http://upos-sz-mirrorkodo.bilivideo.com/upgcxcode/85/68/65666885/65666885-1-30280.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&uipk=5&nbs=1&deadline=1600435516&gen=playurl&os=kodobv&oi=1034817494&trid=32853c85fdf14868b9ef0b7c85928e11u&platform=pc&upsig=3d2f6f966cc95304dd840c378c4f1fd6&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,platform&mid=0&orderid=0,3&agrr=0&logo=80000000",
	}

	file, err := VideoDo(videoData, "")
	if err != nil {
		t.Error(err)
	}
	t.Log(file)

}
