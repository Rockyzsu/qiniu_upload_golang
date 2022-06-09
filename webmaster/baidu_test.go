package webmaster

import (
	"fmt"
	"testing"
)

func TestSpider(t *testing.T) {
	//print(GetPushUrlsWP(1))
	//fmt.Println("End")
	result := GetPushUrlsBD(1)
	for _, i := range result {
		fmt.Println(i)
	}
}
