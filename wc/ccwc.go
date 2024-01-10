package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func iswspace(c rune)bool{

	return (c==' ' || c=='\f'|| c=='\n' || c=='\r' || c=='\t' || c=='\v')
}
// -w
func wordCount(str string) []int {
	cnt:= []int{}	
	// currCnt:=0;
	totalCnt:=0;
	isWord:=false
	for _,char := range str{
		if iswspace(char){
			isWord=true;
			// fmt.Printf("char:%c\n",char)
		}else if(isWord){
			isWord=false;
			totalCnt++;
		}
	}
	cnt = append(cnt, totalCnt+1)
	return cnt
}

// -l
func lineCount(str string) []int {
	cnt:= []int{}	
	lines := strings.Split(str,"\n")
	cnt = append(cnt, len(lines)-1)
	return cnt
}

// -c
func byteCount(str string) []int{
	cnt:= []int{}	
	cnt = append(cnt, len(str))
	return cnt
}

// default
func allCount(s string) []int{
	cnt := []int{0,0,0}

	return cnt
}

func makeResult(cnt []int, fileName string)string{


	temp:="    "
	
	for i:=0 ;i<len(cnt);i++{ 	
		numString := strconv.Itoa(cnt[i])
		temp+=numString+" "
	}
	temp+=fileName
	return temp
}
func main(){

	var fileName string
	var option string
	args := os.Args[1:]
	
	if(len(args)>1){
		option = args[0]
		fileName = args[1]
	}else{
		option = ""
		fileName = args[0];
	}

	fileInput,err := os.ReadFile(fileName);

	if err != nil {
		fmt.Errorf("File is not valid",err)
	}
	var result string
	var cnt []int
	switch option{

		case "":
			cnt=allCount(string(fileInput))
		case "-w":
			cnt=wordCount(string(fileInput))
		case "-c":
			cnt=byteCount(string(fileInput))
		case "-l":
			cnt=lineCount(string(fileInput))
	}

	result = makeResult(cnt,fileName)

	fmt.Println(result)

}