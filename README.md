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
	wallhaven "bitbucket.org/dogancelik/wallhaven-go/search"
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

Note: If you are using Windows, you need to install WallpaperChanger from [here (github.com/philhansen/WallpaperChanger)](https://github.com/philhansen/WallpaperChanger/blob/master/WallpaperChanger.exe)

* **Set first result as wallpaper:**  
	```
	cli "dog" --set
	```

* **Get search results as URLs:**  
	```
	cli "dog" --all
	```
