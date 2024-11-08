# Resumemk

This a simple golang CLI tool that allows you to generate html files from markdown files.
It also enables you to add your own custom css.

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
