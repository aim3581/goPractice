package main

import "testing"

//go build -gcflags '-m -l'
// run this to see which varibles allocated on heap

const creations = 20000000

//go test -bench . -benchmem -a
//run above command to see benchmark
func BenchmarkTestCopyIt(t *testing.B) {
	for i := 0; i < creations; i++ {
		_ = CreateCopy()
	}
}

//go test -run=BenchmarkTestCopyIt -trace copy_trace.out
//go tool trace copy_trace.out

func BenchmarkTestPointerIt(t *testing.B) {
	for i := 0; i < creations; i++ {
		_ = CreatePointer()
	}
}

//go test -run=BenchmarkTestPointerIt -trace pointer_trace.out
//go tool trace pointer_trace.out
