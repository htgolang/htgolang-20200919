package math

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1+2 != 3")
	}
}

func TestSub(t *testing.T) {
	if Sub(2, 1) != 1 {
		t.Error("2-1 != 1")
	}
}

func TestMul(t *testing.T) {
	if Mul(10, 2) != 20 {
		t.Error("10*2 != 20")
	}
}

func TestDiv(t *testing.T) {
	t.Run("10/2 == 5", func(t *testing.T) {
		if Div(10, 2) != 5 {
			t.Error("10/2 != 5")
		}
	})

	t.Run("10/0 == 0", func(t *testing.T) {
		if Div(10, 0) != 0 {
			t.Error("10/0 != 0")
		}
	})
}

func BenchmarkFact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fact(10)
	}
}

func BenchmarkFactV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactV1(10)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("start")
	m.Run()
	fmt.Println("end")
}
