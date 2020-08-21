// Copyright (c) Victor Hurdugaci (https://victorhurdugaci.com). All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var includes arrayFlags
	var packageName string
	flag.Var(&includes, "include", "Glob pattern for the files to include. Must be subfolder of")
	flag.StringVar(&packageName, "package", "", "The name of the package. Optional. Default: folder name")
	flag.Parse()

	workDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	matches := sortedStringSet{}
	for _, include := range includes {
		include = filepath.Join(workDir, include)
		m, err := filepath.Glob(include)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s", include, err)
			os.Exit(1)
		}
		matches.Add(m...)
	}

	code, err := generateGoCode(generateOptions{
		WorkingDir: workDir,
		FilePaths:  matches,
		Package:    packageName,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "code generation failed: %s", err)
		os.Exit(1)
	}

	embedrGoFilePath := filepath.Join(workDir, "embedr.go")
	f, err := os.Create(embedrGoFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open embedr.go: %s", err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString(code)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write embedr.go: %s", err)
		os.Exit(1)
	}
}
