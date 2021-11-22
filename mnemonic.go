package arm64

type Mnemonic int

const (
	InvalidMnemonic Mnemonic = iota
	ADC
	ADCS
	ADD
	ADDS
	NGC
	NGCS
	SBC
	SBCS
	SUB
	SUBS
)

func (m Mnemonic) String() string {
	if m >= 0 && int(m) < len(mnemonicStrings) {
		return mnemonicStrings[m]
	}
	return invalid
}

var mnemonicStrings = []string{
	invalid,
	"adc",
	"adcs",
	"add",
	"adds",
	"ngc",
	"ngcs",
	"sbc",
	"sbcs",
	"sub",
	"subs",
}
