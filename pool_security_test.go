package yamux

import "testing"

func TestPoolPutSmallCapacitiesNoPanic(t *testing.T) {
	for _, size := range []int{0, 1, 2, 4, 8, 16, 32, 33, 64, 65536, 131072} {
		var b []byte
		if size > 0 {
			b = make([]byte, size)
		}
		poolPut(b)
	}
}
