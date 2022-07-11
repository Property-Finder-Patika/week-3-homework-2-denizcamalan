package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
const LOG_FILE = "/Users/denizcamalan/go/src/PropertyFinder-Patika/week-3-homework-2-denizcamalan/example/create_rectangle/rect_log.txt"

type Rectangle struct{
	a int
	b int
}

func main(){

	//rec := Rectangle{a: 5, b: 7} // give area and circumfrance
	rec := Rectangle{a: 5, b: -7} // give Negative lenght or height error
	err := Err(rec)
	if err != nil {
		log.Println("Calculation failed: ", err)
	}else{
		fmt.Println("Rectangle's area is =",Area(rec.a,rec.b))
		fmt.Println("Rectangle's circumference is =",Circumference(rec.a,rec.b))
	}	
}

func Area(a,b int) int{return a*b} // area func

func Circumference(a,b int) (int){return 2*a*b} // circumference func

func Err(r Rectangle) (error){  // error handling func

	if r.a < 0 || r.b < 0{
		return errors.New("Negative lenght or height: "+ fmt.Sprintf("lenght = %d, height = %d", r.b,r.a))
	}else{
		return nil
	}
}

func init() { 
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644) // write logs
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}