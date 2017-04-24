package main

import (
	"bufio"
	"fmt"
	"github.com/844196/ncipher/lib"
	"github.com/jessevdk/go-flags"
	"github.com/mitchellh/cli"
	"io"
	"os"
)

const (
	version = "0.1.0"
)

// sub commands options
type ConverterOptions struct {
	Seed      string `short:"s" long:"seed" description:"Specifies seed" default:"にゃんぱす"`
	Delimiter string `short:"d" long:"delimiter" description:"Specifies delimiter" default:"〜"`
}

var convOpts ConverterOptions
var subParser = flags.NewParser(&convOpts, flags.Default^flags.HelpFlag^flags.PrintErrors)

// common
type ConvertType int

const (
	CONVERT_TYPE_ENCODE ConvertType = 0
	CONVERT_TYPE_DECODE ConvertType = 1
)

func convert(ct ConvertType, path string, conv *ncipher.Converter, inS io.Reader, outS io.Writer) error {
	var fp io.Reader
	var err error
	if path == "-" {
		fp = inS
	} else {
		fp, err = os.Open(path)
		if err != nil {
			return err
		}
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		switch ct {
		case CONVERT_TYPE_ENCODE:
			fmt.Fprintln(outS, conv.Encode(scanner.Text()))
		case CONVERT_TYPE_DECODE:
			fmt.Fprintln(outS, conv.Decode(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// encoder
type Encoder struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

func (e *Encoder) Help() string {
	return "Usage: encode " + subParser.Usage
}

func (e *Encoder) Synopsis() string {
	return "Encode file"
}

func (e *Encoder) Run(args []string) int {
	files, err := subParser.ParseArgs(args)
	if err != nil {
		fmt.Fprintln(e.errStream, err)
		return 1
	}

	opts := ncipher.Options{
		Seed:      convOpts.Seed,
		Delimiter: convOpts.Delimiter,
	}
	conv, err := ncipher.NewConverter(&opts)
	if err != nil {
		fmt.Fprintln(e.errStream, err)
		return 1
	}

	for _, path := range files {
		err := convert(CONVERT_TYPE_ENCODE, path, conv, e.inStream, e.outStream)
		if err != nil {
			fmt.Fprintln(e.errStream, err)
			return 1
		}
	}

	return 0
}

// decoder
type Decoder struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

func (d *Decoder) Help() string {
	return "Usage: encode " + subParser.Usage
}

func (d *Decoder) Synopsis() string {
	return "Decode file"
}

func (d *Decoder) Run(args []string) int {
	files, err := subParser.ParseArgs(args)
	if err != nil {
		fmt.Fprintln(d.errStream, err)
		return 1
	}

	opts := ncipher.Options{
		Seed:      convOpts.Seed,
		Delimiter: convOpts.Delimiter,
	}
	conv, err := ncipher.NewConverter(&opts)
	if err != nil {
		fmt.Fprintln(d.errStream, err)
		return 1
	}

	for _, path := range files {
		err := convert(CONVERT_TYPE_DECODE, path, conv, d.inStream, d.outStream)
		if err != nil {
			fmt.Fprintln(d.errStream, err)
			return 1
		}
	}

	return 0
}

func main() {
	c := cli.NewCLI("ncipher", version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"encode": func() (cli.Command, error) {
			return &Encoder{os.Stdin, os.Stdout, os.Stderr}, nil
		},
		"decode": func() (cli.Command, error) {
			return &Decoder{os.Stdin, os.Stdout, os.Stderr}, nil
		},
	}

	subParser.Usage = "[--seed=SEED] [--delimiter=DELIMITER] <FILE>"

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(exitStatus)
}
