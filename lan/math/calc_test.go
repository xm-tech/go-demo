package math

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	if addRet := Add(1, 2); addRet != 3 {
		t.Errorf("1 + 1 expedted be 3, but %d got", addRet)
	}

	if addRet := Add(-10, -30); addRet != -40 {
		t.Errorf("-10 + -30 expected -40, but %d got", addRet)
	}
}

// Table driven tests
func TestAdd2(t *testing.T) {
	cases := []struct{ A, B, Expected int }{
		{1, 2, 3},
		{-1, -2, -3},
		{-1, -1, -2},
		{0, 0, 0},
	}

	for _, c := range cases {
		// sub tests
		subTestName := fmt.Sprintf("%d + %d", c.A, c.B)
		t.Run(subTestName, func(t *testing.T) {
			actual := c.A + c.B
			if actual != c.Expected {
				t.Fatal("failure")
			}
		})
	}
}

func TestMul(t *testing.T) {
	if mulRet := Mul(3, 4); mulRet != 12 {
		t.Errorf("3 * 4 expected 12, but %d got", mulRet)
	}
}

type addCase struct{ A, B, Expected int }

// Test with helpers
func TestAdd3(t *testing.T) {
	createAddTestCase(t, &addCase{3, 4, 7})
	createAddTestCase(t, &addCase{-1, -1, -2})
	createAddTestCase(t, &addCase{-1, 1, 0})
}

func createAddTestCase(t *testing.T, c *addCase) {
	if addRet := Add(c.A, c.B); addRet != c.Expected {
		t.Fatalf("%d + %d expected  %d, but got %d", c.A, c.B, c.Expected, addRet)
	}
}

func TestAbs(t *testing.T) {
	if abs := Abs(float64(-1)); abs != 1 {
		t.Fatalf("Abs(-1) expected %f, bug got %f", float64(1), abs)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i+1)
	}
}
