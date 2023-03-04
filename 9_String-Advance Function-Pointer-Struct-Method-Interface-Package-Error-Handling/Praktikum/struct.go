package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Score []int
}

func (s Student) AverageScore() float64 {
	sum := 0
	for _, score := range s.Score {
		sum += score
	}
	return float64(sum) / float64(len(s.Score))
}

func (s Student) MinScore() int {
	min := s.Score[0]
	for _, score := range s.Score {
		if score < min {
			min = score
		}
	}
	return min
}

func (s Student) MaxScore() int {
	max := s.Score[0]
	for _, score := range s.Score {
		if score > max {
			max = score
		}
	}
	return max
}

func main() {
	students := []Student{
		{Name: "Rizky", Score: []int{80}},
		{Name: "Kobar", Score: []int{75}},
		{Name: "Ismail", Score: []int{70}},
		{Name: "Umam", Score: []int{75}},
		{Name: "Yopan", Score: []int{60}},
	}

	var totalScore int
	for _, student := range students {
		totalScore += student.Score[0]
	}
	averageScore := float64(totalScore) / float64(len(students))

	fmt.Printf("Average score: %.2f\n", averageScore)

	sort.Slice(students, func(i, j int) bool {
		return students[i].AverageScore() < students[j].AverageScore()
	})

	fmt.Printf("Minimum score: %d (achieved by %s)\n", students[0].MinScore(), students[0].Name)
	fmt.Printf("Maximum score: %d (achieved by %s)\n", students[len(students)-1].MaxScore(), students[len(students)-1].Name)
}
