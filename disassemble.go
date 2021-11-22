package arm64

import (
	"errors"
	"fmt"
)

var ErrNotImplemented = errors.New("not implemented")

func Disassemble(ins uint32) (Instruction, error) {
	var d decoded
	if err := decode(ins, &d); err != nil {
		return Instruction{}, fmt.Errorf("failed to decode %#08x: %w", ins, err)
	}

	switch d.iclass {
	case iclass_addsub_carry:
		r := flagr(d.sf, W, X)
		if d.Rd == 31 && d.op == 1 {
			m := flagm(d.S, NGC, NGCS)
			return New(m).WithOperands(r.At(d.Rn), r.At(d.Rm)), nil
		} else {
			m := flagm(d.op, flagm(d.S, ADC, ADCS), flagm(d.S, SBC, SBCS))
			return New(m).WithOperands(r.At(d.Rd), r.At(d.Rn), r.At(d.Rm)), nil
		}
	default:
		return Instruction{}, fmt.Errorf("failed to diassemble %#08x: %w", ins, ErrNotImplemented)
	}
}

func flagr(flag uint32, off, on RegisterType) RegisterType {
	if flag == 1 {
		return on
	}
	return off
}

func flagm(flag uint32, off, on Mnemonic) Mnemonic {
	if flag == 1 {
		return on
	}
	return off
}
