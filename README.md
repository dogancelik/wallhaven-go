# wallhaven-go

A library for searching Wallhaven, and a partially-working CLI tool.
This library is experimental.

## Install

```
go get bitbucket.org/dogancelik/wallhaven-go
```

## Usage

```go
package main

import (
	"fmt"
	wallhaven "bitbucket.org/dogancelik/wallhaven-go"
)

func main() {
	myOptions := wallhaven.Options{
		Term: "dog",
		Categories: wallhaven.Categories{
			General: true,
			Anime: true,
			People: true
		},
		Purity: wallhaven.Purity{
			Sfw: true,
			Sketchy: false,
			Nsfw: false
		},
		Resolutions: wallhaven.ParseResolutions("1920x1080+"),
		Sorting: "random",
		Page: 1
	}

	// you can omit every field in myOptions

	var results []wallhaven.Result
	var err error
	results, err = wallhaven.Search(myOptions)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(results)
	}
}
```

## CLI

CLI is buggy because Go is buggy [[1]](https://github.com/golang/go/issues/14575), [[2]](https://github.com/golang/go/issues/16131), [[3]](https://github.com/golang/go/issues/17149).

Note: If you are using Windows, you need to install WallpaperChanger from [here (github.com/philhansen/WallpaperChanger)](https://github.com/philhansen/WallpaperChanger/blob/master/WallpaperChanger.exe) and put it beside `wallhaven.exe`

```sh
$ wallhaven
wallhaven v0.1.1
Usage of wallhaven:
  -all           show all results as URLs (default false)
  -c string     categories (available: [g][a][p]) (default "gap")
  -p string     purity (available: [w][s][n]) (default "w")
  -page int     page (default: 1) (default 1)
  -r string     resolutions (example: 1920x1080+)
  -s string     sorting (available: random, relevance, date_added, views) (default "random")
  -set           set first result as wallpaper (default true)
  -t string     search term
  -v             show version number
```

### Set first result as wallpaper

```
wallhaven -t=dog --set
```

### Get search results as URLs

```
wallhaven -t=dog --all
```

### Change wallpaper change command

Open `wallhaven.json` file and find your Operating System, edit the command of your OS and wallhaven will run it that.
`Windows`, `Linux`, `Mac` are JSON array of string fields, so you can execute commands one by one.

#### Example

If we want to execute several commands on Windows, we can do something like this below:

```json
{
  "Linux": [
    "gsettings set org.gnome.desktop.background picture-uri file://__WALL__"
  ],
  "Windows": [
    "WallpaperChanger.exe __WALL__",
    "D:\\MyTools\\TakeScreenshot.exe --desktop",
    "D:\\MyScripts\\CleanDesktop.bat"
  ],
  "Mac": [
    "osascript -e 'tell application \"Finder\" to set desktop picture to POSIX file __WALL__'"
  ]
}
```
