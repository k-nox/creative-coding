package tui

import (
	"errors"
	"strings"

	"github.com/charmbracelet/huh"
)

func newForm() *huh.Form {
	var want string
	var lettercase string
	return huh.NewForm(
		huh.NewGroup(
			letterOrPunctuationSelect(&want),
		),
		huh.NewGroup(
			lettercaseSelect(&want, &lettercase),
		),
		huh.NewGroup(
			charSelect(&lettercase),
		),
		huh.NewGroup(
			doneConfirm(&done),
		),
	)
}

func charSelect(lettercase *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		TitleFunc(func() string {
			switch *lettercase {
			case "UPPERCASE":
				return "PICK YOUR LETTER!"
			case "lowercase":
				return "pick your letter."
			default:
				return "Pick your punctuation."
			}
		}, lettercase).
		OptionsFunc(func() []huh.Option[string] {
			switch *lettercase {
			case "UPPERCASE":
				return options(letters)
			case "lowercase":
				return options(strings.ToLower(letters))
			default:
				return options(punctuation)
			}
		}, lettercase).
		Key("input")
}

func lettercaseSelect(want *string, lettercase *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		TitleFunc(func() string {
			switch *want {
			case "letter":
				return "Do you want an uppercase or lowercase letter?"
			default:
				return "This is a filler, just hit continue."
			}
		}, want).
		OptionsFunc(func() []huh.Option[string] {
			switch *want {
			case "letter":
				return huh.NewOptions("UPPERCASE", "lowercase")
			default:
				return huh.NewOptions("continue", "Continue")
			}
		}, want).
		Validate(func(s string) error {
			switch *want {
			case "letter":
				return nil
			default:
				if s == "Continue" {
					return errors.New("no, I said to hit 'continue', not 'Continue'")
				}
				return nil
			}
		}).
		Value(lettercase)
}

func letterOrPunctuationSelect(want *string) *huh.Select[string] {
	return huh.NewSelect[string]().
		Title("Do you want a letter or punctuation?").
		Options(
			huh.NewOption("letter", "letter"),
			huh.NewOption("punctuation", "punctuation"),
		).
		Value(want)
}

func doneConfirm(done *bool) *huh.Confirm {
	return huh.NewConfirm().
		Title("Are you done yet?").
		Affirmative("!No").
		Negative("!Yes").
		Value(done)
}

func options(from string) []huh.Option[string] {
	opts := []huh.Option[string]{}
	for _, char := range from {
		opts = append(opts, huh.NewOption(string(char), string(char)))
	}
	return opts
}
