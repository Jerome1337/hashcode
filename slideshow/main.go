package main

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
)

type SlidesParams struct {
	PhotosAmount int
	Photos []Photo
	VPhotos []VPhoto
}

type Photo struct {
	ID []int
	TagsAmount int
	Tags []string
}

type VPhoto struct {
	ID int
	TagsAmount int
	Tags []string
}

type SlideShow struct {
	SlidesAmount int
	SlidesPhotoIDs [][]int
}

func main() {
	filepath.Walk("input", func(file string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			params := readEntryFile(file)
			slideShow := params.solve()
			writeSubmissionFile(f.Name(), slideShow)
		}

		return nil
	})
}

func (s SlidesParams) solve() SlideShow {
	var slideShow SlideShow
	// var testedPhotos []int

	if len(s.VPhotos) != 0 {
		for _, cp := range s.VPhotos {
			previousInterest := 0
			// testedPhotos = append(testedPhotos, cp.ID)

			for _, p := range s.VPhotos {
				// if sliceExists(testedPhotos, p.ID) {
				// 	continue
				// }

				// testedPhotos = append(testedPhotos, p.ID)
				commonTags := 0
				interest := 0

				for _, t := range p.Tags {
					if sliceExists(cp.Tags, t) {
						commonTags++
					}

					if commonTags == 0 {
						continue
					}

					interest += interestFactor(commonTags, cp.TagsAmount, p.TagsAmount)
				}

				if interest > 0 && previousInterest < interest {
					previousInterest = interest
					tags := removeDuplicatedVal(append(cp.Tags, p.Tags...))

					s.Photos = append(s.Photos, Photo{
						ID: []int{cp.ID, p.ID},
						TagsAmount: len(tags),
						Tags: tags,
					})

					break
				}
			}
		}
	}

	log.Println("VERTICAL SOLVED")


	for _, cp := range s.Photos {
		previousInterest := 0
		// testedPhotos = []int{}
		// testedPhotos = append(testedPhotos, cp.ID...)

		for _, p := range s.Photos {
			// if sliceExists(testedPhotos, p.ID) {
			// 	continue
			// }

			// testedPhotos = append(testedPhotos, p.ID...)
			commonTags := 0
			interest := 0

			for _, t := range p.Tags {
				if sliceExists(cp.Tags, t) {
					commonTags++
				}

				if commonTags == 0 {
					continue
				}

				interest += interestFactor(commonTags, cp.TagsAmount, p.TagsAmount)
			}

			if interest > 0 && previousInterest < interest {
				previousInterest = interest
				slideShow = SlideShow{
					SlidesPhotoIDs: append(slideShow.SlidesPhotoIDs, cp.ID),
				}

				slideShow.SlidesAmount = len(slideShow.SlidesPhotoIDs)

				break
			}
		}
	}

	log.Println("ALL SOLVED")

	return slideShow
}

func interestFactor(factors ...int) int {
	min := factors[0]

	for _, f := range factors {
		if f < min {
			min = f
		}
	}

	return min
}

func removeDuplicatedVal(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
			if _, value := keys[entry]; !value {
					keys[entry] = true
					list = append(list, entry)
			}
	}

	return list
}


func sliceExists(slice interface{}, items ...interface{}) bool {
	s := reflect.ValueOf(slice)

	for i := 0; i < s.Len(); i++ {
		for _, item := range items {
			if s.Index(i).Interface() == item {
				return true
			}
		}
	}

	return false
}
