package main

// #include <simd.c>
import "C"
import (
	"fmt"
	"github.com/bjwbell/gensimd/simd"
	"time"
	"unsafe"
)

var a, b []float32
var x, y []simd.F32x4

var arrLen = 100000000

func init() {
	// Create slices to be processed
	for i := 0; i < arrLen; i++ {
		a = append(a, float32(i))
		b = append(b, float32(i))
	}

	// Convert a and b to F32x4 slices so this conversion is not done in the performance loop
	for i := 0; i < arrLen; i += 4 {
		x = append(x, simd.F32x4{a[i], a[i+1], a[i+2], a[i+3]})
		y = append(y, simd.F32x4{b[i], b[i+1], b[i+2], b[i+3]})
	}
}

func main() {
	simdIntrinsics()
	simdGensimd()
	unrolled()
	unrolled_nobouncchecking()
}

func simdIntrinsics() {
	start := time.Now()
	C.add_arrays((*C.float)(unsafe.Pointer(&a[0])), (*C.float)(unsafe.Pointer(&b[0])), C.int(len(a)))
	fmt.Printf("SIMD Intrinsics - Bil ops/second: %v\n", float64(arrLen)/time.Since(start).Seconds()/1000000000)
}

func simdGensimd() {
	start := time.Now()
	sum := float32(0)
	for i := 0; i < len(a); i += 4 {
		a := simd.MulF32x4(simd.F32x4{a[i], a[i+1], a[i+2], a[i+3]}, simd.F32x4{b[i], b[i+1], b[i+2], b[i+3]})
		sum += a[0] + a[1] + a[2] + a[3]
	}
	fmt.Printf("SIMD gensimd - Bil ops/second: %v\n", float64(arrLen)/time.Since(start).Seconds()/1000000000)
}

func unrolled() {
	start := time.Now()
	sum := float32(0)
	for i := 0; i < len(a); i += 4 {
		sum += a[i] * b[i]
		sum += a[i+1] * b[i+1]
		sum += a[i+2] * b[i+2]
		sum += a[i+3] * b[i+3]
	}
	fmt.Printf("Unrolled - Bil ops/second: %v\n", float64(arrLen)/time.Since(start).Seconds()/1000000000)
}

func unrolled_nobouncchecking() {
	start := time.Now()
	sum := float32(0)
	for i := 0; i < len(a) && i < len(b); i += 4 {
		aTmp := a[i : i+4 : i+4]
		bTmp := b[i : i+4 : i+4]
		s0 := aTmp[0] * bTmp[0]
		s1 := aTmp[1] * bTmp[1]
		s2 := aTmp[2] * bTmp[2]
		s3 := aTmp[3] * bTmp[3]
		sum += s0 + s1 + s2 + s3
	}
	fmt.Printf("Unrolled no bound checking - Bil ops/second: %v\n", float64(arrLen)/time.Since(start).Seconds()/1000000000)
}
