# Resumemk

This a simple golang CLI tool that allows you to  easily convert your resume from markdown to pdf or html.
It uses [goldmark](https://github.com/yuin/goldmark) to convert the markdown to html and [wkhtmltopdf](https://wkhtmltopdf.org/) to convert the html to pdf.


## Features
- Convert resume from markdown to pdf or html
- Supports native html rendering
- Supports custom styles
- Easy to use

## Installation
First of all, you need to install wkhtmltopdf.
So go to their [website](https://wkhtmltopdf.org/downloads.html) and download the latest version for your OS.
For arch, you can just run:

```bash
yay -S wkhtmltopdf-static
```

## Usage

```bash
./resumemk -f exemple.md -t 'Resume'  -native 
```

## Example

```bash
resumemk -i resume.md -o resume.pdf
```

## License

MIT

# Inspiration
- https://www.youtube.com/watch?v=zktKckI4180
