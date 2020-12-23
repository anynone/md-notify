package md_notify

import (
	"errors"
	"log"
	"os"
)

type MdManager struct {
	folder string
	Papers map[string]*MarkDown // 使用文件名md5作为key
	SortIndex []string // 排序号对应文件名md5,可能多个
	TitleIndex map[string]string // 标题倒排索引
	ClassIndex map[string]string // 类别倒排索引
	ExtraTagIndex map[string][]string // 其他标签的倒排索引
}

func (mdm MdManager) New(folder string) *MdManager{
	return &MdManager{
		folder:folder,
		Papers: map[string]*MarkDown{},
		SortIndex: []string{},
		TitleIndex: map[string]string{},
		ClassIndex: map[string]string{},
		ExtraTagIndex: map[string][]string{}}
}

func (mdm *MdManager) acceptMdFile(finfo os.FileInfo) error{
	log.Println(finfo.Name())
	file, e := os.Open(mdm.folder + string(os.PathSeparator) + finfo.Name())
	if e != nil {
		return e
	}
	one := &MarkDown{}
	md := one.Load(file)
	mdm.Papers[md.Id] = md

	// 排序索引更新
	if md.Content.Sort >= len(mdm.SortIndex) {
		mdm.SortIndex = append(mdm.SortIndex, md.Id)
	}else{
		if md.Content.Sort == 0 {
			mdm.SortIndex =  append([]string{md.Id}, mdm.SortIndex...)
		}else{

			inserted := false
			for index, mdownId := range mdm.SortIndex[md.Content.Sort:] {
				if mdm.Papers[mdownId].Content.Sort > md.Content.Sort || (mdm.Papers[mdownId].Content.Sort == md.Content.Sort && mdm.Papers[mdownId].LastUpdate < md.LastUpdate)  {
					mdm.SortIndex = append(mdm.SortIndex[0:md.Content.Sort + index], append([]string{md.Id}, mdm.SortIndex[md.Content.Sort + index:]...)...)
					inserted = true
					break
				}
			}

			if !inserted {
				mdm.SortIndex = append(mdm.SortIndex, md.Id)
			}
		}


	}

	return nil
}

func (mdm * MdManager) InitAllMarkdown() (*MdManager, error){

	infos, e := mdm.AllMarkdownFile()
	if e != nil {
		return nil, e
	}

	for _, value := range infos {
		e := mdm.acceptMdFile(value)
		if e != nil {
			return nil, e
		}
	}

	return mdm,nil

}

func (mdm *MdManager) ListMarkDown()[]MarkDown{
	mds := []MarkDown{}
	for _, value := range mdm.Papers {
		mds = append(mds, *value)
	}

	return mds
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



