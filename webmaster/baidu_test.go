package webmaster

import (
	"fmt"
	"testing"
)

func TestSpider(t *testing.T) {

	result := GetPushUrlsBD(1)
	for _, i := range result {
		fmt.Println(i)
	}
}
