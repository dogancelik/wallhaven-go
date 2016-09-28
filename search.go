package wallhaven

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imdario/mergo"
)

const searchUrl string = "https://alpha.wallhaven.cc/search"
const imageUrl string = "https://wallpapers.wallhaven.cc/wallpapers/full/wallhaven-"

type Result struct {
	Id           int
	Resolution   string
	Favorites    int
	PageUrl      string
	ThumbnailUrl string
	ImageUrl     string
}

func createQuery(opt *Options) string {
	query := "?"
	query += "q=" + url.QueryEscape(opt.Term)
	query += "&categories=" + opt.Categories.ToBits()
	query += "&purity=" + opt.Purity.ToBits()
	query += "&resolutions=" + url.QueryEscape(strings.Join(opt.Resolutions, ","))
	query += "&sorting=" + opt.Sorting
	query += "&page=" + strconv.Itoa(opt.Page)
	return query
}

func Search(opt *Options) ([]Result, error) {
	mergo.Merge(opt, getDefaultOptions())

	results := []Result{}
	url := searchUrl + createQuery(opt)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return results, err
	}

	doc.Find("#thumbs figure").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("data-wallpaper-id")
		idAsInt, _ := strconv.Atoi(id)

		favs := s.Find(".wall-favs").Text()
		favsAsInt, _ := strconv.Atoi(favs)

		pageUrl, _ := s.Find(".preview").Attr("href")
		thumbUrl, _ := s.Find(".lazyload").Attr("data-src")
		result := Result{
			Id:           idAsInt,
			Resolution:   s.Find(".wall-res").Text(),
			Favorites:    favsAsInt,
			PageUrl:      pageUrl,
			ThumbnailUrl: thumbUrl,
			ImageUrl:     imageUrl + id + ".jpg",
		}
		results = append(results, result)
	})
	return results, nil
}
