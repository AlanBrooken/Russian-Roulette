package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	min, max := 1, 7
	rand.Seed(time.Now().UnixNano())
	deathshot := rand.Intn(max-min) + min
	deathshot2 := rand.Intn(max-min) + min
	deathshot3 := rand.Intn(max-min) + min
	deathshot4 := rand.Intn(max-min) + min
	deathshot5 := rand.Intn(max-min) + min
	//Каждый deathshot - это доп. пуля в барабане, увеличивающая шанс смерти

	fmt.Println("Choose how many bullets will be in the cylinder") //Выбор сложности игры

	for { //цикл для возможности повторного ввода данных для игрока
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
		}

		input = strings.TrimSuffix(input, "\n")

		bullets, err := strconv.Atoi(input) //конвертация первого ввода данных игрока из строки в число
		if err != nil {
			fmt.Println("Please choose a number from 1 to 5")
			continue
		}

		if bullets < 1 || bullets > 5 {
			fmt.Println("Please choose a number from 1 to 5")
			continue
		} else if bullets == 1 {
			fmt.Println("You've chosen", bullets, "bullet. Type ROLL to start")
		} else if bullets >= 2 || bullets <= 5 {
			fmt.Println("You've chosen", bullets, "bullets. Type ROLL to start")
		} //Условия игры выбраны, теперь будет генерация пуль и шансов на выживание

		if bullets == 2 {
			for {
				if deathshot2 == deathshot {
					deathshot2 = rand.Intn(max-min) + min //Пересоздаю второй deathshot, если он совпадает с первым
					//Такая стратегия повторяется со всеми возможными выборами условий игры
					//Думаю, ччто есть более простой код, но я его не придумал
				}
				break
			}
			start, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
			}

			start = strings.TrimSuffix(start, "\n") //Считываение второго ввода данных, в данном случае -
			//это команда для старта игры

			if start == "ROLL" || start == "roll" {
				rolling()
				roll := rand.Intn(max-min) + min             //Создается случайное число
				if roll == deathshot || roll == deathshot2 { //Если оно совпадает с deathshot, то игрок погибает
					fmt.Println("Bang! You're dead")
					break
				}
				fmt.Println("You're lucky. Do you want to play one more time?")
				break
			}
			fmt.Println("Incorrect command. Type ROLL to start the game")
			continue
		} else if bullets == 3 {
			for {
				if deathshot3 == deathshot || deathshot3 == deathshot2 {
					deathshot3 = rand.Intn(max-min) + min
				}
				break
			}
			start, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
			}

			start = strings.TrimSuffix(start, "\n")

			if start == "ROLL" || start == "roll" {
				rolling()
				roll := rand.Intn(max-min) + min

				if roll == deathshot || roll == deathshot2 || roll == deathshot3 {
					fmt.Println("Bang! You're dead")
					break
				}
				fmt.Println("You're lucky. Do you want to play one more time?")
				break
			}
			fmt.Println("Incorrect command. Type ROLL to start the game")
			continue
		} else if bullets == 4 {
			for {
				if deathshot4 == deathshot || deathshot4 == deathshot2 || deathshot4 == deathshot3 {
					deathshot4 = rand.Intn(max-min) + min
				}
				break
			}
			start, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
			}

			start = strings.TrimSuffix(start, "\n")

			if start == "ROLL" || start == "roll" {
				rolling()
				roll := rand.Intn(max-min) + min

				if roll == deathshot || roll == deathshot2 || roll == deathshot3 || roll == deathshot4 {
					fmt.Println("Bang! You're dead")
					break
				}
				fmt.Println("You're lucky. Do you want to play one more time?")
				break
			}
			fmt.Println("Incorrect command. Type ROLL to start the game")
			continue
		} else if bullets == 5 {
			for {
				if deathshot5 == deathshot || deathshot5 == deathshot2 || deathshot5 == deathshot3 || deathshot5 == deathshot4 {
					deathshot5 = rand.Intn(max-min) + min
				}
				break
			}
			start, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
			}

			start = strings.TrimSuffix(start, "\n")

			if start == "ROLL" || start == "roll" {
				rolling()
				roll := rand.Intn(max-min) + min

				if roll == deathshot || roll == deathshot2 || roll == deathshot3 || roll == deathshot4 || roll == deathshot5 {
					fmt.Println("Bang! You're dead")
					break
				}
				fmt.Println("You're lucky. Do you want to play one more time?")
				break
			}
			fmt.Println("Incorrect command. Type ROLL to start the game")
			continue
		} else if bullets == 1 {
			start, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
			}

			start = strings.TrimSuffix(start, "\n")

			if start == "ROLL" || start == "roll" {
				rolling()
				roll := rand.Intn(max-min) + min

				if roll == deathshot {
					fmt.Println("Bang! You're dead")
					break
				}
				fmt.Println("You're lucky. Do you want to play one more time?")
				break
			}
			fmt.Println("Incorrect command. Type ROLL to start the game")
			continue
		}

	}
}

func rolling() {
	fmt.Println("rolling")
	fmt.Println("3")
	fmt.Println("2")
	fmt.Println("1")

}
