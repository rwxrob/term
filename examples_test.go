// Copyright 2022 Robert S. Muhlestein
// SPDX-License-Identifier: Apache-2.0

package term_test

import (
	"fmt"

	"github.com/rwxrob/term"
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
