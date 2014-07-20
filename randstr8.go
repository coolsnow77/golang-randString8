package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(100))
	//fmt.Println(randStr())
	fmt.Println(rstr8(randStr()))
}

func randStr()(s string) {
	slice :=make([]rune,0)

	for x:='A'; x<='Z'; x++ {
		slice = append(slice,x)
	}

	for x:='a'; x<='z'; x++ {
		slice =append(slice,x)
	}

	for x:='0'; x<='9' ;x++ {
		slice= append(slice,x)
	}

	s = string(slice)
	return
}

func rstr8(str string) (s string) {
	leng := 8
	stmp :=[]rune(str)
	for x:=0 ; x<len(stmp) ; x++ {
		stmp[x] = stmp[rand.Intn(62)]
	}

	for x:=0;x<leng; x++ {
		s+= string(stmp[x])
	}
	return
}

