# A new way to encode binary numbers
I Ching or "Classic of Change" is a Chinese book about telling the future or something. It doesn't matter.

What matters is that I Ching defines a set of 64 "figures", called hexagrams, each consisting of 6 stacked horizontal lines, where each line is either a solid, unbroken line (Yang) or a line with a single gap (Yin).

Each hexagram can therefore be easily used to encode a number up to 6bits of entropy, this library implements a function (`Itoiching`) to convert an arbitrary 64bit unsigned integer into its I Ching representation and a function (`Ichingtoi`) to do the reverse.

## Conventions

Each I Ching character represents a 6 bit value, strings are ordered with the most significant 6 bits of the number first. Withing a I Ching character the top line represents the most significant bit of the sextet.
The unbroken line (Yand) corresponds to 0, the broken line corresponds to 1. We do not use the King Wen sequence.

## Example

```
  iching.Itoiching(23) → ䷢
```

## Example programs

This library comes with two example programs (available under the `cmd` directory). Ichingtest takes a number as argument on the command line and converts it to its representation in decimal, octal and I Ching.
The other program is ichingdump and it's similar to hexdump (or od) but using I Ching encoding for all numbers.

![Crystal clear](crystalclear.png?raw=true "Screenshot of Ichingdump")
