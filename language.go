package main

var languages map[string]map[string]string = map[string]map[string]string{
	"en": {
		"help":            "quiz {[command=value]}:\nCommands:\n\nfilename=problems.csv\nshuffle=true\nlanguage=en\ntimer=30\n",
		"result":          "You got %d questions right from %d in total!",
		"show_question":   "\nQuestion #%d: %s\n",
		"input_question":  "Your answer: ",
		"timeout_message": "\nTimeout!",
	},
	"pt": {
		"help":            "quiz {[command=value]}:\nCommandos: \n  filename=problems.csv\n  shuffle=true\n  language=en\n  time=30",
		"result":          "Voçẽ acertou %d questões de %d no total!",
		"show_question":   "\nQuestão #%d: %s\n",
		"input_question":  "Sua resposta: ",
		"timeout_message": "\nTempo acabou!",
	},
}

func GetText(language, text string) string {
	value, ok := languages[language]
	if !ok {
		return languages["en"][text]
	}
	return value[text]
}
