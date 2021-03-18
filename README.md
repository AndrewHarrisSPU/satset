# YouTube

[![YouTube](https://img.youtube.com/vi/XdQCL0HmNL4/0.jpg)](https://www.youtube.com/watch?v=XdQCL0HmNL4)

# satset

## How to build it:
From the repository directory, with Go installed and available :
```
go build
```
generates the executable `satset`.

## How to run it:

After building, and from the repository directory (ensuring the presence of `satset` and `textbook.txt`, taken from Kleinberg and Tardos), the following command solves a 3-SAT problem, for one term per clause:
```
./satset -input textbook.txt
(  v11 |  v12 |  v13 ) &
(  v21 |  v22 |  v23 ) &
(  vk1 |  vk2 |  vk3 ) &
( ~v11 | ~v21 ) &
( ~v13 | ~vk3 ) &
( ~v23 | ~vk1 )
 v11: true
 v22: true
 vk1: true
```

## satset --help

```
Usage of ./satset:
  -dot
    	write a .dot file of the solution (experimental)
  -input string
    	a plaintext file containing a list of terms in CNF
  -seconds int
    	how many seconds before giving up (default 1)
```

# Reduction to Independent Set

## How it works:

Reducing to Independent Set, 3-SAT clauses are converted to to 3-vertex, 3-edge triangles and additional constraint edges. This is a polynomial-time conversion - with *n* vertices, we couldn't have more than *O(n^2)* edges.

Solving 3-SAT approximately is "remarkably simple" (Kleinberg and Tardos, pg.725). Simple means simple: just flip the bits randomly. Gratuitously, this program sends off one hundred go routines (fire up the cores!) to flip their own bits, so it's even simple in parallel. (In the larger scale, the first 7/8ths are simple, but the last 1/8th isn't, and solving the last 1/8th might unsolve the first 7/8ths ...)

To solve this Independent Set problem, we can insist on one satisfied term per clause.

## Experimentally, using dot and Graphviz to visualize solutions
Solutions may emit a .dot file, which can be rendered to a .png image or .svg asset with the `dot` tool (part of [ Graphviz ]( https://graphviz.org )). For example, after `satset` has found a solution:
```
dot -Tpng textbook.dot -o textbook.png
```
generates a .png representation of the solution. The visualization red circles to indicate parts of a solution set. Dashed lines represent a contradiction, and a contradiction may run through a few steps.

A quick gallery:

- textbook (Given in Kleinberg and Tardos, pg. 461)
```
v11 | v12 | v13
v21 | v22 | v23
vk1 | vk2 | vk3

~v11 | ~v21
~v13 | ~vk3
~v23 | ~vk1
```
![ textbook ]( textbook.png?raw=true )

- slides1 (CSC3430-NP slides, 35-37)
```
~x11 | x21 | x31
~x22 | x12 | x32
~x13 | x23 | x43

~x11 | x12
~x13 | x12
x21 | ~x22
x23 | ~x22
```
![ slides1 ]( slides1.png?raw=true )

- slides2 ( CSC3430-NP slides, 38-40)
```
x11 | ~x11 | ~x21
x32 | x22 | x42
~x13 | ~x33 | ~x43

~x21 | x22
~x13 | x11
x32 | ~x33
x42 | ~x43
```
![ slides2 ]( slides2.png?raw=true )