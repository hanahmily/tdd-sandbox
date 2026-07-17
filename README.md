# tdd-sandbox

A minimal Go module used as a throwaway target for exercising vajra's
`milestone-tdd` workflow end to end (assign → human approval → implement →
review → summarize).

It intentionally ships only `go.mod` + this README: the workflow's `assign`
step writes the requirements, the RED unit/e2e tests and the protected-file
list, and `implement` makes them pass. Nothing here is meant to be kept.

Suggested minimal task:

> Add a package `calc` with `Add(a, b int) int`. Follow strict TDD: write a
> failing unit test and a failing end-to-end test first, then implement until
> green.
