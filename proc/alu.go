package proc

/*

Z80 = 8-BIT

shifts (left, right)
or
and
xor
sub
add


 */

 type flags struct {
 	Z, N, H, C int
 }

 type ALU struct {
 	A, B int
 	FlagsIn, FlagsOut flags
 	Operation int
}

func (alu *ALU) process() {

}