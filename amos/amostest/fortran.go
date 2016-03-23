package amostest

/*
void zs1s2_(double * ZRR, double * ZRI, double * S1R, double * S1I, double * S2R, double * S2I, int* NZ, double *ASCLE, double * ALIM, int * IUF);
*/
import "C"
import "unsafe"

func Zs1s2Fort(ZRR, ZRI, S1R, S1I, S2R, S2I float64, NZ int, ASCLE, ALIM float64, IUF int) (ZRRout, ZRIout, S1Rout, S1Iout, S2Rout, S2Iout float64, NZout int, ASCLEout, ALIMout float64, IUFout int) {
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
	return
}
