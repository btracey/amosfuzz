package amostest

/*
void zs1s2_(double * ZRR, double * ZRI, double * S1R, double * S1I, double * S2R, double * S2I, int* NZ, double *ASCLE, double * ALIM, int * IUF);
void zseri_(double * ZR, double * ZI, double * FNU, int * KODE, int * N, double * YR, double * YI, int * NZ, double * tol, double * elim, double * alim);
void zasyi_(double * ZR, double * ZI, double * FNU, int * KODE, int * N, double * YR, double * YI, int * NZ,double * RL, double * tol, double * elim, double * alim);
void zuchk_(double * YR, double * YI, int * NZ, double * ASCLE, double * TOL);
void zairy_(double * ZR, double * ZI, int * ID, int * KODE, double * AIR, double * AII, int * NZ, int * IERR);
*/
import "C"
import "unsafe"

func ZairyFort(ZR, ZI float64, ID, KODE int) (AIR, AII float64, NZ int) {
	var IERR int
	pzr := (*C.double)(&ZR)
	pzi := (*C.double)(&ZI)
	pid := (*C.int)(unsafe.Pointer(&ID))
	pkode := (*C.int)(unsafe.Pointer(&KODE))

	pair := (*C.double)(&AIR)
	paii := (*C.double)(&AII)
	pnz := (*C.int)(unsafe.Pointer(&NZ))
	pierr := (*C.int)(unsafe.Pointer(&IERR))
	C.zairy_(pzr, pzi, pid, pkode, pair, paii, pnz, pierr)

	NZ = int(*pnz)
	return AIR, AII, NZ
}

func ZasyiFort(ZR, ZI, FNU float64, KODE, N int, YR, YI []float64, NZ int, RL, TOL, ELIM, ALIM float64) (
	ZRout, ZIout, FNUout float64, KODEout, Nout int, YRout, YIout []float64, NZout int, RLout, TOLout, ELIMout, ALIMout float64) {

	pzr := (*C.double)(&ZR)
	pzi := (*C.double)(&ZI)
	pfnu := (*C.double)(&FNU)
	pkode := (*C.int)(unsafe.Pointer(&KODE))
	pn := (*C.int)(unsafe.Pointer(&N))
	pyr := (*C.double)(&YR[0])
	pyi := (*C.double)(&YI[0])
	pnz := (*C.int)(unsafe.Pointer(&NZ))
	prl := (*C.double)(&RL)
	ptol := (*C.double)(&TOL)
	pelim := (*C.double)(&ELIM)
	palim := (*C.double)(&ALIM)

	C.zasyi_(pzr, pzi, pfnu, pkode, pn, pyr, pyi, pnz, prl, ptol, pelim, palim)
	KODE = int(*pkode)
	N = int(*pn)
	NZ = int(*pnz)
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, RL, TOL, ELIM, ALIM
}

func ZuchkFort(YR, YI float64, NZ int, ASCLE, TOL float64) (YRout, YIout float64, NZout int, ASCLEout, TOLout float64) {
	pyr := (*C.double)(&YR)
	pyi := (*C.double)(&YI)
	pnz := (*C.int)(unsafe.Pointer(&NZ))
	pascle := (*C.double)(&ASCLE)
	ptol := (*C.double)(&TOL)

	C.zuchk_(pyr, pyi, pnz, pascle, ptol)
	return YR, YI, NZ, ASCLE, TOL
}

func Zs1s2Fort(ZRR, ZRI, S1R, S1I, S2R, S2I float64, NZ int, ASCLE, ALIM float64, IUF int) (
	ZRRout, ZRIout, S1Rout, S1Iout, S2Rout, S2Iout float64, NZout int, ASCLEout, ALIMout float64, IUFout int) {

	pzrr := (*C.double)(&ZRR)
	pzri := (*C.double)(&ZRI)
	ps1r := (*C.double)(&S1R)
	ps1i := (*C.double)(&S1I)
	ps2r := (*C.double)(&S2R)
	ps2i := (*C.double)(&S2I)
	pnz := (*C.int)(unsafe.Pointer(&NZ))
	pascle := (*C.double)(&ASCLE)
	palim := (*C.double)(&ALIM)
	piuf := (*C.int)(unsafe.Pointer(&IUF))

	C.zs1s2_(pzrr, pzri, ps1r, ps1i, ps2r, ps2i, pnz, pascle, palim, piuf)
	return ZRR, ZRI, S1R, S1I, S2R, S2I, NZ, ASCLE, ALIM, IUF
}

func ZseriFort(ZR, ZI, FNU float64, KODE, N int, YR, YI []float64, NZ int, TOL, ELIM, ALIM float64) (
	ZRout, ZIout, FNUout float64, KODEout, Nout int, YRout, YIout []float64, NZout int, TOLout, ELIMout, ALIMout float64) {
	pzr := (*C.double)(&ZR)
	pzi := (*C.double)(&ZI)
	pfnu := (*C.double)(&FNU)
	pkode := (*C.int)(unsafe.Pointer(&KODE))
	pn := (*C.int)(unsafe.Pointer(&N))
	pyr := (*C.double)(&YR[0])
	pyi := (*C.double)(&YI[0])
	pnz := (*C.int)(unsafe.Pointer(&NZ))
	ptol := (*C.double)(&TOL)
	pelim := (*C.double)(&ELIM)
	palim := (*C.double)(&ALIM)

	C.zseri_(pzr, pzi, pfnu, pkode, pn, pyr, pyi, pnz, ptol, pelim, palim)
	KODE = int(*pkode)
	N = int(*pn)
	NZ = int(*pnz)
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
}
