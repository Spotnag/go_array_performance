#include <xmmintrin.h>

void add_arrays(float* a, float* b, int len) {
    __m128 va, vb, vsum, total = _mm_setzero_ps();
    for (int i = 0; i < len; i += 4) {
        va = _mm_load_ps(a + i);
        vb = _mm_load_ps(b + i);
        vsum = _mm_mul_ps(va, vb);
        total = _mm_add_ps(total, vsum);
//        _mm_store_ps(a + i, vsum);
    }
    _mm_store_ps(a, total);
}