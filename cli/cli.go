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

func saveImage(imageUrl string) string {
	parsedUrl, errParse := url.Parse(imageUrl)
	if errParse != nil {
		panic(errParse)
	}
	savePath := filepath.Join(os.TempDir(), path.Base(parsedUrl.Path))

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

const version string = "v0.1.1"

func printVersion() {
	cmdName := filepath.Base(os.Args[0])
	fmt.Println(cmdName + " " + version)
}

func main() {
	term := flag.String("t", "", "search term")
	category := flag.String("c", "gap", "categories (available: [g][a][p])")
	purity := flag.String("p", "w", "purity (available: [w][s][n])")
	res := flag.String("r", "", "resolutions (example: 1920x1080+)")
	sort := flag.String("s", "random", "sorting (available: random, relevance, date_added, views)")
	page := flag.Int("page", 1, "page (default: 1)")
	set := flag.Bool("set", true, "set first result as wallpaper (default: true)")
	all := flag.Bool("all", false, "show all results as URLs (default: false)")
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
		panic(errSearch)
	}

	if *all == true {
		for _, v := range results {
			fmt.Println(v.ImageUrl)
		}
	}

	if *set == true {
		imageUrl := results[0].ImageUrl
		savePath := saveImage(imageUrl)
		fmt.Println(savePath)
		settings, errLoad := loadSettings(getSettingsPath())
		if errLoad != nil {
			panic(errLoad)
		}

		errSet := SetWallpaper(savePath, settings)
		if errSet != nil {
			panic(errSet)
		}
	}
}
