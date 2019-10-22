package main

import (
    "bufio"
    "fmt"
//    "log"
    "os"
    "strings"
)


func main() {

    if len(os.Args) < 2 {
      fmt.Println("./walkByLine <filename>")
      fmt.Println("Syntax error : go run walkByLine <filename>")
      return
    }

    dir, err := os.Getwd()
    if err != nil {
	fmt.Println(err)
    }
//    fmt.Println("Current DIR",dir)
 
    //file, err := os.Open("TS29518_Namf_Communication.yaml")
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    result := make([]string,100)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "$ref"){
           tmpStr1 := strings.Split(line,"#")
           finalStr := strings.Split(tmpStr1[0],"'")[1]
           if strings.Contains(finalStr, "yaml") {
//             fmt.Println(line)
//             fmt.Println(finalStr)
               resStr := strings.Split(tmpStr1[0],"'")[0] + "'"+ dir +"/"+ finalStr +"#"+ tmpStr1[1]
               fmt.Println(resStr)
//               result = append(result,resStr) 
	   } else {
               fmt.Println(line)
	       result = append(result,line)
           }
        } else {
           fmt.Println(line)
	   result = append(result,line)
        }
    }
   
    
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
