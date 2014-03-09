package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func getB(inputVectors [][]float64) int {
	n := len(inputVectors) / 100
	b := int(math.Ceil(math.Log2(float64(n))))
	return b
}

func getOneRandomVector(size int) []float64 {
	var result []float64
	for j := 0; j < size; j++ {
		newFloat := rand.Float64()
		if rand.Int()%2 == 0 {
			newFloat = newFloat * -1.0
		}
		result = append(result, newFloat)
	}
	return result
}

func getRandomVectors(b int, size int) [][]float64 {
	var randomVectors [][]float64
	for i := 0; i < b; i++ {
		newVector := getOneRandomVector(size)
		randomVectors = append(randomVectors, newVector)
	}
	return randomVectors
}

func sum(list []float64) float64 {
	result := 0.0
	for _, value := range list {
		result += value
	}
	return result
}

func naturalDotProduct(first []float64, second []float64) int {
	var products []float64
	for i := 0; i < len(first); i++ {
		products = append(products, first[i]*second[i])
	}
	s := sum(products)
	if s >= 0 {
		return 1
	}
	return 0
}

func binToInt(input []int) int {
	sum := 0
	for i, num := range input {
		sum += num * int(math.Pow(float64(2), float64(i)))
	}
	return sum
}

func getKeyFromVector(inputVector []float64, randomVectors [][]float64) int {
	var newDot []int
	for _, r := range randomVectors {
		newDot = append(newDot, naturalDotProduct(inputVector, r))
	}
	return binToInt(newDot)
}

func getDotProducts(inputVectors [][]float64, randomVectors [][]float64) []int {
	var result []int
	for _, vector := range inputVectors {
		newKey := getKeyFromVector(vector, randomVectors)
		result = append(result, newKey)
	}
	return result
}

func createMap(keys []int, identifiers []string) map[int][]string {
	m := make(map[int][]string)
	for i, key := range keys {
		val, _ := m[key]
		m[key] = append(val, identifiers[i])
	}
	for _, value := range m {
		sort.Strings(value)
	}
	return m
}

func stringToFloat(strs []string) []float64 {
	var result []float64
	for _, str := range strs {
		i, _ := strconv.Atoi(strings.Trim(str, " "))
		result = append(result, float64(i))
	}
	return result
}

func getData(filename string) [][]float64 {
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	var vectors [][]float64
	for _, line := range lines {
		vectors = append(vectors, stringToFloat(strings.Split(line, ",")))
	}
	return vectors
}

func valuesToValueMap(values []string) map[string]int {
	newMap := make(map[string]int)
	for _, key := range values {
		val, _ := newMap[key]
		newMap[key] = val + 1
	}
	return newMap
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	rawData := getData("testingData.txt")
	var inputVectors [][]float64
	var names []string
	for _, v := range rawData {
		inputVectors = append(inputVectors, v[:len(v)-1])
		names = append(names, strconv.FormatFloat(v[len(v)-1], 'f', -1, 64))
	}
	b := getB(inputVectors)
	size := len(inputVectors[0])
	randomVectors := getRandomVectors(b, size)
	dotProducts := getDotProducts(inputVectors, randomVectors)
	m := createMap(dotProducts, names)
	testVectors := getData("pokerData.txt")
	testVector := testVectors[0]
	testKey := getKeyFromVector(testVector[:len(testVector)-1], randomVectors)
	fmt.Println(b)
	fmt.Println(size)
	fmt.Println(len(m))
	fmt.Println(len(testVectors))
	result := m[testKey]
	fmt.Println(valuesToValueMap(result))
	for key, _ := range m {
		fmt.Println(valuesToValueMap(m[key]))
	}
}
