package main

import (
"fmt"
    "io/ioutil"
	"log"
   "strings"
   "os"
   "strconv"
)

func main() {

	
	fmt.Println("Enter the input list filename (with extension,comma delimited)")
	var fname string
	fmt.Scanln(&fname)
	fmt.Println("Enter the starting number to append")
	var snum string
	fmt.Scanln(&snum)
	fmt.Println("Enter the ending number to append")
	var endnum string
	fmt.Scanln(&endnum)

fmt.Println("Creating the wordlist.txt")
createList(snum,endnum,fname)
}

func createList(param1 string,param2 string ,filename string)  {

	startval,_:= strconv.Atoi(param1)
	lastval,_:=strconv.Atoi(param2)
	lastval ++
   
	f, err := os.Create("wordlist.txt")
   
	if err != nil {
	   log.Fatal(err)
   }
   
   defer f.Close()
   
	   content, err := ioutil.ReadFile(filename)
   
		if err != nil {
			 log.Fatal(err)
		}
   
		cnt := string(content)
		arr1 := strings.Split(cnt,",");
   
		for _, baseword := range arr1 {
		   //fmt.Println(baseword)
		   for i:=startval ; i < lastval ; i++ {
			   //fmt.Println(i)
			   f.WriteString(baseword+strconv.Itoa(i)+"\n")
			   }
	   }
   fmt.Println("wordlist.txt created successfully")
}