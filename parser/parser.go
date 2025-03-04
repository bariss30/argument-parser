package parser

import (
	"fmt"
	"os"
)

type Argument struct {
	Name        string
	Description string
	Required    bool
	Default     string
	Value       string
}

type ArgumentParser struct {
	Title       string
	Description string
	Usage       string
	Arguments   []Argument
}

func NewArgumentParser(title, description, usage string) ArgumentParser {
	return ArgumentParser{
		Title:       title,
		Description: description,
		Usage:       usage,
		Arguments:   []Argument{},
	}
}

func (dd *ArgumentParser) AddArgument(name, description string, required bool, defaultValue string) {
	arg := Argument{
		Name:        name,
		Description: description,
		Required:    required,
		Default:     defaultValue,
	}
	dd.Arguments = append(dd.Arguments, arg)
}

func (ap *ArgumentParser) Parse() error {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		ap.Help()
		return nil
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		found := false //Geçerlimi

		for j := range ap.Arguments {
			if ("--"+ap.Arguments[j].Name) == arg || ("-"+ap.Arguments[j].Name) == arg {
				found = true
				if i+1 < len(args) && args[i+1][0] != '-' {
					ap.Arguments[j].Value = args[i+1]
					i++
				} else {
					ap.Arguments[j].Value = "true"
				}
			}
		}

		if !found {
			return fmt.Errorf("bilinmeyen argüman: %s", arg)
		}
	}

	for _, arg := range ap.Arguments {
		if arg.Required && arg.Value == "" {
			return fmt.Errorf("zorunlu argüman eksik: --%s", arg.Name)
		}
	}

	return nil
}

func (ap *ArgumentParser) Help() {
	fmt.Println(ap.Title)
	fmt.Println(ap.Description)
	fmt.Println(ap.Usage)
	fmt.Println("\nKullanılabilir Argümanlar:")
	for _, arg := range ap.Arguments {
		fmt.Printf("--%s: %s\n", arg.Name, arg.Description)
	}
}

func (ap *ArgumentParser) Get(name string) string {
	for _, arg := range ap.Arguments {
		if arg.Name == name {
			if arg.Value != "" {
				return arg.Value
			}
			return arg.Default
		}
	}
	return ""
}
