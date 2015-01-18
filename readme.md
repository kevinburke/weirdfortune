# weirdfortune(6)

Have you ever wished the fortunes in the Unix [`fortune`][fortune] command were
actually funny? Well here is your chance.

## Installation

```sh
git clone git@github.com:kevinburke/weirdfortune.git
cd weirdfortune
./install.sh
```

## Alternate Go Installation

I rewrote the runner in Go since the Python version takes about 120ms to
execute. The Go version executes in about 12ms. To install it, run:

```sh
go install github.com/kevinburke/weirdfortune
```

Note it will look in `/usr/local/games` by default for the fortunes; if they are
located in a different directory you'll have to pass arguments to it.

I can probably distribute binaries if there's enough demand.

## Usage

```sh
$ weirdfortune
```
```
@thrillbo: dont hate me because i'm beautiful, hate me what you can do for your country
```

[fortune]: http://en.wikipedia.org/wiki/Fortune_(Unix)
