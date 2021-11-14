package main

import (
	"fmt"
)

func printLines(ls []line) {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	// Gathering widths
	var wLabel, wOp, wArgs int
	for _, line := range ls {
		wLabel = max(wLabel, len(line.label))
		wOp = max(wOp, len(line.op))
		wArgs = max(wArgs, len(line.args))
	}

	for _, l := range ls {
		switch {
		case "" == l.op && "" == l.comment:
			fmt.Println("")
		case "" == l.op:
			fmt.Println(l.comment)
		case "" == l.comment:
			fmt.Printf("%-*s  %-*s  %-*s  %s\n", wLabel, l.label, wOp, l.op, wArgs, l.args, l.comment)
		default:
			fmt.Printf("%-*s  %-*s  %s\n", wLabel, l.label, wOp, l.op, l.args)
		}
	}
}
