package main

import "github.com/AkifhanIlgaz/gophercises/task/cmd"

//    ~/.bashrc
//export PATH="~/go/bin:$PATH"

func main() {
	cmd.RootCommand.Execute()
}
