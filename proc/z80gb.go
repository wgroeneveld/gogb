
package proc

import (
	"github.com/wgroeneveld/gogb/util"
	"reflect"
	"strings"
)

type register struct {
	A, B, C, D, E, H, L, F int
	SP, PC, I, R int
}

type Z80 struct {
	Clock, Cycles int
	Halted bool

	Reg register
}

var alu = ALU{}

func (cpu *Z80) Nop() int {
	return 1
}

func (cpu *Z80) Halt() int {
	cpu.Halted = true
	return 0
}

func (cpu *Z80) Ld(x string, y string) int {
	fields := util.StrToFields(&cpu.Reg, x, y)

	// TODO (HL) specific part, or define another function (lefthand, righthand ops?)
	fields[0].SetInt(fields[1].Int())

	return 1
}

// by convention: _ separates arguments for opcodes
// key values represent binary opcodes coming from the Gameboy ROM
// http://imrannazar.com/Gameboy-Z80-Opcode-Map
var opcodes = [...]string {
	// 0		1			2			3			4			5			6			7			8			9			A			B			C			D			E			F
	// row 0x
	"nop",

	// row 1x

	// row 2x
	// row 3x
	// row 4x
	"ld_b_b",	"ld_b_c", 	"ld_b_d", 	"ld_b_e",	"ld_b_h", 	"ld_b_l", 	"ld_b_hl",	"ld_b_a", 	"ld_c_b",	"ld_c_c",	"ld_c_d",	"ld_c_e", 	"ld_c_h",	"ld_c_l",	"ld_c_hl",	"ld_c_a",
	// row 5x
	"ld_d_b",	"ld_d_c",	"ld_d_d",	"ld_d_e", 	"ld_d_h",	"ld_d_l",	"ld_d_hl",	"ld_d_a", 	"ld_e_b",	"ld_e_c",	"ld_e_d",	"ld_e_e", 	"ld_e_h",	"ld_e_l",	"ld_e_hl",	"ld_e_a",
	// row 6x
	"ld_h_b",	"ld_h_c",	"ld_h_d",	"ld_h_e", 	"ld_h_h",	"ld_h_l",	"ld_h_hl",	"ld_h_a", 	"ld_l_b",	"ld_l_c",	"ld_l_d",	"ld_l_e", 	"ld_l_h",	"ld_l_l",	"ld_l_hl",	"ld_l_a",
	// row 7x
	"ld_hl_b",	"ld_hl_c",	"ld_hl_d",	"ld_hl_e",	"ld_hl_h",	"ld_hl_l",	"halt",		"ld_hl_a",	"ld_a_b",	"ld_a_c",	"ld_a_d",	"ld_a_e",	"ld_a_h",	"ld_a_l",	"ld_a_hl",	"ld_a_a",
	// row 8x
	// row 9x
	// row Ax
	"and_a_b",	"and_a_c",	"and_a_d",	"and_a_e",	"and_a_h",	"and_a_l",	"and_a_hl",	"and_a__a",	"xorr_b",	"xorr_c",	"xorr_d",	"dorr_e",	"xorr_h",	"xorr_l",	"xorhl",	"xorr_a",
	// row Bx
	// row Cx
	// row Dx
	// row Ex
	// row Fx
}

func (cpu *Z80) execute(opcode string) {
	cpu.Reg.R = (cpu.Reg.R + 1) & 127

	split := strings.Split(opcode, "_")
	//args := split[1:]

	op := strings.Title(split[0])
	method := reflect.ValueOf(cpu).MethodByName(op)

	result := 0
	if !method.IsValid() {
		result = cpu.callAlu(strings.ToUpper(split[1]), strings.ToUpper(split[2]), op)
	} else {
		in := make([]reflect.Value, len(split)-1)
		for i := 0; i < method.Type().NumIn(); i++ {
			in[i] = reflect.ValueOf(strings.ToUpper(split[i+1]))
		}

		result = method.Call(in)[0].Interface().(int)
	}

	cpu.Cycles += result
}

func (cpu *Z80) callAlu(x string, y string, op string) int {
	fields := util.StrToFields(&cpu.Reg, x, y)
	a, b := int(fields[0].Int()), int(fields[1].Int())
	alu.A = a
	alu.B = b
	alu.Operation = op

	alu.Process()
	// TODO iets met resultaat van alu doen?

	// All ALU processing takes only one cycle!
	return 1
}