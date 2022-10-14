# What to do

In this CodeKata I wanted to write a sudoku solver. The last time I solved a sudoku I thought this should be pretty strate forward... well... it's a lot of matrix, numbers, calculation chaos and with the "simple" way that I also solve sudokus myself it's not even possible to solve an Level "Evil" Sudoku from sudoku.com with my algorithm. But let's see what I have done!

# Solve Sudoku - My way

The "custom" way would be my way, so here is how I solve sudokus (kind of). This is how my algorithm works:

1. Board-Analysis: Check every cell if there is only one possible solution
   1. I do this by starting with 1, 2, 3, 4, 5, 6, 7, 8, 9 as possible solutions and remove everything I find in the row, the column and in the 3x3 section of the cell I'm looking at
   2. If there is only one possible solution left, then thats it. I update the board and restart the Board-Analysis. (I do this because the board updated and it could affect other cells that I already looked at. Restarting is a much simpler but less effective way than looking recursivly at only the affected cells.)
   3. A simple Sudoku can already be solved with this attempt. For anything a little harder the next step is executed
2. Possible-Solutions-Analysis: Check every row, cell and 3x3 section if there is a missing value that only fits in 1 cell
   1. In the Board-Analysis we sometimes miss a solution because one cell might could have 2 or more possible solutions. Anyway if we look at a row, column or 3x3 section it can happen, that a cell can hold 2 possible solutions but there is no other cell in the row, column or 3x3 section that can hold on of those 2 possible solutions. In that case the cell is the only possible holder for one of its 2 solutions but it must hold the solution that no one else can hold
3. Finally this starts over and over in a loop until one of the following events occure:
   1. The board has no empty cells anymore and is valid => Victory!
   2. The board becomes invalid => Whoops! Go bughunting
   3. The board does not change between two iterations anymore => Oh no! The algorithm found its limit...

# Solve Sudoku - The brute force way

Well, we have computers that are pretty powerfull, so why should we think if we can just try a million times per seconds?
For this solution I took [this Medium Article](https://medium.com/@nticaric/writing-a-sudoku-solver-in-go-3f6539eafcc5) as inspriation. It's called "Backtrace" and works pretty simple:

1. Find the first empty cell in the board and put a number in it (e.g. start with 1)
2. Check if the board is still valid
   1. If the board is still valid try the the same with the next empty cell
   2. If the board is not valid anymore reset the cell that invalidated the board and try the next number (e.g. 2)
3. Do this recursively until the board has no more empty cells and is valid
