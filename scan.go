package main

import (
	"bufio"
	"io"
	"strings"
)

func scanFormula( r io.Reader ) formula {
	var f formula
	f.literal = make( map[ string ]bool )

	// scan lines
	scanner := bufio.NewScanner( r )
	for scanner.Scan() {
		// scan each line as a clause
		c := scanClause( strings.NewReader( scanner.Text() ), f.literal )
		if len( c ) > 0 {
			f.cs = append( f.cs, c )
		}
		// length-2 constraint edges don't contribute to Independent Set
		if len( c ) == 3 {
			f.k += 1
		}
	}

	return f
}

// Gets the job done. Not pretty.
func scanClause( r io.Reader, vars map[ string ]bool ) clause {
	var c clause

	scanner := bufio.NewScanner( r )
	scanner.Split( termScan )
	for scanner.Scan() {
		// scan a term
		label := scanner.Text()

		// observe and strip negation
		valence := true
		if label[ 0 : 1 ] == "~" {
			valence = false
			label = label[ 1 : ]
		}

		// when a term is first seen, create a bit of backing state
		if _, mapped := vars[ label ]; !mapped {
			vars[ label ] = false
		}

		// 
		c = append( c, term{ label, valence })
	}

	return c
}

func isLabel( ch byte ) bool {
	return	( ch >= 'a' && ch <= 'z' ) ||
			( ch >= 'A' && ch <= 'Z' ) ||
			( ch >= '0' && ch <= '9' ) ||
			ch == '~'
}

// Just pull terms out
func termScan( data []byte, atEOF bool )( adv int, token []byte, err error ){
	var start int

	for start = 0; start < len( data ); start += 1 {
		ch := data[ start ]
		if isLabel( ch ){
			break
		}
	}

	for i := start; i < len( data ); i++ {
		ch := data[ i ]
		if !isLabel( ch ){
			return i, data[ start : i ], nil
		}
	}

	if atEOF && len( data ) > start {
		return len( data ), data[ start : ],nil
	}

	return start, nil, nil
}

