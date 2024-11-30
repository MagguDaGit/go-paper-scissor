## TODO: 
1. Look for optimizations to how the games are played 
2. Make a web server
3. Make other types of games: 
    - Queued games, where a number of rounds is specified, and executed in quick succession 
    - Interactive games: Each round is played one by one and each round waits for a input by user
4. Make other types of opponents: 
    - Opponent where the moves are based on a algorithm  for the best possible chance to win (Might need to make "opponent" logic for this)
    - Opponent where the moves are random ✅






## Optimizations 

### Random games:
These are the optimizations done for random moves simulation: 
Paramenters: 
Two players who plays random moves 
Playing  1 000 000 games

-  Currently before further optimizations, 1 million games are played and summarized in 0,118 seconds
-  After optimizations the random array of moves to a global static array, the new runtime is 0,112
-  After further optimizations by reusing player objects: 0,1033 (one million games)



