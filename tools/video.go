//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/16 5:22 下午

package tools

import (
	"bilibili-spirder/global"
	"bilibili-spirder/utils"
	"bufio"
	"errors"
	"fmt"
	"github.com/guonaihong/gout"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type VideoData struct {
	Video string `json:"video"`
	Audio string `json:"audio"`
}

func httpDo(uri, ranges, referer string) (string, error) {
	global.Logger.Info("url:", uri)
	var httpCode int
	var respHeader string
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	files := strings.Split(u.Path, "/")
	file := files[len(files)-1]

	dir := fmt.Sprintf("%d%d%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	dir = fmt.Sprintf("%s%s%s", utils.Config.Path.Video, string(os.PathSeparator), dir)

	file = fmt.Sprintf("%s%s%s", dir, string(os.PathSeparator), file)

	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(dir, os.ModePerm); err != nil {
				global.Logger.Error("创建目录失败，err:", err.Error())
				return "", err
			}
		}
	}

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()

	err = gout.GET(uri).
		Debug(utils.Config.Http.Debug).
		SetHeader(gout.H{
			"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36",
			"accept":          "*/*",
			"origin":          "https://www.bilibili.com",
			"sec-fetch-site":  "cross-site",
			"sec-fetch-mode":  "cors",
			"sec-fetch-dest":  "empty",
			"referer":         referer,
			"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7",
		}).
		//SetCookies(
		//	&http.Cookie{
		//		Name:  "_uuid",
		//		Value: "634AB149-0A9D-C0AA-7308-E0BBACA9784859372infoc",
		//	},
		//	&http.Cookie{
		//		Name:  "buvid3",
		//		Value: "3AB718B2-AFAD-4342-AE81-0D99843F32A7143106infoc",
		//	},
		//	&http.Cookie{
		//		Name:  "CURRENT_FNVAL",
		//		Value: "80",
		//	},
		//	&http.Cookie{
		//		Name:  "blackside_state",
		//		Value: "1",
		//	},
		//	&http.Cookie{
		//		Name:  "sid",
		//		Value: "9ios624d",
		//	},
		//	&http.Cookie{
		//		Name:  "rpdid",
		//		Value: "|(u|JR)u~RmR0J'uY||kJmYRu",
		//	},
		//	&http.Cookie{
		//		Name:  "bfe_id",
		//		Value: "6f285c892d9d3c1f8f020adad8bed553",
		//	},
		//).
		BindBody(w).
		BindHeader(&respHeader).
		SetTimeout(1 * time.Hour).
		Code(&httpCode).
		Do()
	global.Logger.Info("down success,file:", file)
	if err != nil {
		return "", err
	}
	if httpCode != 200 {
		return "", errors.New("http code：" + strconv.Itoa(httpCode))
	}

	return file, nil
}

func downloadVideo(video VideoData, referer string) (*VideoData, error) {
	global.Logger.Info("开始下载中...")
	videoFile, err := httpDo(video.Video, "0-10", referer)
	if err != nil {
		global.Logger.Error("下载video失败，err:", err.Error())
		return nil, err
	}
	global.Logger.Infof("视频:%s下载成功，音频开始下载中....", videoFile)
	audioFile, err := httpDo(video.Audio, "0-10", referer)
	if err != nil {
		return nil, err
	}
	global.Logger.Infof("音频:%s下载成功", audioFile)

	return &VideoData{
		Video: videoFile,
		Audio: audioFile,
	}, nil
}

func syntheticVideo(videoData *VideoData) (string, error) {
	/*
	 ffmpeg合并音视频命令：
	 ffmpeg -i 233373153-1-30064.m4s -i 233373153-1-30280.m4s -c:v copy -c:a aac -strict experimental 233373153.mp4
	*/
	videoFiles := strings.Split(videoData.Video, "/")
	file := strings.Split(videoFiles[len(videoFiles)-1], "-")[0]
	mp4 := fmt.Sprintf("%s%s%s%s%s.mp4", utils.Config.Path.Video, string(os.PathSeparator), videoFiles[len(videoFiles)-2], string(os.PathSeparator), file)
	shell := fmt.Sprintf("ffmpeg -i %s -i %s -c:v copy -c:a aac -strict experimental %s", videoData.Video, videoData.Audio, mp4)
	global.Logger.Infof("执行命令：%v", shell)
	cmd1 := exec.Command("/bin/bash", "-c", shell)
	err := cmd1.Run()
	if err != nil {
		return "", err
	}
	return mp4, nil

}

func VideoDo(video VideoData, referer string) (string, error) {
	global.Logger.Infof("开始下载 %s的 video",referer)
	videoData, err := downloadVideo(video, referer)
	if err != nil {
		return "", err
	}
	global.Logger.Infof("文件：%#v 下载成功，开始合并操作", *videoData)
	mp4, err := syntheticVideo(videoData)
	global.Logger.Infof("文件：%#v 合并成功", *videoData)
	if err != nil {
		return "", nil
	}

	return mp4, nil

}
