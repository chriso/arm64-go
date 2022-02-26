package arm64

import (
	"errors"
)

var ErrInvalidOperand = errors.New("invalid operand")

// Assemble assembles an instruction.
func Assemble(i *Instruction) (uint32, error) {
	switch m := i.Mnemonic; m {
	case ADC, ADCS, SBC, SBCS:
		rd, rn, rm, err := unpackRegister3(i)
		if err != nil {
			return 0, err
		}
		if i.Operands[3] != nil || !rn.Type().GP() {
			return 0, ErrInvalidOperand
		}
		sf := flag(rd.Type() == X)
		op := flag(m == SBC || m == SBCS)
		s := flag(m == ADCS || m == SBCS)
		return encode_addsub_carry(sf, op, s, rm.Index(), rn.Index(), rd.Index())
	case ADD, ADDS, SUB, SUBS:
		rd, rn, err := unpackRegister2(i)
		if err != nil {
			return 0, err
		}
		if i.Operands[3] != nil {
			return 0, ErrInvalidOperand
		}
		sf := flag(rd.Type() == X)
		op := flag(m == SUB || m == SUBS)
		s := flag(m == ADDS || m == SUBS)
		switch r := i.Operands[2].(type) {
		case ExtShiftRegister:
			if !rn.Type().GP() || r.Amount > 4 {
				return 0, ErrInvalidOperand
			}
			option := r.ExtShift
			if option == LSL {
				option = UXTW
			}
			return encode_addsub_ext(sf, op, s, 0, r.Index(), uint32(option), uint32(r.Amount), rn.Index(), rd.Index())
		case Immediate:
			if r < 0 || r > 0xFFF {
				return 0, ErrInvalidOperand
			}
			return encode_addsub_imm(sf, op, s, 0, uint32(r), rn.Index(), rd.Index())
		case ExtShiftImmediate:
			if r.Immediate < 0 || r.Immediate > 0xFFF {
				return 0, ErrInvalidOperand
			}
			if r.ExtShift != LSL || r.Amount != 12 {
				return 0, ErrInvalidOperand
			}
			return encode_addsub_imm(sf, op, s, 1, uint32(r.Immediate), rn.Index(), rd.Index())
		default:
			return 0, ErrInvalidOperand
		}
	case CMN, CMP:
		rn, err := unpackRegister(i.Operands[0])
		if err != nil {
			return 0, err
		}
		sf := flag(rn.Type() == X)
		op := flag(m == CMP)
		switch r := i.Operands[1].(type) {
		case ExtShiftRegister:
			if !rn.Type().GP() || r.Amount > 4 {
				return 0, ErrInvalidOperand
			}
			option := r.ExtShift
			if option == LSL {
				option = UXTW
			}
			return encode_addsub_ext(sf, op, 1, 0, r.Index(), uint32(option), uint32(r.Amount), rn.Index(), 31)
		case Immediate:
			if r < 0 || r > 0xFFF {
				return 0, ErrInvalidOperand
			}
			return encode_addsub_imm(sf, op, 1, 0, uint32(r), rn.Index(), 31)
		case ExtShiftImmediate:
			if r.Immediate < 0 || r.Immediate > 0xFFF {
				return 0, ErrInvalidOperand
			}
			if r.ExtShift != LSL || r.Amount != 12 {
				return 0, ErrInvalidOperand
			}
			return encode_addsub_imm(sf, op, 1, 1, uint32(r.Immediate), rn.Index(), 31)
		default:
			return 0, ErrNotImplemented
		}
	case MOV:
		rd, rn, err := unpackRegister2(i)
		if err != nil {
			return 0, err
		}
		sf := flag(rd.Type() == X)
		if rd.Index() == 31 || rn.Index() == 31 {
			return encode_addsub_imm(sf, 0, 0, 0, 0, rn.Index(), rd.Index())
		}
		return 0, ErrNotImplemented
	case NGC, NGCS:
		rn, rm, err := unpackRegister2(i)
		if err != nil {
			return 0, err
		}
		if i.Operands[2] != nil || !rn.Type().GP() {
			return 0, ErrInvalidOperand
		}
		sf := flag(rn.Type() == X)
		s := flag(m == NGCS)
		return encode_addsub_carry(sf, 1, s, rm.Index(), rn.Index(), 31)
	}

	return 0, ErrNotImplemented
}

func flag(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

func unpackRegister(op Operand) (Register, error) {
	switch r := op.(type) {
	case Register:
		return r, nil
	case ExtShiftRegister:
		return r.Register, nil
	default:
		return 0, ErrInvalidOperand
	}
}

func unpackRegister2(i *Instruction) (a, b Register, err error) {
	if a, err = unpackRegister(i.Operands[0]); err != nil {
		return
	}
	if b, err = unpackRegister(i.Operands[1]); err != nil {
		return
	}
	if a.Type() != b.Type() {
		err = ErrInvalidOperand
	}
	return
}

func unpackRegister3(i *Instruction) (a, b, c Register, err error) {
	if a, err = unpackRegister(i.Operands[0]); err != nil {
		return
	}
	if b, err = unpackRegister(i.Operands[1]); err != nil {
		return
	}
	if c, err = unpackRegister(i.Operands[2]); err != nil {
		return
	}
	if a.Type() != b.Type() || b.Type() != c.Type() {
		err = ErrInvalidOperand
	}
	return
}
