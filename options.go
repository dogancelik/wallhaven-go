package wallhaven

const SortRandom = "random"
const SortRelevance = "relevance"
const SortDateAdded = "date_added"
const SortViews = "views"

type Options struct {
	Term        string
	Categories  Categories
	Purity      Purity
	Resolutions []string
	Sorting     string
	Page        int
}

func GetDefaultOptions() Options {
	opt := Options{}
	opt.Term = ""
	opt.Categories.General = true
	opt.Categories.Anime = true
	opt.Categories.People = true
	opt.Purity.Sfw = true
	opt.Purity.Sketchy = false
	opt.Purity.Nsfw = false
	opt.Resolutions = make([]string, 0)
	opt.Sorting = SortRandom
	opt.Page = 1
	return opt
}
