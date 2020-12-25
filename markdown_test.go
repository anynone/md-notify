package mdnotify

import (
	"io/ioutil"
	"testing"
)

func TestMarkDown_ParseContent(t *testing.T) {
	bytes, _ := ioutil.ReadFile("md_files/test1.md")
	down := MarkDown{}
	content := down.ParseContent(bytes)
	t.Log(content, content.Extra)
}
