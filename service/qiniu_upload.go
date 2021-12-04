/*
 * #Date: 2021-11-23 01:33:18
 * #Author: Rocky Chen
 * #LastEditors: Rocky Chen
 * #LastEditTime: 2021-11-23 02:05:38
 * #FilePath: \qiniu_web_gin\service\qiniu_upload.go
 * #Description:
 */
package service

import (
	"context"
	"fmt"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
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
	fileInfo, err := os.Stat(this.path)
	if err != nil {
		return ""
	}
	return fileInfo.Name()
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
	this.cfg.Zone = &storage.ZoneHuanan
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

func UploadImg(path, source string) (string, error) {
	uploader := Uploader{}
	formUploader, user, ret, upToken, putExtra := uploader.NewUpload()

	fileInfo := NewFileInfo(path)

	if !fileInfo.FileExit() {
		return "", fmt.Errorf("File not found!")
	}

	err := formUploader.PutFile(
		context.Background(),
		&ret,
		upToken,
		fileInfo.QiNiuPath(source),
		fileInfo.path,
		&putExtra)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return "", fmt.Errorf("上传失败")
	}

	return user.space + ret.Key, nil
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

func CheckExist(filename string) bool {
	q_file := QiniuFileInfo{}
	return q_file.Exist(filename)
}
