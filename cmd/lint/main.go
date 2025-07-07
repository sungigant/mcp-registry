/*
Copyright © 2025 Docker, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lint [<server-name>,...]")
		os.Exit(1)
	}

	names := os.Args[1:]

	for _, name := range names {
		if err := lint(name); err != nil {
			fmt.Printf("❌ %s: %v\n", name, err)
		} else {
			fmt.Printf("✅ %s\n", name)
		}
	}

}

func lint(name string) error {
	serverYaml, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	var server map[string]interface{}
	if err := yaml.Unmarshal(serverYaml, &server); err != nil {
		return err
	}

	serverBytes, err := json.Marshal(server)
	if err != nil {
		return err
	}

	var validate SchemaJson
	if err := validate(serverBytes); err != nil {
		return err
	}

	return nil
}
