# Performance 

## CPU benchmarks 

### Iteration 1 - Random games
Altough this is iteration 1, I have made quite some changes since starting. A lot of optimizations have been performed. 
The main tradeoff is that starting this I tried to make it really readable and object oriented. While I'd argue this is still mainly the case, 
I have changed some code to use small datatypes, comparison using ints rather than strings, and reusing objects where I can.

Using the built in benchmarking and pprof analytics tool I will note down the performance of this for tracking 

After some opimizations, mostly lookups by int8 values and reusing objects as much as possible the performance was improved. 

Results 1 000 000 games, 10 Iteration
Total games played  = 10 000 000
Total duration: 1,145s
``` 
BenchmarkPlayRandomGames-24           10         103292585 ns/op
PASS
ok      rubrumcreation.com/go-paper-scissor/tests       1.145s
```

Further optimizations would compromize the readability and overall "neatness". But if it's neccessarry later I will have to make changes.

### Random games:
These are the optimizations done for random moves simulation: 
Paramenters: 
Two players who plays random moves 
Playing  1 000 000 games

-  Currently before further optimizations, 1 million games are played and summarized in 0,118 seconds
-  After optimizations the random array of moves to a global static array, the new runtime is 0,112
-  After further optimizations by reusing player objects: 0,1033 (one million games)
-  Trying to use a map is actually slower than using a switch case to check for values. Even though it looks less "elegant" than a map lookup for example. Using the map lookup took 0,131 seconds

This switch 

```
switch {
	case m1 == m2:
		return 0 // draw
	case (m1.GetType() == models.ROCK && m2.GetType() == models.SCISSOR) ||
		(m1.GetType() == models.PAPER && m2.GetType() == models.ROCK) ||
		(m1.GetType() == models.SCISSOR && m2.GetType() == models.PAPER):
		return 1 // move 1 wins
	default:
		return 2 // move 2 wins
	}

```
is more performant than: 
```

var outcomes = map[models.MoveType]map[models.MoveType]int{
    models.ROCK: {
        models.ROCK:    0,
        models.PAPER:   2,
        models.SCISSOR: 1,
    },
    models.PAPER: {
        models.ROCK:    1,
        models.PAPER:   0,
        models.SCISSOR: 2,
    },
    models.SCISSOR: {
        models.ROCK:    2,
        models.PAPER:   1,
        models.SCISSOR: 0,
    },
}
```
### API 

### Iteration 1 - Barebones webserver 
Current state of the webserver: 
For 10 random games requested by the api, using: curl -o /dev/null -s -w 'Total: %{time_total}s\n'  localhost:8080/random/test 
The total time to response was: 0,000425s 

Using browser (chrome) the total time waiting for server response was
0,00142 seconds

For now this is OK. Testing the performance of the webserver itself is trick when you take latency, internet connection speed etc. 
It can be tweaked, but not neccessary for now. 
