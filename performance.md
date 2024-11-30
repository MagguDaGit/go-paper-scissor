# Performance 

## CPU benchmarks 

### Iteration 1 - Random games
Altough this is iteration 1, I have made quite some changes since starting. A lot of optimizations have been performed. 
The main tradeoff is that starting this I tried to make it really readable and object oriented. While I'd argue this is still mainly the case, 
I have changed some code to use small datatypes, comparison using ints rather than strings, and reusing objects where I can.

Using the built in benchmarking and pprof analytics tool I will note down the performance of this here for tracking 

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
