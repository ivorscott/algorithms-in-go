# Algorithms in Go

Good code is 
1. readable
2. scalable

There are many ways to solve a problem, some are more efficient than others.

## Big O

Big O is about measuring the scalability of code in terms of time and space
complexity. Scalability is a concern for large companies with lots of users 
because as inputs grow, bad algorithms can slow an application down.

Big O asks: 

- As the input increases, how much does an algorithm slow down.

![Screenshot 2023-11-19 at 13.07.58.png](..%2F..%2F..%2Fvar%2Ffolders%2F3t%2Fpmlsll2j6sb5fyz4j34xlc7h0000gn%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_BFnN6C%2FScreenshot%202023-11-19%20at%2013.07.58.png)


### Time complexities

- O(1) - constant time [assignments,retrieval,return statements] (one operation)
  - The number of steps required to execute an algorithm does not change with the change in
  the input size. For example:
    - push and pop operations of a stack
    - return the first element of a list
    - getting an element from a hash table
- O(n) - linear time [loops] (as the input increases so does the number of operations)
  - Execution time of an algorithm is directly proportional to its input size.
    - Searching for an element in array
    - Traversal of a linked list to find the maximum or minimum element
- O(log n) - Logarithmic time (searching algorithms have log n if they are sorted)
  - Execution time of an algorithm proportional to the logarithm of the input size. Half of the input is 
    - pruned out without traversing it at each algorithm stage.
      - Binary search
- O(n log n) - Log Linear time (usually sorting operations)
- O(n^2) - Quadratic time [double loop] (every element in a collection needs to be compared to every other element)
  - Execution of an algorithm is directly proportional to the square root of the input size.
  Each element is compared to all other elements in these algorithms. For example, in two nested loops,
  - if the outer loop runs n times and the inner loops runs n times, then the body of the inner loop will execute n^2 times.
- O(2^n) - Exponential time [complex full search] (recursive algorithms that solves a problem of size N)
  - all possible permutation
- O(n!) - Factorial time (adding a loop for every element)


- O(1) - Excellent/Best
- O(log n) - Good
- O(n) - Fair
- O(n log n) - Bad
- O(n^2), O(2^n) and O(n!) - Horrible/Worst

Iterating through half a collection is stall `O(n)`. 

O(a*b) - Two separate collections (give an identifier for each input)

### Calculating time complexity

You can identify the cost of each line but you never need to do this is an interview.
Normally you only need to determine the worst case big O notation for the program. However,
it's good to see how things are calculated. 

Remember to drop the constants to simply terms in the end. For example: given 6 constant time operations and
3 linear time operations is you would have 

```O(6 + 3n) Or simply: O(n).```

Once we drop the constants we see`O(n)` linear time complexity is the worst case. Also given a program with
`O(n + log n)` we drop `log n` as a non-dominant term. This is because we only care for the upper bound. So 
`log n` can be omitted because it grows more slowly than `n`. We only keep the dominant term.

### Rules

1. Determine The Worst Case
2. Remove Constants
3. Use Different Terms For Inputs
4. Drop Non-dominants

<hr>

What is the time complexity of the following program:

```go 
package main

func doSomething(input []string) int {
	foo := 100 //O(1)
	
	for i := 0; i < len(input); i++ { //O(n) 
		doSomethingElse() //O(n)
		foo = foo + 50 //O()n
	}
	return foo //O(1)
}

func main() {
    doSomething([]string{"a","b","c"})
}

// Simplify terms
// Big O(2 + 3n) = O(n)
```

What is the time complexity of the following program:

```go 
package main

func doSomething(input, input2 []string) {
	for _, v := range input { //O(n) 
		print(v) //O(n)
	}

	for _, v := range input2 { //O(n)
		print(v) //O(n)
	}
}

func main() {
    doSomething([]string{"a","b","c"},[]string{"d","e","f"})
}

// Different terms for inputs
// Big O(a + b)

// If the loops were nested:
// Big O(a * b)
```

### What causes time in a function

- Operations (+,-, \*, /)
- Comparisons (<, >, ===)
- Looping (for, while)
- Outside Function call (function())


### Space Complexity

We look at the total size of memory used relative to the input. We try to determine how 
much memory is being used.

heap - store variables
stack - keep track of function calls


What is the space complexity of the following program:

```go 
package main

func doSomething(input []string) {
	for i := 0; i < len(input); i++ { //O(1) (creates 1 variable)
		print("test")
	}
}

func main() {
    doSomething([]string{"a","b","c"})
}

// Big O(1)
```
What is the space complexity of the following program:

```go 
package main

func doSomething(input []string) []string {
	var results []string //O(1) 
	for i := 0; i < len(input); i++ { //O(1) (creates 1 variable)
		results = append(results, "test") //O(n) (appends n strings)
	}
	return results
}

func main() {
    doSomething([]string{"a","b","c"})
}

// Big O(2 + n) = O(n)
```

### Master Theorem

The Master Theorem helps us determine the time complexity (asymptotic notation) of a recurrence relation.
A recurrence relation is an equation that uses recursion.

![master-theorem-definition.png](doc%2Fimg%2Fmaster-theorem-definition.png)

There are 3 cases: 

1. Work to split/recombine a problem is dwarfed by subproblems.
   i.e. the recursion tree is leaf-heavy

2. Work to split/recombine a problem is comparable to subproblems.

3. Work to split/recombine a problem dominates subproblems.
   i.e. the recursion tree is root-heavy.

![master-theorem.png](doc%2Fimg%2Fmaster-theorem.png)

Resources

- https://www.educative.io/courses/grokking-the-coding-interview
- https://www.educative.io/courses/decode-the-coding-interview-go
- https://github.com/TomorrowWu/golang-algorithms
- https://github.com/Hemant-Jain-Author/Data-Structures-Algorithms-In-Go
- https://www.freecodecamp.org/news/big-o-cheat-sheet-time-complexity-chart/
- https://en.wikipedia.org/wiki/Master_theorem_(analysis_of_algorithms)