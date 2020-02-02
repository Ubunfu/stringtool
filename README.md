# stringtool
Cobra app for experimenting with Go and stuff like that

1000 YEARS ARCTIC CODE VAULT!

[![CircleCI](https://circleci.com/gh/Ubunfu/stringtool.svg?style=svg)](https://circleci.com/gh/Ubunfu/stringtool)

## Help
View the built-in man page for this app by running ```stringtool help```.  To see the man page for a specific command, run ```stringtool help <command>```.

## Sub-Commands

### reverse
```
Reverse a string lexigraphically

Usage:
  stringtool reverse [flags]

Aliases:
  reverse, r

Flags:
  -h, --help            help for reverse
  -s, --string string   String to reverse
```

### enumerate
```
Enumerate all alpha-numeric strings of the given length parameters.
		Brute-force style.

Usage:
  stringtool enumerate [flags]

Aliases:
  enumerate, e

Flags:
  -b, --begin string     Starting point for enumeration. E.g. Ry4
  -e, --end string       Ending point for enumeration. E.g. ccc
  -f, --flushInterval int   Number of strings that will be enumerated before flushing to disk (default 10000)
  -h, --help             help for enumerate
  -x, --max-length int   Maximum string length (default 3)
  -n, --min-length int   Minimum string length (default 1)
  -o, --output string    Output file for enumerated strings (default "strings.out")
```

### hash
```
hash strings from a file and writes the {string, hash} pairs to a file

Usage:
  stringtool hash [flags]

Aliases:
  hash, h

Flags:
  -a, --algorithm string   Algorithm to use for hashing the strings: [ md5 | sha1 | sha512 ]
  -e, --encoding string    Encoding to use for writing the hashed strings: [ hex | base64 ] (default "hex")
  -f, --flushInterval int   Number of strings that will be hashed before flushing to disk (default 10000)
  -h, --help               help for hash
  -i, --in-file string     File with strings to hash
  -o, --out-file string    File to write the strings and hashes (default "hashes.out")
  -r, --rounds int         Number of rounds of hashing to perform on the input strings (default 1)
```
