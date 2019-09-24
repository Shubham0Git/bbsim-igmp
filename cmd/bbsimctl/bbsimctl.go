/*
 * Portions copyright 2019-present Open Networking Foundation
 * Original copyright 2019-present Ciena Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/opencord/bbsim/internal/bbsimctl/commands"
	"github.com/opencord/bbsim/internal/bbsimctl/config"
	"os"
	"path"
)

func main() {

	parser := flags.NewNamedParser(path.Base(os.Args[0]),
		flags.HelpFlag|flags.PassDoubleDash|flags.PassAfterNonOption)
	_, err := parser.AddGroup("Global Options", "", &config.GlobalOptions)
	if err != nil {
		panic(err)
	}

	commands.RegisterConfigCommands(parser)
	commands.RegisterOltCommands(parser)
	commands.RegisterONUCommands(parser)

	_, err = parser.ParseArgs(os.Args[1:])
	if err != nil {
		_, ok := err.(*flags.Error)
		if ok {
			real := err.(*flags.Error)
			if real.Type == flags.ErrHelp {
				os.Stdout.WriteString(err.Error() + "\n")
				return
			}
		}

		//corderror, ok := err.(corderrors.CordCtlError)
		//if ok {
		//	if corderror.ShouldDumpStack() || config.GlobalOptions.Debug {
		//		os.Stderr.WriteString("\n" + corderror.Stack())
		//	}
		//}

		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err.Error())

		os.Exit(1)
	}
}