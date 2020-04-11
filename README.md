# go-monty-hall

> Suppose you're on a game show, and you're given the choice of three doors: Behind one door is a car; behind the others, goats. You pick a door, say No. 1, and the host, who knows what's behind the doors, opens another door, say No. 3, which has a goat. He then says to you, "Do you want to pick door No. 2?" Is it to your advantage to switch your choice?

The answer:

## Thou shalt switch the door!

While there are mathematical explanations on [Wikipedia](https://en.wikipedia.org/wiki/Monty_Hall_problem), proving it empirically is more satisfying and convincing to me. Before we explain the solution, here are certain assumptions:
  1. The host must always open a door that was not picked by the contestant.
  2. The host must always open a door to reveal a goat and never the car.
  3. The host must always offer the chance to switch between the originally chosen door and the remaining closed door.

## The Experiment:

We create a function which will take a parameter stating whether we swap our choice (`switch` is a reserved golang keyword :P) We generate two random numbers (less than 3), one is where the prize is and other is our choice.

```golang
func win(swap bool) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prize := r.Int() % 3
	choice := r.Int() % 3
    ...
```

There are are two possibilities once you choose a door
  1. You chose correct door
  2. You chose a door with the goat.

### 1. When you picked the correct door initially

```golang
        ...
	if prize == choice {
		if swap {
			return false
		} else {
			return true
		}
	}
    ...
```

This is a rather easy one. If we swap on this case, because we chose correct for initially, we lose. If we don't swap, we win.

### 2. You chose a door with the goat

This is a tricky one. Go back to assumptions 1 and 2 we stated earlier. As per assumption 1, host does not open the door you chose (which is wrong in this case). Out of other two doors, one of which is correct and one is not. According to assumption 2, host will open the door to reveal goat. So the door that remains (apart from one you chose and one the host opened), has the car. Code would look something like this

```golang
    ...
    if swap {
        return true
    }
    return false
}
```

Because you chose goat door initially, and host also revealed the goat door, if you swap to the only remaining door, you win. Else, you lose.

## The Trial

We would run a 100k trials, **with swap**, and see how many time we would win overall.

```golang
func main() {
	trials := 100000
	successSwap := 0
	swap := true
	for i := 0; i < trials; i++ {
		if win(swap) {
			successSwap++
		}
	}
	fmt.Printf("Success with swap: %v is %d%%",swap, successSwap*100.0/trials)
}
```

The result: we win 66% of times, if we swap.
```
$> go build -o play main.go
$> ./play
Success with swap: true is 66%
```
