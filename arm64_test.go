package arm64

import "testing"

func TestARM64(t *testing.T) {
	for _, test := range []struct {
		raw uint32
		ins Instruction
		str string
	}{
		{0b_0_0_0_11010000_00010_000000_00001_00000, New(ADC).WithOperands(W0, W1, W2), "adc w0, w1, w2"},
		{0b_0_0_1_11010000_00010_000000_00001_00000, New(ADCS).WithOperands(W0, W1, W2), "adcs w0, w1, w2"},
		{0b_1_0_0_11010000_11101_000000_11111_10101, New(ADC).WithOperands(X21, XZR, X29), "adc x21, xzr, x29"},
		{0b_0_1_0_11010000_00111_000000_00101_00011, New(SBC).WithOperands(W3, W5, W7), "sbc w3, w5, w7"},
		{0b_1_1_1_11010000_00011_000000_00010_00000, New(SBCS).WithOperands(X0, X2, X3), "sbcs x0, x2, x3"},
		{0b_1_1_0_11010000_00010_000000_00001_11111, New(NGC).WithOperands(X1, X2), "ngc x1, x2"},
		{0b_0_1_1_11010000_00010_000000_00001_11111, New(NGCS).WithOperands(W1, W2), "ngcs w1, w2"},
	} {
		t.Run(test.str, func(t *testing.T) {
			ins, err := Disassemble(test.raw)
			if err != nil {
				t.Error(err)
			}
			assertInstruction(t, &ins, test.str, test.raw)
			assertInstruction(t, &test.ins, test.str, test.raw)
		})
	}
}

func assertInstruction(t *testing.T, i *Instruction, expectString string, expectEncoded uint32) {
	if str := i.String(); str != expectString {
		t.Errorf("unexpected string: got %v, expected %v", str, expectString)
	}
	encoded, err := Assemble(i)
	if err != nil {
		t.Errorf("failed to assemble %s: %v", expectString, err)
	}
	if encoded != expectEncoded {
		t.Errorf("unexpected encoding: got %#08x, expected %#08x", encoded, expectEncoded)
	}
}
