package md_notify

import (
	"io/ioutil"
	"testing"
)

func TestMarkDown_ParseContent(t *testing.T) {
	bytes, _ := ioutil.ReadFile("/home/yuechenxing/go/src/md-notify/md_files/test1.md")
	down := MarkDown{}
	content := down.ParseContent(bytes)
	t.Log(content, content.Extra)
}
