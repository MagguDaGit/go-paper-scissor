## TODO: 
1. Look for optimizations to how the games are played 

2. Make a web server (Barebones) ✅ 
    2.1 - Make it for different types of games (For now you can simluate an amount of random games, you can play against random moves, but there is alot of random) 
    2.2 - Make it FASTER (Games takes 0,00007 seconds, but the response takes 2 seconds? ) ✅ (won't fix, this was firefox delaying on first response. The expected speed is achived using curl 
or chrome)

3. Make other types of games: 
    3.1 - Queued games, where a number of rounds is specified, and executed in quick succession ✅ - This is done for /simulation/random
    3.2 - Interactive games: Each round is played one by one and each round waits for a input by user - ✅ This is done for /play/random

4. Make other types of opponents: 
    4.1 - Opponent where the moves are based on a algorithm  for the best possible chance to win (Might need to make "opponent" logic for this)
    4.2 - Opponent where the moves are random ✅

5. Add basic logger service: 
    5.1 - Add basic info logging for api endpoints ✅ 
    5.2 - Add basic error logging for service endpoint ✅
    5.3 - Add store solution for logs

6. Speed up games by using coroutines 
    6.1 - Add the games in chunks, and delegate the chunk of games to a coroutines - ✅ This is now done when games are over 100 000 
    6.2 - Goal: Simulate 1 billion games. Right now this causes an out of memory exception. Need to look at the memory usage 

7. Stress test the api by thousands of request
    7.1 - Actions will be created by the outcome of said test

8. Add a error object to return to client - ✅ 


## Optimizations 
Look in performance.md ()[./performance.md]



