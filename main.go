package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
// RUN IT <icon src="AllIcons.Actions.Execute"/>. RUN IT <icon src="AllIcons.Actions.Execute"/>. RUN IT <icon src="AllIcons.Actions.Execute"/>. RUN IT <icon src="AllIcons.Actions.Execute"/>. RUN IT <icon src="AllIcons.Actions.Execute"/>. RUN IT <icon src="AllIcons.Actions.Execute"/>. RUN IT <icon src="AllIcons.Actions.Execute"/>. RUN IT <icon src="AllIcons.Actions.Execute"/>. RUN IT. RUN IT. RUN IT. RUN IT. RUN IT. RUN IT. RUN IT. RUN IT.
//

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 100; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		//fmt.Println("i =", 100/i)
		love(i)
	}
}

func some() {

}

func some1() {

}

func some2() {

}

func some3() {

}

func some4() {

}

func some5() {

}

func love(value_of_love int) {
	fmt.Println("I'm love you soooooo much. and I love sex with you. A LOT OF SEEEEX. И жопку твою люблю.")
	fmt.Printf("Love loading... %d! \n", value_of_love)
	if value_of_love == 100 {
		fmt.Println("Love loading complete!")
		fmt.Println("Время спонтанного секса!")
	}
}
