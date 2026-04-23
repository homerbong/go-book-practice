package display

import "fmt"

func SectionTitle(sectionTitle string) {
	fmt.Println()
	fmt.Println("#################################################")
	fmt.Println("#", sectionTitle)
	fmt.Println("#################################################")
	fmt.Println()
}
