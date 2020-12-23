package md_notify

import (
	"log"
	"testing"
)

func TestMdManager_AllMarkdownFile(t *testing.T){
	manager := MdManager{}.New("/home/yuechenxing/go/src/md-notify/md_files")
	infos, e := manager.AllMarkdownFile()
	if e != nil {
		t.Fatal(e)
	}

	for _, value := range infos {
		e := manager.acceptMdFile(value)
		if e != nil {
			t.Fatal(e)
		}
	}

	log.Println(manager.Papers)
}

