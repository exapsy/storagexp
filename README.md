# Storage read/write experiment

This is a simple experiment to test the storage read/write speed
on Memory and Disk.

## Memory

The memory read/write speed is tested using a string hashmap

## Disk

The disk read/write speed is tested creating a file with
name `<key>` and content `<value>`.

That way we can avoid any complexities from file format and marshalling,
to test purely only the read/write performance on disk and nothing else.


## Results

```
Memory Repo
        Write Time taken: 13.95933ms
        Read Time taken: 25.408422ms

        Total Time: 39.367752ms
        Average Time: 19.683876ms
--------------------------------
File Repo
        Write Time taken: 15.859788131s
        Read Time taken: 59.67467994s

        Total Time: 1m15.534468071s
        Average Time: 37.767234035s
--------------------------------
```
