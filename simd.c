#include <xmmintrin.h>

void add_arrays(float* a, float* b, int len) {
    __m128 va, vb, vsum;
    for (int i = 0; i < len; i += 4) {
        va = _mm_load_ps(a + i);
        vb = _mm_load_ps(b + i);
        vsum = _mm_mul_ps(va, vb);
        _mm_store_ps(a + i, vsum);
    }
}