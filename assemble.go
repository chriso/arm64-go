package arm64

import (
	"errors"
)

var ErrInvalidOperand = errors.New("invalid operand")

func Assemble(i *Instruction) (uint32, error) {
	switch m := i.Mnemonic; m {
	case ADC, ADCS, SBC, SBCS:
		rd, rn, rm, err := unpackreg3(i)
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
		rd, rn, rm, err := unpackreg3(i)
		if err != nil {
			return 0, err
		}
		e, ok := i.Operands[2].(ExtShiftRegister)
		if !ok || i.Operands[3] != nil || !rn.Type().GP() || e.Amount > 4 {
			return 0, ErrInvalidOperand
		}
		option := e.ExtShift
		if option == LSL {
			option = UXTW
		}
		sf := flag(rd.Type() == X)
		op := flag(m == SUB || m == SUBS)
		s := flag(m == ADDS || m == SUBS)
		return encode_addsub_ext(sf, op, s, 0, rm.Index(), uint32(option), uint32(e.Amount), rn.Index(), rd.Index())
	case NGC, NGCS:
		rn, rm, err := unpackreg2(i)
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

func unpackreg(op Operand) (Register, error) {
	switch r := op.(type) {
	case Register:
		return r, nil
	case ExtShiftRegister:
		return r.Register, nil
	default:
		return 0, ErrInvalidOperand
	}
}

func unpackreg2(i *Instruction) (a, b Register, err error) {
	if a, err = unpackreg(i.Operands[0]); err != nil {
		return
	}
	if b, err = unpackreg(i.Operands[1]); err != nil {
		return
	}
	if a.Type() != b.Type() {
		err = ErrInvalidOperand
	}
	return
}

func unpackreg3(i *Instruction) (a, b, c Register, err error) {
	if a, err = unpackreg(i.Operands[0]); err != nil {
		return
	}
	if b, err = unpackreg(i.Operands[1]); err != nil {
		return
	}
	if c, err = unpackreg(i.Operands[2]); err != nil {
		return
	}
	if a.Type() != b.Type() || b.Type() != c.Type() {
		err = ErrInvalidOperand
	}
	return
}
