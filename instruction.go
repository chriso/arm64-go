package arm64

type Instruction struct {
	Mnemonic Mnemonic
	Operands [MaxOperands]Operand
}

const MaxOperands = 3

func New(mnemonic Mnemonic) Instruction {
	return Instruction{Mnemonic: mnemonic}
}

func (i Instruction) WithOperands(operands ...Operand) Instruction {
	for n, operand := range operands {
		if n >= MaxOperands {
			panic("exceeded max operands")
		}
		i.Operands[n] = operand
	}
	return i
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
