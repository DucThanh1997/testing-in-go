package sort

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestBubleSort(t *testing.T) {
	// 3 buoc de tao test case
	// init
	elements := []int{9, 7, 5, 1, 2, 1, 6, 8, 0}
	// execution
	// BubbleSort(elements)
	// validiation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 9, len(elements))
	assert.EqualValues(t, 0, elements[len(elements)-1])
	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)
	go func() {
		BubbleSort(elements)
		timeoutChan <- false 
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		timeoutChan <- true
	}()

	if <- timeoutChan {
		t.Error("lâu quá")
	}
	if elements[0] != 0 {
		t.Error("args: sai oi")
	}

	if elements[len(elements)-1] != 9 {
		t.Error(("sai 2"))
	}
}

// chỉ số coverage là khi cái test của bạn chạy vào toàn bộ các code trong cái function muốn test nhưng bạn
// vẫn cần cái thằng validiation

func TestSort(t *testing.T) {
	// 3 buoc de tao test case
	// init
	elements := []int{9, 7, 5, 1, 2, 1, 6, 8, 0}
	// execution
	Sort(elements)
	// validiation

	if elements[0] != 0 {
		t.Error(("args: sai oi"))
	}

	if elements[len(elements)-1] != 9 {
		t.Error(("sai 2"))
	}
}

func getElements(n int) []int {
	result :=make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}

func BenchmarkBuubleSort(b *testing.B) {
	elements := getElements(100000)
	// execution
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}

}

func BenchmarkSort(b *testing.B) {
	elements := getElements(100000)
	// execution
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}

}

// chạy benchmark bằng lệnh này go test -bench=.

// kết quả chạy benchmark

// BenchmarkBuubleSort-4   	60003000	        24.1 ns/op	       0 B/op	       0 allocs/op

// số lần chạy là 60003000
// tốn 24.1 nano second cho 1 lần chạy


// BenchmarkSort-4   	     589	   1856661 ns/op	     171 B/op	       1 allocs/op

// 16433996800 35365544