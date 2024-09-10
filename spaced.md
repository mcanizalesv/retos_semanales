# Spaced repetition example

```go
// Original list
[learn, run, sleep, eat, walk]
```

### Round 1

> Question: learn
> User answer: correr

The answer is wrong, `learn` moves from index 0 to index 1.

```go
// List updated
[run, learn, sleep, eat, walk]
```

### Round 2

> Question: run
> User answer: correr

The answer is correct, Since it's the first correct answer, it moves 2^1 places. `run` moves from index 0 to index 2.

```go
// List updated
[learn, sleep, run, eat, walk]
```

### Round 3

> Question: learn
> User answer: dormir

The answer is wrong again. `learn` moves from index 0 to index 1.

```go
// List updated
[sleep, learn, run, eat, walk]
```

### Round 4

> Question: sleep
> User answer: dormir

The answer is correct. `sleep` moves from index 0 to index 2 (since it's the first correct answer, it moves 2^1 places).

```go
// List updated
[learn, run, sleep, eat, walk]
```

### Round 5

> Question: learn
> User answer: correr

The answer is wrong again. `learn` moves from index 0 to index 1.

```go
// List updated
[run, learn, sleep, eat, walk]
```

### Round 6

> Question: run
> User answer: correr

The answer is correct again. Since it's the second correct answer, it moves 2^2 places. `run` moves from index 0 to index 4.

```go
// List updated
[learn, sleep, eat, walk, run]
```

### Round 7

> Question: learn
> User answer: aprender

Finally, the answer is correct. Since it's the first correct answer, it moves 2^1 places. `learn` moves from index 0 to index 2.

```go
// List updated
[sleep, eat, learn, walk, run]
```

### Round 8

> Question: sleep
> User answer: aprender

Now, the answer is incorrect. `sleep` resets and moves from index 0 to index 1 to restart the learning process.

```go
// List updated
[eat, sleep, learn, walk, run]
```

### Extra

Let's imagine we get to `run` again and the answer is correct:

```go
// List before update
[run, eat, sleep, learn, walk]
```

> Question: run
> User answer: correr

The answer is correct again. Since it's the third correct answer, it should move 2^3 = 8 places. However, there is no index 8, so `run` moves to the end of the slice.

```go
// List after update
[eat, sleep, learn, walk, run]
```
