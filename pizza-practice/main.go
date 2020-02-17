package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	MaxSlice int
	PizzaTypeNum int
	PizzaSliceNum []int
)

func main() {
	filepath.Walk("input", func(file string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			MaxSlice = 0
			PizzaTypeNum = 0
			PizzaSliceNum = []int{}
			readEntryFile(file)
			writeSubmissionFile(f.Name(), solveProblem())
		}

		return nil
	})
}

func readEntryFile(file string) {
	openedFile, _ := os.Open(file)
	scanner := bufio.NewScanner(openedFile)
	isFirstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		lineContent := strings.Split(line, " ")

		if isFirstLine {
			for i, v := range lineContent {
				intv := convertStringToInt(v)

				if i == 0 {
					MaxSlice = intv
				}

				PizzaTypeNum = intv
			}

			isFirstLine = !isFirstLine
			continue
		}

		for _, v := range lineContent {
			PizzaSliceNum = append(PizzaSliceNum, convertStringToInt(v))
		}
	}
}

func solveProblem() []int {
	var (
		slicesSum = 0
		pizzaType = make([]int, 0)
	)

	for i := len(PizzaSliceNum)-1; i >= 0; i-- {
		tmp := slicesSum + PizzaSliceNum[i]

		if tmp >= MaxSlice {
			continue
		}

		slicesSum = tmp
		pizzaType = append(pizzaType, i)
	}

	return pizzaType
}

func writeSubmissionFile(outputFile string, pizzaType []int) {
	f, err := os.Create(fmt.Sprintf("submission/%s", strings.Replace(outputFile, ".in", ".sub", 1)))

	if err != nil {
		log.Fatal("Cannot create file", err)
	}

	vt := []string{}
	for i := len(pizzaType)-1; i >= 0; i-- {
		vt = append(vt, strconv.Itoa(pizzaType[i]))
	}

	lines := []string{strconv.Itoa(len(pizzaType)), strings.Join(vt, " ")}

	for _, line := range lines {
		fmt.Fprintln(f, line)

		if err != nil {
			log.Fatal("Error writting string", err)
		}
	}


	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("File written successfully")
}

func convertStringToInt(r string) int {
	i, err := strconv.Atoi(r)

	if err != nil {
		log.Fatal("Error converting rune", err)
	}

	return i
}
