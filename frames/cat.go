package main

import (
	"fmt"
	"time"
)

func main() {
	frames := []string{
		`
         /\_/\  
        ( o.o ) 
         > ^ <   
		`,
		`
         /\_/\  
        ( -.- ) 
         > ^ <   
		`,
		`
         /\_/\  
        ( o.o ) 
         > ^ <  
        `,
		`
         /\_/\  
        ( o.o ) 
         > ^ <  
        `,
	}

	for {
		for _, frame := range frames {
			fmt.Print("\033[H\033[2J") // Clear screen
			fmt.Println(frame)
			time.Sleep(500 * time.Millisecond) // Delay for frame change
		}
	}
}
