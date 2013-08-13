package main

import (
	"github.com/npadmana/mpi"
	"log"
)

func TestInt64() {
	if n1 := mpi.TypeSize(mpi.MPI_i64); n1 != 8 {
		log.Printf("FAIL : TestInt64 : Size of MPI_LONG=%d, needs to be 8 for code to work", n1)
	} else {
		log.Println("PASS : TestInt64")
	}
}

func main() {
	mpi.Initialize()
	TestInt64()
	mpi.Finalize()
}
