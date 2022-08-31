package services

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	// 3 buoc de tao test case
	// init
	elements := []int{9, 7, 5, 1, 2, 1, 6, 8, 0}
	// execution
	fmt.Println("element: ", elements)
	Sort(elements)
	// validiation
	fmt.Println("element: ", elements)
	if elements[0] != 0 {
		t.Error(("args: sai oi"))
	}

	if elements[len(elements) - 1] != 9 {
		t.Error(("sai 2"))
	}
} 

// chỉ số coverage là khi cái test của bạn chạy vào toàn bộ các code trong cái function muốn test nhưng bạn 
// vẫn cần cái thằng validiation