package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)
var correct,game int
var gues string
var err error

func main(){

	number := rangeIn(1000,9999) // generate random secret number between 1000-9999.

	for correct != 4 {	// when all gues numbers are correct, stop

		fmt.Println("Guess the 4 digit number :)")
		fmt.Scan(&gues)	// must enter a 4 digit number

		game,correct,err = feedback(number,gues) // give feedback about your entered guesing number
		if err != nil {
			log.Println(err)	// if your number not 4 digit or int print error
		}else{
			if game != -1 { 	// if number is true
				fmt.Printf("+%d\n",correct)
			}else if correct == 0 {
				fmt.Println(game) // if any number of guessing is not correct
			}else {
				fmt.Printf("%d+%d\n",game,correct) // if number not true but part of true
			}
		}
	}
}

func feedback(n int, gues string) (result,count int,err error){
	n_str := strconv.Itoa(n)
	rune_gues := []rune(gues)
	if len(rune_gues) == 4 && check(rune_gues){ // check gues number is 4 digit and number or not
		for i,value := range n_str{
			if value == rune_gues[i] { 
				count++ 
				err = nil
			}
			if value != rune_gues[i] {
				result = -1
				err = nil
			}
		}
	}else if !check(rune_gues) {
		return result,count,errors.New("It is not a number ! " + fmt.Sprintf("gues = %s\n",gues))
	}else{
		return result,count,errors.New("It is not 4 digit number ! " + fmt.Sprintf("gues = %s\n",gues))
	}
	return result,count,err
}

func rangeIn(low, hi int) int {
    return low + rand.Intn(hi-low)
}

func check(str []rune) bool {
	for _,value := range str{
		_,err := strconv.Atoi(string(value))
		fmt.Println(string(value))
		if  err != nil {
			return false
		}
	}
	return true
}