// Copyright 2022 Robert S. Muhlestein
// SPDX-License-Identifier: Apache-2.0

package term_test

import (
	"fmt"

	"github.com/rwxrob/term"
	"github.com/rwxrob/term/esc"
)

func ExampleRed() {
	term.Red = "<red>"
	term.Reset = "<reset>"
	fmt.Println(term.Red + "simply red" + term.Reset)
	term.AttrOff()
	fmt.Println(term.Red + "simply red" + term.Reset)
	// Output:
	// <red>simply red<reset>
	// simply red
}

func ExampleStripNonPrint() {
	some := esc.Bold + "not bold" + esc.Reset
	fmt.Println(term.StripNonPrint(some))
	// Output;
	// not bold
}
