package main

import (
	"fmt"
	"testing"
)

func TestCSSX( test *testing.T ){
	cs := []clause{
		clause{[ 3 ]term{
			{ false, "x", 1 },
			{ false, "x", 0 },
			{ true, "x", 2 }}},
		clause{[ 3 ]term{
			{ false, "x", 2 },
			{ false, "x", 1 },
			{ true, "x", 3 }}},
		clause{[ 3 ]term{
			{ true, "x", 1 },
			{ false, "x", 0 },
			{ false, "x", 3 }}},
	}

	// println( "\n clauses\n" )
	// cs, k, ys := gen3SAT( kcs, kns )
	// for _, c := range cs {
	// 	fmt.Println( c )
	// }

	println( "\n PCNF\n" )
	script := cnf( cs, 4, 0 )
	fmt.Println( script )
}

// func TestKClause( test *testing.T ){

// 	cs := []clause{
// 		clause{ [ 3 ]term{
// 			{ true, "x", 1 },
// 			{ false, "x", 2 },
// 			{ false, "x", 3 }}},
// 		clause{ [ 3 ]term{
// 			{ false, "x", 1 },
// 			{ true, "x", 2 },
// 			{ false, "x", 3 }}},
// 		clause{ [ 3 ]term{
// 			{ true, "x", 1 },
// 			{ true, "x", 2 },
// 			{ true, "x", 3 }}},
// 		clause{ [ 3 ]term{
// 			{ true, "x", 1 },
// 			{ true, "x", 2 },
// 			{ false, "x", 4 }}},
// 		clause{ [ 3 ]term {
// 			{ true, "x", 2 },
// 			{ false, "x", 3 },
// 			{ false, "x", 4 }}},
// 		}

// 	println( cnf( cs, 0 ))
// }

func TestVC( test *testing.T ){
	adjItoB := func( list []int ) []bool {
		bits := make( []bool, len( list ))
		for i := range list {
			if list[ i ] != 0 {
				bits[ i ] = true
			}
		}

		return bits
	}

	v1 := []int{ 0, 1, 1, 0, 0, 0 }
	v2 := []int{ 1, 0, 1, 1, 1, 1 }
	v3 := []int{ 1, 1, 0, 0, 0, 0 }
	v4 := []int{ 0, 1, 0, 0, 0, 0 }
	v5 := []int{ 0, 1, 0, 0, 0, 0 }	
	v6 := []int{ 0, 1, 0, 0, 0, 0 }	


	adj := [][]bool{
		adjItoB( v1 ),
		adjItoB( v2 ),
		adjItoB( v3 ),
		adjItoB( v4 ),
		adjItoB( v5 ),
		adjItoB( v6 ),
	}

	println( "\n kclauses\n" )
	kcs, kns := genKSAT( adj )

	println( "\n clauses\n" )
	cs, k, ys := gen3SAT( kcs, kns )
	for _, c := range cs {
		fmt.Println( c )
	}

	println( "\n PCNF\n" )
	script := cnf( cs, k, ys )

	fmt.Println( script )
}

func TestPetersen( test *testing.T ){
	// adj literals easier with 0s and 1s
	adjItoB := func( list []int ) []bool {
		bits := make( []bool, len( list ))
		for i := range list {
			if list[ i ] != 0 {
				bits[ i ] = true
			}
		}

		return bits
	}

	// v1 := []int{ 0, 1, 0, 0, 1, 0, 1, 0, 0, 0 }
	// v2 := []int{ 1, 0, 1, 0, 0, 0, 0, 1, 0, 0 }
	// v3 := []int{ 0, 1, 0, 1, 0, 0, 0, 0, 1, 0 }
	// v4 := []int{ 0, 0, 1, 0, 1, 0, 0, 0, 0, 1 }
	// v5 := []int{ 1, 0, 0, 1, 0, 1, 0, 0, 0, 0 }
	// v6 := []int{ 0, 0, 0, 0, 1, 0, 0, 1, 1, 0 }
	// v7 := []int{ 1, 0, 0, 0, 0, 0, 0, 0, 1, 1 }
	// v8 := []int{ 0, 1, 0, 0, 0, 1, 0, 0, 0, 1 }
	// v9 := []int{ 0, 0, 1, 0, 0, 1, 1, 0, 0, 0 }
	// v10 := []int{ 0, 0, 0, 1, 0, 0, 1, 1, 0, 0 }

	v1 := []int{ 0, 1, 0, 0, 1, 1, 0, 0, 0, 0 }
	v2 := []int{ 1, 0, 1, 0, 0, 0, 1, 0, 0, 0 }
	v3 := []int{ 0, 1, 0, 1, 0, 0, 0, 1, 0, 0 }
	v4 := []int{ 0, 0, 1, 0, 1, 0, 0, 0, 1, 0 }
	v5 := []int{ 1, 0, 0, 1, 0, 0, 0, 0, 0, 1 }
	v6 := []int{ 1, 0, 0, 0, 0, 0, 0, 1, 1, 0 }
	v7 := []int{ 0, 1, 0, 0, 0, 0, 0, 0, 1, 1 }
	v8 := []int{ 0, 0, 1, 0, 0, 1, 0, 0, 0, 1 }
	v9 := []int{ 0, 0, 0, 1, 0, 1, 1, 0, 0, 0 }
	v10 := []int{ 0, 0, 0, 0, 1, 0, 1, 1, 0, 0 }

	adj := [][]bool{
		adjItoB( v1 ),
		adjItoB( v2 ),
		adjItoB( v3 ),
		adjItoB( v4 ),
		adjItoB( v5 ),
		adjItoB( v6 ),
		adjItoB( v7 ),
		adjItoB( v8 ),
		adjItoB( v9 ),
		adjItoB( v10 ),
	}

	println( "\n kclauses\n" )
	kcs, kns := genKSAT( adj )
	for _, k := range kcs {
		fmt.Println( k )
	}
	for _, k := range kns {
		fmt.Println( k )
	}

	script := kSATcnf( kcs, kns )

	// println( "\n clauses\n" )
	// cs, k, ys := gen3SAT( kcs, kns )
	// for _, c := range cs {
	// 	fmt.Println( c )
	// }

	// println( "\n PCNF\n" )
	// script := cnf( cs, k, ys )

	fmt.Println( script )

}