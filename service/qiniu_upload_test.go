package service

import "testing"

func TestDeleteImage(t *testing.T) {
	//t.Log("Pass")
	url := "http://xximg.30daydo.com/picgo/1635745295943.jpg"
	if !DeleteImage(url) {
		t.Log("Failed")
	}
}
