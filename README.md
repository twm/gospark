# spark
### sparklines for your shell

See? Here's a graph of your productivity gains after using spark: ▁▂▃▅▇

This is a compatible reimplementation of the [original spark][original] in the
[Go programming language][golang].  Install it with:

```sh
› go install github.com/twm/gospark
```

This will give you a `gospark` command.

To run the tests, compile it by typing `go build -o spark` and then run
`./test`.  Note that one of the tests fails, likely due to slightly different
rounding.  As the result is visually indistinguishable I haven't really
bothered to fix it.

The original README follows.

## install

spark is a [shell script][bin], so drop it somewhere and make sure it's added
to your `$PATH`. It's helpful if you have a super-neat collection of dotfiles,
[like mine][dotfiles].

## usage

Just run `spark` and pass it a list of numbers (comma-delimited, spaces,
whatever you'd like). It's designed to be used in conjunction with other
scripts that can output in that format.

    spark 0 30 55 80 33 150
    ▁▂▃▅▂▇

Invoke help with `spark -h`.

## cooler usage

There's a lot of stuff you can do.

Number of commits to the github/github Git repository, by author:

```sh
› git shortlog -s |
      cut -f1 |
      spark
  ▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▃▁▁▁▁▁▁▁▁▂▁▁▅▁▂▁▁▁▂▁▁▁▁▁▁▁▁▂▁▁▁▁▁▁▁▁▁▁▁▁▁▁
```

Magnitude of earthquakes over 1.0 in the last 24 hours:

```sh
› curl http://earthquake.usgs.gov/earthquakes/catalogs/eqs1day-M1.txt --silent | 
  sed '1d' |
  cut -d, -f9 |
  spark
  ▅▆▂▃▂▂▂▅▂▂▅▇▂▂▂▃▆▆▆▅▃▂▂▂▁▂▂▆▁▃▂▂▂▂▃▂▆▂▂▂▁▂▂▃▂▂▃▂▂▃▂▂▁▂▂▅▂▂▆▆▅▃▆
```

Code visualization. The number of characters of `spark` itself, by line, ignoring empty lines:

```sh
› awk '{ print length($0) }' spark |
  grep -Ev 0 |
  spark
  ▁▁▁▁▅▁▇▁▁▅▁▁▁▁▁▂▂▁▃▃▁▁▃▁▃▁▂▁▁▂▂▅▂▃▂▃▃▁▆▃▃▃▁▇▁▁▂▂▂▇▅▁▂▂▁▇▁▃▁▇▁▂▁▇▁▁▆▂▁▇▁▂▁▁▂▅▁▂▁▆▇▇▂▁▂▁▁▁▂▂▁▅▁▂▁▁▃▁▃▁▁▁▃▂▂▂▁▁▅▂▁▁▁▁▂▂▁▁▁▂▂
```

Since it's just a shell script, you could pop it in your prompt, too:

```
ruby-1.8.7-p334 in spark/ on master with history: ▂▅▇▂
›
```

## wicked cool usage

Sounds like a wiki is a great place to collect all of your 
[wicked cool usage][wiki] for spark.

## ▇▁ ⟦⟧ ▇▁

This is a [@holman][holman] joint.

[original]: https://github.com/holman/spark
[golang]:   http://golang.org/
[dotfiles]: https://github.com/holman/dotfiles 
[bin]:      https://github.com/holman/spark/blob/master/spark
[wiki]:     https://github.com/holman/spark/wiki/Wicked-Cool-Usage
[holman]:   https://twitter.com/holman
