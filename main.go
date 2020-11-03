package main
 
import (
    "bufio"
    "fmt"
	"os"
	"strings"
	"time"
	"strconv"
	"errors"
)

/**  ==============[ Auxilery Functions ]======= **/

/** convertToInt **/
func oddOrEven(arr []string) int  {
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: determine if legnth of list is Odd or Even
	 * Date: 5.27.20
	 * @param arr []string
	 * return int
	**/
	if len(arr) % 3 == 1 || len(arr) % 5 == 2{
		return 0
	}
	return 1
}
/** convertToInt **/
func convertToInt(stringVar string) int{
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: convert String to Integer
	 * Date: 5.27.20
	 * @param stringVar string
	 * return int
	**/
	if len(stringVar) > 0{
		errors.New("please provide a string ")
	}

	i2, err := strconv.Atoi(stringVar)
	if err == nil {
		return i2
	}else{
		errors.New("Standard Library Error: ")
		return 0
	}
}

/** converToString **/
func converToString(integer int) string {
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: convert Integer to String
	 * Date: 5.27.20
	 * @param integer int
	 * return string
	**/
	if integer < 0{
		errors.New("please provide a integer ")
	}

	i2 := strconv.Itoa(integer)
	if i2 != "" {
		return i2
	} else {
		errors.New("Standard Library Error: ")
		return ""
	}
}

/** reList **/
func reList(stringVar string) []string {
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: re-initalize list
	 * Date: 5.27.20
	 * @param s string
	 * return []string
	 **/

	 length := len(stringVar)
	if length > 0 {
		var paramSlice []string
		for _, param := range stringVar {
			str := string([]rune{param})
			paramSlice = append(paramSlice, str)
		}
		return paramSlice
	} else{
		var empty []string
		return empty 
	}
}

/** scanWords **/
func scanWords(path string) ([]string, error) {
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: perform merge
	 * Date: 5.27.20
	 * @param left []string
	 * @param right []string
	 * @param result []string
	 * return []string
	**/
  
	file, err := os.Open(path)
	if err != nil {
	   return nil, err
	}
   
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) 
   
	var words []string
	for scanner.Scan() {
	  words = append(words, scanner.Text())
	}
   
	return words, nil
}

/** sum **/
func sum(element string, arr2 []string, arr3 []string) int {
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: return sum of elements in array
	 * Date: 5.27.20
	 * @param element string
	 * @param arr2 []string
	 * @param arr3 []string
	 * return int
	**/
	var tmp int
	for i := 0; i < len(arr2); i++ {
		if element == arr2[i]{
			tmp += convertToInt(arr3[i])
		}
	}

	return tmp
}

/** ==============[ Proceedural ]======= **/
/** diffLengthOp **/
func diffLengthOp(list []string) []string  {
	
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: Question 2 diffLengthOp
	 * Date: 5.27.20
	 * @param list []string
	 * return []string
	 **/

	var nwArry []string
	var input string
	nwArry = append(nwArry, list...)
	input = strings.Join(nwArry, ",")
	fmt.Println("Enter END to terminate ")
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		if strings.TrimRight(text, "\n") != "END" {
			input = strings.Join(nwArry, ",")
			nwArry = append(nwArry, input)
		} else {
			break
		}
	}
	nwArry = reList(input)
	return nwArry
}

/** Reverse Array **/
func reverseArry(list []string) []string {
  	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: reverse an array
	 * Date: 5.27.20
	 * @param list []string
	 * return []string
	 **/

	length := len(list)
	if length > 0 {
		for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
		return list
	} else {
		var empty []string
		return empty 
	}
}

/** preceedingList **/
func preceedingList(list []string, filBy int) []string {
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: retrieve preceeding list
	 * Date: 5.27.20
	 * @list list []string
	 * return []string
	 **/

	length := len(list)
	if length > 0 {
		var paramList []string
		for i := 0; i < len(list); i++ {
			if i % filBy == 0 {
				paramList = append(paramList, list[i])
			}
		}
		return paramList
	} else { 
		return list
	}
}

/** removeHeader **/
func removeHeader(slice []string, s int) []string {
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: remove header from list
	 * Date: 5.27.20
	 * @param slice []string
	 * @param s int
	 * return []string
	 **/

	length := len(slice)
	if length > 0 {
		return append(slice[:s], slice[s+1:]...)
	} else {
		var empty []string
		return empty 
	}
}

