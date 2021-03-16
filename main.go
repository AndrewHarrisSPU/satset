package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var path = flag.String( "i", "bool.txt", "a plaintext file containing a list of terms in CNF" )

	flag.Parse()

	fin, err := os.Open( *path )
	if err != nil {
		panic( err )
	}

	dst := strings.TrimSuffix( *path, filepath.Ext( *path ))
	fout, err := os.Create( dst + ".dot" )
	if err != nil {
		panic( err )
	}

	f := ReadFormula( fin )
	f.solve( 3 )
	fout.WriteString( f.dot() )
}