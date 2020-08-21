// Copyright (c) Victor Hurdugaci (https://victorhurdugaci.com). All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type generateOptions struct {
	WorkingDir string
	FilePaths  sortedStringSet
	Package    string
}

func generateGoCode(opts generateOptions) (string, error) {
	filesToInclude := opts.FilePaths.All()

	packageName := opts.Package
	if packageName == "" {
		packageName = filepath.Dir(opts.WorkingDir)
	}

	normalizedFilePaths := make([]string, len(opts.FilePaths))
	for i, filePath := range filesToInclude {
		normalizedFilePath, err := normalizePath(opts.WorkingDir, filePath)
		if err != nil {
			return "", fmt.Errorf("%s: %s", filePath, err)
		}

		normalizedFilePaths[i] = normalizedFilePath
	}

	sb := strings.Builder{}
	sb.WriteString("// Code generated by embedr; DO NOT EDIT.\n\n")
	sb.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	sb.WriteString("import (\n")
	sb.WriteString("	\"github.com/victorhurdugaci/embedr\"\n")
	sb.WriteString(")\n\n")
	sb.WriteString("func init() {\n")

	for i, filePath := range filesToInclude {
		fileContent, err := fileAsString(filePath)
		if err != nil {
			return "", fmt.Errorf("%s: %s", filePath, err)
		}
		sb.WriteString(fmt.Sprintf(
			"	embedr.Add(\"%s\", \"%s\")\n",
			normalizedFilePaths[i],
			fileContent))
	}

	sb.WriteString("}")

	return sb.String(), nil
}

func normalizePath(workDir string, path string) (string, error) {
	if !filepath.IsAbs(workDir) {
		return "", fmt.Errorf("Workdir must be abs: %s", workDir)
	}

	if !filepath.IsAbs(path) {
		return "", fmt.Errorf("Path must be abs: %s", workDir)
	}

	if !strings.HasPrefix(path, workDir) {
		return "", fmt.Errorf("Path must be relative to %s", workDir)
	}

	return filepath.Rel(workDir, path)
}

func fileAsString(absPath string) (string, error) {
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}