package main

import (
	"fmt"
	"time"
)

func main() {

	cur_tim := time.Now()

	fmt.Println("curt=", cur_tim)

	formated_t := cur_tim.Format("02/01/2006 , 15:04:05 ,Monday")
	formated_t1 := cur_tim.Format("02/01/2006 , 3:05 PM ,Monday")
	fmt.Println("Frmt_t=", formated_t)
	fmt.Println("Frmt_t1=", formated_t1)

	t2 := "08-11-2024"
	layout_str := "02-01-2006"

	ct2, _ := time.Parse(layout_str, t2)
	fmt.Println("ct2=", ct2)
	fmt.Printf("type of ct2 : %T \n \n adding one day incur time \n", ct2)

	newDate := cur_tim.Add(24 * time.Hour)

	formated_t2 := newDate.Format("02/01/2006 , 15:04:05 ,Monday")

	fmt.Println("newDate : ", formated_t2)

}
