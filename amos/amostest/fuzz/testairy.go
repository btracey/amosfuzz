package fuzz

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/btracey/amosfuzz/amos"
	"github.com/btracey/amosfuzz/amos/amostest"

	"github.com/gonum/floats"
)

func Fuzz(data []byte) int {
	n := 8
	x, ok := F64S(data, n)
	if !ok {
		return -1
	}
	i1, ok := Int(data, 1)
	if !ok {
		return -1
	}
	data = data[1:]
	i2, ok := Int(data, 1)
	if !ok {
		return -1
	}
	data = data[1:]

	zs1s2test(x, []int{i1, i2})
	return 0
}

func zs1s2test(x []float64, is []int) {
	ZRR := x[0]
	ZRI := x[1]
	S1R := x[2]
	S1I := x[3]
	S2R := x[4]
	S2I := x[5]
	ASCLE := x[6]
	ALIM := x[7]

	NZ := is[0]
	IUF := is[1]

	ZRRfort, ZRIfort, S1Rfort, S1Ifort, S2Rfort, S2Ifort, NZfort, ASCLEfort, ALIMfort, IUFfort :=
		amostest.Zs1s2Fort(ZRR, ZRI, S1R, S1I, S2R, S2I, NZ, ASCLE, ALIM, IUF)
	ZRRamos, ZRIamos, S1Ramos, S1Iamos, S2Ramos, S2Iamos, NZamos, ASCLEamos, ALIMamos, IUFamos :=
		amos.Zs1s2(ZRR, ZRI, S1R, S1I, S2R, S2I, NZ, ASCLE, ALIM, IUF)

	SameF64Approx("zrr", ZRRfort, ZRRamos, 1e-14, 1e-14)
	SameF64Approx("zri", ZRIfort, ZRIamos, 1e-14, 1e-14)
	SameF64Approx("s1r", S1Rfort, S1Ramos, 1e-14, 1e-14)
	SameF64Approx("s1i", S1Ifort, S1Iamos, 1e-14, 1e-14)
	SameF64Approx("s2r", S2Rfort, S2Ramos, 1e-14, 1e-14)
	SameF64Approx("s2i", S2Ifort, S2Iamos, 1e-14, 1e-14)
	SameF64Approx("ASCLE", ASCLEfort, ASCLEamos, 1e-14, 1e-14)
	SameF64Approx("ALIM", ALIMfort, ALIMamos, 1e-14, 1e-14)
	SameInt("iuf", IUFfort, IUFamos)
	SameInt("nz", NZfort, NZamos)
}

func SameInt(str string, c, native int) {
	if c != native {
		panic(fmt.Sprintf("Case %s: Int mismatch. c = %v, native = %v.", str, c, native))
	}
}

func SameF64Approx(str string, c, native, absTol, relTol float64) {
	if math.IsNaN(c) && math.IsNaN(native) {
		return
	}
	if !floats.EqualWithinAbsOrRel(c, native, absTol, relTol) {
		cb := math.Float64bits(c)
		nb := math.Float64bits(native)
		same := floats.EqualWithinAbsOrRel(c, native, absTol, relTol)
		panic(fmt.Sprintf("Case %s: Float64 mismatch. c = %v, native = %v\n cb: %v, nb: %v\n%v,%v,%v", str, c, native, cb, nb, same, absTol, relTol))
	}
}

func F64S(data []byte, l int) ([]float64, bool) {
	var ok bool
	x := make([]float64, l)
	for i := range x {
		x[i], ok = F64(data)
		if !ok {
			return nil, false
		}
		data = data[8:]
	}
	return x, true
}

func Int(data []byte, b int) (n int, ok bool) {
	if len(data) < b {
		return 0, false
	}
	if b == 1 {
		return int(data[0]), true
	}
	if b == 2 {
		return int(binary.LittleEndian.Uint16(data[:2:2])), true
	}
	panic("not coded")
}

func F64(data []byte) (float64, bool) {
	if len(data) < 8 {
		return math.NaN(), false
	}
	uint64 := binary.LittleEndian.Uint64(data[:8:8])
	float64 := math.Float64frombits(uint64)
	return float64, true
}
