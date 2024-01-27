package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	ArrayList    map[string][]int
	ExpectedList map[string][]int
}

const (
	maxInt32 = int32(^uint32(0) >> 1)
	minInt32 = -maxInt32 - 1
)

var testData = &TestData{
	ArrayList: map[string][]int{
		"empty":                            {},
		"oneElement":                       {0},
		"twoElementsSorted":                {5, 144},
		"twoElementsUnsorted":              {10, -7},
		"moreThanOneElement":               {1, 3, 4, 2},
		"moreThanOneElementWithRepetition": {1, 4, 4, 2},
		"moreThanOneElement2":              {7, 7, 1, 0, 99, -5, 10},
		"sameElement":                      {1, 1, 1, 1},
		"negativeNumbers":                  {-5, -2, -10, -1, -3},
		"descendingOrder":                  {5, 4, 3, 2, 1},
		"randomOrder":                      {9, 2, 7, 4, 1, 6, 3, 8, 5},
		"duplicateElements":                {2, 2, 1, 1, 3, 3, 4, 4},
		"largeArray":                       {-1, -10000, -12345, -2032, -23, 0, 0, 0, 0, 10, 10000, 1024, 1024354, 155, 174, 1955, 2, 255, 3, 322, 4741, 96524},
		"singleNegativeElement":            {-7},
		"arrayWithZeroes":                  {0, -2, 0, 3, 0},
		"ascendingOrder":                   {1, 2, 3, 4, 5},
		"descendingOrderWithDuplicates":    {5, 5, 4, 3, 3, 2, 1},
		"boundaryValues":                   {2147483648, -2147483647},
		"mixedSignNumbers":                 {-1, 0, 1, -2, 2},
	},
	ExpectedList: map[string][]int{
		"empty":                            {},
		"oneElement":                       {0},
		"twoElementsSorted":                {5, 144},
		"twoElementsUnsorted":              {-7, 10},
		"moreThanOneElement":               {1, 2, 3, 4},
		"moreThanOneElementWithRepetition": {1, 2, 4, 4},
		"moreThanOneElement2":              {-5, 0, 1, 7, 7, 10, 99},
		"sameElement":                      {1, 1, 1, 1},
		"negativeNumbers":                  {-10, -5, -3, -2, -1},
		"descendingOrder":                  {1, 2, 3, 4, 5},
		"randomOrder":                      {1, 2, 3, 4, 5, 6, 7, 8, 9},
		"duplicateElements":                {1, 1, 2, 2, 3, 3, 4, 4},
		"largeArray":                       {-1, -10000, -12345, -2032, -23, 0, 0, 0, 0, 10, 10000, 1024, 1024354, 155, 174, 1955, 2, 255, 3, 322, 4741, 96524},
		"singleNegativeElement":            {-7},
		"arrayWithZeroes":                  {-2, 0, 0, 0, 3},
		"ascendingOrder":                   {1, 2, 3, 4, 5},
		"descendingOrderWithDuplicates":    {1, 2, 3, 3, 4, 5, 5},
		"boundaryValues":                   {-2147483647, 2147483648},
		"mixedSignNumbers":                 {-2, -1, 0, 1, 2},
	},
}

func TestBubble(t *testing.T) {

	for testCase, array := range testData.ArrayList {
		t.Run(testCase, func(t *testing.T) {
			actual := Bubble(array)
			assert.ElementsMatch(t, actual, testData.ExpectedList[testCase])
		})

	}
}
func TestSelection(t *testing.T) {

	for testCase, array := range testData.ArrayList {
		t.Run(testCase, func(t *testing.T) {
			actual := Selection(array)
			assert.ElementsMatch(t, actual, testData.ExpectedList[testCase])
		})

	}
}

func TestInsertion(t *testing.T) {

	for testCase, array := range testData.ArrayList {
		t.Run(testCase, func(t *testing.T) {
			actual := Insertion(array)
			assert.ElementsMatch(t, actual, testData.ExpectedList[testCase])
		})

	}

}
func TestQuickLomuto(t *testing.T) {

	for testCase, array := range testData.ArrayList {
		t.Run(testCase, func(t *testing.T) {
			actual := QuickLomuto(array, 0, len(array)-1)
			assert.ElementsMatch(t, actual, testData.ExpectedList[testCase])
		})

	}

}
func TestQuickSimple(t *testing.T) {

	for testCase, array := range testData.ArrayList {
		t.Run(testCase, func(t *testing.T) {
			actual := QuickSimple(array)
			assert.ElementsMatch(t, actual, testData.ExpectedList[testCase])
		})

	}

}
