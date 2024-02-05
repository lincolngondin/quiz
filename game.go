package main

import (
	"fmt"
	"log"
	"time"
)

type Game struct {
	Correct int
	Wrong   int
	Total   int
}

func NewGame(total int) *Game {
	return &Game{
		Correct: 0,
		Wrong:   0,
		Total:   total,
	}
}

func (game *Game) UpdateScore(right bool) {
	if right {
		game.Correct++
	} else {
		game.Wrong++
	}
}
func (game *Game) ShowGameResult(){
	fmt.Printf("You got %d questions right from %d in total!", game.Correct, game.Total)
}

func InitGame(filename string, timer int, language string, shuffle bool) {
	questions, err := GetAllQuestionsFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

    channelTime := time.After(time.Second*time.Duration(timer))
    finalizado := make(chan bool)
    var game *Game = NewGame(len(questions))

    go func(game *Game,fim chan bool) {
        var answer string
        var actualQuestion Question
        for questionNumber := 0; questionNumber < len(questions); questionNumber++{
            actualQuestion = questions[questionNumber]
            fmt.Printf("\nQuestion #%d: %s\n", questionNumber+1, actualQuestion.Statement)
            fmt.Printf("Your answer: ")
            fmt.Scanln(&answer)
            game.UpdateScore(answer == actualQuestion.Answer)
        }
        fim<-true
    }(game, finalizado)

    select {
        case <-channelTime:
            fmt.Println("\nTimeout!")
        case <-finalizado:
            fmt.Println("Terminou normal!")
    }
    game.ShowGameResult()
}

