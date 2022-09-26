package service

import "testing"

func TestDeleteImage(t *testing.T) {
	/*

	 */
	url := "http://www.com/picgo/1635745295943.jpg"
	if !DeleteImage(url) {
		t.Log("Failed")
	}
}

func TestUpload(t *testing.T) {
	file := "/home/xda/Pictures/dc321c62125c521796893b1eff0762f0.png"
	_, err := UploadImg("/pico/testing.png", file)
	if err != nil {
		t.Log("Failed")
	} else {
		t.Log("Passed")
	}

}
