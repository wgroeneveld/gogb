package proc

import (
	"errors"
	"reflect"
)

/*
Z80 ALU Unit

	(A) arithmetic operations:
Add, Sub, Div, Mul, Inc, Decr

	(L) Logical operations:
And, Or, Nand, Nor, Xor, Xnor, Not, (relational ops)
 */

 type flags struct {
 	Carry, Halfcarry, Zero, Negative int
 }

 type ALU struct {
 	A, B 				int
 	FlagsIn, FlagsOut 	flags
 	Operation 			string
}

 func (alu *ALU) Add() {
	alu.FlagsOut.Zero = 1
 }

func (alu *ALU) Process() {
	if alu.Operation == "" {
		panic(errors.New("Unknown operation!"))
	}

	method := reflect.ValueOf(alu).MethodByName(alu.Operation)
	if !method.IsValid() {
		panic(errors.New("Invalid operation: " + alu.Operation))
	}

	method.Call(nil)
}