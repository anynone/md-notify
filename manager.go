package md_notify

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type MdManager struct {
	folder string
	Papers map[string]*MarkDown // 使用文件名md5作为key
	SortIndex []string // 排序号对应文件名md5,可能多个
	TitleIndex map[string]string // 标题倒排索引
	ClassIndex map[string]map[string]uint8 // 类别倒排索引
	//ExtraTagIndex map[string][string]string // 其他标签的倒排索引
}

func (mdm MdManager) New(folder string) *MdManager{
	return &MdManager{
		folder:folder,
		Papers: map[string]*MarkDown{},
		SortIndex: []string{},
		TitleIndex: map[string]string{},
		ClassIndex: map[string]map[string]uint8{},
		//ExtraTagIndex: map[string][]string{}
	}
}

func (mdm *MdManager) acceptMdFile(finfo os.FileInfo) error{
	//log.Println(finfo.Name())
	file, e := os.Open(mdm.folder + string(os.PathSeparator) + finfo.Name())
	if e != nil {
		return e
	}
	one := &MarkDown{}
	md := one.Load(file)

	_, ok := mdm.Papers[md.Id]
	if ok {
		// 清除和当前加入文件的信息
		mdm.clearMdInfo(md)
	}

	mdm.Papers[md.Id] = md
	if _, ok := mdm.ClassIndex[md.Content.Class]; !ok {
		mdm.ClassIndex[md.Content.Class] = map[string]uint8{}
	}
	mdm.ClassIndex[md.Content.Class][md.Id] = 0

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

func (mdm *MdManager) MarkdownContent(id string) string {
	md, ok := mdm.Papers[id]
	if !ok {
		return ""
	}

	bytes, _ := ioutil.ReadFile(md.FileName)

	if bytes == nil {
		bytes = []byte{}
	}


	sarr := strings.SplitAfter(string(bytes), "[content]")

	return sarr[1]
}

func (mdm *MdManager) ListByClass(value string, start int, len int) []MarkDown {
	tagPapers, ok := mdm.ClassIndex[value]
	if !ok {
		return []MarkDown{}
	}
	// 按照sortindex排序
	ret := []MarkDown{}
	for _, mdId := range mdm.SortIndex {
		if _, ok := tagPapers[mdId]; ok {
			ret = append(ret, *mdm.Papers[mdId])
		}
	}

	return ret
}

func (mdm *MdManager) clearMdInfo(down *MarkDown) {
	// 清除文章
	delete(mdm.Papers, down.Id)
	// 清除排序
	mdSortKey := -1
	for key, value := range mdm.SortIndex {
		if value == down.Id {
			mdSortKey = key
		}
	}
	if mdSortKey >= 0 {
		 indexLen := len(mdm.SortIndex)
		if indexLen <= 1 {
			mdm.SortIndex = []string{}
		}else{
			if mdSortKey == 0 {
				mdm.SortIndex = mdm.SortIndex[1:]
			}

			if mdSortKey == indexLen -1 {
				mdm.SortIndex = mdm.SortIndex[0:mdSortKey]
			}

			if mdSortKey >0 && mdSortKey < indexLen - 1 {
				mdm.SortIndex = append(mdm.SortIndex[0:mdSortKey], mdm.SortIndex[mdSortKey +1 :]...)
			}
		}
	}

	// 清除类别
	delete(mdm.ClassIndex[down.Content.Class], down.Id)
}



