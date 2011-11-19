// gospark
// https://github.com/twm/gospark
// © 2011 Thomas W. Most, MIT license
//
// Go implementation of spark, for when you want to replace a 1.6 KiB shell
// script with a 1.1 MiB statically-linked executable.  This implementation
// accepts all of the inputs the original does, but may produce slightly
// different output because it uses floating-point math.
//
// This is probably not good Go code.  I did this to learn more of the
// language.
package main

import (
	"io/ioutil"
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var usage = `
  USAGE:
    spark [-h] VALUE,...

  EXAMPLES:
    spark 1 5 22 13 53
    ▁▁▃▂█
    spark 0,30,55,80,33,150
    ▁▂▃▄▂█
    echo 9 13 5 17 1 | spark
    ▄▆▂█▁
`
var help = flag.Bool("h", false, "request help")
var ticks = [...]string{"▁","▂","▃","▄","▅","▆","▇","█"}

// Split s by commas and whitespace, outputting any valid numbers found
func appendNums(nums []float64, s string) []float64 {
	for _, word := range strings.Fields(s) {
		for _, part := range strings.Split(word, ",") {
			num, err := strconv.Atof64(part)
			if err == nil {
				nums = append(nums, float64(num))
			}
		}
	}
	return nums
}

func main() {
	flag.Parse()
	if *help {
		fmt.Print(usage)
		os.Exit(2)
	}

	nums := make([]float64, 0, flag.NArg())

	if flag.NArg() == 0 {
		if buf, err := ioutil.ReadAll(os.Stdin); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading: %v", err)
			os.Exit(1)
		} else {
			nums = appendNums(nums, string(buf))
		}
	} else {
		for _, arg := range flag.Args() {
			nums = appendNums(nums, arg)
		}
	}

	if len(nums) == 0 {
		return
	}

	min := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	spread := max - min
	if spread < .00000001 {
		spread = 1
	}
	scale := float64(len(ticks) - 1) / spread

	for _, v := range nums {
		index := (v - min) * scale
		fmt.Print(ticks[int(index)])
	}
	fmt.Println()
}
