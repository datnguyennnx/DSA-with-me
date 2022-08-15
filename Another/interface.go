package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ eat, move, speak string }
type Bird struct{ eat, move, speak string }
type Snake struct{ eat, move, speak string }

func (c Cow) Eat() string   { return c.eat }
func (c Cow) Move() string  { return c.move }
func (c Cow) Speak() string { return c.speak }

func (c Bird) Eat() string   { return c.eat }
func (c Bird) Move() string  { return c.move }
func (c Bird) Speak() string { return c.speak }

func (c Snake) Eat() string   { return c.eat }
func (c Snake) Move() string  { return c.move }
func (c Snake) Speak() string { return c.speak }
func main() {

	cow := Cow{eat: "grass", move: "walk", speak: "moo"}
	bird := Bird{eat: "worms", move: "fly", speak: "peep"}
	snake := Snake{eat: "mice", move: "slither", speak: "hsss"}

	var nameType map[string]string

	nameType = make(map[string]string)

	for {
		fmt.Printf("\n>please enter the request -> name & infomation: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		request := scanner.Text()

		//fmt.Printf("request is: %s\n", request)

		if len(strings.Split(request, " ")) != 3 {
			fmt.Println("Command input number is not equal to 3")
			continue
		}

		command := strings.Split(request, " ")[0]
		fmt.Printf("command: %s\n", command)

		switch command {
		case "newanimal":

			name := strings.Split(request, " ")[1]
			animalType := strings.Split(request, " ")[2]

			fmt.Printf("name: %s\n", name)
			fmt.Printf("animalType: %s\n", animalType)

			//adding name checking
			_, ok := nameType[name]
			if ok == false {
				nameType[name] = animalType
				fmt.Println("Created it!")
			} else {
				fmt.Printf("Animal name %s is taken, please try another name\n", name)
			}
		case "query":

			name := strings.Split(request, " ")[1]
			info := strings.Split(request, " ")[2]

			fmt.Printf("name: %s\n", name)
			fmt.Printf("info: %s\n", info)

			switch nameType[name] {
			case "cow":
				switch info {
				case "eat":
					cow.Eat()
				case "move":
					cow.Move()
				case "speak":
					cow.Speak()
				default:
					fmt.Printf("%s is not in (eat, move, speak), please try again\n", info)
					continue
				}
			case "bird":
				switch info {
				case "eat":
					bird.Eat()
				case "move":
					bird.Move()
				case "speak":
					bird.Speak()
				default:
					fmt.Printf("%s is not in (eat, move, speak), please try again\n", info)
					continue
				}
			case "snake":
				switch info {
				case "eat":
					snake.Eat()
				case "move":
					snake.Move()
				case "speak":
					snake.Speak()
				default:
					fmt.Printf("%s is not in (eat, move, speak), please try again\n", info)
					continue
				}
			default:
				fmt.Printf("%s is not in (cow, bird, snake), please try again\n", nameType[name])
				continue
			}
		default:
			fmt.Printf("%s is not in (newanimal, query), please try again\n", command)
			continue
		}
	}

}
