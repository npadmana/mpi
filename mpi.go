// Package mpi wraps any MPI commands I might need.
//
package mpi

/*
#cgo pkg-config: ompi

#include <stdlib.h>
#include "mpi.h"

MPI_Op mpiop(int i) {
	MPI_Op retval;
	switch(i) {
	case 0 :
		retval = MPI_SUM;
		break;
	default :
		MPI_Abort(MPI_COMM_WORLD,1);
	}
	return retval;
}

MPI_Datatype mpitype(int i) {
	MPI_Datatype retval;
	switch(i) {
	case 0 :
		retval = MPI_LONG;
		break;
	default :
		MPI_Abort(MPI_COMM_WORLD,1);
	}
	return retval;
}

*/
import "C"

import (
	"errors"
	"os"
	"unsafe"
)

type Comm C.MPI_Comm

var (
	SUM = C.mpiop(0)
)

var (
	LONG = C.mpitype(0)
)

// Initialize initializes the MPI environment
func Initialize() error {
	// Allocate space for argc and argv
	argc := C.int(len(os.Args))
	argv := make([](*C.char), argc)
	// Copy os.Args into argv
	for i, gstr := range os.Args {
		argv[i] = C.CString(gstr)
	}
	ptrargv := &argv[0]

	perr := C.MPI_Init(&argc, &ptrargv)
	if perr != 0 {
		return errors.New("Error initializing MPI")
	}

	// update os.Args
	os.Args = os.Args[0:0]
	for i := 0; i < int(argc); i++ {
		os.Args = append(os.Args, C.GoString(argv[i]))
		C.free(unsafe.Pointer(argv[i]))
	}

	return nil
}

// Finalize finalizes the MPI environment
func Finalize() error {
	perr := C.MPI_Finalize()
	if perr != 0 {
		return errors.New("Error initializing MPI")
	}
	return nil
}

// AllReduceInt64 : MPI_Allreduce for int64
func AllReduceInt64(comm C.MPI_Comm, in, out *int64, n int, op C.MPI_Op) {
	C.MPI_Allreduce(unsafe.Pointer(&in), unsafe.Pointer(&out), C.int(n), op, LONG, comm)
}

// Abort calls MPI_Abort
func Abort(comm Comm, err int) error {
	perr := C.MPI_Abort(comm, C.int(err))
	if perr != 0 {
		return errors.New("Error aborting!?!!")
	}
	return nil
}

// Barrier calls MPI_Barrier
func Barrier(comm Comm, err int) error {
	perr := C.MPI_Barrier(comm)
	if perr != 0 {
		return errors.New("Error calling Barrier")
	}
	return nil
}

// Rank returns the MPI_Rank
func Rank(comm Comm, err int) (int, error) {
	var r C.int
	perr := C.MPI_Comm_rank(comm, &r)
	if perr != 0 {
		return -1, errors.New("Error calling MPI_Comm_rank")
	}
	return int(r), nil
}

// Size returns the MPI_Size
func Size(comm Comm, err int) (int, error) {
	var r C.int
	perr := C.MPI_Comm_size(comm, &r)
	if perr != 0 {
		return -1, errors.New("Error calling MPI_Comm_size")
	}
	return int(r), nil
}
