package main

import(
	"fmt"
	"flag"
	"strconv"
	"strings"
	"reflect"
)

func checkCommonAfterSumInArray(sum float64, arr []float64) []float64{
	m := make(map[float64]bool)
	commonarrayArray := make([]float64, 0)
	for i := 0; i < len(arr); i++ {
			m[arr[i]] = true
	}

	for _, item := range arr{
		if item == sum {
			commonarrayArray = append(commonarrayArray, item)
		}
	}
	return commonarrayArray
}

// get this func from internet
func inArray(val interface{}, array interface{}) (exists bool, index int) {
    exists = false
    index = -1

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }

    return
}

func checkCommonInArrays(first , second []float32) []float64{
	commonarrayArray := make([]float64, 0)
	for _, it := range first{
		if exist, index := inArray(it, second); exist{
			commonarrayArray = append(commonarrayArray, first(index))
		}
	}
	return commonarrayArray

}

func sumArray(items []float32) float64{
	var sumItems float64

	for _, item := range first{
		sumItems+=item
	}
	return sumItems
}

func changeItemType(arr []string) []float64{
	resArray := make([]float64, 0)
	for _, item :=range arr{
		if it, err:= strconv.ParseFloat(item, 64); err == nil{
			resArray = append(resArray, it)
		}
	}
	return resArray
}

//splitTokenString split string to token pares
func splitIncomeString(inputString string) []string {
	s := strings.Split(inputString, ",")
	return s
}

func main()  {
	first := flag.String("first", "", "first array ")
	second := flag.String("second", "", "seccond array")
	flag.Parse()
	prepFirst := splitIncomeString(*first)
	prepSecond := splitIncomeString(*second)
	first = changeItemType(prepFirst)
	second = changeItemType(prepSecond)
	resArray := checkCommonInArrays(first, seccond) 
	sumFirst := sumArray(first)
	sumSecond:= sumArray(first)
	checkCommonFirst := checkCommonAfterSumInArray(sumFirst, second)
	checkCommonSecond := checkCommonAfterSumInArray(sumSecond, first)
	resArray = append(resArray, checkCommonFirst...)
	resArray = append(resArray, checkCommonSecond...)
	fmt.Println("Common Points:", resArray)
}