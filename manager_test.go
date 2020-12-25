package mdnotify

import (
	"log"
	"testing"
)

func TestMdManager_AllMarkdownFile(t *testing.T){
	manager := MdManager{}.New("mdnotify/md_files")
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
	manager := MdManager{}.New("mdnotify/md_files")
	mdManager, e := manager.InitAllMarkdown()
	if e != nil {
		t.Fatal(e)
	}

	t.Log(mdManager)
}

func TestMdManager_ListMarkDown(t *testing.T) {
	manager := MdManager{}.New("mdnotify/md_files")
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

func TestMdManager_MarkdownContent(t *testing.T) {
	manager := MdManager{}.New("/home/yuechenxing/go/src/mdnotify/md_files")
	manager.InitAllMarkdown()
	t.Log(manager.SortIndex)

	content := manager.MarkdownContent("b968bc4af7d9aae06bccaabad60bf35a")
	t.Log(content)
}

func TestMdManager_ListByTag(t *testing.T) {
	manager := MdManager{}.New("/home/yuechenxing/go/src/mdnotify/md_files")
	manager.InitAllMarkdown()

	list := manager.ListByClass("类别1", 0, 10)
	t.Log(list)
}

