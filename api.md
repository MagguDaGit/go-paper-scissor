# API docs 


## Routes 

### /simulation/random?amount={amount}
Returns a summary of a simulation of games where all moves used are pseudo random. 
A amount can be specified. If no amount is specifies the default is 100 



### /play/random?move={move}
Will play a game, where the opponent uses random moves. 
Will play a singe game and return the result as a summary object 
The move query parameter is required to play, if not included a error will be returned


### /play/smart 
Not yet implemented, need to have a thunk about this one...                                                                                                                                                                                                                              

