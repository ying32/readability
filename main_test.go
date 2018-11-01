package readability

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	test, err := NewReadability("http://wd.leiting.com/home/news/news_detail.php?id=599")
	if err != nil {
		fmt.Println("failed.", err)
		return
	}
	test.Parse()
	fmt.Println(test.Title)
	fmt.Println(test.Content)
}
