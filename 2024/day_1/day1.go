package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseLists() (*[]int, *[]int) {
	filename := "listsInput.txt"
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	var historianList, officeList []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splitStrings := strings.Split(line, "   ")

		historianId, err := strconv.Atoi(splitStrings[0])
		if err != nil {
			log.Fatal(err)
		}

		officeId, err := strconv.Atoi(splitStrings[1])
		if err != nil {
			log.Fatal(err)
		}

		historianList = append(historianList, historianId)
		officeList = append(officeList, officeId)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &historianList, &officeList
}

func selectionSort(unsortedList *[]int) *[]int {
	sortedList := make([]int, len(*unsortedList))
	copy(sortedList, *unsortedList)

	for i := 0; i < len(sortedList); i++ {
		minIndex := i
		for j := i + 1; j < len(sortedList); j++ {
			if sortedList[j] < sortedList[minIndex] {
				sortedList[j], sortedList[i] = sortedList[i], sortedList[j]
			}
		}
	}
	return &sortedList
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func compare(historianList *[]int, officeList *[]int) (difference *int) {
	var diffSum int
	for i := 0; i < len(*historianList); i++ {
		diff := abs((*officeList)[i] - (*historianList)[i])
		//fmt.Printf("%d %d difference: %d\n", (*historianList)[i], (*officeList)[i], diff)
		diffSum += diff
	}
	return &diffSum
}

func calculateSimilarityScore(historianList *[]int, officeList *[]int) *int {
	var similarityScore int
	for len(*historianList) != 0 {
		i := 0
		k := 0
		//================================== i ======================================//
		for ; i < len(*historianList) && (*historianList)[i] == (*historianList)[0]; i++ {
		}
        i--

		//==================================j =====================================//

		for j := 0; j < len(*officeList); j++ {
			if (*officeList)[j] == (*historianList)[i] {
				if j != 0 {
					*officeList = (*officeList)[j:]
				}

				//=================================== k ================================//
				for ; k < len(*officeList) && (*officeList)[k] == (*historianList)[i]; k++ {
				}
			}
		}

		//fmt.Printf("%d: %dx left %dx right\n", (*historianList)[i], i+1, k)
		similarityScore += (i + 1) * ((*historianList)[i] * (k))
		*historianList = (*historianList)[i+1:]
	}

	return &similarityScore
}

func main() {
	historianList, officeList := parseLists()

	sortedHistorianList := selectionSort(historianList)
	sortedOfficeList := selectionSort(officeList)

	difference := compare(sortedHistorianList, sortedOfficeList)
	fmt.Println(*difference) // result part 1

	similarityScore := calculateSimilarityScore(sortedHistorianList, sortedOfficeList)
	fmt.Println(*similarityScore)

}
