package util

import "testing"

type someTest struct {
	Prop1, Prop2 int
}

func TestReflectUtil(t *testing.T) {
	testInstance := someTest{ 1, 2}

	t.Run("StrToFields converts string to callable reflective fields", func(t *testing.T) {

		fields := StrToFields(&testInstance, "Prop1", "Prop2")

		if fields[0].Int() != 1 {
			t.Errorf("Unable to access Prop1 with value 1")
		}
		if fields[1].Int() != 2 {
			t.Errorf("Unable to access Prop2 with value 2")
		}
	})
}