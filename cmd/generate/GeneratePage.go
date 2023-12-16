package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"text/template"
	"unicode/utf8"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"jirku.sk/mcmamina/pkg/text"
)

func GeneratePage(c *cli.Context) error {
	params := pageParameters{}
	var err error
	params.Name = c.String("name")
	if params.Name == "" {
		name := promptui.Prompt{
			Label: "Provide Page name",
			Validate: func(v string) error {
				if v == "" {
					return errors.New("validation of the name")
				}
				return nil
			},
		}
		params.Name, err = name.Run()
		if err != nil {
			return err
		}
		params.Name = text.ToCamelCase(params.Name)
	}

	params.Path = c.String("path")
	if params.Path == "" {
		path := promptui.Prompt{
			Label: "Provide Page Path",
			Validate: func(v string) error {
				if v == "" {
					return errors.New("empty path")
				}
				_, err := url.Parse(v)
				if err != nil {
					return errors.New("invalid path")
				}
				return nil
			},
		}
		params.Path, err = path.Run()
		if err != nil {
			return err
		}
	}
	params.Title = c.String("title")
	if params.Title == "" {
		title := promptui.Prompt{
			Label: "Provide Page Title",
			Validate: func(v string) error {
				if v == "" {
					return errors.New("empty title")
				}
				_, err := url.Parse(v)
				if err != nil {
					return errors.New("invalid title")
				}
				return nil
			},
		}
		params.Title, err = title.Run()
		if err != nil {
			return err
		}
	}

	if err = generatePageTempl(c, params); err != nil {
		return err
	}
	if err = generatePageTs(c, params); err != nil {
		return err
	}
	if err = generatePageCSS(c, params); err != nil {
		return err
	}
	if err = generatePageHandler(c, params); err != nil {
		return err
	}

	return nil
}

func generatePageTempl(c *cli.Context, params pageParameters) error {
	var pageTpl bytes.Buffer
	err := template.Must(template.New("pageTempl").Parse(pageTemplTmpl)).Execute(&pageTpl, params)
	if err != nil {
		return err
	}
	printBox(fmt.Sprintf("👉 %s...", params.PageTemplPath()))
	if c.Bool("dry-run") {
		printHighlightedGo(pageTpl)
		if err != nil {
			return err
		}
		return nil
	}
	err = writeToFile(params.PageTemplPath(), pageTpl)
	if err != nil {
		return err
	}
	return nil
}
func generatePageTs(c *cli.Context, params pageParameters) error {
	var buffer bytes.Buffer
	err := template.Must(template.New("pageTs").Parse(pageTs)).Execute(&buffer, params)
	if err != nil {
		return err
	}

	printBox(fmt.Sprintf("👉 %s...", params.PageTsPath()))
	if c.Bool("dry-run") {
		printHighlightedTs(buffer)
		if err != nil {
			return err
		}
		return nil
	}
	err = writeToFile(params.PageTsPath(), buffer)
	if err != nil {
		return err
	}
	return nil
}
func generatePageCSS(c *cli.Context, params pageParameters) error {
	var buffer bytes.Buffer
	err := template.Must(template.New("pageCss").Parse(pageCss)).Execute(&buffer, params)
	if err != nil {
		return err
	}

	printBox(fmt.Sprintf("👉 %s...", params.PageCSSPath()))
	if c.Bool("dry-run") {
		printHighlightedCss(buffer)
		if err != nil {
			return err
		}
		return nil
	}
	err = writeToFile(params.PageCSSPath(), buffer)
	if err != nil {
		return err
	}
	return nil
}
func generatePageHandler(c *cli.Context, params pageParameters) error {
	var buffer bytes.Buffer
	err := template.Must(template.New("pageHandler").Parse(pageHandlerGo)).Execute(&buffer, params)
	if err != nil {
		return err
	}

	printBox(fmt.Sprintf("👉 %s...", params.PageHandlerPath()))
	if c.Bool("dry-run") {
		printHighlightedGo(buffer)
		if err != nil {
			return err
		}
		return nil
	}
	err = writeToFile(params.PageHandlerPath(), buffer)
	if err != nil {
		return err
	}
	return nil
}

func printBox(text string) {
	textWidth := utf8.RuneCountInString(text)
	// Adjust the box width for characters that take up more than one visual space
	// Assuming each emoticon or wide character takes up 2 spaces
	extraWidth := 0
	for _, runeValue := range text {
		if utf8.RuneLen(runeValue) > 1 {
			extraWidth++
		}
	}

	top := "┌" + strings.Repeat("─", textWidth+extraWidth+4) + "┐"
	middle := "│  " + text + "  │"
	bottom := "└" + strings.Repeat("─", textWidth+extraWidth+4) + "┘"

	fmt.Println(top)
	fmt.Println(middle)
	fmt.Println(bottom)
}

func writeToFile(path string, buffer bytes.Buffer) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = buffer.WriteTo(file)
	if err != nil {
		return err
	}
	return nil
}

func printHighlighted(code bytes.Buffer, lexer chroma.Lexer) error {
	iterator, err := lexer.Tokenise(nil, code.String())
	if err != nil {
		return err
	}
	formatter := formatters.Get("terminal256")
	style := styles.Get("dracula")
	err = formatter.Format(os.Stdout, style, iterator)
	if err != nil {
		return err
	}
	return nil
}
func printHighlightedGo(code bytes.Buffer) error {
	lexer := lexers.Get("go")
	return printHighlighted(code, lexer)
}
func printHighlightedTs(code bytes.Buffer) error {
	lexer := lexers.Get("TypeScript")
	return printHighlighted(code, lexer)
}
func printHighlightedCss(code bytes.Buffer) error {
	lexer := lexers.Get("CSS")
	return printHighlighted(code, lexer)
}
