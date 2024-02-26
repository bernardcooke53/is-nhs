/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const (
	L = 1000000000
	R = 10000000000
)

func intToDigits(n int) []int {
	ret := []int{}

	for n > 0 {
		ret = append(ret, n%10)
		n /= 10
	}
	slices.Reverse(ret)
	return ret
}

func isValidNHSNumber(n int) bool {
	if n < L || n >= R {
		return false
	}

	digits := intToDigits(n)

	leadingDigits, r := digits[:len(digits)-1], n%10

	leadingWeightedSum := 0

	for i, digit := range leadingDigits {
		leadingWeightedSum += digit * (10 - i)
	}

	checkDigit := 11 - (leadingWeightedSum % 11)
	if checkDigit == 11 {
		return 0 == r
	} else {
		return checkDigit == r
	}
}

var rootCmd = &cobra.Command{
	Short: "Check if a number is a valid NHS number or not",
	Long:  `Check if a number is a valid NHS number; exit code 0 if valid, 1 if not.`,
	Use:   `is-nhs [OPTIONS] [INPUT]`,
	Run:   Run,
	Args:  cobra.RangeArgs(0, 1),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Run(cmd *cobra.Command, args []string) {
	quiet, err := cmd.PersistentFlags().GetBool("quiet")
	cobra.CheckErr(err)

	var input string
	if len(args) == 1 {
		input = args[0]
	}

	if input == "-" || len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		cobra.CheckErr(err)
		if input == "\n" {
			_ = cmd.Usage()
			os.Exit(2)
		}
	}

	input = strings.Trim(input, "\n ")
	inputN, err := strconv.Atoi(input)
	cobra.CheckErr(err)

	if isValidNHSNumber(inputN) {
		if !quiet {
			fmt.Printf("%s is a valid NHS Number\n", input)
		}
	} else {
		if !quiet {
			fmt.Printf("%s is not a valid NHS Number\n", input)
		}
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().
		BoolP("quiet", "q", false, "suppress output")
}
