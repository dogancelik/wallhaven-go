package main

import (
	"fmt"

	"../search"
)

func getChar(run rune) string {
	return fmt.Sprintf("%c", run)
}

func parseCategories(opt *search.Options, cat string) {
	for _, r := range cat {
		switch getChar(r) {
		case "g":
			opt.Categories.General = true
		case "a":
			opt.Categories.Anime = true
		case "p":
			opt.Categories.People = true
		}
	}
}

func parsePurity(opt *search.Options, pur string) {
	for _, r := range pur {
		switch getChar(r) {
		case "w":
			opt.Purity.Sfw = true
		case "s":
			opt.Purity.Sketchy = true
		case "n":
			opt.Purity.Nsfw = true
		}
	}
}
