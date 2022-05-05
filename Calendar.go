package main

import "fmt"
import "time"

func main() {
	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 Mon"), "         ", dt.Format("15:04:05"))
	arr := [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"};
	for i := 0; i < len(arr); i++ {
		if (int(dt.Weekday()) == i) {
			fmt.Print("\x1b[1;31m", arr[i], "\x1b[0m", "  ");
		} else {
			fmt.Print(arr[i], "  ");
		}
	}
	fmt.Println();
}
