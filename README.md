# depoch

Convert an epoch timestamp to a human readable date.

## Install

```sh
go install github.com/dhulian/depoch
```

## Use

```sh
$ depoch 1696617071
2023-10-06T12:31:11-06:00

# from STDIN
echo "1696617071" | depoch
2023-10-06T12:31:11-06:00
```
