// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// goreplacer is a simple command line program for efficiently replacing bytes within a stream of bytes.
//
// Usage
//
// Use `goreplacer help` to see full help documentation.
//
//	goreplacer [--lines] [--replacements] '[[OLD1, NEW1], [OLD2, NEW2], ...]' [-|FILE]
//
// Examples
//
//	# show the
//	echo 'hello world '| goreplacer -r '[["world", "planet"]]'
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/spatialcurrent/goreplacer/pkg/replacer"
)

const (
	flagLines         = "lines"
	flagMaximumNumber = "max"
	flagReplacements  = "replacements"

	NoLimit = -1
)

func initFlags(flag *pflag.FlagSet) {
	flag.BoolP(flagLines, "l", false, "process as lines")
	flag.IntP(flagMaximumNumber, "n", NoLimit, "maximum number of replacements")
	flag.StringP(flagReplacements, "r", "", "replacements encoded as 2-dimensional JSON array, i.e., [[OLD1, NEW1], [OLD2, NEW2], ...]")
}

func initReader(path string) (io.Reader, error) {
	if path == "-" {
		return os.Stdin, nil
	}
	pathExpanded, err := homedir.Expand(path)
	if err != nil {
		return nil, fmt.Errorf("error expanding input path %q: %w", path, err)
	}
	file, err := os.Open(pathExpanded)
	if err != nil {
		return nil, fmt.Errorf("error opening input file at path %q: %w", path, err)
	}
	return file, nil
}

func main() {

	rootCommand := &cobra.Command{
		Use:                   "goreplacer [--lines] [--max MAX] [--replacements] '[[OLD1, NEW1], [OLD2, NEW2], ...]' [-|FILE]",
		DisableFlagsInUseLine: true,
		DisableFlagParsing:    false,
		Short: `goreplacer is a simple tool for for efficiently replacing bytes within a stream of bytes.
Replacements are encoded as a 2 dimensional JSON array of a series of replacement pairs ([OLD, NEW]).
If the positional argument is "-" or not given, then reads from stdin.`,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			v := viper.New()

			if errorBind := v.BindPFlags(cmd.Flags()); errorBind != nil {
				return errorBind
			}

			v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
			v.AutomaticEnv()

			if len(args) > 1 {
				return errors.New("only one input file allowed")
			}

			replacementsString := v.GetString(flagReplacements)
			if len(replacementsString) == 0 {
				return fmt.Errorf("replacements are missing")
			}

			replacements := make([]replacer.Replacement, 0)
			err := json.Unmarshal([]byte(replacementsString), &replacements)
			if err != nil {
				return fmt.Errorf("error unmarshaling replacements: %w", err)
			}

			r, err := replacer.New(replacements...)
			if err != nil {
				return fmt.Errorf("error creating replacer: %w", err)
			}

			lines := v.GetBool(flagLines)

			path := "-"
			if len(args) == 1 {
				path = args[0]
			}

			reader, err := initReader(path)
			if err != nil {
				return fmt.Errorf("error initializing reader: %w", err)
			}

			max := v.GetInt(flagMaximumNumber)
			if max == 0 {
				return errors.New("maximum number is zero, expecting negative one or a value greater than zero")
			}

			if lines {
				scanner := bufio.NewScanner(reader)
				if max > 0 {
					for scanner.Scan() {
						in := scanner.Bytes()
						out, errReplaceBytes := r.ReplaceBytesN(in, max)
						if errReplaceBytes != nil {
							return fmt.Errorf("error replacing bytes %q: %w", in, errReplaceBytes)
						}
						if _, errWrite := os.Stdout.Write(append(out, '\n')); errWrite != nil {
							return fmt.Errorf("error writing output to stdout %q: %w", string(out), errWrite)
						}
					}
				} else {
					for scanner.Scan() {
						in := scanner.Bytes()
						out, errReplaceBytes := r.ReplaceBytes(in)
						if errReplaceBytes != nil {
							return fmt.Errorf("error replacing bytes %q: %w", in, errReplaceBytes)
						}
						if _, errWrite := os.Stdout.Write(append(out, '\n')); errWrite != nil {
							return fmt.Errorf("error writing output to stdout %q: %w", string(out), errWrite)
						}
					}
				}

				if errScanner := scanner.Err(); errScanner != nil {
					return fmt.Errorf("error scanning input: %w", errScanner)
				}
				return nil
			}

			in, err := ioutil.ReadAll(reader)
			if err != nil {
				return fmt.Errorf("error reading input: %w", err)
			}

			if max > 0 {
				out, err := r.ReplaceBytesN(in, max)
				if err != nil {
					return fmt.Errorf("error replacing bytes %q: %w", in, err)
				}
				if _, err := os.Stdout.Write(append(out, '\n')); err != nil {
					return fmt.Errorf("error writing output to stdout %q", string(out))
				}
				return nil
			}

			out, err := r.ReplaceBytes(in)
			if err != nil {
				return fmt.Errorf("error replacing bytes %q: %w", in, err)
			}
			if _, err := os.Stdout.Write(append(out, '\n')); err != nil {
				return fmt.Errorf("error writing output to stdout %q", string(out))
			}

			return nil
		},
	}
	initFlags(rootCommand.Flags())

	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "goreplacer: "+err.Error())
		fmt.Fprintln(os.Stderr, "Try goreplacer --help for more information.")
		os.Exit(1)
	}
}
