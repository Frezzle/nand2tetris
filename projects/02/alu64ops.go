package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// TODO fix this later, for some reason some ops show correct output according to the nand2tetris book,
// but some others don't, e.g. "x - 1" and "y" doesn't seem to work.

func main() {
	if len(os.Args) < 3 {
		log.Fatal("x and y arg required.")
	}

	x, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal("x must be an int.")
	}

	y, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal("y must be an int.")
	}

	performAllALUOperations(x, y)
}

func performAllALUOperations(x, y int64) {
	// print header
	fmt.Println("zx nx zy ny f  no | out\t\top")

	// loop through all combinations of micro-operations
	for op := byte(0); op < 64; op++ {
		zx, nx, zy, ny, f, no := getMicroOperations(op)

		// pre-process inputs
		if zx != 0 {
			x = 0
		}
		if nx != 0 {
			x = ^x
		}
		if zy != 0 {
			y = 0
		}
		if ny != 0 {
			y = ^y
		}

		// calculate output
		var output int64
		if f != 0 {
			output = x + y
		} else {
			output = x & y
		}

		// post-process output
		if no != 0 {
			output = ^output
		}

		// print inputs and output
		opDescription := microOpsCombos[microOps{zx, nx, zy, ny, f, no}]
		fmt.Printf("%d  %d  %d  %d  %d  %d  | %d\t\t%s\n", zx, nx, zy, ny, f, no, output, opDescription)
	}
}

func getMicroOperations(opCode byte) (zx, nx, zy, ny, f, no byte) {
	no = opCode & 0b000001
	f = opCode & 0b000010 >> 1
	ny = opCode & 0b000100 >> 2
	zy = opCode & 0b001000 >> 3
	nx = opCode & 0b010000 >> 4
	zx = opCode & 0b100000 >> 5
	return
}

type microOps struct {
	zx byte
	nx byte
	zy byte
	ny byte
	f  byte
	no byte
}

var microOpsCombos = map[microOps]string{
	// the ones in the nand2tetris book
	{1, 0, 1, 0, 1, 0}: "always 0",
	{1, 1, 1, 1, 1, 1}: "always 1",
	{1, 1, 1, 0, 1, 0}: "always -1",
	{0, 0, 1, 1, 0, 0}: "x",
	{1, 1, 0, 0, 0, 0}: "y",
	{0, 0, 1, 1, 0, 1}: "!x",
	{1, 1, 0, 0, 0, 1}: "!y",
	{0, 0, 1, 1, 1, 1}: "-x",
	{1, 1, 0, 0, 1, 1}: "-y",
	{0, 1, 1, 1, 1, 1}: "x + 1",
	{1, 1, 0, 1, 1, 1}: "y + 1",
	{0, 0, 1, 1, 1, 0}: "x - 1",
	{1, 1, 0, 0, 1, 0}: "y - 1",
	{0, 0, 0, 0, 1, 0}: "x + y",
	{0, 1, 0, 0, 1, 1}: "x - y",
	{0, 0, 0, 1, 1, 1}: "y - x",
	{0, 0, 0, 0, 0, 0}: "x & y",
	{0, 1, 0, 1, 0, 1}: "x | y",
	// other interesting ones
}
