package types

import (
	"fmt"
	"testing"
	"time"
)

func TestMedian(t *testing.T) {
	// create a map where key is time in UnixNanoSeconds and value is voting power
	m := make(map[time.Time]int64)
	t1 := Now()
	t2 := t1.Add(10 * time.Minute)
	t3 := t2.Add(10 * time.Minute)
	m[t1] = 100
	m[t2] = 1000
	m[t3] = 501

	for k, v := range m {
		fmt.Println("k:", k, " v:", v)
	}

	fmt.Println("Median:", WeightedMedian(m))
}
