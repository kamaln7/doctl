/*
Copyright 2018 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package commands

import (
	"fmt"

	"github.com/digitalocean/doctl"
	"github.com/spf13/cobra/doc"
)

const defaultManPath = "/usr/local/share/man/man1"

// ManDocs generates the mandocs command
func ManDocs() *Command {
	cmd := CmdBuilder(DoitCmd, generateManDocs, "mandocs", "generate man docs", Writer)
	AddStringFlag(cmd, doctl.ArgDest, "", defaultManPath, "path to store the generated man files in")

	return cmd
}

func generateManDocs(c *CmdConfig) error {
	dest, err := c.Doit.GetString(c.NS, doctl.ArgDest)
	if err != nil {
		return err
	}

	header := &doc.GenManHeader{
		Title:   "doctl",
		Section: "1",
	}
	err = doc.GenManTree(DoitCmd.Command, header, dest)
	if err != nil {
		return err
	}

	fmt.Fprintf(c.Out, "generated man docs for `doctl` in %s\n", dest)
	return nil
}
