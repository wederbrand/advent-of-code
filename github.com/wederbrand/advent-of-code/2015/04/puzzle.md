\--- Day 4: The Ideal Stocking Stuffer ---
----------

Santa needs help [mining](https://en.wikipedia.org/wiki/Bitcoin#Mining) some AdventCoins (very similar to [bitcoins](https://en.wikipedia.org/wiki/Bitcoin)) to use as gifts for all the economically forward-thinking little girls and boys.

To do this, he needs to find [MD5](https://en.wikipedia.org/wiki/MD5) hashes which, in [hexadecimal](https://en.wikipedia.org/wiki/Hexadecimal), start with at least *five zeroes*. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: `1`, `2`, `3`, ...) that produces such a hash.

For example:

* If your secret key is `abcdef`, the answer is `609043`, because the MD5 hash of `abcdef609043` starts with five zeroes (`000001dbbfa...`), and it is the lowest such number to do so.
* If your secret key is `pqrstuv`, the lowest number it combines with to make an MD5 hash starting with five zeroes is `1048970`; that is, the MD5 hash of `pqrstuv1048970` looks like `000006136ef...`.

Your puzzle answer was `117946`.

\--- Part Two ---
----------

Now find one that starts with *six zeroes*.

Your puzzle answer was `3938038`.

Both parts of this puzzle are complete! They provide two gold stars: \*\*

At this point, you should [return to your Advent calendar](/2015) and try another puzzle.

Your puzzle input was `ckczppom`.

You can also [Shareon [Bluesky](https://bsky.app/intent/compose?text=I%27ve+completed+%22The+Ideal+Stocking+Stuffer%22+%2D+Day+4+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F4) [Twitter](https://twitter.com/intent/tweet?text=I%27ve+completed+%22The+Ideal+Stocking+Stuffer%22+%2D+Day+4+%2D+Advent+of+Code+2015&url=https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F4&related=ericwastl&hashtags=AdventOfCode) [Mastodon](javascript:void(0);)] this puzzle.