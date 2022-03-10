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
)

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

func CheckExist(filename string) bool {
	q_file := QiniuFileInfo{}
	return q_file.Exist(filename)
}

func DeleteImage(url string) bool {
	// url 全路径

	fileInfo := QiniuFileInfo{}
	err := fileInfo.Delete(url)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func Walkaround() []string {
	fileInfo := QiniuFileInfo{}
	imageList := fileInfo.Walk()

	imageList = SortName(imageList)
	return imageList

}
