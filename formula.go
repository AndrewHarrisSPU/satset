package main

import (
	"fmt"
	"math/rand"
	"time"
)

// DEFINITIONS
// A CNF formula is a collection of clauses
type formula struct {
	k		int
	bits 	map[ string ]bool
	cs		[]clause
}

// a clause is collection of (hopefully 2 or 3) terms
type clause []term

// a term is a label referring to a literal value, and a valence (true: non-negated, false: negated)
type term struct {
	label	string
	valence	bool
}

// "remarkably simple" (Kleinberg and Tardos, pg.725)
func ( f *formula ) solve( k int ) {
	rand.Seed( time.Now().UnixNano() )

	f.k = k + 1
	for !f.eval() || f.k != k {
		fmt.Printf( "f.k: %d\n", f.k )
		for i := range f.bits {
			// fmt.Printf( "%s: %v\n", k, f.bits[ k ])
			f.bits[ i ] = rand.Int() % 2 == 0
		}
	} 

	fmt.Println( f )

	fmt.Println( "solution:" )
	for i, v := range f.bits {
		fmt.Printf( "\t%s: %v\n", i, v )
	}
}

// Working in CNF, the formula is the AND of all clauses
func ( f *formula ) eval() bool {
	f.k = 0
	for _, c := range f.cs {
		result := false
		for _, t := range c {
			literal := f.bits[ t.label ]
			eval := literal == t.valence
			result = result || eval
			if len( c ) == 3 && eval {
				f.k += 1
			}
		}
		if !result {
			fmt.Println( c )
			return false
		}
	}

	return true
}

// We can emit G = ( V, E ) as a dot file ...
func ( f formula ) dot() string {
	out := "strict graph {\n\trankdir = LR;\n"

	for _, c := range f.cs {
		vs := []string{}
		for _, t := range c {
			vs = append( vs, t.String() )

			var color, fontcolor, pos, style string = "color = black", "", "", ""

			if f.bits[ t.label ] == t.valence {
				color = "color = red"
			}

			if len( c ) == 2 {
				color = "color = gray"
				fontcolor = "fontcolor = gray"
				pos = "pos = \"-10,0!\""
				style = "style = dotted"
			}

			out += fmt.Sprintf( "\tnode [ %s %s %s %s ] \"%s\"\n", color, fontcolor, pos, style, t )
		}

		if len( c ) == 3 {
			out += fmt.Sprintf( "\t\"%s\" -- \"%s\"\n", c[ 0 ], c[ 1 ])		
			out += fmt.Sprintf( "\t\"%s\" -- \"%s\"\n", c[ 1 ], c[ 2 ])
			out += fmt.Sprintf( "\t\"%s\" -- \"%s\"\n", c[ 2 ], c[ 0 ])
		} else {
			out += fmt.Sprintf( "\t\"%s\" -- \"%s\" [ style = dotted ]\n", c[ 0 ], c[ 1 ])
		}
	}

	// the tricky part ...
	for label := range f.bits {
		pos, neg := "", ""
		for _, c := range f.cs {
			for _, t := range c {
				if t.label == label {
					if t.valence {
						pos = " " + label
					} else {
						neg = "~" + label
					}
				}
			}
		}
		if pos != "" && neg != "" {
			out += fmt.Sprintf( "\t\"%s\" -- \"%s\" [ style = dotted ]\n", pos, neg )
		}
	}

	// for _, e := range E {
	// 	out += fmt.Sprintf( "\t\"%s\" -- \"%s\"\n", e[ 0 ], e[ 1 ])
	// }

	out += "}"

	return out
}

// STRING
func ( f formula ) String() string {
	out := ""
	for _, c := range f.cs {
		out += fmt.Sprintf( "%s &\n", c )
	}
	return out
}

func ( c clause ) String() string {
	out := "( "
	for _, t := range c {
		out += fmt.Sprintf( "%s | ", t )
	}
	// out = out[ : len( out ) - 3 ] + " ) &"
	return out
}

func ( t term ) String() string {
	var negate string
	if t.valence == false {
		negate = "~"
	} else {
		negate = " "
	}

	return fmt.Sprintf( "%s%s", negate, t.label )
}
