package mathx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSumTable(t *testing.T) {
	for _, c := range []struct{ a, b, want int }{{2, 3, 5}, {10, -5, 5}, {0, 0, 0}, {-2, -3, -5}} {
		if got := Sum(c.a, c.b); got != c.want {
			t.Fatalf("Sum(%d,%d)=%d, want %d", c.a, c.b, got, c.want)
		}
	}
}
func TestDivide(t *testing.T) {
	for _, c := range []struct {
		name       string
		a, b, want int
		bad        bool
	}{{"positive", 10, 2, 5, false}, {"negative", -9, 2, -4, false}, {"zero numerator", 0, 3, 0, false}, {"divide by zero", 1, 0, 0, true}} {
		t.Run(c.name, func(t *testing.T) {
			v, e := Divide(c.a, c.b)
			if c.bad {
				require.Error(t, e)
			} else {
				require.NoError(t, e)
				assert.Equal(t, c.want, v)
			}
		})
	}
}

var sink int

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = Sum(i, 456)
	}
}
