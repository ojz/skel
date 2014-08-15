package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var source string
var dest string
var nonce string
var name string

func init() {
	flag.StringVar(&source, "source", "", "Where to find the template files.")
	flag.StringVar(&dest, "dest", "", "Where to save the generated files.")
	flag.StringVar(&nonce, "nonce", "", "The word which should be replaced (defaults to the basename of the source).")
	flag.StringVar(&name, "name", "", "The name of the new project (defaults to the basename of the destination).")
}

func replace(old string) string {
	// @TODO it would be nice if we could support casing:
	// - find (case insensitive) matches of nonce
	// - determine casing (lower, UPPER, Capital, CamelCase, lowerCamel)
	// - replace by the corresponding casing of name
	// (+ setup casings of name in advance)
	return strings.Replace(old, nonce, name, -1)
}

func main() {
	flag.Parse()

	if source == "" || dest == "" {
		flag.Usage()
		return
	}

	if nonce == "" {
		nonce = path.Base(source)
	}

	if name == "" {
		name = path.Base(dest)
	}

	copydir(source, dest)
}

func copydir(source string, dest string) {

	// first create the wanted directory:
	os.Mkdir(dest, os.ModePerm)

	// copy the files in it
	var fp *os.File
	var err error

	fp, err = os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	var files []os.FileInfo
	files, err = fp.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			copydir(source+"/"+f.Name(), dest+"/"+replace(f.Name()))
		} else {
			copyfile(source+"/"+f.Name(), dest+"/"+replace(f.Name()))
		}
	}
}

func copyfile(infile string, outfile string) {
	in, _ := ioutil.ReadFile(infile)
	out := replace(string(in))
	ioutil.WriteFile(outfile, []byte(out), os.ModePerm)
}
