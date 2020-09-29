//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/15 12:28 下午

package tools

import (
	"bilibili-spirder/utils"
	"errors"
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/tidwall/gjson"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

type SearchData struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Tag         string `json:"tag"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Pic         string `json:"pic"`
}

//搜索列表抓取
//https://api.bilibili.com/x/web-interface/search/type?context=&page=50&order=&keyword=FPV%20%E7%A9%BF%E8%B6%8A%E6%9C%BA&duration=&tids_1=&tids_2=&__refresh__=true&_extra=&search_type=video&highlight=1&single_column=0
func Search(keyword string, pageNo, pageSize int) ([]SearchData, int64, error) {
	var data []SearchData
	uri := fmt.Sprintf("https://api.bilibili.com/x/web-interface/search/type?context=&page=%d&order=&keyword=%s&duration=&tids_1=&tids_2=&__refresh__=true&_extra=&search_type=video&highlight=1&single_column=0", pageNo, url.QueryEscape(keyword))
	body, err := httpGetDo(uri, gout.H{})
	if err != nil {
		return data, 0, err
	}
	code := gjson.Get(body, "code").Int()
	if code != 0 {
		return data, 0, errors.New("code不为空")
	}
	numPages := gjson.Get(body, "data.numResults").Int()
	if numPages <= 0 {
		return data, 0, errors.New("列表数据集为0")
	}
	result := gjson.Get(body, "data.result").Array()

	data, err = parseSearchData(result)
	if err != nil {
		return data, 0, err
	}

	return data, numPages, nil
}

//解析搜索列表数据，获取视频播放页面地址
func parseSearchData(result []gjson.Result) ([]SearchData, error) {
	d := make([]SearchData, 0)
	for _, vd := range result {
		d = append(d, SearchData{
			Title:       vd.Get("title").String(),
			Author:      vd.Get("author").Str,
			Url:         vd.Get("arcurl").Str,
			Tag:         vd.Get("tag").Str,
			Description: vd.Get("description").Str,
			Pic:         fmt.Sprintf("http:%s", vd.Get("pic").Str),
		})
	}
	return d, nil
}

//执行http请求
func httpGetDo(url string, query gout.H) (string, error) {
	var body string
	httpCode := 0

	err := gout.GET(url).
		Debug(utils.Config.Http.Debug).
		SetQuery(query).
		SetHeader(gout.H{
			"authority":                 "api.bilibili.com",
			"upgrade-insecure-requests": 1,
			"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36",
			"sec-fetch-site":            "none",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-dest":            "document",
			"accept-language":           "accept-language",
			//"referer":                   "https://search.bilibili.com/all?keyword=FPV%E7%A9%BF%E8%B6%8A%E6%9C%BA",
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
		BindBody(&body).
		SetTimeout(time.Duration(utils.Config.Http.Timeout) * time.Second).
		Code(&httpCode).
		Do()
	if err != nil {
		return "", err
	}
	if httpCode != 200 {
		return "", errors.New("http code:" + strconv.FormatInt(int64(httpCode), 10))
	}
	return body, nil
}

//获取视频页面的视频地址
func ParseVideoData(url string) (VideoData, error) {
	var video VideoData
	url = fmt.Sprintf("%s?from=search", url)
	//global.Logger.Info("开始解析：", url)
	body, err := httpGetDo(url, nil)
	if err != nil {
		return video, err
	}
	pattern := `<script>window.__playinfo__=(?s:(.+?))</script>`
	re := regexp.MustCompile(pattern)
	d := re.FindStringSubmatch(body)
	if d == nil {
		return video, errors.New("解析内容失败~")
	}
	videos := gjson.Get(d[1], "data.dash.video").Array()
	for _, id := range []int64{64, 32, 16} {
		if uri := getVideoPath(videos, id); uri != "" {
			video.Video = uri
			break
		}
	}

	audios := gjson.Get(d[1], "data.dash.audio").Array()
	for _, id := range []int64{30280, 30216, 30232} {
		if uri := getAudioPath(audios, id); uri != "" {
			video.Audio = uri
		}
	}
	if video.Video == "" || video.Audio == "" {
		return video, errors.New("页面：" + url + " 的video或audio未获取成功")
	}
	return video, nil
}

func getVideoPath(videos []gjson.Result, id int64) string {
	var uri string
	for _, v := range videos {
		if v.Get("codecid").Int() == 7 && v.Get("id").Int() == id {
			uri = v.Get("base_url").Str
		}
	}
	return uri
}

func getAudioPath(audios []gjson.Result, id int64) string {
	var uri string
	for _, v := range audios {
		if v.Get("id").Int() == id {
			uri = v.Get("base_url").Str
		}
	}
	return uri
}
