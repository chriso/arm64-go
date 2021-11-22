package arm64

type Mnemonic int

const (
	InvalidMnemonic Mnemonic = iota
	ADC
	ADCS
	SBC
	SBCS
	NGC
	NGCS
)

func (m Mnemonic) String() string {
	if m >= 0 && int(m) < len(mnemonicStrings) {
		return mnemonicStrings[m]
	}
	return "<invalid>"
}

var mnemonicStrings = []string{
	"<invalid>",
	"adc",
	"adcs",
	"sbc",
	"sbcs",
	"ngc",
	"ngcs",
}
