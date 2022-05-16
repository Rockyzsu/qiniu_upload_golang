package webmaster

import (
	"fmt"
	"testing"
)

func TestSpider(t *testing.T) {
	print(GetPushUrlsWP(1))
	fmt.Println("End")
}
