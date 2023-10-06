# depoch

Convert an epoch timestamp to a human readable ([RFC3339](https://ijmacd.github.io/rfc3339-iso8601/)) timestamp.

## Install

```sh
 go install github.com/dhulihan/depoch@latest
```

## Use

```sh
$ depoch 1696617071
2023-10-06T12:31:11-06:00

# from STDIN
echo "1696617071" | depoch
2023-10-06T12:31:11-06:00
```

## FAQ

* Why not just use `date`? Because it does not support STDIN, and the CLI flags differ between BSD/GNU.
