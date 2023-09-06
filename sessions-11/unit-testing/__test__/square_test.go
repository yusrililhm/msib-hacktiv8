package test

import (
	"log"
	"math"
	"testing"
	"unit-test/square"

	"github.com/stretchr/testify/assert"
)

var persegi square.SquareInterface = square.SetSisi(6)

func TestKeliling(t *testing.T) {
	t.Logf("keliling : %.2f", persegi.Keliling())
}

func TestValidasiSisi(t *testing.T)  {
	if persegi.Area() == 0 {
		t.Errorf("sisi tidak boleh 0")
	}
}

// benchmark
func BenchmarkVolume(b *testing.B) {
	for i := 0; i < b.N; i++ {
		persegi.Volume()
	}
}

// testify
func TestLuasTestify(t *testing.T)  {
	assert.Equal(t, persegi.Volume(), 216.00, "volume salah")
	log.Println(math.Pow(6, 3))
}
