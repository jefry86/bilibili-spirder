//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/15 3:40 下午

package tools

import (
	"testing"
)

func TestSearch(t *testing.T) {
	keyword := "FPV穿超机"
	pageSize := 20
	pageNo := 1
	data, nun, err := Search(keyword, pageNo, pageSize)
	if err != nil {
		t.Error(err)
	}
	t.Log(nun)
	t.Logf("%#v", data)
}

func TestParseVideoData(t *testing.T) {
	url := "https://www.bilibili.com/video/av839681814?from=search"
	data, err := ParseVideoData(url)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
