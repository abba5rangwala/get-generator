package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/abba5/get-generator/bootstrap"
	"github.com/abba5/get-generator/parser"
	"os"
	"path/filepath"
	"strings"
)

var genBuildFlags = flag.String("gen_build_flags", "", "build flags when running the generator while bootstrapping")
var allStructs = flag.Bool("all", false, "generate getters for all structs in a file")
var leaveTemps = flag.Bool("leave_temps", false, "do not delete temporary files")
var noformat = flag.Bool("noformat", false, "do not run 'gofmt -w' on output file")
var specifiedName = flag.String("output_filename", "", "specify the filename of the output")

func main() {
	flag.Parse()
	fname := flag.Args()

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
		return fmt.Errorf("error parsing %v: %v", fname, err)
	}

	var outName string
	if fInfo.IsDir() {
		outName = filepath.Join(fname, p.PkgName+"_getter.go")
	} else {
		if s := strings.TrimSuffix(fname, ".go"); s == fname {
			return errors.New("filename must end in '.go'")
		} else {
			outName = s + "_getter.go"
		}
	}

	if *specifiedName != "" {
		outName = *specifiedName
	}

	var trimmedGenBuildFlags string
	if *genBuildFlags != "" {
		trimmedGenBuildFlags = strings.TrimSpace(*genBuildFlags)
	}

	g := bootstrap.Generator{
		GenBuildFlags: trimmedGenBuildFlags,
		PkgPath:       p.PkgPath,
		PkgName:       p.PkgName,
		Types:         p.StructNames,
		LeaveTemps:    *leaveTemps,
		OutName:       outName,
		NoFormat:      *noformat,
	}

	if err := g.Run(); err != nil {
		return fmt.Errorf("bootstrap failed: %v", err)
	}

	return nil
}
