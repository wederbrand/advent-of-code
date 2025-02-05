\--- Day 7: Some Assembly Required ---
----------

This year, Santa brought little Bobby Tables a set of wires and [bitwise logic gates](https://en.wikipedia.org/wiki/Bitwise_operation)! Unfortunately, little Bobby is a little under the recommended age range, and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a [16-bit](https://en.wikipedia.org/wiki/16-bit) signal (a number from `0` to `65535`). A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations. A gate provides no signal until all of its inputs have a signal.

The included instructions booklet describes how to connect the parts together: `x AND y -> z` means to connect wires `x` and `y` to an AND gate, and then connect its output to wire `z`.

For example:

* `123 -> x` means that the signal `123` is provided to wire `x`.
* `x AND y -> z` means that the [bitwise AND](https://en.wikipedia.org/wiki/Bitwise_operation#AND) of wire `x` and wire `y` is provided to wire `z`.
* `p LSHIFT 2 -> q` means that the value from wire `p` is [left-shifted](https://en.wikipedia.org/wiki/Logical_shift) by `2` and then provided to wire `q`.
* `NOT e -> f` means that the [bitwise complement](https://en.wikipedia.org/wiki/Bitwise_operation#NOT) of the value from wire `e` is provided to wire `f`.

Other possible gates include `OR` ([bitwise OR](https://en.wikipedia.org/wiki/Bitwise_operation#OR)) and `RSHIFT` ([right-shift](https://en.wikipedia.org/wiki/Logical_shift)). If, for some reason, you'd like to *emulate* the circuit instead, almost all programming languages (for example, [C](https://en.wikipedia.org/wiki/Bitwise_operations_in_C), [JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_Operators), or [Python](https://wiki.python.org/moin/BitwiseOperators)) provide operators for these gates.

For example, here is a simple circuit:

```
123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i

```

After it is run, these are the signals on the wires:

```
d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456

```

In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to *wire `a`*?

Your puzzle answer was `46065`.

\--- Part Two ---
----------

Now, take the signal you got on wire `a`, override wire `b` to that signal, and reset the other wires (including wire `a`). What new signal is ultimately provided to wire `a`?

Your puzzle answer was `14134`.

Both parts of this puzzle are complete! They provide two gold stars: \*\*

At this point, you should [return to your Advent calendar](/2015) and try another puzzle.

If you still want to see it, you can [get your puzzle input](7/input).

You can also [Shareon [Bluesky](https://bsky.app/intent/compose?text=I%27ve+completed+%22Some+Assembly+Required%22+%2D+Day+7+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F7) [Twitter](https://twitter.com/intent/tweet?text=I%27ve+completed+%22Some+Assembly+Required%22+%2D+Day+7+%2D+Advent+of+Code+2015&url=https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F7&related=ericwastl&hashtags=AdventOfCode) [Mastodon](javascript:void(0);)] this puzzle.