## TODO: 
1. Look for optimizations to how the games are played 
2. Make a web server (Barebones) ✅ 
    2.1 - Make it for different types of games, 
    2.2 - Make it FASTER (Games takes 0,00007 seconds, but the response takes 2 seconds? )
3. Make other types of games: 
    - Queued games, where a number of rounds is specified, and executed in quick succession 
    - Interactive games: Each round is played one by one and each round waits for a input by user
4. Make other types of opponents: 
    - Opponent where the moves are based on a algorithm  for the best possible chance to win (Might need to make "opponent" logic for this)
    - Opponent where the moves are random ✅

## Optimizations 
Look in performance.md ()[./performance.md]


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


