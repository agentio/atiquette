package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// for direct .json lexicon files or directories containing lexicon .json files,
// return one flat list of all paths to .json files
// this assumes that all .json files are lexicons!
func expandLexiconArgs(args []string) ([]string, error) {
	var out []string
	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			out, err = collectJSONFiles(arg, out)
			if err != nil {
				return nil, err
			}
		} else if strings.HasSuffix(arg, ".json") {
			out = append(out, arg)
		}
	}
	return out, nil
}

// collect all the JSON files in a specified directory and its children
func collectJSONFiles(dir string, out []string) ([]string, error) {
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".json") {
			out = append(out, path)
		}
		return nil
	})
	if err != nil {
		return out, err
	}
	return out, nil
}
