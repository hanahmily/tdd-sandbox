# Package `difference` — interface contract

Owned by the tech lead. The coder implements against this contract; do not
change these signatures without a justified revision request.

## Public API

```go
package difference

// Sub returns the arithmetic difference of a and b, i.e. a - b.
// It is a pure function: no state, no error, defined for all int inputs.
func Sub(a, b int) int
```

### Semantics

- `Sub(a, b)` MUST equal `a - b` for every pair of `int` operands.
- Must be correct across the three operand classes the milestone requires:
  positive, negative, and zero.

## CLI entry point (for end-to-end use)

`cmd/difference` is a thin command wrapper used by the e2e test to exercise
`Sub` through a real built binary:

- Usage: `difference <a> <b>` where `<a>` and `<b>` are base-10 integers.
- On success: prints `Sub(a, b)` followed by a newline to stdout, exit 0.
- On wrong argument count or unparseable operands: message to stderr, exit 2.

The command wrapper is production source and is intentionally NOT protected —
the coder may shape its internals as long as the observable contract above
holds.
