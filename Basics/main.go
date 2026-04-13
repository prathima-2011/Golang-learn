package main
import ("fmt"; "path")

func main() {
	dir, file := path.Split("Basics/main.go");
	fmt.Println("Directory:", dir);
	fmt.Println("File:", file);
}