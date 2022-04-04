// Copyright 2022 Robert S. Muhlestein
// SPDX-License-Identifier: Apache-2.0

package term

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/rwxrob/term/esc"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	Reset      string
	Bright     string
	Bold       string
	Dim        string
	Italic     string
	Under      string
	Blink      string
	BlinkF     string
	Reverse    string
	Hidden     string
	Strike     string
	BoldItalic string
	Black      string
	Red        string
	Green      string
	Yellow     string
	Blue       string
	Magenta    string
	Cyan       string
	White      string
	BBlack     string
	BRed       string
	BGreen     string
	BYellow    string
	BBlue      string
	BMagenta   string
	BCyan      string
	BWhite     string
	HBlack     string
	HRed       string
	HGreen     string
	HYellow    string
	HBlue      string
	HMagenta   string
	HCyan      string
	HWhite     string
	BHBlack    string
	BHRed      string
	BHGreen    string
	BHYellow   string
	BHBlue     string
	BHMagenta  string
	BHCyan     string
	BHWhite    string
	X          string
	B          string
	I          string
	U          string
	BI         string
)

// WinSizeStruct is the exact struct used by the ioctl system library.
type WinSizeStruct struct {
	Row, Col       uint16
	Xpixel, Ypixel uint16
}

// WinSize is 80x24 by default but is detected and set to a more
// accurate value at init() time on systems that support ioctl
// (currently) and can be updated with WinSizeUpdate on systems that
// support it. This value can be overriden by those wishing a more
// consistent value or who prefer not to fill the screen completely when
// displaying help and usage information.
var WinSize WinSizeStruct

var interactive bool

func init() { SetInteractive(DetectInteractive()) }

// SetInteractive forces the interactive internal state affecting output
// including calling AttrOn (true) or AttrOff (false).
func SetInteractive(to bool) {
	interactive = to
	if to {
		AttrOn()
	} else {
		AttrOff()
	}
}

// IsInteractive returns the internal interactive state set by
// SetInteractive. The default is that returned by DetectInteractive set
// at  init() time.
func IsInteractive() bool { return interactive }

// DetectInteractive returns true if the output is to an interactive
// terminal (not piped in any way).
func DetectInteractive() bool {
	if f, _ := os.Stdout.Stat(); (f.Mode() & os.ModeCharDevice) != 0 {
		return true
	}
	return false
}

// AttrAreOn contains the state of the last AttrOn/AttrOff call.
var AttrAreOn bool

// AttrOff sets all the terminal attributes to zero values (empty strings).
// Note that this does not affect anything in the esc subpackage (which
// contains the constants from the VT100 specification). Sets the
// AttrAreOn bool to false.
func AttrOff() {
	AttrAreOn = false
	Reset = ""
	Bright = ""
	Bold = ""
	Dim = ""
	Italic = ""
	Under = ""
	Blink = ""
	BlinkF = ""
	Reverse = ""
	Hidden = ""
	Strike = ""
	BoldItalic = ""
	Black = ""
	Red = ""
	Green = ""
	Yellow = ""
	Blue = ""
	Magenta = ""
	Cyan = ""
	White = ""
	BBlack = ""
	BRed = ""
	BGreen = ""
	BYellow = ""
	BBlue = ""
	BMagenta = ""
	BCyan = ""
	BWhite = ""
	HBlack = ""
	HRed = ""
	HGreen = ""
	HYellow = ""
	HBlue = ""
	HMagenta = ""
	HCyan = ""
	HWhite = ""
	BHBlack = ""
	BHRed = ""
	BHGreen = ""
	BHYellow = ""
	BHBlue = ""
	BHMagenta = ""
	BHCyan = ""
	BHWhite = ""
	X = ""
	B = ""
	I = ""
	U = ""
	BI = ""
}

// AttrOn sets all the terminal attributes to zero values (empty strings).
// Note that this does not affect anything in the esc subpackage (which
// contains the constants from the VT100 specification). Sets the
// AttrAreOn bool to true.
func AttrOn() {
	AttrAreOn = true
	Reset = esc.Reset
	Bright = esc.Bright
	Bold = esc.Bold
	Dim = esc.Dim
	Italic = esc.Italic
	Under = esc.Under
	Blink = esc.Blink
	BlinkF = esc.BlinkF
	Reverse = esc.Reverse
	Hidden = esc.Hidden
	Strike = esc.Strike
	Black = esc.Black
	Red = esc.Red
	Green = esc.Green
	Yellow = esc.Yellow
	Blue = esc.Blue
	Magenta = esc.Magenta
	Cyan = esc.Cyan
	White = esc.White
	BBlack = esc.BBlack
	BRed = esc.BRed
	BGreen = esc.BGreen
	BYellow = esc.BYellow
	BBlue = esc.BBlue
	BMagenta = esc.BMagenta
	BCyan = esc.BCyan
	BWhite = esc.BWhite
	HBlack = esc.HBlack
	HRed = esc.HRed
	HGreen = esc.HGreen
	HYellow = esc.HYellow
	HBlue = esc.HBlue
	HMagenta = esc.HMagenta
	HCyan = esc.HCyan
	HWhite = esc.HWhite
	BHBlack = esc.BHBlack
	BHRed = esc.BHRed
	BHGreen = esc.BHGreen
	BHYellow = esc.BHYellow
	BHBlue = esc.BHBlue
	BHMagenta = esc.BHMagenta
	BHCyan = esc.BHCyan
	BHWhite = esc.BHWhite
	X = esc.Reset
	B = esc.Bold
	I = esc.Italic
	U = esc.Under
	BI = esc.BoldItalic
}

// Read reads a single line of input and chomps the \r?\n. Also see
// ReadHidden.
func Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// ReadHidden disables the cursor and echoing to the screen and reads
// a single line of input. Leading and trailing whitespace are removed.
// Also see Read.
func ReadHidden() string {
	byt, err := terminal.ReadPassword(0)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(byt))
}

// Prompt prints the given message if the terminal IsInteractive and
// reads the string by calling Read. The argument signature is identical
// and passed to to fmt.Printf().
func Prompt(form string, args ...any) string {
	if IsInteractive() {
		fmt.Printf(form, args...)
	}
	return Read()
}

// PromptHidden prints the given message if the terminal IsInteractive
// and reads the string by calling ReadHidden (which does not echo to
// the screen). The argument signature is identical and passed to to
// fmt.Printf().
func PromptHidden(form string, args ...any) string {
	if IsInteractive() {
		fmt.Printf(form, args...)
	}
	return ReadHidden()
}

// StripNonPrint remove non-printable runes, e.g. control characters in
// a string that is meant for consumption by terminals that support
// control characters.
func StripNonPrint(s string) string {
	return strings.Map(
		func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, s)
}

// EmphFromLess sets Italic, Bold, BoldItalic, and Under from the
// LESS_TERMCAP_us, _md, _mb, and _us environment variables
// respectively. This is a long used way to provide color to UNIX man
// pages dating back to initial color terminals. UNIX users frequently
// set these to provide color to man pages and more. Observes AttrAreOn
// and will simply return if set to false.
func EmphFromLess() {
	if !AttrAreOn {
		return
	}
	var x string
	x = os.Getenv("LESS_TERMCAP_us")
	if x != "" {
		Italic = x
	}
	x = os.Getenv("LESS_TERMCAP_md")
	if x != "" {
		Bold = x
	}
	x = os.Getenv("LESS_TERMCAP_mb")
	if x != "" {
		BoldItalic = x
	}
	x = os.Getenv("LESS_TERMCAP_us")
	if x != "" {
		Under = x
	}
}
