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

func main() {
	nRuns := 1000000
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
		n := rand.Intn(5) + 1
		yr := make([]float64, n+1)
		yi := make([]float64, n+1)
		for j := range yr {
			yr[j] = randnum()
			yi[j] = randnum()
		}
		//tol := 1e-16 * (float64(rand.Intn(1000)) + 1)
		tol := 1e-14
		_ = tol
		_ = id
		_ = kode

		// Correctly tested.
		zs1s2test(x, is)
		// zuchktest(x, is, tol)
		// zkscltest(x, is, tol, n, yr, yi)
		// zmlritest(x, is, tol, n, yr, yi, kode)
		// zseritest(x, is, tol, n, yr, yi, kode)
		// zasyitest(x, is, tol, n, yr, yi, kode)
		// zbknutest(x, is, tol, n, yr, yi, kode)
		// zacaiest(x, is, tol, n, yr, yi, kode)
		// zairytest(x, kode, id)

	}
}

func zshchtest(x []float64) {
	zr := x[0]
	zi := x[1]
	cshr := x[2]
	cshi := x[3]
	cchr := x[4]
	cchi := x[5]

	ZRfort, ZIfort, CSHRfort, CSHIfort, CCHRfort, CCHIfort := amostest.Zshch(zr, zi, cshr, cshi, cchr, cchi)
	ZRamos, ZIamos, CSHRamos, CSHIamos, CCHRamos, CCHIamos := amos.Zshch(zr, zi, cshr, cshi, cchr, cchi)

	SameF64("zshch zr", ZRfort, ZRamos)
	SameF64("zshch zi", ZIfort, ZIamos)
	SameF64("zshch cshr", CSHRfort, CSHRamos)
	SameF64("zshch cshi", CSHIfort, CSHIamos)
	SameF64("zshch cshr", CCHRfort, CCHRamos)
	SameF64("zshch cchi", CCHIfort, CCHIamos)
}

func zairytest(x []float64, kode, id int) {
	ZR := x[0]
	ZI := x[1]
	KODE := kode
	ID := id

	AIRfort, AIIfort, NZfort := amostest.ZairyFort(ZR, ZI, ID, KODE)
	AIRamos, AIIamos, NZamos := amos.Zairy(ZR, ZI, ID, KODE)

	SameF64("zairy air", AIRfort, AIRamos)
	SameF64("zairy aii", AIIfort, AIIamos)
	SameInt("zairy nz", NZfort, NZamos)
}

func zkscltest(x []float64, is []int, tol float64, n int, yr, yi []float64) {
	ZRR := x[0]
	ZRI := x[1]
	FNU := x[2]
	NZ := is[1]
	ELIM := x[3]
	ASCLE := x[4]
	RZR := x[6]
	RZI := x[7]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRRfort, ZRIfort, FNUfort, Nfort, YRfort, YIfort, NZfort, RZRfort, RZIfort, ASCLEfort, TOLfort, ELIMfort :=
		amostest.ZksclFort(ZRR, ZRI, FNU, n, yrfort[1:], yifort[1:], NZ, RZR, RZI, ASCLE, tol, ELIM)
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
	ZRRamos, ZRIamos, FNUamos, Namos, YRamos, YIamos, NZamos, RZRamos, RZIamos, ASCLEamos, TOLamos, ELIMamos :=
		amos.Zkscl(ZRR, ZRI, FNU, n, yramos, yiamos, NZ, RZR, RZI, ASCLE, tol, ELIM)

	SameF64("zkscl zrr", ZRRfort, ZRRamos)
	SameF64("zkscl zri", ZRIfort, ZRIamos)
	SameF64("zkscl fnu", FNUfort, FNUamos)
	SameInt("zkscl n", Nfort, Namos)
	SameInt("zkscl nz", NZfort, NZamos)
	SameF64("zkscl rzr", RZRfort, RZRamos)
	SameF64("zkscl rzi", RZIfort, RZIamos)
	SameF64("zkscl ascle", ASCLEfort, ASCLEamos)
	SameF64("zkscl tol", TOLfort, TOLamos)
	SameF64("zkscl elim", ELIMfort, ELIMamos)

	SameF64S("zkscl yr", YRfort2, YRamos)
	SameF64S("zkscl yi", YIfort2, YIamos)
}

