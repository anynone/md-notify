package md_notify

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"md-notify/inter_struct"
	"os"
	"strconv"
	"strings"
)

type MarkDown struct {
	Id string
	FileName string
	LastUpdate int64
	Content inter_struct.Content
}

//读取markdown
func (md *MarkDown) Load(file *os.File) *MarkDown{
	hash := md5.New()
	hash.Write([]byte(file.Name()))
	md.Id = hex.EncodeToString(hash.Sum(nil))

	md.FileName = file.Name()
	info, _ := file.Stat()
	md.LastUpdate = info.ModTime().Unix()
	bytes, _ := ioutil.ReadAll(file)

	md.Content = md.ParseContent(bytes)
	return md
}

func (md *MarkDown) ParseContent(bytes []byte) inter_struct.Content {

	mdContent := inter_struct.Content{
		Title:    "",
		Class:    "",
		Sort:     0,
		SubTitle: "",
		Image:    "",
		Extra: map[string]string{},
	}
	content := string(bytes)

	split := strings.SplitAfter(content, "[content]")

	config := strings.ReplaceAll(split[0], "[content]", "")

	// 分行
	configSlice := strings.Split(config, "\n")

	// 遍历,逐个写入
	for _, value := range configSlice {
		kv := strings.Split(value, "=")
		if len(kv) < 2 {
			continue
		}
		switch kv[0] {
		case "title":
			mdContent.Title = kv[1]
		case "class":
			mdContent.Class = kv[1]
		case "sort":
			num, e := strconv.Atoi(kv[1])
			if e != nil {
				continue
			}
			mdContent.Sort = num
		case "sub_title":
			mdContent.SubTitle = kv[1]
		case "image":
			mdContent.Image = kv[1]
		default:
			mdContent.Extra[kv[0]] = kv[1]
		}
	}

	return mdContent
}

func (md *MarkDown) MarkdownContent() []byte{
	bytes, e := ioutil.ReadFile(md.FileName)
	if e != nil {
		return []byte{}
	}

	content := strings.SplitAfter(string(bytes), "[content]")
	if len(content) < 2 {
		return []byte{}
	}

	return []byte(content[1])
}


