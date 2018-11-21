package proc

import "testing"

func TestAlu(t *testing.T) {

	t.Run("Add 8bit, no carry", func(t *testing.T) {
		alu := ALU {
			A: 1,
			B: 2,
			Operation: "Add",
		}

		result := alu.Process()

		if alu.FlagsOut.C == true {
			t.Errorf("No carry expected")
		}

		if alu.FlagsOut.Z == true {
			t.Errorf("No zero expected")
		}

		if alu.Z != 3 {
			t.Errorf("expected result of A + B to be 3, but got %d", alu.Z)
		}

		if result != 1 {
			t.Errorf("expected one cycle to have passed for ALU operation - actual cycles: %d", result)
		}
	})

	t.Run("Add 8bit, carry", func(t *testing.T) {
		alu := ALU {
			A: bit8 - 2,
			B: 3,
			Operation: "Add",
		}

		alu.Process()

		if alu.FlagsOut.C == false {
			t.Errorf("Carry expected")
		}

		if alu.Z != 0 {
			t.Errorf("expected result of A + B to be 0, but got %d", alu.Z)
		}
	})

	t.Run("Add 8bit, no carry but max 8bit value", func(t *testing.T) {
		alu := ALU {
			A: bit8 - 2,
			B: 2,
			Operation: "Add",
		}

		alu.Process()

		if alu.FlagsOut.C == true {
			t.Errorf("No arry expected")
		}

		if alu.Z != bit8 {
			t.Errorf("expected result of A + B to be max 8bit val, but got %d", alu.Z)
		}
	})

	t.Run("Add 8bit, zero", func(t *testing.T) {
		alu := ALU {
			A: 0,
			B: 0,
			Operation: "Add",
		}

		alu.Process()
		if alu.FlagsOut.Z == false {
			t.Errorf("Zero expected")
		}
	})

}
