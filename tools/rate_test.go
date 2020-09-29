//@Description todo
//@Author 凌云  jefry52@gmail.com
//@DateTime 2020/9/18 4:32 下午

package tools

import (
	"testing"
)

func TestNewRate(t *testing.T) {
	rate := NewRate(10)
	for {
		rate.Get()
		t.Log("get")
	}
}
