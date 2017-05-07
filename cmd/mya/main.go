package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/844196/ncipher"
	"github.com/mitchellh/cli"
)

const (
	Name    = "mya"
	Version = "0.1.0"
)

type stream struct {
	in       io.ReadCloser
	out, err io.Writer
}

type encoding func(string) (string, error)

type encodingFactory func(*ncipher.Config) (encoding, error)

type myaCmd struct {
	name   string
	stream stream
	encodingFactory
}

func (m *myaCmd) Help() string {
	u := fmt.Sprintf("Usage: %s [options] <FILE>...\n", m.name)

	myaCmdOpts.VisitAll(func(f *flag.Flag) {
		u += fmt.Sprintf("  -%s\t%s (default \"%s\")\n", f.Name, f.Usage, f.DefValue)
	})

	return u
}

func (m *myaCmd) Synopsis() string {
	return ""
}

func (m *myaCmd) Run(args []string) int {
	var err error

	myaCmdOpts.Parse(args)

	cnf := ncipher.Config{
		Seed:      seed,
		Delimiter: delimiter,
	}
	enc, err := m.encodingFactory(&cnf)
	if err != nil {
		fmt.Fprintln(m.stream.err, err)
		return 1
	}

	for _, path := range myaCmdOpts.Args() {
		err = func() error {
			var fp io.ReadCloser

			if path == "-" {
				fp = m.stream.in
			} else {
				fp, err = os.Open(path)
				if err != nil {
					return err
				}
			}
			defer fp.Close()

			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				out, err := enc(scanner.Text())
				if err != nil {
					return err
				}
				fmt.Fprintln(m.stream.out, out)
			}
			if err := scanner.Err(); err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			fmt.Fprintln(m.stream.err, err)
			return 1
		}
	}

	return 0
}

var (
	myaCmdOpts = flag.NewFlagSet("", flag.ExitOnError)
	seed       string
	delimiter  string
)

func init() {
	myaCmdOpts.StringVar(&seed, "s", ncipher.StdConfig.Seed, "seed value")
	myaCmdOpts.StringVar(&delimiter, "d", ncipher.StdConfig.Delimiter, "delimiter value")
}

func main() {
	stdStream := stream{os.Stdin, os.Stdout, os.Stderr}

	cmd := cli.NewCLI(Name, Version)
	cmd.Args = os.Args[1:]
	cmd.Commands = map[string]cli.CommandFactory{
		"encode": func() (c cli.Command, e error) {
			c = &myaCmd{
				name:   "encode",
				stream: stdStream,
				encodingFactory: func(cnf *ncipher.Config) (encoding, error) {
					enc, err := ncipher.NewEncoding(cnf)
					if err != nil {
						return nil, err
					}
					return func(src string) (dst string, err error) {
						return enc.Encode(src), err
					}, nil
				},
			}
			return
		},
		"decode": func() (c cli.Command, e error) {
			c = &myaCmd{
				name:   "decode",
				stream: stdStream,
				encodingFactory: func(cnf *ncipher.Config) (encoding, error) {
					enc, err := ncipher.NewEncoding(cnf)
					if err != nil {
						return nil, err
					}
					return func(src string) (dst string, err error) {
						return enc.Decode(src)
					}, nil
				},
			}
			return
		},
	}

	stat, err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(stat)
}
