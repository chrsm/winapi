package user

type (
	ModKey     uint
	VirtualKey uint
)

const (
	ModAlt      ModKey = 0x0001
	ModControl         = 0x0002
	ModNoRepeat        = 0x4000
	ModShift           = 0x0004
	ModWin             = 0x0008
	ModNone            = 0x0000
)

const (
	VirtKeyZero  VirtualKey = 0x30
	VirtKeyOne              = 0x31
	VirtKeyTwo              = 0x32
	VirtKeyThree            = 0x33
	VirtKeyFour             = 0x34
	VirtKeyFive             = 0x35
	VirtKeySix              = 0x36
	VirtKeySeven            = 0x37
	VirtKeyEight            = 0x38
	VirtKeyNine             = 0x39

	VirtKeyBackspace  = 0x08
	VirtKeyTab        = 0x09
	VirtKeyEnter      = 0x0D
	VirtKeyShift      = 0x10
	VirtKeyCtrl       = 0x11
	VirtKeyAlt        = 0x12
	VirtKeyPauseBreak = 0x13
	VirtKeyCapsLock   = 0x14
	VirtKeyEscape     = 0x1B
	VirtKeyPageUp     = 0x21
	VirtKeyPageDown   = 0x22
	VirtKeyEnd        = 0x23
	VirtKeyHome       = 0x24
	VirtKeyLeftArrow  = 0x25
	VirtKeyUpArrow    = 0x26
	VirtKeyRightArrow = 0x27
	VirtKeyDownArrow  = 0x28
	VirtKeyInsert     = 0x2D
	VirtKeyDelete     = 0x2E

	VirtKeyA = 0x41
	VirtKeyB = 0x42
	VirtKeyC = 0x43
	VirtKeyD = 0x44
	VirtKeyE = 0x45
	VirtKeyF = 0x46
	VirtKeyG = 0x47
	VirtKeyH = 0x48
	VirtKeyI = 0x49
	VirtKeyJ = 0x4A
	VirtKeyK = 0x4B
	VirtKeyL = 0x4C
	VirtKeyM = 0x4D
	VirtKeyN = 0x4E
	VirtKeyO = 0x4F
	VirtKeyP = 0x50
	VirtKeyQ = 0x51
	VirtKeyR = 0x52
	VirtKeyS = 0x53
	VirtKeyT = 0x54
	VirtKeyU = 0x55
	VirtKeyV = 0x56
	VirtKeyW = 0x57
	VirtKeyX = 0x58
	VirtKeyY = 0x59
	VirtKeyZ = 0x5A

	VirtKeyLeftWin  = 0x5B
	VirtKeyRightWin = 0x5C
	VirtKeySelect   = 0x5D

	VirtKeyNumpadZero   = 0x60
	VirtKeyNumpadOne    = 0x61
	VirtKeyNumpadTwo    = 0x62
	VirtKeyNumpadThree  = 0x63
	VirtKeyNumpadFour   = 0x64
	VirtKeyNumpadFive   = 0x65
	VirtKeyNumpadSix    = 0x66
	VirtKeyNumpadSeven  = 0x67
	VirtKeyNumpadEight  = 0x68
	VirtKeyNumpadNine   = 0x69
	VirtKeyMultiply     = 0x6A
	VirtKeyAdd          = 0x6B
	VirtKeySubtract     = 0x6D
	VirtKeyDecimalPoint = 0x6E
	VirtKeyDivide       = 0x6F

	VirtKeyF1  = 0x70
	VirtKeyF2  = 0x71
	VirtKeyF3  = 0x72
	VirtKeyF4  = 0x73
	VirtKeyF5  = 0x74
	VirtKeyF6  = 0x75
	VirtKeyF7  = 0x76
	VirtKeyF8  = 0x77
	VirtKeyF9  = 0x78
	VirtKeyF10 = 0x79
	VirtKeyF11 = 0x7A
	VirtKeyF12 = 0x7B

	VirtKeyNumLock    = 0x90
	VirtKeyScrollLock = 0x91

	VirtKeySemiColon    = 0xBA
	VirtKeyEqualSign    = 0xBB
	VirtKeyComma        = 0xBC
	VirtKeyDash         = 0xBD
	VirtKeyPeriod       = 0xBE
	VirtKeyForwardSlash = 0xBF
	VirtKeyGraveAccent  = 0xC0
	VirtKeyOpenBracket  = 0xDB
	VirtKeyBackSlash    = 0xDC
	VirtKeyCloseBracket = 0xDD
	VirtKeySingleQuote  = 0xDE

	VirtKeyPrintScreen = 0x2C
	VirtKeySnapshot    = VirtKeyPrintScreen
)

var modKeyMap = map[string]ModKey{
	"Alt":      ModAlt,
	"Control":  ModControl,
	"NoRepeat": ModNoRepeat,
	"Shift":    ModShift,
	"Win":      ModWin,
	"None":     ModNone,
}

