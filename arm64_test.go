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

		{0b_0_0_0_01011001_00011_000_000_00010_00001, New(ADD).WithOperands(W1, W2, W3.ExtShift(UXTB, 0)), "add w1, w2, w3, uxtb"},
		{0b_1_0_1_01011001_00011_001_011_00010_00001, New(ADDS).WithOperands(X1, X2, X3.ExtShift(UXTH, 3)), "adds x1, x2, x3, uxth #0x3"},
		{0b_1_0_0_01011001_00011_010_010_00010_00001, New(ADD).WithOperands(X1, X2, X3.ExtShift(UXTW, 2)), "add x1, x2, x3, uxtw #0x2"},
		{0b_1_1_0_01011001_00011_011_001_00010_00001, New(SUB).WithOperands(X1, X2, X3.ExtShift(UXTX, 1)), "sub x1, x2, x3, uxtx #0x1"},
		{0b_0_0_0_01011001_00000_100_100_00000_00000, New(ADD).WithOperands(W0, W0, W0.ExtShift(SXTB, 4)), "add w0, w0, w0, sxtb #0x4"},
		{0b_0_0_0_01011001_00001_101_011_00001_00001, New(ADD).WithOperands(W1, W1, W1.ExtShift(SXTH, 3)), "add w1, w1, w1, sxth #0x3"},
		{0b_1_0_0_01011001_00010_110_010_00010_00010, New(ADD).WithOperands(X2, X2, X2.ExtShift(SXTW, 2)), "add x2, x2, x2, sxtw #0x2"},
		{0b_1_1_1_01011001_00011_111_001_11111_11111, New(SUBS).WithOperands(SP, SP, X3.ExtShift(SXTX, 1)), "subs sp, sp, x3, sxtx #0x1"},
		{0b_0_0_1_01011001_11111_010_000_11111_11111, New(ADDS).WithOperands(WSP, WSP, WZR.ExtShift(UXTW, 0)), "adds wsp, wsp, wzr, uxtw"},
		{0b_0_0_0_01011001_11111_010_100_11111_11111, New(ADD).WithOperands(WSP, WSP, WZR.ExtShift(LSL, 4)), "add wsp, wsp, wzr, lsl #0x4"},
		{0b_1_0_0_01011001_11111_010_100_11111_11111, New(ADD).WithOperands(SP, SP, XZR.ExtShift(LSL, 4)), "add sp, sp, xzr, lsl #0x4"},
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
