# Readme

I know there are algorithms to solve these puzzles (likely in a better way), but I went at this without researching any
of them to test my own capabilities to take a reasonably complex logical problem into some algorithms.

# TODO

* Board Org
    * ~~Create Board Object~~
    * ~~Create Cell Object~~
    * ~~Create Column Getter~~
    * ~~Create Row Getter~~
    * ~~Create Sector Getter~~
    * Create Segment Getter
        * Parameters (SectorX, SectorY, direction (Y or X), number int (1-3))
        * Will retrieve a segment of 3 cells within a sector
* Solutions
    * ~~All-But-One Filler~~
        * ~~if you have a row, column, or sector with 8/9 values, you can be certain that the remaining cell must be the
          remaining value~~
    * ~~Cross Fill on 2 sectors with value~~
        * ~~If you have 2 segments with a value, and there is only 1 possible cell in the 3 related segment~~
        * ~~eg: the 3 in [3,8] is horizontally cross filled because of the 3s in [5,9] and [7,7] leaving only that
          cell~~
    * Cross Fill on 1 sector with value, 1 sector with blocked segment
        * Example: (in current state)
            * The 2 in [8,2] blocks cells [4,2], [6,2], and [3,2] from being 2
            * This means that the top middle sector must have a 2 in either [4,3] or [5,3] - which are both in row 3
            * The 2 in [1,4] blocks cells [1,3] and [1,1] from being 2
            * This leaves [2,3] and [3,1] as possible 2's in the top left sector
            * because the top middle sector's 2 is definitely going to be in row 3, we can further eliminate [2,3] from
              possible 2s, leaving only [3,1]
          