package main

import "fmt"

func main() {

	length := 10.0 //Length of the rectangle
	width := 5.0   //Width of the rectangle

	area := length * width

	fmt.Printf("The area of the rectangle is: %.2f square units.\n", area)

}

/*
1) mkdir test01
2) cd test01
// create files and run check it works
3) git init
4) git add 
5) git commit -m "First comment"
6) git status
7) Create a github.com repository- same name as folder
8) git remote add origin https://github.com/dayanajoseph/test01.git
9) git remote set-url origin https://dayanajoseph:$GITHUB_TOKEN@github.com/dayanajoseph/test01.git
--
*/
