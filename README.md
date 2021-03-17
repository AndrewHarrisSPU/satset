# satset

## How to build it:
No external dependencies were used. From the repository directory, with Go installed:
```
go build
```
generates the executable `satset`.

## How to run it:

To demonstrate solving an independent set problem:
```
./satset -input textbook.txt -k 2 -s 10 -dot
solving for:  vk2 v12 v13 v21 vk1 v11 v22 v23 vk3
stumped!
```

We can try a bigger 'k':
```
./satset -input textbook.txt -k 3 -s 10 -dot
solving for:  v11 v12 v13 v21 vk3 v22 v23 vk1 vk2
 vk2: true
 v11: true
 v22: true
```

```
Usage of ./satset:
  -dot
    	write a .dot file of the solution. OPTIONAL
  -input string
    	a plaintext file containing a list of terms in CNF
  -k int
    	solve for k terms evaluating true. 0 for anything that solves. (default -1)
  -s int
    	how many seconds before giving up (default 1)
```

## dot and Graphviz
Solutions may emit a .dot file, which can be rendered to a .png image or .svg asset with the `dot` tool (part of [ Graphviz ]( https://graphviz.org )). 

I'm not sure this is entirely robust. The visualization uses dashed lines to represent contradictions and red circles to indicate parts of a solution set.

A quick gallery:

[ slides ]( slides.png )
[ simple ]( simple.png )
[ textbook ]( textbook.png )