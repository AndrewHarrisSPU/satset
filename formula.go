package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed( time.Now().UnixNano() )
}

// TYPES
// A CNF formula is a collection of clauses
// The map to booleans is suboptimal
type formula struct {
	k			int
	solution	map[ term ]struct{}
	literal 	map[ string ]bool
	cs			[]clause
}

// a clause is collection of (hopefully 2 or 3) terms
type clause []term

// a term is a label referring to a literal value, and a valence (true: non-negated, false: negated)
type term struct {
	label	string
	valence	bool
}

// Copy will point to same clauses, but new literals
func ( f formula ) freshBits() formula {
	solution := make( map[ term ]struct{} )
	literal := make( map[ string ]bool )
	for label := range f.literal {
		literal[ label ] = false
	}

	return formula{
		k: f.k,
		solution: solution,
		literal: literal,
		cs: f.cs,
	}
}

// LOGIC
// We just flip bits until we're satisfied ... Could be a long time. Or forever.
// "remarkably simple" (Kleinberg and Tardos, pg.725)
func ( f *formula ) solve( result chan *formula ) {
	// while we don't evaluate true, or we don't have a k-ary solution
	for !f.eval(){
		f.solution = make( map[ term ]struct{} )
		// flip the bits
		for i := range f.literal {
			f.literal[ i ] = rand.Int() % 2 == 0
		}
	}

	result <- f
}

// Working in CNF, the formula is the AND of all clauses
// bits aren't mutated here, just evaluated
func ( f *formula ) eval() bool {
	// per clause
	for _, c := range f.cs {
		evalClause := false
		// per term
		for _, t := range c {
			evalTerm := f.literal[ t.label ] == t.valence
			// if the literal bit  term matches the valence of the term
			if evalTerm {
				// Only counting the length-3 clauses as 'real' ...
				if len( c ) == 3 && !evalClause {
					f.solution[ t ] = struct{}{}
				}
				// this whole clause is true
				evalClause = evalClause || true
			}
		}

		// no true terms
		if !evalClause {
			return false
		}
	}

	if !( len( f.solution ) == f.k ) {
		return false
	}

	return true
}

// DOT
// We can emit G = ( V, E ) as a dot file ...
func ( f formula ) dot() string {
	out := "strict graph {\n\trankdir = LR;\n"

	for _, c := range f.cs {
		vs := []string{}
		for _, t := range c {
			vs = append( vs, t.String() )

			var color, fontcolor, pos, style string = "color = black", "", "", ""

			if _, ok := f.solution[ t ]; ok {
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
	for label := range f.literal {
		pos, neg := "", ""
		for _, c := range f.cs {
			foundPos, foundNeg := false, false
			for _, t := range c {
				if t.label == label {
					if t.valence {
						foundPos = true
					} else {
						foundNeg = true
					}
				}
			}
			if foundPos && !foundNeg {
				pos = " " + label
			}
			if !foundPos && foundNeg {
				neg = "~" + label
			}
		}
		if pos != "" && neg != "" {
			out += fmt.Sprintf( "\t\"%s\" -- \"%s\" [ style = dotted ]\n", pos, neg )
		}
	}

	out += "}"

	return out
}

// STRING
func ( f formula ) String() string {
	out := ""
	for _, c := range f.cs {
		out += fmt.Sprintf( "%s &\n", c )
	}
	out = out[ : len( out ) - 2 ]
	return out
}

func ( c clause ) String() string {
	out := "( "
	for _, t := range c {
		out += fmt.Sprintf( "%s | ", t )
	}
	out = out[ : len( out ) - 3 ] + " )"
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
