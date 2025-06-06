# Nine cards puzzle solver

Finds solutions to the nine cards puzzle by Portel—or at least [the light bulbs variation of it](https://www.portel.ee/en/pood/lightbulbs).

The puzzle has nine square tiles, which must be arranged in a 3-by-3 grid such that every light bulb that neighbours another tile is complete.

## An invalid arrangement from Portel's website

![image](https://github.com/user-attachments/assets/1e665c1e-5d06-46ea-ac80-3446849439c1)

## A valid solution, found using this tool

![image](https://github.com/user-attachments/assets/890b4908-36ab-4d8f-b8c6-fe939bf7094c)

## How many arrangements?

There are nearly 24 billion possible arrangements;

| Variable                                                                                              | Quantity           | Total            |
|-------------------------------------------------------------------------------------------------------|--------------------|------------------|
| 9 cards can be arranged in 9 different positions                                                      | 9!                 | 362,880          |
| 4 orientations of each of the 9 cards/positions                                                       | 4<sup>9</sup>      | 262,144          |
| Total                                                                                                 | 9! ✖ 4<sup>9</sup> | = 95,126,814,720 |
| Each arrangement can be specified in four different orientations, meaning only ¼ of those are unique. | 9! ✖ 4<sup>8</sup> | = 23,781,703,680 |

A brute force of trying all 24B arrangements would take a long time, [even for a computer](https://github.com/dcreemer/scramble#:~:text=naive%20golang%20solution%20took%20almost%2022%20minutes%20to%20find%20all%20of%20the%20solutions).

## Approach

The optimal approach is to;

1. For a given position (start with the first)
2. Check if the first unused tile fits
3. If it does, move to the next position (recursively) and repeat
4. If it does not fit, or if after checking the following positions there are no valid solutions with this tile in this position, rotate it a quarter turn to the next orientation
5. Once all four orientations have been exhausted, try the next unused tile in this position

The implementation iterates through orientations (rotation) before iterating over tiles to make the code comments clearer.

As an optimisation and deduplication step, the middle tile can be never-rotated. This avoids discovering duplicate solutions that differ only be rotating the entire solution.

There may be some micro optimisation by changing the order positions are iterated through to optimise comparing edges vs placing tiles. This implementation simply goes top left (position 0) to bottom right (9). However filling these positions first may be faster;

|   |   |   |
|---|---|---|
| X |   | X |
|   | X |   |
| X |   | X |

Or perhaps;

|   |   |   |
|---|---|---|
| X | X |   |
| X |   |   |
|   |   | &nbsp; |


## Solutions

There are 3 unique solutions for this light bulbs variation of the puzzle. It doesn't currently support other puzzles.

![image](https://github.com/user-attachments/assets/e7e3877c-c210-43f9-8596-c9fb30f11b44)

Tile number by position, including the orientation of North/top edge of tile.

| Position | Tile | Orientation |
|----------|------|-------------|
| 0        | 1    | N           |
| 1        | 4    | W           |
| 2        | 3    | E           |
| 3        | 7    | W           |
| 4        | 5    | N           |
| 5        | 9    | S           |
| 6        | 2    | E           |
| 7        | 8    | W           |
| 8        | 6    | S           |

| Position | Tile | Orientation |
|----------|------|-------------|
| 0        | 2    | N           |
| 1        | 1    | W           |
| 2        | 4    | E           |
| 3        | 8    | W           |
| 4        | 5    | N           |
| 5        | 9    | S           |
| 6        | 3    | E           |
| 7        | 7    | W           |
| 8        | 6    | S           |

| Position | Tile | Orientation |
|----------|------|-------------|
| 0        | 1    | W           |
| 1        | 7    | S           |
| 2        | 9    | S           |
| 3        | 5    | N           |
| 4        | 2    | N           |
| 5        | 6    | S           |
| 6        | 4    | S           |
| 7        | 8    | W           |
| 8        | 3    | S           |

### Positions reference

|   |   |   |
|---|---|---|
| 0 | 1 | 2 |
| 3 | 4 | 5 |
| 6 | 7 | 8 |
