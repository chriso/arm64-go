package arm64

import "strconv"

type RegisterType uint32

const (
	InvalidRegisterType RegisterType = iota
	W
	X
)

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
		return "<invalid>"
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

	WSP = RegisterSP(WZR)
	SP  = RegisterSP(XZR)
)

func (Register) operand() {}

func (r Register) Type() RegisterType {
	return RegisterType(r >> 5)
}

func (r Register) Index() uint32 {
	return uint32(r) & 31
}

func (r Register) String() string {
	switch r {
	case WZR:
		return "wzr"
	case XZR:
		return "xzr"
	case InvalidRegister:
		return "<invalid>"
	default:
		return r.Type().String() + strconv.Itoa(int(r.Index()))
	}
}

func (r Register) SP() RegisterSP {
	return RegisterSP(r)
}

type RegisterSP Register

func (RegisterSP) operand() {}

func (sp RegisterSP) String() string {
	switch sp {
	case WSP:
		return "wsp"
	case SP:
		return "sp"
	default:
		return Register(sp).String()
	}
}
