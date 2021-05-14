package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

func main() {
	var qs = []*survey.Question{
		{
			Name: "action",
			Prompt: &survey.Select{
				Message: ">>> Select ACTION:",
				Options: []string{
					"feat:    	new features, big improvement",
					"fix:     	bug fixes",
					"refactor:	code refactoring, without any visual user change",
					"chore:   	any else but the app",
					"test:    	any kind of tests (functionnal, unitary...)",
					"style:   	cosmetic changes (spaces vs tab...)",
					"docs:    	documentation",
				},
			},
		},
		{
			Name: "topic",
			Prompt: &survey.Select{
				Message: ">>> Select TOPIC:",
				Options: []string{
					"dev:         developpers (API changes, refactors...)",
					"usr:         final users (UI changes)",
					"pkg:         packagers   (packaging changes)",
					"test:        testers     (test only related changes)",
				},
			},
		},
		{
			Name: "message",
			Prompt: &survey.Input{
				Message: ">>> Type Commit message:",
			},
			Validate: survey.Required,
		},
		{
			Name: "tag",
			Prompt: &survey.Select{
				Message: ">>> Select TOPIC:",
				Options: []string{
					"refactor:    obviously for refactoring code only'",
					"minor:       a very meaningless change (a typo, adding a comment)'",
					"cosmetic:    cosmetic driven change (re-indentation, 80-col...)'",
					"wip:         partial functionality but complete subfunctionality.'}",
				},
			},
		},
		{
			Name: "ticketnum",
			Prompt: &survey.Input{
				Message: ">>> Select ID-ISSUE:",
				Default: "nt",
			},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Action    string `survey:"action"`
		Topic     string `survey:"topic"`
		Message   string `survey:"message"`
		Tag       string `survey:"tag"`
		TicketNum string `survey:"ticketnum"`
	}{}

	err := survey.Ask(qs, &answers)

	if err == terminal.InterruptErr {
		fmt.Println("interrupted")

		os.Exit(0)
	} else if err != nil {
		panic(err)
	}

	fmt.Printf(
		"\n%s(%s): %s #%s [!%s]\n\n",
		strings.Split(answers.Action, ":")[0],
		strings.Split(answers.Topic, ":")[0],
		answers.Message,
		strings.Split(answers.Tag, ":")[0],
		answers.TicketNum,
	)

	fmt.Scanf("h")
}