var virtualKeyMap = map[string]VirtualKey{
	"Zero":         VirtKeyZero,
	"One":          VirtKeyOne,
	"Two":          VirtKeyTwo,
	"Three":        VirtKeyThree,
	"Four":         VirtKeyFour,
	"Five":         VirtKeyFive,
	"Six":          VirtKeySix,
	"Seven":        VirtKeySeven,
	"Eight":        VirtKeyEight,
	"Nine":         VirtKeyNine,
	"Backspace":    VirtKeyBackspace,
	"Tab":          VirtKeyTab,
	"Enter":        VirtKeyEnter,
	"Shift":        VirtKeyShift,
	"Ctrl":         VirtKeyCtrl,
	"Alt":          VirtKeyAlt,
	"PauseBreak":   VirtKeyPauseBreak,
	"CapsLock":     VirtKeyCapsLock,
	"Escape":       VirtKeyEscape,
	"PageUp":       VirtKeyPageUp,
	"PageDown":     VirtKeyPageDown,
	"End":          VirtKeyEnd,
	"Home":         VirtKeyHome,
	"LeftArrow":    VirtKeyLeftArrow,
	"UpArrow":      VirtKeyUpArrow,
	"RightArrow":   VirtKeyRightArrow,
	"DownArrow":    VirtKeyDownArrow,
	"Insert":       VirtKeyInsert,
	"Delete":       VirtKeyDelete,
	"A":            VirtKeyA,
	"B":            VirtKeyB,
	"C":            VirtKeyC,
	"D":            VirtKeyD,
	"E":            VirtKeyE,
	"F":            VirtKeyF,
	"G":            VirtKeyG,
	"H":            VirtKeyH,
	"I":            VirtKeyI,
	"J":            VirtKeyJ,
	"K":            VirtKeyK,
	"L":            VirtKeyL,
	"M":            VirtKeyM,
	"N":            VirtKeyN,
	"O":            VirtKeyO,
	"P":            VirtKeyP,
	"Q":            VirtKeyQ,
	"R":            VirtKeyR,
	"S":            VirtKeyS,
	"T":            VirtKeyT,
	"U":            VirtKeyU,
	"V":            VirtKeyV,
	"W":            VirtKeyW,
	"X":            VirtKeyX,
	"Y":            VirtKeyY,
	"Z":            VirtKeyZ,
	"LeftWin":      VirtKeyLeftWin,
	"RightWin":     VirtKeyRightWin,
	"SelectKey":    VirtKeySelect,
	"NumpadZero":   VirtKeyNumpadZero,
	"NumpadOne":    VirtKeyNumpadOne,
	"NumpadTwo":    VirtKeyNumpadTwo,
	"NumpadThree":  VirtKeyNumpadThree,
	"NumpadFour":   VirtKeyNumpadFour,
	"NumpadFive":   VirtKeyNumpadFive,
	"NumpadSix":    VirtKeyNumpadSix,
	"NumpadSeven":  VirtKeyNumpadSeven,
	"NumpadEight":  VirtKeyNumpadEight,
	"NumpadNine":   VirtKeyNumpadNine,
	"Multiply":     VirtKeyMultiply,
	"Add":          VirtKeyAdd,
	"Subtract":     VirtKeySubtract,
	"DecimalPoint": VirtKeyDecimalPoint,
	"Divide":       VirtKeyDivide,
	"F1":           VirtKeyF1,
	"F2":           VirtKeyF2,
	"F3":           VirtKeyF3,
	"F4":           VirtKeyF4,
	"F5":           VirtKeyF5,
	"F6":           VirtKeyF6,
	"F7":           VirtKeyF7,
	"F8":           VirtKeyF8,
	"F9":           VirtKeyF9,
	"F10":          VirtKeyF10,
	"F11":          VirtKeyF11,
	"F12":          VirtKeyF12,
	"NumLock":      VirtKeyNumLock,
	"ScrollLock":   VirtKeyScrollLock,
	"SemiColon":    VirtKeySemiColon,
	"EqualSign":    VirtKeyEqualSign,
	"Comma":        VirtKeyComma,
	"Dash":         VirtKeyDash,
	"Period":       VirtKeyPeriod,
	"ForwardSlash": VirtKeyForwardSlash,
	"GraveAccent":  VirtKeyGraveAccent,
	"OpenBracket":  VirtKeyOpenBracket,
	"BackSlash":    VirtKeyBackSlash,
	"CloseBracket": VirtKeyCloseBracket,
	"SingleQuote":  VirtKeySingleQuote,
}

// GetVirtualKeyByName takes a the name of a key and returns the VK_* code for it.
func GetVirtualKeyByName(name string) VirtualKey {
	return virtualKeyMap[name]
}

// GetModKeyByName takes the name of a modifier key and returns the ModKey* code for it.
func GetModKeyByName(name string) ModKey {
	return modKeyMap[name]
}
