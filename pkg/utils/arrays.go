package utils

import (
	"log"
	"strconv"
	"strings"
)

func MapTo2D(vs []string, sep string) [][]string {
	vsm := make([][]string, len(vs))
	for i, v := range vs {
		vsm[i] = strings.Split(v, sep)
	}

	return vsm
}

func MapToInt(vs []string) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		var err error
		vsm[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Error processing map to int")
		}
	}
	return vsm
}

func Filter(vs []string, f func(string) bool) []string {
	vsm := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsm = append(vsm, v)
		}
	}
	return vsm
}

func FilterInt(vi []int, f func(int) bool) []int {
	vim := make([]int, 0)
	for _, v := range vi {
		if f(v) {
			vim = append(vim, v)
		}
	}

	return vim
}

func FindIntIndexes(vi []int, f func(int) bool) []int {
	vim := make([]int, 0)

	for i, v := range vi {
		if f(v) {
			vim = append(vim, i)
		}
	}

	return vim
}

func Every(vi []int, f func(int) bool) bool {
	out := true

	for _, v := range vi {
		if !f(v) {
			return false
		}
	}

	return out
}
