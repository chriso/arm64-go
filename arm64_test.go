package arm64

import "testing"

func TestARM64(t *testing.T) {
	for _, test := range []struct {
		raw uint32
		ins Instruction
		str string
	}{
		// Add/Subtract/Negate with Carry (ADC, ADCS, SBC, SBCS, NGC, NGCS)
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/ADC--Add-with-Carry-
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/ADCS--Add-with-Carry--setting-flags-
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/SBC--Subtract-with-Carry-
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/SBCS--Subtract-with-Carry--setting-flags-
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/NGC--Negate-with-Carry--an-alias-of-SBC-
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/NGCS--Negate-with-Carry--setting-flags--an-alias-of-SBCS-
		// 0b_{sf}_{op}_{S}_{11010000}_{Rm(5)}_{000000}_{Rn(5)}_{Rd(5)}
		{0b_0_0_0_11010000_00010_000000_00001_00000, New(ADC, W0, W1, W2), "adc w0, w1, w2"},
		{0b_0_0_1_11010000_00010_000000_00001_00000, New(ADCS, W0, W1, W2), "adcs w0, w1, w2"},
		{0b_1_0_0_11010000_11101_000000_11111_10101, New(ADC, X21, XZR, X29), "adc x21, xzr, x29"},
		{0b_0_1_0_11010000_00111_000000_00101_00011, New(SBC, W3, W5, W7), "sbc w3, w5, w7"},
		{0b_1_1_1_11010000_00011_000000_00010_00000, New(SBCS, X0, X2, X3), "sbcs x0, x2, x3"},
		{0b_1_1_0_11010000_00010_000000_00001_11111, New(NGC, X1, X2), "ngc x1, x2"},
		{0b_0_1_1_11010000_00010_000000_00001_11111, New(NGCS, W1, W2), "ngcs w1, w2"},

		// Add/Subtract (ADD, ADDS, SUB, SUBS) - Extended Register
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/ADD--extended-register---Add--extended-register--
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/ADDS--extended-register---Add--extended-register---setting-flags-
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/SUB--extended-register---Subtract--extended-register--
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/SUBS--extended-register---Subtract--extended-register---setting-flags-
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/CMP--extended-register---Compare--extended-register---an-alias-of-SUBS--extended-register--
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/CMN--extended-register---Compare-Negative--extended-register---an-alias-of-ADDS--extended-register--
		// 0b_{sf}_{op}_{S}_{01011001}_{Rm(5)}_{option(3)}_{imm3(3)}_{Rn(5)}_{Rd(5)}
		{0b_0_0_0_01011001_00011_000_000_00010_00001, New(ADD, W1, W2, W3.ExtShift(UXTB, 0)), "add w1, w2, w3, uxtb"},
		{0b_1_0_1_01011001_00011_001_011_00010_00001, New(ADDS, X1, X2, X3.ExtShift(UXTH, 3)), "adds x1, x2, x3, uxth #0x3"},
		{0b_1_0_0_01011001_00011_010_010_00010_00001, New(ADD, X1, X2, X3.ExtShift(UXTW, 2)), "add x1, x2, x3, uxtw #0x2"},
		{0b_1_1_0_01011001_00011_011_001_00010_00001, New(SUB, X1, X2, X3.ExtShift(UXTX, 1)), "sub x1, x2, x3, uxtx #0x1"},
		{0b_0_0_0_01011001_00000_100_100_00000_00000, New(ADD, W0, W0, W0.ExtShift(SXTB, 4)), "add w0, w0, w0, sxtb #0x4"},
		{0b_0_0_0_01011001_00001_101_011_00001_00001, New(ADD, W1, W1, W1.ExtShift(SXTH, 3)), "add w1, w1, w1, sxth #0x3"},
		{0b_1_0_0_01011001_00010_110_010_00010_00010, New(ADD, X2, X2, X2.ExtShift(SXTW, 2)), "add x2, x2, x2, sxtw #0x2"},
		{0b_0_0_0_01011001_11111_010_100_11111_11111, New(ADD, WSP, WSP, WZR.ExtShift(LSL, 4)), "add wsp, wsp, wzr, lsl #0x4"},
		{0b_1_0_0_01011001_11111_010_100_11111_11111, New(ADD, SP, SP, XZR.ExtShift(LSL, 4)), "add sp, sp, xzr, lsl #0x4"},
		{0b_1_1_1_01011001_00011_111_001_11111_11111, New(CMP, SP, X3.ExtShift(SXTX, 1)), "cmp sp, x3, sxtx #0x1"},
		{0b_0_0_1_01011001_11111_010_000_11111_11111, New(CMN, WSP, WZR.ExtShift(UXTW, 0)), "cmn wsp, wzr, uxtw"},

		// Add/Subtract (ADD, ADDS, SUB, SUBS) - Immediate
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/ADD--immediate---Add--immediate--
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/ADDS--immediate---Add--immediate---setting-flags-
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/SUB--immediate---Subtract--immediate--
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/SUBS--immediate---Subtract--immediate---setting-flags-
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/MOV--to-from-SP---Move-between-register-and-stack-pointer--an-alias-of-ADD--immediate--
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/CMP--immediate---Compare--immediate---an-alias-of-SUBS--immediate--
		// https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/CMN--immediate---Compare-Negative--immediate---an-alias-of-ADDS--immediate--
		// 0b_{sf}_{op}_{S}_{100010}_{sh}_{imm12}_{Rn(5)}_{Rd(5)}
		{0b_1_0_0_100010_0_111111111111_00010_00001, New(ADD, X1, X2, Immediate(4095)), "add x1, x2, #0xfff"},
		{0b_1_0_1_100010_1_111111111111_00010_00001, New(ADDS, X1, X2, Immediate(4095).ExtShift(LSL, 12)), "adds x1, x2, #0xfff, lsl #0xc"},
		{0b_0_1_1_100010_0_111111111111_00010_00001, New(SUBS, W1, W2, Immediate(4095)), "subs w1, w2, #0xfff"},
		{0b_1_1_0_100010_1_000000000001_00010_00001, New(SUB, X1, X2, Immediate(1).ExtShift(LSL, 12)), "sub x1, x2, #0x1, lsl #0xc"},
		{0b_1_0_0_100010_0_000000000010_11111_11111, New(ADD, SP, SP, Immediate(2)), "add sp, sp, #0x2"},
		{0b_1_0_0_100010_0_000000000000_00001_11111, New(MOV, SP, X1), "mov sp, x1"},
		{0b_0_0_0_100010_0_000000000000_11111_00010, New(MOV, W2, WSP), "mov w2, wsp"},
		{0b_0_1_1_100010_0_111111111111_00001_11111, New(CMP, W1, Immediate(4095)), "cmp w1, #0xfff"},
		{0b_1_1_1_100010_1_000000000001_00010_11111, New(CMP, X2, Immediate(1).ExtShift(LSL, 12)), "cmp x2, #0x1, lsl #0xc"},
		{0b_0_0_1_100010_0_111111111111_00001_11111, New(CMN, W1, Immediate(4095)), "cmn w1, #0xfff"},
		{0b_1_0_1_100010_1_000000000001_00010_11111, New(CMN, X2, Immediate(1).ExtShift(LSL, 12)), "cmn x2, #0x1, lsl #0xc"},
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
