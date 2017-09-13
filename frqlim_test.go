package frqlim

import (
	"testing"
	"time"
	"fmt"
)

func Test(t *testing.T) {
	lim := New(4, time.Second)
	for i := 0; i < 20; i++ {
		fmt.Println(lim.Do())
		time.Sleep(time.Millisecond * 200)
	}
}

func Benchmark(b *testing.B) {
	lim := New(100, time.Second)
	for i:=0;i<b.N ;i++  {
		lim.Do()
	}
}