func zbknutest(x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
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
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort, ELIMfort, ALIMfort :=
		amostest.ZbknuFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, tol, ELIM, ALIM)
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
		amos.Zbknu(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, tol, ELIM, ALIM)

	SameF64("zbknu zr", ZRfort, ZRamos)
	SameF64("zbknu zr", ZIfort, ZIamos)
	SameF64("zbknu fnu", FNUfort, FNUamos)
	SameInt("zbknu kode", KODEfort, KODEamos)
	SameInt("zbknu n", Nfort, Namos)
	SameInt("zbknu nz", NZfort, NZamos)
	SameF64("zbknu tol", TOLfort, TOLamos)
	SameF64("zbknu elim", ELIMfort, ELIMamos)
	SameF64("zbknu alim", ALIMfort, ALIMamos)

	SameF64S("zbknu yr", YRfort2, YRamos)
	SameF64S("zbknu yi", YIfort2, YIamos)
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
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, RLfort, TOLfort, ELIMfort, ALIMfort :=
		amostest.ZasyiFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, RL, tol, ELIM, ALIM)
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

	SameF64("zasyi zr", ZRfort, ZRamos)
	SameF64("zasyi zr", ZIfort, ZIamos)
	SameF64("zasyi fnu", FNUfort, FNUamos)
	SameInt("zasyi kode", KODEfort, KODEamos)
	SameInt("zasyi n", Nfort, Namos)
	SameInt("zasyi nz", NZfort, NZamos)
	SameF64("zasyi rl", RLfort, RLamos)
	SameF64("zasyi tol", TOLfort, TOLamos)
	SameF64("zasyi elim", ELIMfort, ELIMamos)
	SameF64("zasyi alim", ALIMfort, ALIMamos)

	SameF64S("zasyi yr", YRfort2, YRamos)
	SameF64S("zasyi yi", YIfort2, YIamos)
}

func zmlritest(x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort :=
		amostest.ZmlriFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, tol)
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
	ZRamos, ZIamos, FNUamos, KODEamos, Namos, YRamos, YIamos, NZamos, TOLamos :=
		amos.Zmlri(ZR, ZI, FNU, KODE, n, yramos, yiamos, NZ, tol)

	SameF64("zmlri zr", ZRfort, ZRamos)
	SameF64("zmlri zi", ZIfort, ZIamos)
	SameF64("zmlri fnu", FNUfort, FNUamos)
	SameInt("zmlri kode", KODEfort, KODEamos)
	SameInt("zmlri n", Nfort, Namos)
	SameInt("zmlri nz", NZfort, NZamos)
	SameF64("zmlri tol", TOLfort, TOLamos)

	SameF64S("zmlri yr", YRfort2, YRamos)
	SameF64S("zmlri yi", YIfort2, YIamos)
}

