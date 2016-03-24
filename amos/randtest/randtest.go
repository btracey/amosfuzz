package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/btracey/amosfuzz/amos"
	"github.com/btracey/amosfuzz/amos/amostest"
	"github.com/gonum/floats"
)

func randnum() float64 {
	r := 2e4
	if rand.Float64() > 0.99 {
		return 0
	}
	return rand.Float64()*r - r/2
}

func main() {
	nRuns := 100000
	for i := 0; i < nRuns; i++ {
		fmt.Println("\nrun = ", i)
		x := make([]float64, 8)
		for j := range x {
			x[j] = randnum()
		}

		is := make([]int, 3)
		for j := range is {
			is[j] = rand.Intn(1000)
		}
		kode := rand.Intn(2) + 1
		id := rand.Intn(2)
		n := rand.Intn(1) + 1
		yr := make([]float64, n+1)
		yi := make([]float64, n+1)
		for j := range yr {
			yr[j] = randnum()
			yi[j] = randnum()
		}
		tol := 1e-16 * float64(rand.Intn(1000))
		_ = tol
		_ = id
		//zs1s2test(x, is)
		//zuchktest(x, is, tol)
		//zseritest(x, is, tol, n, yr, yi, kode) // This has an infinite loop if the fortran code is called
		//zasyitest(x, is, tol, n, yr, yi, kode)
		zairytest(x, kode, id)
	}
}

func zairytest(x []float64, kode, id int) {
	ZR := x[0]
	ZI := x[1]
	KODE := kode
	ID := id

	AIRfort, AIIfort, NZfort := amostest.ZairyFort(ZR, ZI, ID, KODE)
	AIRamos, AIIamos, NZamos := amos.Zairy(ZR, ZI, ID, KODE)

	SameF64Approx("zairy air", AIRfort, AIRamos, 1e-4, 1e-4)
	SameF64Approx("zairy aii", AIIfort, AIIamos, 1e-4, 1e-4)
	SameInt("zairy nz", NZfort, NZamos)
}

func zasyitest(x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	ELIM := x[3]
	ALIM := x[4]
	RL := x[5]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	fmt.Println("Starting fortran")
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, RLfort, TOLfort, ELIMfort, ALIMfort :=
		amostest.ZasyiFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, RL, tol, ELIM, ALIM)
	fmt.Println("done fortran")
	YRfort2 := make([]float64, len(yrfort))
	YRfort2[0] = yrfort[0]
	copy(YRfort2[1:], YRfort)
	YIfort2 := make([]float64, len(yifort))
	YIfort2[0] = yifort[0]
	copy(YIfort2[1:], YIfort)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, RLamos, TOLamos, ELIMamos, ALIMamos :=
		amos.Zasyi(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, RL, tol, ELIM, ALIM)

	SameF64Approx("zasyi zr", ZRfort, ZRamos, 1e-14, 1e-14)
	SameF64Approx("zasyi zr", ZIfort, ZIamos, 1e-14, 1e-14)
	SameF64Approx("zasyi fnu", FNUfort, FNUamos, 1e-14, 1e-14)
	SameInt("zasyi kode", KODEfort, KODEamos)
	SameInt("zasyi n", Nfort, Namos)
	SameInt("zasyi nz", NZfort, NZamos)
	SameF64Approx("zasyi rl", RLfort, RLamos, 1e-14, 1e-14)
	SameF64Approx("zasyi tol", TOLfort, TOLamos, 1e-14, 1e-14)
	SameF64Approx("zasyi elim", ELIMfort, ELIMamos, 1e-14, 1e-14)
	SameF64Approx("zasyi alim", ALIMfort, ALIMamos, 1e-14, 1e-14)

	SameF64SApprox("zasyi yr", YRfort2, YRamos, 1e-14, 1e-14)
	SameF64SApprox("zasyi yi", YIfort2, YIamos, 1e-14, 1e-14)
}

