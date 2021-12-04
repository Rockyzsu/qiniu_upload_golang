/*
 * #Date: 2021-11-23 02:07:18
 * #Author: Rocky Chen
 * #LastEditors: Rocky Chen
 * #LastEditTime: 2021-11-23 02:11:25
 * #FilePath: \qiniu_web_gin\service\config.go
 * #Description:
 */
package service

import "os"

type UserInfo struct {
	accessKey string
	secretKey string
	bucket    string
	space     string
}

func NewUserInfo() *UserInfo {

	return &UserInfo{
		accessKey: os.Getenv("qiniu_access_key"),
		secretKey: os.Getenv("qiniu_secret_key"),
		bucket:    "golandspace",
		space:     "http://xximg.30daydo.com/",
	}
}