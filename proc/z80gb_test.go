
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
