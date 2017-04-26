package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/mitchellh/cli"
)

type IoStream struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

type Converter struct {
	Usage       string
	Description string
	Invoker     func(string, *ncipher.Options) (string, error)
	IoStream    IoStream
}

func (c *Converter) Help() string {
	return c.Usage
}

func (c *Converter) Synopsis() string {
	return c.Description
}

func (c *Converter) Run(args []string) int {
	files, err := subParser.ParseArgs(args)
	if err != nil {
		fmt.Fprintln(c.IoStream.errStream, err)
		return 1
	}

	opts, err := ncipher.NewOptions(convOpts.Seed, convOpts.Delimiter)
	if err != nil {
		fmt.Fprintln(c.IoStream.errStream, err)
		return 1
	}

	for _, path := range files {
		err = func() error {
			var fp io.Reader
			var err error
			if path == "-" {
				fp = c.IoStream.inStream
			} else {
				fp, err = os.Open(path)
				if err != nil {
					return err
				}
				defer fp.Close()
			}

			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				result, err := c.Invoker(scanner.Text(), opts)
				if err != nil {
					return err
				}
				fmt.Fprintln(c.IoStream.outStream, result)
			}
			if err := scanner.Err(); err != nil {
				return err
			}

			return nil
		}()
		if err != nil {
			fmt.Fprintln(c.IoStream.errStream, err)
			return 1
		}
	}

	return 0
}

func main() {
	stdStream := IoStream{
		inStream:  os.Stdin,
		outStream: os.Stdout,
		errStream: os.Stderr,
	}

	c := cli.NewCLI("ncipher", version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"encode": func() (cli.Command, error) {
			return &Converter{
				Usage:       "Usage: decode [--seed=SEED] [--delimiter=DELIMITER] <FILE>...",
				Description: "Encode file",
				Invoker: func(s string, o *ncipher.Options) (string, error) {
					return ncipher.Encode(s, o)
				},
				IoStream: stdStream,
			}, nil
		},
		"decode": func() (cli.Command, error) {
			return &Converter{
				Usage:       "Usage: encode [--seed=SEED] [--delimiter=DELIMITER] <FILE>...",
				Description: "Decode file",
				Invoker: func(s string, o *ncipher.Options) (string, error) {
					return ncipher.Decode(s, o)
				},
				IoStream: stdStream,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(exitStatus)
}
