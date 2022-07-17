package bridge

import "fmt"

func main() {
	darkTheme := new(DarkTheme)
	about := About{theme: darkTheme}
	careers := Careers{theme: darkTheme}
	fmt.Println(about.getContent())
	fmt.Println(careers.getContent())
}
