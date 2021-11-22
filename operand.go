package arm64

type Operand interface {
	String() string

	operand()
}
