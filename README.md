# Monty Hall Stats
Calculate the stats of the Monty Hall problem.

## What is the Monty Hall problem?
> Suppose you're on a game show, and you're given the choice of three doors: Behind one door is a car; behind the others, goats. You pick a door, say No. 1, and the host, who knows what's behind the doors, opens another door, say No. 3, which has a goat. He then says to you, "Do you want to pick door No. 2?" Is it to your advantage to switch your choice?
- [Wikipedia](https://en.wikipedia.org/wiki/Monty_Hall_problem)

Your intuition might say there would be no reason to switch the doors since the odds should be 1/2 of recieving the car since there would only be 2 doors left, but that assumption would be wrong as the chance is actually locked in at the begining of the game, so keeping the door would result in a 1/3 chance of recieving the car, while switching the door results in a 2/3 chance of winning the car. 

## Config

Parameters to change

In the begining of the main function 2 variables are present, defining the outcome of the game

Those are `switchDoor` and `gamesToPlay`

Change switchDoor to false, if you don't want to switch doors and change gamesToPlay to if you want to play it a diffrent amount of games each time.


## Running
To run the program you need to have go installed, and then simply run `go run .` in your prefered command line inside the project directory.

## Output
To save you the time of actually running this i will present the outputs of 2 runs of the program with 1000000 games played in each run below

If I chose to keep the original door
```
In 100000 games
Games won: 22263
Games lost: 77737
This gave you an average of 22.26% percent chance of winning if you kept the original door
```

If i chose to switch doors
```
In 100000 games
Games won: 66624
Games lost: 33376
This gave you an average of 66.62% percent chance of winning if you switched the door
```


