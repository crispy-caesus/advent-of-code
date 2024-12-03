package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parse(filename string) *string {
    contents, _ := os.ReadFile(filename)
    stringConents := string(contents)
    return &stringConents
}

func regex(text *string) *[]string {
    r, _ := regexp.Compile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
    matches := r.FindAllString(*text, -1)

    return &matches
}

func calculate(matches *[]string) *int {
    var result int
    r, _ := regexp.Compile(`[0-9]{1,3}`)

    for i := range *matches {
        numbers := r.FindAllString((*matches)[i], 2)
        num1, _ := strconv.Atoi(numbers[0])
        num2, _ := strconv.Atoi(numbers[1])
        result += num1 * num2
    }
    
    return &result
}

func part1() *int {
    text := parse("input.txt")
    matches := regex(text)
    result := calculate(matches)
    return result
}


func main() {
    resultPart1 := part1()
    fmt.Printf("part: %d\n", *resultPart1)
}


