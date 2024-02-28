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
	for i := 0; i < arrLen; i++ {
		a = append(a, float32(i))
		b = append(b, float32(i))
	}

	for i := 0; i < arrLen; i += 4 {
		x = append(x, simd.F32x4{a[i], a[i+1], a[i+2], a[i+3]})
		y = append(y, simd.F32x4{b[i], b[i+1], b[i+2], b[i+3]})
	}
}

func main() {
	simdIntrinsics()
	simdGensimd()
}

func simdIntrinsics() {
	start := time.Now()
	C.add_arrays((*C.float)(unsafe.Pointer(&a[0])), (*C.float)(unsafe.Pointer(&b[0])), C.int(len(a)))
	fmt.Printf("SIMD Intrinsics: %v\n", float64(arrLen)/time.Since(start).Seconds()/1000000000)
}

func simdGensimd() {
	start := time.Now()
	sum := float32(0)
	for i := 0; i < len(a); i += 4 {
		a := simd.MulF32x4(simd.F32x4{a[i], a[i+1], a[i+2], a[i+3]}, simd.F32x4{b[i], b[i+1], b[i+2], b[i+3]})
		sum += a[0] + a[1] + a[2] + a[3]
	}
	fmt.Printf("SIMD gensimd: %v\n", float64(arrLen)/time.Since(start).Seconds())
}
