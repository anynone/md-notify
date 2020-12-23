package md_notify

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"md-notify/inter_struct"
	"os"
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
	return inter_struct.Content{
		Title:    "",
		Class:    "",
		Sort:     0,
		SubTitle: "",
		Image:    "",
		Extra:    nil,
	}	
}


