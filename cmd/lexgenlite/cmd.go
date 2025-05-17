package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/agentio/atiquette/cmd/lexgenlite/gen"
	"github.com/spf13/cobra"
)

func Cmd(ctx context.Context) *cobra.Command {
	var outdir string
	var genServer bool
	var genHandlers bool
	var typesImport []string
	var packageName string
	var buildFile string
	var maxStringLength int

	cmd := &cobra.Command{
		Use:   "lexgen LEXICONS",
		Short: "Generate support code for lexicons",
		RunE: func(cmd *cobra.Command, args []string) error {
			paths, err := expandLexiconArgs(args)
			if err != nil {
				return err
			}

			var schemas []*gen.Schema
			for _, arg := range paths {
				s, err := gen.ReadSchema(arg)
				if err != nil {
					// TODO: maybe this shouldn't be an error
					return fmt.Errorf("failed to read file %q: %w", arg, err)
				}
				// skip non-Lexicon JSON files
				if s.Lexicon == 0 {
					continue
				}
				schemas = append(schemas, s)
			}

			if buildFile == "" {
				return errors.New("--build-file is required")
			}

			var packages []gen.Package

			blob, err := os.ReadFile(buildFile)
			if err != nil {
				return fmt.Errorf("--build-file error, %w", err)
			}
			packages, err = gen.ParsePackages(blob)
			if err != nil {
				return fmt.Errorf("--build-file error, %w", err)
			}
			if len(packages) == 0 {
				return errors.New("--build-file must specify at least one Package{}")
			}

			if genServer {
				pkgname := packageName
				if outdir == "" {
					return fmt.Errorf("must specify output directory (--outdir)")
				}
				defmap := gen.BuildExtDefMap(schemas, packages)
				_ = defmap
				paths := typesImport
				importmap := make(map[string]string)
				for _, p := range paths {
					parts := strings.Split(p, ":")
					importmap[parts[0]] = parts[1]
				}
				handlers := genHandlers
				if err := gen.CreateHandlerStub(pkgname, importmap, outdir, schemas, handlers); err != nil {
					return err
				}
			} else {
				return gen.Run(schemas, packages, maxStringLength)
			}

			return nil
		},
	}
	cmd.Flags().StringVar(&outdir, "outdir", "", "output directory")
	cmd.Flags().BoolVar(&genServer, "gen-server", false, "")
	cmd.Flags().BoolVar(&genHandlers, "gen-handlers", false, "")
	cmd.Flags().StringSliceVar(&typesImport, "types-import", nil, "")
	cmd.Flags().StringVar(&packageName, "package", "schemagen", "")
	cmd.Flags().StringVar(&buildFile, "build-file", "", "")
	cmd.Flags().IntVar(&maxStringLength, "max-string-length", 8192, "")
	return cmd
}
