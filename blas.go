package gorgonia

import (
	"sync"

	tf64 "github.com/chewxy/gorgonia/tensor/f64"
	"github.com/gonum/blas"
	"github.com/gonum/blas/native"
)

var blasdoor sync.Mutex
var whichblas BLAS

// BLAS represents all the possible implementations of BLAS.
// The default is Gonum's Native
type BLAS interface {
	blas.Float32
	blas.Float64
	// blas.Complex64
	// blas.Complex128
}

// only blastoise.Implementation() and cubone.Implementation() are batchedBLAS -
// they both batch cgo calls (and cubone batches cuda calls)
type batchedBLAS interface {
	WorkAvailable() int
	DoWork()
	BLAS
}

// Use defines which BLAS implementation gorgonia should use.
// The default is Gonum's Native. These are the other options:
//		Use(blastoise.Implementation())
//		Use(cubone.Implementation())
//		Use(cgo.Implementation)
// Note the differences in the brackets. The blastoise and cubone ones are functions.
func Use(b BLAS) {
	// close the blast door! close the blast door!
	blasdoor.Lock()
	// open the blast door! open the blast door!
	defer blasdoor.Unlock()
	// those lines were few of the better additions to the Special Edition. There, I said it. The Special Edition is superior. Except Han still shot first in my mind.

	whichblas = b
	f64 := b.(blas.Float64)
	tf64.Use(f64)

	// TODO:
	// float32
}

// WhichBLAS() returns the BLAS that gorgonia uses.
func WhichBLAS() BLAS { return whichblas }

func init() {
	whichblas = native.Implementation{}
}
