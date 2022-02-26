package arm64

import (
	"fmt"
	"strconv"
)

const invalid = "<invalid>"

type Operand interface {
	String() string

	operand()
}

type RegisterType uint32

const (
	_ RegisterType = iota
	W
	X
)

func (t RegisterType) GP() bool {
	return t == W || t == X
}

func (t RegisterType) At(i uint32) Register {
	if i < 0 || i >= 32 {
		return InvalidRegister
	}
	return Register(uint32(t)<<5 + i)
}

func (t RegisterType) String() string {
	switch t {
	case W:
		return "w"
	case X:
		return "x"
	default:
		return invalid
	}
}

type Register uint32

const (
	W0 = Register(W<<5) + iota
	W1
	W2
	W3
	W4
	W5
	W6
	W7
	W8
	W9
	W10
	W11
	W12
	W13
	W14
	W15
	W16
	W17
	W18
	W19
	W20
	W21
	W22
	W23
	W24
	W25
	W26
	W27
	W28
	W29
	W30
	WZR
)

const (
	X0 = Register(X<<5) + iota
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	X29
	X30
	XZR
)

const (
	InvalidRegister Register = 0

	spFlag = 1 << 31
	WSP    = WZR | spFlag
	SP     = XZR | spFlag
)

func (Register) operand() {}

func (r Register) Type() RegisterType {
	return RegisterType(r>>5) & 7
}

func (r Register) Index() uint32 {
	return uint32(r) & 31
}

func (r Register) String() string {
	switch r {
	case WSP:
		return "wsp"
	case SP:
		return "sp"
	case WZR:
		return "wzr"
	case XZR:
		return "xzr"
	case InvalidRegister:
		return invalid
	default:
		return r.Type().String() + strconv.Itoa(int(r.Index()))
	}
}

func (r Register) SP() Register {
	return r | spFlag
}

func (r Register) ExtShift(e ExtShift, amount uint8) ExtShiftRegister {
	return ExtShiftRegister{r, e, amount}
}

type ExtShiftRegister struct {
	Register
	ExtShift ExtShift
	Amount   uint8
}

func (e ExtShiftRegister) operand() {}

func (e ExtShiftRegister) String() string {
	str := e.Register.String() + ", " + e.ExtShift.String()
	if e.Amount > 0 {
		str += fmt.Sprintf(" #%#x", e.Amount)
	}
	return str
}

type ExtShift uint8

const (
	UXTB ExtShift = iota
	UXTH
	UXTW
	UXTX
	SXTB
	SXTH
	SXTW
	SXTX
	LSL
	LSR
	ASR
	ROR
)

func (e ExtShift) String() string {
	switch e {
	case UXTB:
		return "uxtb"
	case UXTH:
		return "uxth"
	case UXTW:
		return "uxtw"
	case UXTX:
		return "uxtx"
	case SXTB:
		return "sxtb"
	case SXTH:
		return "sxth"
	case SXTW:
		return "sxtw"
	case SXTX:
		return "sxtx"
	case LSL:
		return "lsl"
	case LSR:
		return "lsr"
	case ASR:
		return "ast"
	case ROR:
		return "ror"
	default:
		return invalid
	}
}

type Immediate int64

func (i Immediate) operand() {}

func (i Immediate) ExtShift(e ExtShift, amount uint8) ExtShiftImmediate {
	return ExtShiftImmediate{i, e, amount}
}

func (i Immediate) String() string {
	return fmt.Sprintf("#%#x", int64(i))
}

type ExtShiftImmediate struct {
	Immediate
	ExtShift ExtShift
	Amount   uint8
}

func (e ExtShiftImmediate) operand() {}

func (e ExtShiftImmediate) String() string {
	str := e.Immediate.String() + ", " + e.ExtShift.String()
	if e.Amount > 0 {
		str += fmt.Sprintf(" #%#x", e.Amount)
	}
	return str
}
