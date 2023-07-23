/*

 MIT License

 (C) Copyright 2023 Hewlett Packard Enterprise Development LP

 Permission is hereby granted, free of charge, to any person obtaining a
 copy of this software and associated documentation files (the "Software"),
 to deal in the Software without restriction, including without limitation
 the rights to use, copy, modify, merge, publish, distribute, sublicense,
 and/or sell copies of the Software, and to permit persons to whom the
 Software is furnished to do so, subject to the following conditions:

 The above copyright notice and this permission notice shall be included
 in all copies or substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 OTHER DEALINGS IN THE SOFTWARE.

*/

package get

import (
	"github.com/Cray-HPE/gru/pkg/cmd/cli/bios"
	"github.com/Cray-HPE/gru/pkg/cmd/cli/power"
	"github.com/spf13/cobra"
)

// NewCommand creates the `get` subcommand.
func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "get",
		Short: "Shortcut for getting certain information from RedFish.",
		Long:  `Shortcut for getting certain information from RedFish.`,
		Run: func(c *cobra.Command, args []string) {
		},
	}
	c.AddCommand(
		bios.NewGetCommand(),
		power.NewGetCommand(),
	)
	return c
}
