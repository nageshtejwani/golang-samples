/*
Race Condition is defined as the condition where the systemâ€™s substantive
behavior is dependent on the sequence or timing of other uncontrollable events.

Below an example of two successive executions of the application:

Test #1								Test #2

[Write function] Item:	1			[Write function] Item:	1
[Read function] Item:	1			[Read function] Item:	1
[Write function] Item:	2			[Read function] Item:	1	<----- Race Condition
[Read function] Item:	2			[Write function] Item:	2
[Write function] Item:	3			[Read function] Item:	2	<----- Race Condition
[Read function] Item:	3			[Write function] Item:	3
[Read function] Item:	3			[Read function] Item:	3
[Write function] Item:	4			[Write function] Item:	4
[Read function] Item:	4			[Read function] Item:	4
[Write function] Item:	5			[Write function] Item:	5
[Write function] Item:	6			[Read function] Item:	5	<----- Race Condition
[Read function] Item:	5			[Write function] Item:	7	<----- Race Condition
[Write function] Item:	7			[Read function] Item:	7
[Read function] Item:	7			[Read function] Item:	7
[Write function] Item:	8			[Write function] Item:	8
[Read function] Item:	8			[Write function] Item:	9	<----- Race Condition

*/

package main

import (
	"fmt"
	"strconv"
	"time"
)

var item uint64

func Write() {

	for {
		item = item + 1
		fmt.Println("[Write function] Item:\t" + strconv.FormatUint(item, 10))
		time.Sleep(1 * time.Second)
	}

}

func Read() {

	for {
		fmt.Println("[Read function] Item:\t" + strconv.FormatUint(item, 10))
		time.Sleep(1 * time.Second)
	}

}

// func main() {

// 	item = 0

// 	go Write()

// 	go Read()

// 	// Wait until some signal is captured.
// 	sigC := make(chan os.Signal, 1)
// 	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
// 	<-sigC

// }
