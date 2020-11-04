package main
import "fmt"

func main(){
	maximumGroups:= 12
	maximumObjects:= 12
	for numberOfObjects:=1; numberOfObjects <= maximumObjects; numberOfObjects++{
		for numberOfGroups := 1; numberOfGroups <= maximumGroups; numberOfGroups++{
			fmt.Println(numberOfObjects * numberOfGroups)
		}
	}
}