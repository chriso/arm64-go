package arm64

import (
	"errors"
	"fmt"
)

var ErrNotImplemented = errors.New("not implemented")
var ErrUndefined = errors.New("undefined")

// Disassemble disassembles a raw 32-bit instruction.
func Disassemble(ins uint32) (Instruction, error) {
	var d decoded
	if err := decode(ins, &d); err != nil {
		return Instruction{}, fmt.Errorf("failed to decode %#08x: %w", ins, err)
	}

	switch d.iclass {
	case iclass_addsub_carry:
		r := chooseRegister(d.sf, X, W)
		if d.Rd == 31 && d.op == 1 {
			m := chooseMnemonic(d.S, NGCS, NGC)
			return New(m, r.At(d.Rn), r.At(d.Rm)), nil
		} else {
			m := chooseMnemonic(d.op, chooseMnemonic(d.S, SBCS, SBC), chooseMnemonic(d.S, ADCS, ADC))
			return New(m, r.At(d.Rd), r.At(d.Rn), r.At(d.Rm)), nil
		}
	case iclass_addsub_ext:
		r := chooseRegister(d.sf, X, W)
		e := ExtShift(d.option)
		amount := uint8(d.imm3)
		if amount > 4 {
			return Instruction{}, ErrUndefined
		}
		if (d.Rd == 31 || d.Rn == 31) && e == UXTW && amount != 0 {
			e = LSL
		}
		if d.S == 1 && d.Rd == 31 {
			m := chooseMnemonic(d.op, CMP, CMN)
			return New(m, r.At(d.Rn).SP(), r.At(d.Rm).ExtShift(e, amount)), nil
		}
		m := chooseMnemonic(d.op, chooseMnemonic(d.S, SUBS, SUB), chooseMnemonic(d.S, ADDS, ADD))
		return New(m, r.At(d.Rd).SP(), r.At(d.Rn).SP(), r.At(d.Rm).ExtShift(e, amount)), nil
	case iclass_addsub_imm:
		r := chooseRegister(d.sf, X, W)
		var op Operand
		if d.sh == 1 {
			op = Immediate(d.imm12).ExtShift(LSL, 12)
		} else if d.imm12 == 0 && (d.Rd == 31 || d.Rn == 31) {
			return New(MOV, r.At(d.Rd).SP(), r.At(d.Rn).SP()), nil
		} else {
			op = Immediate(d.imm12)
		}
		if d.S == 1 && d.Rd == 31 {
			m := chooseMnemonic(d.op, CMP, CMN)
			return New(m, r.At(d.Rn).SP(), op), nil
		}
		m := chooseMnemonic(d.op, chooseMnemonic(d.S, SUBS, SUB), chooseMnemonic(d.S, ADDS, ADD))
		return New(m, r.At(d.Rd).SP(), r.At(d.Rn).SP(), op), nil
	default:
		return Instruction{}, fmt.Errorf("failed to diassemble %#08x: %w", ins, ErrNotImplemented)
	}
}

func chooseRegister(flag uint32, on, off RegisterType) RegisterType {
	if flag == 1 {
		return on
	}
	return off
}

func chooseMnemonic(flag uint32, on, off Mnemonic) Mnemonic {
	if flag == 1 {
		return on
	}
	return off
}
