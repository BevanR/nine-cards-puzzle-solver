# Nine cards puzzle solver

Finds solutions to the nine cards puzzle by Portel—or at least [the light bulbs variation of it](https://www.portel.ee/en/pood/lightbulbs).

The puzzle has nine square tiles, which must be arranged in a 3-by-3 grid such that every light bulb that neighbours another tile is complete.

### An invalid arrangement from Portel's website

![image](https://github.com/user-attachments/assets/1e665c1e-5d06-46ea-ac80-3446849439c1)

### A valid solution, found using this tool

![image](https://github.com/user-attachments/assets/890b4908-36ab-4d8f-b8c6-fe939bf7094c)

There are nearly 24 billion possible arrangements;

| Variable                                                                                              | Quantity           | Total            |
|-------------------------------------------------------------------------------------------------------|--------------------|------------------|
| 9 cards can be arranged in 9 different positions                                                      | 9!                 | 362,880          |
| 4 orientations of each of the 9 cards/positions                                                       | 4<sup>9</sup>      | 262,144          |
| Total                                                                                                 | 9! ✖ 4<sup>9</sup> | = 95,126,814,720 |
| Each arrangement can be specified in four different orientations, meaning only ¼ of those are unique. | 9! ✖ 4<sup>8</sup> | = 23,781,703,680 |

A brute force of trying all 24B arrangements would take a long time, [even for a computer](https://github.com/dcreemer/scramble#:~:text=naive%20golang%20solution%20took%20almost%2022%20minutes%20to%20find%20all%20of%20the%20solutions).

The optimal approach is to;

1. For a given position (start with the first)
2. Check if the first unused tile fits
3. If it does, move to the next position (recursively) and repeat
4. If it does not fit, or if after checking the following positions there are no valid solutions with this tile in this position, rotate it a quarter turn to the next orientation
5. Once all four orientations have been exhausted, try the next unused tile in this position

There are 3 unique solutions for this light bulbs variation of the puzzle. It doesn't currently support other puzzles.

![image](https://github.com/user-attachments/assets/e7e3877c-c210-43f9-8596-c9fb30f11b44)
