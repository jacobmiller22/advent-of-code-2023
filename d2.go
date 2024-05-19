package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isGameValid(game string, total_blue int, total_red int, total_green int) bool {
	setstrs := strings.Split(game, "; ")

	for _, set := range setstrs {
		log.Printf("%s", set)
		cubestrs := strings.Split(set, ", ")
		for _, cube := range cubestrs {
			subcubestrs := strings.Split(cube, " ")
			if subcubestrs[1] == "blue" {
				cube_blue, err := strconv.Atoi(subcubestrs[0])
				if err != nil {
					log.Fatalf("Error while parsing blue subcubestr: %v", err)
				}
				if cube_blue > total_blue {
					return false
				}

			} else if subcubestrs[1] == "red" {
				cube_red, err := strconv.Atoi(subcubestrs[0])
				if err != nil {
					log.Fatalf("Error while parsing red subcubestr: %v", err)
				}
				if cube_red > total_red {
					return false
				}

			} else if subcubestrs[1] == "green" {
				cube_green, err := strconv.Atoi(subcubestrs[0])
				if err != nil {
					log.Fatalf("Error while parsing green subcubestr: %v", err)
				}
				if cube_green > total_green {
					return false
				}

			}

		}
	}
	return true
}

func getPower(game string) int {
	setstrs := strings.Split(game, "; ")

	max_cube_blue := 0
	max_cube_red := 0
	max_cube_green := 0

	for _, set := range setstrs {
		log.Printf("%s", set)
		cubestrs := strings.Split(set, ", ")
		for _, cube := range cubestrs {
			subcubestrs := strings.Split(cube, " ")
			if subcubestrs[1] == "blue" {
				cube_blue, err := strconv.Atoi(subcubestrs[0])
				if err != nil {
					log.Fatalf("Error while parsing blue subcubestr: %v", err)
				}

				max_cube_blue = max(max_cube_blue, cube_blue)
			} else if subcubestrs[1] == "red" {
				cube_red, err := strconv.Atoi(subcubestrs[0])
				if err != nil {
					log.Fatalf("Error while parsing red subcubestr: %v", err)
				}

				max_cube_red = max(max_cube_red, cube_red)
			} else if subcubestrs[1] == "green" {
				cube_green, err := strconv.Atoi(subcubestrs[0])
				if err != nil {
					log.Fatalf("Error while parsing green subcubestr: %v", err)
				}

				max_cube_green = max(max_cube_green, cube_green)
			}

		}
	}
	return max_cube_blue * max_cube_red * max_cube_green
}

func main() {
	fp, err := os.Open("input-d2.txt")
	if err != nil {
		log.Fatal("Failed to read input file")
	}

	scanner := bufio.NewScanner(fp)

	if err != nil {
		log.Fatal("Failed to read input file")
	}

	// total_green := 13
	// total_red := 12
	// total_blue := 14

	sum_gameids := 0
	sum_powers := 0

	for scanner.Scan() {
		line := scanner.Text()

		substrs := strings.Split(line, ": ")

		subgamestr := strings.Split(substrs[0], " ")

		game_id, err := strconv.Atoi(subgamestr[1])
		if err != nil {
			log.Fatal("Error while parsing game id")
		}

		// if !isGameValid(substrs[1], total_blue, total_red, total_green) {
		// 	continue
		// }

		sum_gameids += game_id
		log.Print(getPower(substrs[1]))
		sum_powers += getPower(substrs[1])
	}
	log.Printf("sum game ids: %d\n", sum_gameids)
	log.Printf("sum powers: %d\n", sum_powers)

}
