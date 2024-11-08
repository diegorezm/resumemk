package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func getCssStyle(path string) string {
	stylesCss, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading styles.css: %v", err)
		return ""
	}
	return string(stylesCss)
}

func convertToHtml(content []byte, title, stylePath string, native bool) (string, error) {
	htmlTemplate := `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width">
  <title>%s</title>
  <style>
    %s
  </style>
</head>
<body>
  <div id="resume">
    %s
  </div>
</body>
</html>`

	var md goldmark.Markdown
	if native {
		md = goldmark.New(
			goldmark.WithExtensions(
				extension.GFM,
				extension.Typographer,
			),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithUnsafe(),
				html.WithHardWraps(),
				html.WithXHTML(),
			),
		)
	} else {
		md = goldmark.New(
			goldmark.WithExtensions(
				extension.GFM,
				extension.Typographer,
			),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
			goldmark.WithRendererOptions(
				html.WithXHTML(),
			),
		)
	}

	var buf bytes.Buffer

	err := md.Convert(content, &buf)

	if err != nil {
		return "", err
	}

	template := fmt.Sprintf(htmlTemplate, title, getCssStyle(stylePath), buf.String())
	return template, nil
}

func getContentFromFile(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func main() {
	var filePath string
	var outFile string
	var templateTitle string
	var enableNativeHtml bool
	var stylePath string

	flag.StringVar(&filePath, "f", "", "File to convert")
	flag.StringVar(&outFile, "o", "out.html", "Output file")
	flag.StringVar(&templateTitle, "t", "Resume", "Titlte of the resume")
	flag.StringVar(&stylePath, "s", "styles.css", "Path to styles.css")
	flag.BoolVar(&enableNativeHtml, "native", false, "Enable native html")

	flag.Parse()

	if filePath == "" {
		fmt.Println("Please provide a file to convert")
		return
	}

	content, err := getContentFromFile(filePath)

	if err != nil {
		fmt.Printf("Error reading file: %v", err)
	}

	template, err := convertToHtml(content, templateTitle, stylePath, enableNativeHtml)

	if err != nil {
		fmt.Printf("Error converting to html: %v", err)
		return
	}

	if outFile != "" {
		err = os.WriteFile(outFile, []byte(template), 0644)
		if err != nil {
			fmt.Printf("Error writing to file: %v", err)
			return
		}
	}
}
