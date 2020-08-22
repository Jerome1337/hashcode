package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readEntryFile(file string) SlidesParams {
	openedFile, _ := os.Open(file)
	scanner := bufio.NewScanner(openedFile)
	i := 0
	isFirstLine := true
	photosAmount := 0
	photos := []Photo{}
	vPhotos := []VPhoto{}

	for scanner.Scan() {
		line := scanner.Text()

		if isFirstLine {
				photosAmount = convertStringToInt(line)
				isFirstLine = !isFirstLine
				continue
		}

		lineContent := strings.Split(line, " ")

		if lineContent[0] == "H" {
			photos = append(photos, Photo{
				ID: []int{i},
				TagsAmount: convertStringToInt(lineContent[1]),
				Tags: lineContent[2:],
			})
		} else {
			vPhotos = append(vPhotos, VPhoto{
				ID: i,
				TagsAmount: convertStringToInt(lineContent[1]),
				Tags: lineContent[2:],
			})
		}

		i++
	}

	return SlidesParams{
		PhotosAmount: photosAmount,
		Photos: photos,
		VPhotos: vPhotos,
	}
}

func writeSubmissionFile(outputFile string, slideShow SlideShow) {
	f, err := os.Create(fmt.Sprintf("submission/%s", strings.Replace(outputFile, ".txt", ".sub", 1)))

	if err != nil {
		log.Fatal("Cannot create file", err)
	}

	lines := []string{strconv.Itoa(slideShow.SlidesAmount)}
	for _, ids := range slideShow.SlidesPhotoIDs {
		stringID := []string{}
		for _, id := range ids {
			stringID = append(stringID, strconv.Itoa(id))
		}

		lines = append(lines, strings.Join(stringID, " "))
	}

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
		log.Fatal("Error converting string ", err)
	}

	return i
}
