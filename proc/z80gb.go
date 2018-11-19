
package proc

import (
	"reflect"
	"strings"
)

type register struct {
	A, B, C, D, E, H, L, F int
	SP, PC, I, R int
}

type Z80 struct {
	Clock int
	Cycles int	

	Reg register
}

func (cpu *Z80) Nop() int {
	return 1
}

func (cpu *Z80) Ld(x string, y string) int {
	r := reflect.Indirect(reflect.ValueOf(&cpu.Reg))
	fieldX := r.FieldByName(x)
	fieldY := r.FieldByName(y)

	fieldX.SetInt(fieldY.Int())

	return 1
}

var opcodes = [...]string {
	// by convention: _ separates arguments for opcodes 
	"nop",
	"ld_a_b",
}

func (cpu *Z80) execute(opcode string) {
	cpu.Reg.R = (cpu.Reg.R + 1) & 127

	split := strings.Split(opcode, "_")
	//args := split[1:]

	method := reflect.ValueOf(cpu).MethodByName(strings.Title(split[0]))
	in := make([]reflect.Value, len(split) - 1)
	for i := 0; i < method.Type().NumIn(); i++ {
		in[i] = reflect.ValueOf(split[i + 1])
	}

	result := method.Call(in)[0].Interface().(int)

	cpu.Cycles += result
}

