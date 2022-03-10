package service

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"os"
	"strings"
)

type FileInfo struct {
	path     string // 完整路径
	filename string //文件名
}

func (this *FileInfo) FileExit() bool {
	_, err := os.Stat(this.path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (this *FileInfo) GetFilename() string {
	fileInfo, err := os.Stat(this.path) //全路径
	if err != nil {
		return ""
	}
	return fileInfo.Name() // 文件名,没有路径
}

func (this *FileInfo) SetFilename() {
	this.filename = this.GetFilename()
}
func (this *FileInfo) QiNiuPath(source string) string {
	return source + "/" + this.filename
}

func NewFileInfo(path string) *FileInfo {
	fileInfo := &FileInfo{
		path: path,
	}
	fileInfo.SetFilename()
	return fileInfo
}

type Authorition struct {
	user      *UserInfo
	putPolicy storage.PutPolicy
	mac       *qbox.Mac
	cfg       storage.Config
}

func (this *Authorition) init() {
	this.user = NewUserInfo()
	this.putPolicy = storage.PutPolicy{
		Scope: this.user.bucket,
	}

	this.mac = qbox.NewMac(this.user.accessKey, this.user.secretKey)

	this.cfg = storage.Config{}
	// 空间对应的机房
	this.cfg.Zone = &storage.ZoneHuanan // 默认用的华南地区，其他地区的需要在这里改
	// 是否使用https域名
	this.cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	this.cfg.UseCdnDomains = false

}

type Uploader struct {
	Authorition
}

func (this *Uploader) NewUpload() (*storage.FormUploader, *UserInfo, storage.PutRet, string, storage.PutExtra) {
	// 七牛sdk提供
	this.init()
	// 构建表单上传的对象
	upToken := this.putPolicy.UploadToken(this.mac)
	formUploader := storage.NewFormUploader(&this.cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "CentOS Upload",
		},
	}
	return formUploader, this.user, ret, upToken, putExtra
}

type QiniuFileInfo struct {
	Authorition
}

func (this *QiniuFileInfo) Exist(filename string) bool {
	this.init()
	bucketManager := storage.NewBucketManager(this.mac, &this.cfg)
	_, sErr := bucketManager.Stat(this.user.bucket, SOURCE+"/"+filename)
	if sErr != nil {
		//fmt.Println(sErr)
		return false
	} else {
		//fmt.Println(fileInfo)
		//fmt.Println(fileInfo.String())
		//可以解析文件的PutTime
		//fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
		return true
	}

}

func (this *QiniuFileInfo) Delete(url string) error {
	this.init()
	//fmt.Println("space", this.user.space)

	//fmt.Println(url)
	url = strings.Trim(url, "")
	string_list := strings.Split(url, this.user.space)
	//fmt.Println(string_list)
	//fmt.Println(len(string_list))
	if len(string_list) != 2 {
		return fmt.Errorf("%s", "input url error")
	}
	bucketManager := storage.NewBucketManager(this.mac, &this.cfg)
	fmt.Println(string_list[0])
	fmt.Println(string_list[1])
	err := bucketManager.Delete(this.user.bucket, string_list[1]) // 除去域名的
	if err != nil {
		log.Println("删除文件失败")
		log.Println(err)
		return err
	}
	// 一般删除完了cdn还会缓存一段时间,如果需要马上删除,需要到七牛云上刷新一下缓存

	return nil
}

func (this *QiniuFileInfo) Walk() []string {
	// 遍历资源
	this.init()
	bucketManager := storage.NewBucketManager(this.mac, &this.cfg)

	fmt.Println("\nStart walk\n")
	limit := 1000
	prefix_list := []string{"picgo", "sharex", "typora", "upic", "uploadas", "webupload"}
	delimiter := ""
	//初始列举marker为空
	marker := ""

	imageList := make([]string, 0)
	for _, prefix := range prefix_list {
		for {
			entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(this.user.bucket, prefix, delimiter, marker, limit)
			if err != nil {
				fmt.Println("list error,", err)
				break
			}

			//print entries
			for _, entry := range entries {
				imageList = append(imageList, this.user.space+entry.Key)
			}

			if hasNext {
				marker = nextMarker
			} else {
				//list end
				break
			}

		}
	}
	return imageList
}
