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

func TestMdManager_InitAllMarkdown(t *testing.T) {
	manager := MdManager{}.New("/home/yuechenxing/go/src/md-notify/md_files")
	mdManager, e := manager.InitAllMarkdown()
	if e != nil {
		t.Fatal(e)
	}

	t.Log(mdManager)
}

func TestMdManager_ListMarkDown(t *testing.T) {
	manager := MdManager{}.New("/home/yuechenxing/go/src/md-notify/md_files")
	mdManager, e := manager.InitAllMarkdown()
	if e != nil {
		t.Fatal(e)
	}

	for _, mdId := range mdManager.SortIndex {
		t.Log(mdManager.Papers[mdId].FileName, mdManager.Papers[mdId].Content.Sort)
	}
	//down := mdManager.ListMarkDown()
	//t.Log(down)

	//t.Log(string(manager.Papers["12d21a21adb7312b147715b35d8f8b44"].MarkdownContent()))
}

