package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
const LOG_FILE = "/Users/denizcamalan/go/src/PropertyFinder-Patika/week-3-homework-2-denizcamalan/exercises/create_rectangle/rect_log.txt"

type Rectangle struct{
	a int
	b int
}

func init(){
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644) // write logs
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
func main(){

	//rec := Rectangle{a: 5, b: 7} // give area and circumfrance
	rec := Rectangle{a: 15, b: -4} // give Negative lenght or height error
	err := rec.Err()
	if err != nil {
		log.Println("Calculation failed: ", err)
	}else{
		fmt.Println("Rectangle's area is =",rec.Area())
		fmt.Println("Rectangle's circumference is =",rec.Circumference())
	}	
}

func (r Rectangle) Area() int{return r.a*r.b} // area func

func (r Rectangle) Circumference() (int){return 2*r.a*r.b} // circumference func

func (r Rectangle) Err() (error){  // error handling func
	
	if r.a < 0 || r.b < 0{
		return errors.New("Negative lenght or height: "+ fmt.Sprintf("lenght = %d, height = %d", r.b,r.a))
	}else{
		return nil
	}
}
