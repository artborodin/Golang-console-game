
package main

// Игра Угадай слово

import "fmt"

func main() {
	words := []string{"тест", "привет", "мир"}

	
	for _, word := range words {

		letters := []rune(word)      // буквы слова
		guess_map := map[rune]bool{} // карта открытых букв

		fmt.Println("Новое слово, количество букв: ", len(letters))
		for {
			var input string

			fmt.Print("Введите 1 букву ")
			_, err := fmt.Scanf("%s", &input)

			if err != nil {
				fmt.Println("error input", err)
				continue
			}
			if input == "END" {
				fmt.Println()
				break
			}

			// д.б. введен один символ
			if len([]rune(input)) != 1 {
				//
				continue
			}
			// введенная буква
			guess_letter := ([]rune(input))[0]

			for _, ltr := range letters {
				if guess_letter == ltr {
					guess_map[ltr] = true
					//fmt.Println("Ok!")
				}
			}
			// формирование строки с открытыми буквами
			var progress []rune
			for _, ltr := range letters {
				if guess_map[ltr] {
					progress = append(progress, ltr)
				} else {
					progress = append(progress, '*')
				}
			}
			fmt.Println(string(progress))

			if string(progress) == word {
				fmt.Println("слово отгадано!")
				break
			}
		}
	}
}
