package md_notify

import (
	"errors"
	"log"
	"os"
)

type MdManager struct {
	folder string
	Papers map[string]*MarkDown // 使用文件名md5作为key
	SortIndex map[int][]string // 排序号对应文件名md5,可能多个
	TitleIndex map[string]string // 标题倒排索引
	ClassIndex map[string]string // 类别倒排索引
	ExtraTagIndex map[string][]string // 其他标签的倒排索引
}

func (mdm MdManager) New(folder string) *MdManager{
	return &MdManager{
		folder:folder,
		Papers: map[string]*MarkDown{},
		SortIndex: map[int][]string{},
		TitleIndex: map[string]string{},
		ClassIndex: map[string]string{},
		ExtraTagIndex: map[string][]string{}}
}

func (mdm *MdManager) acceptMdFile(finfo os.FileInfo) error{
	file, e := os.Open(mdm.folder + string(os.PathSeparator) + finfo.Name())
	if e != nil {
		return e
	}
	one := &MarkDown{}
	md := one.Load(file)

	if _, ok := mdm.Papers[md.Id]; !ok {
		mdm.Papers[md.Id] = md
	}else{
		log.Println("已经存在,覆盖操作todo")
	}


	return nil
}

func (mdm * MdManager) InitAllMarkdown() (*MdManager, error){

	//allFileInfo, e := md.AllMarkdownFile()
	//if e != nil {
	//	return nil, e
	//}
	//
	//for _, finfo := range allFileInfo {
	//
	//}


	return nil,nil

}

func (mdm *MdManager) AllMarkdownFile() ([]os.FileInfo, error){
	file, e := os.Open(mdm.folder)
	if e != nil {
		return []os.FileInfo{}, e
	}

	fileInfo, e := file.Stat()
	if e != nil {
		return []os.FileInfo{}, e
	}
	if ! fileInfo.IsDir() {
		return []os.FileInfo{}, errors.New(mdm.folder + "不是一个文件夹")
	}

	fileInfos, e := file.Readdir(0)
	if e != nil {
		return []os.FileInfo{}, e
	}

	return fileInfos, nil
}



