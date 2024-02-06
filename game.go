package main

import (
	"fmt"
	"log"
	"math/rand"
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
func (game *Game) ShowGameResult(language string) {
	fmt.Printf(GetText(language, "result"), game.Correct, game.Total)
}

func InitGame(opts QuizOpts) {
	questions, err := GetAllQuestionsFromFile(opts.Filename)
	if err != nil {
		log.Fatal(err)
	}

	var channelTime <-chan time.Time
	if opts.Timer != 0 {
		channelTime = time.After(time.Second * time.Duration(opts.Timer))
	}
	finalizado := make(chan bool)
	var game *Game = NewGame(len(questions))

	go func(game *Game, fim chan bool) {
		var answer string
		var actualQuestion Question

		if opts.Shuffle {
			for i, question := range questions {
				newPos := rand.Intn(len(questions))
				questions[i] = questions[newPos]
				questions[newPos] = question
			}
		}

		for questionNumber := 0; questionNumber < len(questions); questionNumber++ {
			actualQuestion = questions[questionNumber]
			fmt.Printf(GetText(opts.Language, "show_question"), questionNumber+1, actualQuestion.Statement)
			fmt.Printf(GetText(opts.Language, "input_question"))
			fmt.Scanln(&answer)
			game.UpdateScore(answer == actualQuestion.Answer)
		}
		fim <- true
	}(game, finalizado)

	select {
	case <-channelTime:
		fmt.Println(GetText(opts.Language, "timeout_message"))
	case <-finalizado:
	}
	game.ShowGameResult(opts.Language)
}
