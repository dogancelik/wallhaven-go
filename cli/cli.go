package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	".."
)

func saveImage(imageUrl string, saveDirectory string) string {
	parsedUrl, errParse := url.Parse(imageUrl)
	if errParse != nil {
		panic(errParse)
	}

	errMkdir := os.MkdirAll(saveDirectory, os.ModePerm)
	if errMkdir != nil {
		panic(errMkdir)
	}

	savePath := filepath.Join(saveDirectory, path.Base(parsedUrl.Path))

	file, errCreate := os.Create(savePath)
	if errCreate != nil {
		panic(errCreate)
	}

	res, errHttp := http.Get(imageUrl)
	if errHttp != nil {
		panic(errHttp)
	}
	defer res.Body.Close()

	_, errWrite := io.Copy(file, res.Body)
	if errWrite != nil {
		panic(errWrite)
	}
	defer file.Close()

	return savePath
}

const version string = "v0.2.0"

func printVersion() {
	cmdName := filepath.Base(os.Args[0])
	fmt.Println(cmdName + " " + version)
}

func die(code int, msg string, err error) {
	os.Stderr.WriteString(msg)
	if err != nil {
		os.Stderr.WriteString("\nError: " + err.Error())
	}
	os.Exit(code)
}

func main() {
	term := flag.String("t", "", "search term")
	category := flag.String("c", "gap", "categories (available: [g][a][p])")
	purity := flag.String("p", "w", "purity (available: [w][s][n])")
	res := flag.String("r", "", "resolutions (example: 1920x1080+)")
	sort := flag.String("s", "random", "sorting (available: random, relevance, date_added, views)")
	page := flag.Int("page", 1, "search page number")
	set := flag.Bool("set", false, "set first result as wallpaper (default: true)")
	all := flag.Bool("all", false, "show all results as URLs (default: false)")
	saveDir := flag.String("dir", "$TEMP", "save directory")
	version := flag.Bool("v", false, "show version number")
	flag.Parse()

	if *version == true {
		printVersion()
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		printVersion()
		flag.Usage()
		os.Exit(0)
	}

	opt := wallhaven.Options{}
	opt.Term = *term
	parseCategories(&opt, *category)
	parsePurity(&opt, *purity)
	opt.Resolutions = wallhaven.ParseResolutions(*res)
	opt.Sorting = *sort
	opt.Page = *page

	results, errSearch := wallhaven.Search(&opt)

	if errSearch != nil {
		die(1, "Search error", errSearch)
	}

	if *all == true {
		for _, v := range results {
			fmt.Println(v.ImageUrl)
		}
		os.Exit(0)
	}

	if len(results) == 0 {
		die(2, "No results found", nil)
	}

	if *set == true {
		imageUrl := results[0].ImageUrl

		if *saveDir == "$TEMP" {
			*saveDir = os.TempDir()
		}
		savePath := saveImage(imageUrl, *saveDir)

		absPath, errAbs := filepath.Abs(savePath)
		if errAbs != nil {
			fmt.Fprintln(os.Stderr, "Error when getting absolute path:", errAbs.Error())
			fmt.Println(savePath)
		} else {
			fmt.Println(absPath)
		}

		settings, errLoad := loadSettings(getSettingsPath())
		if errLoad != nil {
			die(3, "Loading settings failed", errLoad)
		}

		errSet := SetWallpaper(savePath, settings)
		if errSet != nil {
			die(4, "Setting image as wallpaper failed", errSet)
		}
	}
}