func zseritest(x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	ELIM := x[3]
	ALIM := x[4]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	fmt.Println("Starting fortran")
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort, ELIMfort, ALIMfort :=
		amostest.ZseriFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, tol, ELIM, ALIM)
	fmt.Println("done fortran")
	YRfort2 := make([]float64, len(yrfort))
	YRfort2[0] = yrfort[0]
	copy(YRfort2[1:], YRfort)
	YIfort2 := make([]float64, len(yifort))
	YIfort2[0] = yifort[0]
	copy(YIfort2[1:], YIfort)

	yramos := make([]float64, len(yr))
	copy(yramos, yr)
	yiamos := make([]float64, len(yi))
	copy(yiamos, yi)
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, TOLamos, ELIMamos, ALIMamos :=
		amos.Zseri(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, tol, ELIM, ALIM)

	SameF64Approx("zseri zr", ZRfort, ZRamos, 1e-14, 1e-14)
	SameF64Approx("zseri zr", ZIfort, ZIamos, 1e-14, 1e-14)
	SameF64Approx("zseri fnu", FNUfort, FNUamos, 1e-14, 1e-14)
	SameInt("zseri kode", KODEfort, KODEamos)
	SameInt("zseri n", Nfort, Namos)
	SameInt("zseri nz", NZfort, NZamos)
	SameF64Approx("zseri tol", TOLfort, TOLamos, 1e-14, 1e-14)
	SameF64Approx("zseri elim", ELIMfort, ELIMamos, 1e-14, 1e-14)
	SameF64Approx("zseri elim", ALIMfort, ALIMamos, 1e-14, 1e-14)

	SameF64SApprox("zseri yr", YRfort2, YRamos, 1e-14, 1e-14)
	SameF64SApprox("zseri yi", YIfort2, YIamos, 1e-14, 1e-14)
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

	SameF64Approx("zs1s2 zrr", ZRRfort, ZRRamos, 1e-14, 1e-14)
	SameF64Approx("zs1s2 zri", ZRIfort, ZRIamos, 1e-14, 1e-14)
	SameF64Approx("zs1s2 s1r", S1Rfort, S1Ramos, 1e-14, 1e-14)
	SameF64Approx("zs1s2 s1i", S1Ifort, S1Iamos, 1e-14, 1e-14)
	SameF64Approx("zs1s2 s2r", S2Rfort, S2Ramos, 1e-14, 1e-14)
	SameF64Approx("zs1s2 s2i", S2Ifort, S2Iamos, 1e-14, 1e-14)
	SameF64Approx("zs1s2 ascle", ASCLEfort, ASCLEamos, 1e-14, 1e-14)
	SameF64Approx("zs1s2 alim", ALIMfort, ALIMamos, 1e-14, 1e-14)
	SameInt("iuf", IUFfort, IUFamos)
	SameInt("nz", NZfort, NZamos)
}

func zuchktest(x []float64, is []int, tol float64) {
	YR := x[0]
	YI := x[1]
	NZ := is[0]
	ASCLE := x[2]
	TOL := tol

	YRfort, YIfort, NZfort, ASCLEfort, TOLfort := amostest.ZuchkFort(YR, YI, NZ, ASCLE, TOL)
	YRamos, YIamos, NZamos, ASCLEamos, TOLamos := amos.Zuchk(YR, YI, NZ, ASCLE, TOL)

	SameF64Approx("zuchk yr", YRfort, YRamos, 1e-14, 1e-14)
	SameF64Approx("zuchk yi", YIfort, YIamos, 1e-14, 1e-14)
	SameInt("zuchk nz", NZfort, NZamos)
	SameF64Approx("zuchk ascle", ASCLEfort, ASCLEamos, 1e-14, 1e-14)
	SameF64Approx("zuchk tol", TOLfort, TOLamos, 1e-14, 1e-14)
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

func SameInt(str string, c, native int) {
	if c != native {
		panic(fmt.Sprintf("Case %s: Int mismatch. c = %v, native = %v.", str, c, native))
	}
}

func SameF64SApprox(str string, c, native []float64, absTol, relTol float64) {
	if len(c) != len(native) {
		panic(str)
	}
	fmt.Println("len ", len(c))
	for i, v := range c {
		SameF64Approx(str+"_idx_"+strconv.Itoa(i), v, native[i], absTol, relTol)
	}
}
