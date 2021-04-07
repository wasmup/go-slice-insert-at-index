package main

import (
	"flag"
	"os"
	"testing"
)

func insert(a []int, index int, value int) []int {
	a = append(a[:index+1], a[index:]...) // Step 1+2
	a[index] = value                      // Step 3
	return a
}
func insert2(a []int, index int, value int) []int {
	last := len(a) - 1
	a = append(a, a[last])           // Step 1
	copy(a[index+1:], a[index:last]) // Step 2
	a[index] = value                 // Step 3
	return a
}
func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = insert(a, 2, 42)
	}
}
func BenchmarkInsert2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r = insert2(a, 2, 42)
	}
}

var (
	n    = flag.Int("n", 32, "buffer length")
	a, r []int
)

// We use TestMain to set up the buffer.
func TestMain(m *testing.M) {
	flag.Parse()
	a = make([]int, *n)
	os.Exit(m.Run())
}
