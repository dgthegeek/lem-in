# LEM-IN : Quick outline of the code :

## 1) First transforme the file in to a graph (you can use)

## 2) Do all the necessary check in a way to validate the graph

A validate graph must have :
- A start and end room
- Start room and end room connected
- A valid number of ants
- No cycle etc...

## 3) Find all the existing path in the graph
- We make sure to extract the start room from all the path since it's not part of the path "technically"

## 4) Group the path according to their start room

## 5) Get the smallest paths among the paths
- This will help to make the best choice (of path) during the simulation

## 6) ... some manipulations to get the good ones out of the lot ...

## 7) Launch the ants simulation !!
