package main

import (
	"fmt"
	"sort"
)
func CountLetters(input string) map[string]int {
	count := make(map[string]int)
	for _, letter := range input {
		count[string(letter)]++
	}
	return count
}

func IsPerfectNumber(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func FilterAndSortWords(words []string) []string{
	vowels := []string{}
	consonants := []string{}
	for _, word := range words {
		if isVowel(word[0]) {
			vowels = append(vowels, word)
		} else {
			consonants = append(consonants, word)
		}
	}

	sort.Strings(vowels)
	sort.Reverse(sort.StringSlice(consonants))

	return append(vowels, consonants...)
}

func isVowel(letter byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, vowel := range vowels {
		if letter == vowel {
			return true
		}
	}
	return false
}

type SuperHero struct {
	Name        string
	City        string
	Powers      []string
	Personality string
}

func findFavoriteHero(heroes []SuperHero) SuperHero {
	for _, hero := range heroes {
		if hero.City == "Nova York" {
			for _, power := range hero.Powers {
				if power == "Teia de Aranha" {
					return hero
				}
			}
		}
	}
	return SuperHero{}
}

type Character struct {
	Name            string   
	Village         string   
	SpecialAbilities []string 
	Personality     string   
	Dream           string   
}

type Criteria struct {
	Personality string
	Village     string
	Ability     string
}

func findFavoriteCharacter(characters []Character, criteria Criteria) string{
	for _, character := range characters {
		//Verifica qual o favorito
		if character.Village == criteria.Village &&
			character.Personality == criteria.Personality &&
			contains(character.SpecialAbilities, criteria.Ability) {
			return character.Name
		}
	}
	return "Nenhum personagem preenche seus requisitos"
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func main() {

	input := "banana"
	countLetter := CountLetters(input)
	fmt.Println(countLetter) // Output: map[a:3 b:1 n:2]
//---------------------------------------------------------------------------------------------------
	words := []string{"banana", "uva", "laranja", "abacaxi", "manga", "elefante"}
	filterAndSortWord := FilterAndSortWords(words)
	fmt.Println(filterAndSortWord)
//---------------------------------------------------------------------------------------------------
	characters := []Character{
		{
			Name:             "Naruto Uzumaki",
			Village:          "Konoha",
			SpecialAbilities: []string{"Rasengan", "Shadow Clone Jutsu", "Nine-Tails Chakra Mode"},
			Personality:      "Determined",
		},
		{
			Name:             "Sasuke Uchiha",
			Village:          "Konoha",
			SpecialAbilities: []string{"Sharingan", "Chidori", "Amaterasu"},
			Personality:      "Calm",
		},
		{
			Name:             "Kakashi Hatake",
			Village:          "Konoha",
			SpecialAbilities: []string{"Sharingan", "Lightning Blade", "Ninjutsu mastery"},
			Personality:      "Calm",
		},
	}
	criteria := Criteria{
		Personality: "Calm",
		Village:     "Konoha",
		Ability:     "Chidori",
	}

	favorite := findFavoriteCharacter(characters, criteria)
	fmt.Println("Seu personagem favorito é:", favorite)
//---------------------------------------------------------------------------------------------------
	heroes := []SuperHero{
		{"Superman", "Metrópolis", []string{"Super força", "Visão de calor"}, "Justo"},
		{"Spider-Man", "Nova York", []string{"Teia de Aranha", "Sentido Aranha"}, "Engraçado"},
		{"Batman", "Gotham", []string{"Artes marciais", "Tecnologia"}, "Sombrio"},
	}
	favoriteHero := findFavoriteHero(heroes)
	fmt.Println("Super-Herói favorito:", favoriteHero.Name)
//---------------------------------------------------------------------------------------------------
	fmt.Println(IsPerfectNumber(6))  // true
	fmt.Println(IsPerfectNumber(28)) // true
	fmt.Println(IsPerfectNumber(10)) // false
}