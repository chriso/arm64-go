package arm64

type Instruction struct {
	Mnemonic Mnemonic
	Operands [MaxOperands]Operand
}

const MaxOperands = 4

func New(mnemonic Mnemonic, operands ...Operand) (ins Instruction) {
	ins.Mnemonic = mnemonic
	for n, operand := range operands {
		if n >= MaxOperands {
			panic("exceeded max operands")
		}
		ins.Operands[n] = operand
	}
	return
}

func (i Instruction) String() string {
	s := i.Mnemonic.String()
	for i, op := range i.Operands {
		if op == nil {
			break
		}
		if i == 0 {
		} else {
			s += ","
		}
		s += " "
		s += op.String()
	}
	return s
}
