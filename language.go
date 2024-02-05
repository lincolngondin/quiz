package main

var languages map[string]map[string]string = map[string]map[string]string{
    "en": {
        "help": "quiz {[command=value]}:\nCommands:\n\nfilename=problems.csv\nshuffle=true\nlanguage=en\ntime=30",
    },
    "pt": {
        "help": "quiz {[command=value]}:\nCommandos: \n  filename=problems.csv\n  shuffle=true\n  language=en\n  time=30",
    },
}

func GetText(language, text string) string {
    value, ok := languages[language]
    if !ok {
        return languages["en"][text]
    }
    return value[text]
}
