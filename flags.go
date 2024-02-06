package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type QuizOpts struct {
	Filename string
	Language string
	Shuffle  bool
	Timer    uint
}

func NewDefaultQuizOpts() QuizOpts {
	return QuizOpts{
		Filename: "problems.csv",
		Language: "en",
		Shuffle:  false,
		Timer:    30,
	}
}

func ProcessArguments(args []string) (QuizOpts, bool, error) {
	opts := NewDefaultQuizOpts()
	reg := regexp.MustCompile("^(filename|shuffle|language|timer)=(.*)$")
	if len(os.Args) >= 2 {
		if os.Args[1] == "help" {
			fmt.Println(GetText(opts.Language, "help"))
			return opts, true, nil
		}
		for _, option := range os.Args[1:] {
			find := reg.FindStringSubmatch(option)
			if reg.MatchString(option) {
				switch find[1] {
				case "filename":
					opts.Filename = find[2]
				case "shuffle":
					if find[2] == "true" {
						opts.Shuffle = true
					} else {
						opts.Shuffle = false
					}
				case "timer":
					timer, err := strconv.ParseUint(find[2], 10, 64)
					if err != nil {
						return opts, false, errors.New(fmt.Sprintf("Invalid Value for %s: Must be an number >= 0, 0 is equal no timer!", find[1]))
					}
					opts.Timer = uint(timer)
				case "language":
					availableLanguages := make([]string, 0, 2)
					contain := false
					for language := range languages {
						availableLanguages = append(availableLanguages, language)
						if find[2] == language {
							contain = true
						}
					}
					if !contain {
						return opts, false, errors.New(fmt.Sprintf("Invalid Value for %s: Language not Available!", find[1]))
					}
					opts.Language = find[2]
				}
			} else {
				return opts, false, errors.New(fmt.Sprintf("Invalid Argument %s: See quiz help to see all supported flags!", find[1]))
			}
		}
	}
	return opts, false, nil
}