func zacaiest(x []float64, is []int, tol float64, n int, yr, yi []float64, kode int) {
	ZR := x[0]
	ZI := x[1]
	FNU := x[2]
	KODE := kode
	NZ := is[1]
	MR := is[2]
	ELIM := x[3]
	ALIM := x[4]
	RL := x[5]

	yrfort := make([]float64, len(yr))
	copy(yrfort, yr)
	yifort := make([]float64, len(yi))
	copy(yifort, yi)
	ZRfort, ZIfort, FNUfort, KODEfort, MRfort, Nfort, YRfort, YIfort, NZfort, RLfort, TOLfort, ELIMfort, ALIMfort :=
		amostest.ZacaiFort(ZR, ZI, FNU, KODE, MR, n, yrfort[1:], yifort[1:], NZ, RL, tol, ELIM, ALIM)
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
	ZRamos, ZIamos, FNUamos, KODEamos, MRamos, Namos, YRamos, YIamos, NZamos, RLamos, TOLamos, ELIMamos, ALIMamos :=
		amos.Zacai(ZR, ZI, FNU, KODE, MR, n, yramos, yiamos, NZ, RL, tol, ELIM, ALIM)

	SameF64("zacai zr", ZRfort, ZRamos)
	SameF64("zacai zi", ZIfort, ZIamos)
	SameF64("zacai fnu", FNUfort, FNUamos)
	SameInt("zacai kode", KODEfort, KODEamos)
	SameInt("zacai mr", MRfort, MRamos)
	SameInt("zacai n", Nfort, Namos)
	SameInt("zacai nz", NZfort, NZamos)
	SameF64("zacai rl", RLfort, RLamos)
	SameF64("zacai tol", TOLfort, TOLamos)
	SameF64("zacai elim", ELIMfort, ELIMamos)
	SameF64("zacai elim", ALIMfort, ALIMamos)

	SameF64S("zacai yr", YRfort2, YRamos)
	SameF64S("zacai yi", YIfort2, YIamos)
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
	ZRfort, ZIfort, FNUfort, KODEfort, Nfort, YRfort, YIfort, NZfort, TOLfort, ELIMfort, ALIMfort :=
		amostest.ZseriFort(ZR, ZI, FNU, KODE, n, yrfort[1:], yifort[1:], NZ, tol, ELIM, ALIM)
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

	SameF64("zseri zr", ZRfort, ZRamos)
	SameF64("zseri zi", ZIfort, ZIamos)
	SameF64("zseri fnu", FNUfort, FNUamos)
	SameInt("zseri kode", KODEfort, KODEamos)
	SameInt("zseri n", Nfort, Namos)
	SameInt("zseri nz", NZfort, NZamos)
	SameF64("zseri tol", TOLfort, TOLamos)
	SameF64("zseri elim", ELIMfort, ELIMamos)
	SameF64("zseri elim", ALIMfort, ALIMamos)

	SameF64S("zseri yr", YRfort2, YRamos)
	SameF64S("zseri yi", YIfort2, YIamos)
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

	SameF64("zs1s2 zrr", ZRRfort, ZRRamos)
	SameF64("zs1s2 zri", ZRIfort, ZRIamos)
	SameF64("zs1s2 s1r", S1Rfort, S1Ramos)
	SameF64("zs1s2 s1i", S1Ifort, S1Iamos)
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

	SameF64("zuchk yr", YRfort, YRamos)
	SameF64("zuchk yi", YIfort, YIamos)
	SameInt("zuchk nz", NZfort, NZamos)
	SameF64("zuchk ascle", ASCLEfort, ASCLEamos)
	SameF64("zuchk tol", TOLfort, TOLamos)
}

func SameF64(str string, c, native float64) {
	if math.IsNaN(c) && math.IsNaN(native) {
		return
	}
	if c == native {
		return
	}
	cb := math.Float64bits(c)
	nb := math.Float64bits(native)
	panic(fmt.Sprintf("Case %s: Float64 mismatch. c = %v, native = %v\n cb: %v, nb: %v\n", str, c, native, cb, nb))
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

func SameF64ULP(str string, c, native float64, tol uint) {
	if math.IsNaN(c) && math.IsNaN(native) {
		return
	}
	if !floats.EqualWithinULP(c, native, tol) {
		cb := math.Float64bits(c)
		nb := math.Float64bits(native)
		panic(fmt.Sprintf("Case %s: Float64 mismatch. c = %v, native = %v\n cb: %v, nb: %v\n", str, c, native, cb, nb))
	}
}

func SameInt(str string, c, native int) {
	if c != native {
		panic(fmt.Sprintf("Case %s: Int mismatch. c = %v, native = %v.", str, c, native))
	}
}

func SameF64S(str string, c, native []float64) {
	if len(c) != len(native) {
		panic(str)
	}
	for i, v := range c {
		SameF64(str+"_idx_"+strconv.Itoa(i), v, native[i])
	}
}

func SameF64SApprox(str string, c, native []float64, absTol, relTol float64) {
	if len(c) != len(native) {
		panic(str)
	}
	fmt.Println("in approx")
	for i, v := range c {
		SameF64Approx(str+"_idx_"+strconv.Itoa(i), v, native[i], absTol, relTol)
	}
}
