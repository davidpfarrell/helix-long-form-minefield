# helix-long-form-minefield

### Prolog

So this repository is the result of a coding test I did for an interview several years ago. I've kept as a private repo all that time, and have decided its either time to delete it or make it public.

I'm rather happy with the solution so I decided to make it public (and archvied) and keep it around.

---------
### Helix :: Coding Exercise :: Long-Form :: Minefield

Write a program which takes as input a list of mines composing a 2D minefield; each mine has an X position, a Y position, and an explosive power. All three parameters may be assumed to be single-precision floats; explosive power may not be negative. There may not be more than one mine at the same coordinates.

When a mine in the minefield is triggered at time T=0, it causes all other mines within a straight-line distance less than or equal to its explosive power to be triggered at time T=1. Those mines subsequently trigger additional mines at T=2, and soforth, in a chain reaction.

Have your program determine, for any given input minefield, the mine that, if triggered first, will result in the highest number of explosions occurring during a single time interval. Output the coordinates of the winning mine, the time interval of the peak number of explosions, and the number of explosions during that interval. In case of a tie, output each of the best mines, sorted by X coordinate then Y coordinate.

Assume that the minefield may be large, but not larger than can easily fit in memory; optimize for processing efficiency.

-----------
### License

This project is released under the [Unlicense](https://unlicense.org/) License.  See `LICENSE` file.
