package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var profitCalculator = &cobra.Command{
	Use:   "profit-calculator",
	Short: "Profit Calculator",
	Long: `cli to calculate the profit
	Usage: profit-calculator --taxRate 10 --revenue 100 --expenses 50
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// tax, revenue, expenses
		taxRate, _ := cmd.Flags().GetFloat64("taxRate")
		revenue, _ := cmd.Flags().GetFloat64("revenue")
		expenses, _ := cmd.Flags().GetFloat64("expenses")

		EBT := revenue - expenses
		profit := EBT * (1 - taxRate/100)
		fmt.Println("Profit:", profit)
	},
}

var simpleCalculator = &cobra.Command{
	Use:   "simple-calculator",
	Short: "Simple Calculator",
	Long: `cli to calculate the simple calculator
	Usage: simple-calculator add 1 2
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Usage: simple calculator [operation] [num1] [num2]")
			fmt.Println("Operations: add, sub, mul, div, pow")
			return
		}

		operation := args[0]
		num1, err1 := strconv.ParseFloat(args[1], 64)
		num2, err2 := strconv.ParseFloat(args[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid numbers provided.")
			return
		}

		switch operation {
		case "add":
			fmt.Println(num1 + num2)
		case "sub":
			fmt.Println(num1 - num2)
		case "mul":
			fmt.Println(num1 * num2)
		case "div":
			fmt.Println(num1 / num2)
		}
	},
}

var fizzBuzz = &cobra.Command{
	Use:   "fizzbuzz",
	Short: "FizzBuzz",
	Long: `cli to calculate the fizzbuzz
	Usage: fizzbuzz 15
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a number")
			return
		}
		num, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid number.")
			return
		}
		switch {
		case num%3 == 0 && num%5 == 0:
			fmt.Println("FizzBuzz")
		case num%3 == 0:
			fmt.Println("Fizz")
		case num%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(num)
		}
	},
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "Go Session 2",
		Short: "Profit Calculator",
		Long:  "cli to calculate the profit",
	}

	rootCmd.AddCommand(profitCalculator)
	rootCmd.AddCommand(simpleCalculator)
	rootCmd.AddCommand(fizzBuzz)

	profitCalculator.PersistentFlags().Float64P("taxRate", "t", 0, "tax rate")
	profitCalculator.PersistentFlags().Float64P("revenue", "r", 0, "revenue")
	profitCalculator.PersistentFlags().Float64P("expenses", "e", 0, "expenses")

	simpleCalculator.PersistentFlags().Float64P("num1", "1", 0, "number 1")
	simpleCalculator.PersistentFlags().Float64P("num2", "2", 0, "number 2")
	simpleCalculator.PersistentFlags().StringP("operation", "o", "", "operation")

	fizzBuzz.PersistentFlags().IntP("num", "n", 0, "number")

	rootCmd.Execute()
}
