# Gofont

_A simple google font downloader written in Golang_

## About this project

This project contains my very first lines of golang. I created this project mainly because I want to try out Go (I really hate it so far).

## What does it do?

Just feed it an url to google font api and it will download that font to your desired location!

**Example of a google font api url**

```
https://fonts.googleapis.com/css?family=Charm|Dancing+Script
```

### Example 1: Download fonts from google to a directory

Syntax:

```
$ gofont "<google_font_link>" "<location_for_downloaded_font>"
```

Example:

```
$ gofont "https://fonts.googleapis.com/css?family=Charm|Dancing+Script" "fonts"
```

### Example 2: Download fonts from google and generate a css file for that font

Syntax:

```
$ gofont --cssFile="<location_for_css>" "<google_font_link>" "<location_for_downloaded_font>"
```

Example:

```
$ gofont --cssFile="./styles/fonts.css" "https://fonts.googleapis.com/css?family=Charm|Dancing+Script" "fonts"
```

## Todo

I'm still trying out Go so maybe I'll add some more features such as:

- [x] Generate a CSS file pointing to the location of the downloaded font
- [ ] Make logger more colorful by printing the details with color
- [x] Cry a lot because I don't understand how Go `import` works
- [ ] Search for fonts on google fonts

## License

- [MIT](LICENSE)
