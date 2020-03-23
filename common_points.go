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

func checkCommonInArrays(firstArr []float64, secondArr []float64)[]float64{
	commonarrayArray := make([]float64, 0)
	for _, it := range firstArr{
		if exist, index := inArray(it, secondArr); exist{
			commonarrayArray = append(commonarrayArray, firstArr[index])
		}
	}
	return commonarrayArray

}

func sumArray(items []float64) float64{
	var sumItems float64

	for _, item := range items{
		sumItems+=float64(item)
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
func splitIncomeString(inputString string)[]string{
	s := strings.Split(inputString, ",")
	fmt.Println(s)
	fmt.Printf("Type: %T\n", s)
	return s
}

func main()  {
	first := flag.String("f", "", "first array ")
	second := flag.String("s", "", "second array")
	flag.Parse()
	fmt.Println(*first)
	fmt.Printf("Type of first: %T \n", *first)
	fmt.Println(*second)
	_f := splitIncomeString(*first)
	fmt.Println(_f)
	_s := splitIncomeString(*second)
	fmt.Println(_s)
	// prepFirst := splitIncomeString("asd,qwe")
	// prepSecond := splitIncomeString(*second)
	// fmt.Printf("%T", *prepFirst)
	// fmt.Printf("%T", *prepSecond)
	_first := changeItemType(_f)
	fmt.Println(_first)
	fmt.Printf("Typr %T \n", _first)
	_second := changeItemType(_s)
	fmt.Println(_second)
	fmt.Printf("Typr %T \n", _second)
	resArray := checkCommonInArrays(_first, _second) 
	sumFirst := sumArray(_first)
	sumSecond:= sumArray(_second)
	checkCommonFirst := checkCommonAfterSumInArray(sumFirst, _second)
	checkCommonSecond := checkCommonAfterSumInArray(sumSecond, _first)
	fmt.Println("Typr: ", checkCommonFirst)
	resArray = append(resArray, checkCommonFirst...)
	resArray = append(resArray, checkCommonSecond...)
	fmt.Printf("Common Points: %v\n", resArray)
}