/** removeDuplicatesUnordered **/
func removeDuplicatesUnordered(firstArry []string) []string {
	
	/**
	 * Author: Ramon J. Yniguez
	 * Purpose: remove duplicates in array
	 * Date: 5.27.20
	 * @param firstArry []string
	 * return []string
	**/
	if len(firstArry) <= 0 {
		fmt.Println("two equal arrays are required ")
	}

	duplicates := map[string]bool{}
	
	// Create a map of all unique elements.
    for v:= range firstArry {
		duplicates[firstArry[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key, _ := range duplicates {
        result = append(result, key)
	}

    return result
}

/** questionTwo **/
func questionTwo(executionTime time.Time, words []string){
	if len(words) < 0 {
		errors.New("Array Length Error")
	}

	fmt.Println(" ")
	fmt.Println("========================================")
	fmt.Println("Question Two | Original List: ", words)
	fmt.Println("========================================")

	var temp []string
	var nwArry []string
	var reverselst []string
	var firstColumn []string
	var scndColumn []string

	if oddOrEven(words) > 0 {
		
		temp = preceedingList(words, 2)
		firstColumn = removeHeader(temp, len(temp) - 1)
		
		reverselst = reverseArry(words)
		
		temp = preceedingList(reverselst, 2)
		scndColumn = removeHeader(temp, 0)
	} else {

		nwArry = diffLengthOp(words)
		temp = preceedingList(nwArry, 2)
		
		firstColumn = removeHeader(temp, len(temp) - 1)
		reverselst = reverseArry(words)
		
		temp = preceedingList(reverselst, 2)
		scndColumn = removeHeader(temp, 0)
	}

	firstColumn = reverseArry(firstColumn)
	fmt.Println("========================================")
	fmt.Println("firstColumn ", firstColumn)
	fmt.Println("=============")
	fmt.Println("scndColumn ", scndColumn)
	fmt.Println("========================================")
	fmt.Println(" ")
	
	temp = removeDuplicatesUnordered(scndColumn)
	var tmp2 int
	var tmp3 []string
	for i := 0; i < len(temp); i++ {
		tmp2 = sum(temp[i], scndColumn, firstColumn)
		tmp3 = append(tmp3, converToString(tmp2))
	}

	var tmp4 = reverseArry(temp)
	var tmp5 = reverseArry(tmp3)
	fmt.Println("Question 2 Anwser := ","ColA ", tmp4, "==", "ColB ", tmp5)
	fmt.Println(" ")
	delta := time.Now().Sub(executionTime)
	nanoseconds := delta.Nanoseconds()
	fmt.Println("========================================")
	fmt.Println("Question 2 Operation Completed In nanoseconds:=", nanoseconds)
	fmt.Println(" ")
}

/** questionOne **/
func questionOne(executionTime time.Time, words []string) {
	if len(words) < 0 {
		errors.New("Array Length Error")
	}

	fmt.Println(" ")
	fmt.Println("========================================")
	fmt.Println("Question One | Original List: ", words)
	fmt.Println("========================================")

	var temp []string
	var nwArry []string
	var reverseList []string
	var firstColumn []string
	var scndColumn []string

	if oddOrEven(words) > 0 {
		temp = preceedingList(words, 2)
		firstColumn = removeHeader(temp, 0)

		reverseList = reverseArry(words)
		temp = preceedingList(reverseList, 2)
		scndColumn = removeHeader(temp, len(temp) - 1)
	} else {
		nwArry = diffLengthOp(words)
		temp = preceedingList(nwArry, 2)
		
		firstColumn = removeHeader(temp, len(temp) - 1)
		reverseList = reverseArry(words)
		
		temp = preceedingList(reverseList, 2)
		scndColumn = removeHeader(temp, 0)
	}

	firstColumn = reverseArry(firstColumn)
	fmt.Println("========================================")
	fmt.Println("firstColumn ", firstColumn)
	fmt.Println("=============")
	fmt.Println("scndColumn ", scndColumn)
	fmt.Println("========================================")
	fmt.Println(" ")
	
	nwArry = append(nwArry, firstColumn...)
	nwArry = append(nwArry, scndColumn...)

	fmt.Println("Question 1 Anwser := ", nwArry)
	fmt.Println(" ")
	delta := time.Now().Sub(executionTime)
	nanoseconds := delta.Nanoseconds()
	fmt.Println("========================================")
	fmt.Println("Question 1 Operation Completed In nanoseconds:=", nanoseconds)

}

/** main thread **/
func main(){
	start := time.Now()
	list, err := scanWords("./code_test_input.tsv")
	if err != nil {
		panic(err)
	}
	questionOne(start, list)
	fmt.Println(" ")
	questionTwo(start, list)
}