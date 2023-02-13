package main

import (

    "fmt"
    "strconv"
)

// when you are declaring variable at the package level, you can not use the colon equals syntax you actually have to use the full declaration syntax


// THAT'S HOW YOU CONCATANATE SAME TYPES OF VARIABLE AT ONE PLACE.
// var (
// actorName string = "Elisabeth Sladen"
//  companion string = "Sarah Jane Smith"
//  doctorName int = 3
//  season int = 11
// )

func main(){
    // var i int /*The way you declare variable in go its pretty much the same the way you speak*/
    

    // var i int = 42
    // var j int = 27
    // k := 99
    
    // fmt.Println(i)
    // fmt.Printf("%v , %T" , j , j)
    // fmt.Printf("%v , %T" , k , k)

    // Note:- Variable with the innermost scope actually takes precedence. This is called shadowing.

    // Note:- if we have a local variable that's declared and not used , then that's actually a compile time error.

    // Note:- If we have declared the variable at the package level and is lowercase as scoped to the package then any file in the same package can access that variable. If it is uppercase at the package level, then it's export at the front of package and it is globally visible

    // length of the varible name reflect the life of the variable.

    // best practice in go is to keep these acronyms as all uppercase and the reason is readability
    // for ex:- var theHTTP string = "https://google.com"


    // var i int = 42
    // fmt.Printf("%v , %T\n", i , i)

    // var j float32
    // j = float32(i)
    // fmt.Printf("%v , %T\n" , j , j)

    // while changing the type of varible we need to do the explicit conversion and direct change will show compilation error as go doesn't want to lose the information while changing the types of variables.

//   now for string conv process we use the "strconv"
var i int = 42
fmt.Printf("%v , %T\n",i,i)

var j string
j = strconv.Itoa(i)
fmt.Printf("%v , %T\n" , j , j)
}  