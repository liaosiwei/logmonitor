package shcmd

import (
	"fmt"
	"testing"
	"time"
)

func TestRunWithin(t *testing.T) {
	out, err := RunWithin("sh /Users/baidu/gospace/gowork/src/baidu.com/liaosiwei/shcmd/test.sh", 1*time.Second)
	if err != nil {
		t.Error("cmd failed")
	}
	fmt.Println(out.String())
}
