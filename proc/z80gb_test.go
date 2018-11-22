
package proc

import (
	"testing"
)

func setup() Z80 {
	return Z80 {
		Reg: register { A: 1, B: 2, },
		Cycles: 0,
	}
}

func TestCPUAndALUIntegration(t *testing.T) {
	op("inc16_a_b").test("16-bit Increment splits into 2 8-bit inputs and merges result back").with(Z80 {
		Reg: register { A: 255, B: 2, },
		Cycles: 0,
	}).expects(Z80 {
		Reg: register { A: 255, B: 3, },
		Cycles: 1,
	}).verify(t)

	op("inc16_a_b").test("16-bit Increment overflows first 8-bit into second half").with(Z80 {
		Reg: register { A: 2, B: 255, },
		Cycles: 0,
	}).expects(Z80 {
		Reg: register { A: 3, B: 0, },
		Cycles: 1,
	}).verify(t)

	op("add_a_b").test("ALU output sampled into CPU registers").with(Z80 {
		Reg: register { A: 254, B: 3, },
	}).expects(Z80 {
		Reg: register { A: 1, B: 3, F: 0x30 },
		Cycles: 1,
	}).verify(t)

	op("add_a_b").test("ALU integration test with add").with(Z80 {
		Reg: register { A: 1, B: 2, },
		Cycles: 0,
	}).expects(Z80 {
		Reg: register { A: 3, B: 2, },
		Cycles: 1,
	}).verify(t)
}

func TestExecuteOperations(t *testing.T) {
	op("ld_a_b").with(Z80 {
		Reg: register { A: 1, B: 2, },
		Cycles: 0,
	}).expects(Z80 {
		Reg: register { A: 3, B: 2, },
		Cycles: 1,
	}).verify(t)

	op("halt").with(Z80 {}).expects(Z80 {
		Halted: true,
	}).verify(t)

	op("nop").with(Z80 {
		Reg: register { A: 1, B: 2, },
		Cycles: 0,
	}).expects(Z80 {
		Reg: register { A: 1, B: 2, },
		Cycles: 0,
	}).verify(t)
}

type actual struct {
	op, testname string
	actual, expected Z80
}

func op(opcode string) *actual {
	return &actual{
		op: opcode,
		testname: opcode,
	}
}

func (a *actual) test(testname string) *actual {
	a.testname = testname
	return a
}

func (a *actual) with(z80 Z80) *actual {
	a.actual = z80
	return a
}

func (a *actual) expectsSame() *actual {
	a.expected = a.actual;
	return a;
}

func (a *actual) expects(z80 Z80) *actual {
	a.expected = z80
	return a
}

func (a *actual) verify(t *testing.T) {
	t.Run(a.testname, func(t *testing.T) {
		a.actual.execute(a.op)

		if a.actual.Cycles != a.expected.Cycles {
			t.Errorf("Cycles do not match; expected %d but got %d", a.expected.Cycles, a.actual.Cycles)
		}
		if a.actual.Halted != a.expected.Halted {
			t.Errorf("Halted flag does not match, expected %t but got %t", a.expected.Halted, a.actual.Halted)
		}
		if a.actual.Reg.H != a.expected.Reg.H {
			t.Errorf("Reg flag H does not match; expected %d but got %d", a.expected.Reg.H, a.actual.Reg.H)
		}
		if a.actual.Reg.C != a.expected.Reg.C {
			t.Errorf("Reg flag C does not match; expected %d but got %d", a.expected.Reg.C, a.actual.Reg.C)
		}
		if a.actual.Reg.A != a.expected.Reg.A {
			t.Errorf("Reg flag A does not match; expected %d but got %d", a.expected.Reg.A, a.actual.Reg.A)
		}
		if a.actual.Reg.B != a.expected.Reg.B {
			t.Errorf("Reg flag B does not match; expected %d but got %d", a.expected.Reg.B, a.actual.Reg.B)
		}
		if a.actual.Reg.D != a.expected.Reg.D {
			t.Errorf("Reg flag D does not match; expected %d but got %d", a.expected.Reg.D, a.actual.Reg.D)
		}
		if a.actual.Reg.E != a.expected.Reg.E {
			t.Errorf("Reg flag E does not match; expected %d but got %d", a.expected.Reg.E, a.actual.Reg.E)
		}
		if a.actual.Reg.F != a.expected.Reg.F {
			t.Errorf("Reg flag F does not match; expected %d but got %d", a.expected.Reg.F, a.actual.Reg.F)
		}
		if a.actual.Reg.I != a.expected.Reg.I {
			t.Errorf("Reg flag I does not match; expected %d but got %d", a.expected.Reg.I, a.actual.Reg.I)
		}
		if a.actual.Reg.L != a.expected.Reg.L {
			t.Errorf("Reg flag L does not match; expected %d but got %d", a.expected.Reg.L, a.actual.Reg.L)
		}
		if a.actual.Reg.PC != a.expected.Reg.PC {
			t.Errorf("Reg flag PC does not match; expected %d but got %d", a.expected.Reg.PC, a.actual.Reg.PC)
		}
		if a.actual.Reg.SP != a.expected.Reg.SP {
			t.Errorf("Reg flag SP does not match; expected %d but got %d", a.expected.Reg.SP, a.actual.Reg.SP)
		}
	})
}