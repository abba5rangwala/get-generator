package get_generator

import (
	"flag"
	"fmt"
	"github.com/abba5/get-generator/parser"
	"os"
)

var allStructs = flag.Bool("all", false, "generate marshaler/unmarshalers for all structs in a file")

func main() {
	flag.Parse()
	fname := flag.Args()

	/*gofile := os.Getenv("GOFILE")
	if *processPkg {
		gofile = filepath.Dir(gofile)
	}*/

	/*if len(files) == 0 && gofile != "" {
		files = []string{gofile}
	} else if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}*/

	if err := generate(fname[0]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func generate(fname string) (err error) {
	fInfo, err := os.Stat(fname)
	if err != nil {
		return err
	}

	p := parser.Parser{AllStructs: *allStructs}
	if err := p.Parse(fname, fInfo.IsDir()); err != nil {
		return fmt.Errorf("Error parsing %v: %v", fname, err)
	}
}
