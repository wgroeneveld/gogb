package proc

import (
	"errors"
	"reflect"
)

/*
Z80 ALU Unit

	(A) arithmetic operations:
Add(c), Sub(c), Div, Mul, Inc, Decr

	(L) Logical operations:
And, Or, Nand, Nor, Xor, Xnor, Not, (relational ops)
 */

 type flags struct {
 	Z, N, H, C bool		// 8-bit F register in cpu: ZNHC0000; first 4 are never used
 						// these will be sampled back into F after the operation
 }

 var bitz = 0x80
 var bitn = 0x40
 var bith = 0x20
 var bitc = 0x10

 func (aluFlags *flags) sampleFlagsIn(f int) {
 	aluFlags.Z = f & bitz > 0
 	aluFlags.N = f & bitn > 0
 	aluFlags.H = f & bith > 0
 	aluFlags.C = f & bitc > 0
 }

func (aluFlags flags) sampleFlags() int {
	result := 0
	if aluFlags.Z {
		result |= bitz
	}
	if aluFlags.N {
		result |= bitn
	}
	if aluFlags.H {
		result |= bith
	}
	if aluFlags.C {
		result |= bitc
	}
	return result
}


type ALU struct {
 	A, B, Z 			int
 	FlagsIn, FlagsOut 	flags
 	Operation 			string
}

 var bit8 = 0xff
 var hbit8 = 0x0f
 var hbit16 = 0x0fff
 var bit16 = 0xffff

 func (alu *ALU) Addc() int {
	return 1
 }

 func (alu *ALU) Add() int {
	result := alu.A + alu.B

	alu.FlagsOut = flags {
		Z: result == 0,
		N: false,
		H : (alu.A & hbit8) + (alu.B & hbit8) > hbit8,
		C: result > bit8,
	}

	alu.Z = result & bit8
	return 1
 }

func (alu *ALU) Process() int {
	if alu.Operation == "" {
		panic(errors.New("Unknown operation!"))
	}

	method := reflect.ValueOf(alu).MethodByName(alu.Operation)
	if !method.IsValid() {
		panic(errors.New("Invalid operation: " + alu.Operation))
	}

	return method.Call(nil)[0].Interface().(int)
}