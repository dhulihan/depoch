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
$ echo "1696617071" | depoch
2023-10-06T12:31:11-06:00

# reverse convert a time string to an epoch
# WARNING: time zone may not be maintained, depending on format.
$ echo "Mon Jan 02 15:04:05 -0700 2006" | depoch
1136239445
```

## FAQ

* Why not just use `date`? Because it does not support STDIN, and the CLI flags differ between BSD/GNU.
