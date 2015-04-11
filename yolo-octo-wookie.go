package main

/*
char *message = "Yolo Octo Wookie!";
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.GoString(C.message))
}
