package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type OUTPUT_TYPE string

const (
	PDF  OUTPUT_TYPE = "pdf"
	HTML OUTPUT_TYPE = "html"
)

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

	return fmt.Sprintf(htmlTemplate, title, getCssStyle(stylePath), buf.String()), nil
}

func getContentFromFile(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func getCssStyle(path string) string {
	stylesCss, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading styles.css: %v", err)
		return ""
	}
	return string(stylesCss)
}

func toHtmlFile(outFile, content string) error {
	outPath := fmt.Sprintf("%s.html", outFile)
	err := os.WriteFile(outPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	return nil
}

func convertToPdf(content, outFile, title string) error {
	err := toHtmlFile(outFile, content)
	if err != nil {
		return err
	}

	inFile := fmt.Sprintf("%s.html", outFile)
	pdfFile := fmt.Sprintf("%s.pdf", outFile)

	cmd := exec.Command("wkhtmltopdf", "--title", title, inFile, pdfFile)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Error converting HTML to PDF: %v", err)
	}

	fmt.Printf("PDF generated successfully: %s\n", pdfFile)
	return nil
}

func main() {
	var filePath string
	var outFile string
	var templateTitle string
	var enableNativeHtml bool
	var stylePath string
	var outputType OUTPUT_TYPE
	var outputTypeString string

	flag.StringVar(&filePath, "f", "", "File to convert")
	flag.StringVar(&outFile, "o", "out", "Output file (without extension)")
	flag.StringVar(&outputTypeString, "type", "pdf", "Output type (pdf or html)")
	flag.StringVar(&templateTitle, "t", "Resume", "Title of the resume")
	flag.StringVar(&stylePath, "s", "styles.css", "Path to styles.css")
	flag.BoolVar(&enableNativeHtml, "native", false, "Enable native HTML rendering")

	flag.Parse()

	if outputTypeString != "pdf" && outputTypeString != "html" {
		fmt.Println("Invalid output type. Please use pdf or html")
		return
	}

	outputType = OUTPUT_TYPE(outputTypeString)

	if filePath == "" {
		fmt.Println("Please provide a file to convert")
		return
	}

	content, err := getContentFromFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	template, err := convertToHtml(content, templateTitle, stylePath, enableNativeHtml)
	if err != nil {
		fmt.Printf("Error converting to HTML: %v\n", err)
		return
	}

	if outputType == PDF {
		err = convertToPdf(template, outFile, templateTitle)
		if err != nil {
			fmt.Printf("Error converting to PDF: %v\n", err)
		}
	} else {
		err = toHtmlFile(outFile, template)
		if err != nil {
			fmt.Printf("Error converting to HTML: %v\n", err)
		} else {
			fmt.Printf("HTML generated successfully: %s.html\n", outFile)
		}
	}
}
