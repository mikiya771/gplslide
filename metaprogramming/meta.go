package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func calcPrimeNumbers(max int) []int {
	furui := make([]int, max+1)
	for i := 2; i <= max; i++ {
		furui[i] = i
	}
	furui = furui[2:]
	primes := []int{}
	count := 0
	for len(furui) != 0 {
		count++
		if count%100000 == 0 {
			fmt.Println(len(furui))
		}
		primes = append(primes, furui[0])
		furui = StiffOff(furui[0], furui)
	}
	return primes
}

func StiffOff(start int, furui []int) []int {
	next := []int{}
	for _, num := range furui {
		if num%start != 0 {
			next = append(next, num)
		}
	}
	return next
}

func main() {
	const fileNmae = "const.go"
	const genDir = "main"
	const packageName = "main"
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error on getting directory: error: %v\n", err)
		return
	}
	file := filepath.Join(cwd, genDir, fileNmae)
	_ = os.Remove(file)
	fp, err := os.Create(file)
	if err != nil {
		fmt.Printf("Error on making file: error: %v\n", err)
		return
	}
	defer fp.Close()
	const mMax = 1000000
	primes := calcPrimeNumbers(mMax)
	declarePackage(packageName, fp)
	formatNewLine(fp)
	declareIntSliceVariable("primes", primes, fp)

}
func declarePackage(packageName string, w io.Writer) {
	fmt.Fprintf(w, "package %s\n", packageName)
}
func declareIntSliceVariable(variableName string, array []int, w io.Writer) {
	strArr := []string{}
	for _, num := range array {
		strArr = append(strArr, strconv.Itoa(num))
	}
	str := strings.Join(strArr, ", ")
	fmt.Fprintf(w, "var %s = []int{%s}", variableName, str)
}
func formatNewLine(w io.Writer) {
	fmt.Fprint(w, "\n")
}
