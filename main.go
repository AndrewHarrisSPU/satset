package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// FLAGS
	var path = flag.String( "input", "", "a plaintext file containing a list of terms in CNF" )
	var seconds = flag.Int( "seconds", 1, "how many seconds before giving up" )
	var dot = flag.Bool( "dot", false, "write a .dot file of the solution (experimental)" )
	flag.Parse()

	if *path == "" {
		fmt.Println( "Please use -input to specify an input file. See --help for full options." )
		return
	}

	problem := Load( *path )
	result := Solve( problem, *seconds )

	if result != nil {
		fmt.Println( result )
		// write solution to console
		for t := range result.solution {
			fmt.Printf( "%s: %v\n", t, result.literal[ t.label ])
		}

		// write a soution graph to a file
		if *dot {
			Dot( *path, result )
		}
	} else {
		// better luck next time?
		fmt.Println( "stumped!" )
	}
}

func Load( path string ) formula {
	fin, err := os.Open( path )
	if err != nil {
		panic( err )
	}
	defer fin.Close()

	return scanFormula( fin )
}

func Dot( path string, f *formula ) {
	base := strings.TrimSuffix( path, filepath.Ext( path ))
	fout, err := os.Create( base + ".dot" )
	if err != nil {
		panic( err )
	}
	defer fout.Close()

	fout.WriteString( f.dot() )
}

func Solve( problem formula, seconds int ) *formula {
	// channel with solution status
	// we will pass a correct solution, or nil
	done := make( chan *formula )

	// spin up enough solvers for the runtime scheduler to work with
	for i := 0; i < 100; i++ {
		solver := problem.freshBits()
		go solver.solve( done )
	}

	// put a limit on it ...
	timer := time.NewTimer( time.Duration( seconds ) * time.Second )
	go func() {
		<- timer.C 
		done <- nil
	}()

	// wait for a solution
	return <-done
}