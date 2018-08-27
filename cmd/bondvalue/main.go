package main

import (
	"fmt"

	"github.com/jrmycanady/gogovtbond"
	"github.com/spf13/pflag"
)

func main() {

	var dataFilePath = pflag.StringP("data-file", "d", "./data.txt", "The path to the treasure data file.")
	var series = pflag.StringP("series", "s", "", "The bond series. (I, E, EE, S")
	var rdYear = pflag.IntP("redemtion-year", "y", -1, "The year the bond will be redeemed. (YYYY)")
	var rdMonth = pflag.IntP("redemtion-month", "m", -1, "The month the bond will be redemed. (MM)")
	var iYear = pflag.IntP("issue-year", "Y", -1, "The year the bond was issued. (YYYY)")
	var iMonth = pflag.IntP("issue-month", "M", -1, "The month the bond was issued. (MM)")
	var value = pflag.IntP("value", "v", -1, "The face value of the bond.")

	pflag.Parse()

	b := gogovtbond.BondData{}
	if err := b.LoadFromFile(*dataFilePath); err != nil {
		fmt.Println("Failed to read data file")
	}

	// Convert from EE to data file N type.
	if *series == "EE" {
		*series = "N"
	}
	if *series == "" || *rdYear == -1 || *rdMonth == -1 || *iYear == -1 || *iMonth == -1 || *value == -1 {
		fmt.Println("Missing parameters")
	}

	fmt.Println(b.BondValue(*series, *rdYear, *rdMonth, *iYear, *iMonth, *value))

}
