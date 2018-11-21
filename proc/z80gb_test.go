
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

func TestExecuteOperations(t *testing.T) {

	t.Run("ALU output sampled into CPU registers", func(t *testing.T) {
		cpu := setup()
		cpu.Reg.A = 254
		cpu.Reg.B = 3
		cpu.execute("add_a_b")

		if cpu.Reg.F == 0 {
			t.Errorf("expected flags to be sampled into register, but was 0")
		}

		if cpu.Reg.A != 1 {
			t.Errorf("expected result to be sampled into register A (2), but was %d", cpu.Reg.A)
		}
	})

	t.Run("some ALU operation integration test", func(t *testing.T) {
		cpu := setup()
		cpu.execute("add_a_b")

		if cpu.Cycles != 1 {
			t.Errorf("expected one cycle to have passed for ALU operation - actual cycles: %d", cpu.Cycles)
		}
	})

	t.Run("ld_ab", func(t *testing.T) {
		cpu := setup()
		cpu.execute("ld_a_b")

		if cpu.Reg.A != 2 {
			t.Errorf("expected reg A to be loaded with value of B, instead: %d", cpu.Reg.A)
		}

		if cpu.Reg.B != 2 {
			t.Errorf("expected reg B to be unmodified, instead: %d", cpu.Reg.B)
		}

		if cpu.Cycles != 1 {
			t.Errorf("expected one cycle to have passed with NOP operation - actual cycles: %d", cpu.Cycles)
		}
	})

	t.Run("halt", func(t *testing.T) {
		cpu := setup()
		cpu.execute("halt")

		if cpu.Halted != true {
			t.Errorf("expected cpu to have halted, but did not")
		}
	})

	t.Run("nop", func(t *testing.T) {
		cpu := setup()
		cpu.execute("nop")

		if cpu.Cycles != 1 {
			t.Errorf("expected one cycle to have passed with NOP operation - actual cycles: %d", cpu.Cycles)
		}		
	})


}
