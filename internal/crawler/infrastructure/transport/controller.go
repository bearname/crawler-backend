package transport

import (
	"github.com/bearname/url-extractor/pkg/app"
	"net/http"
	"strconv"
)

type Controller struct {
	BaseController
	crawler *app.Crawler
}

func New(crawler *app.Crawler) *Controller {
	c := new(Controller)
	c.crawler = crawler
	return c
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	urlLink, ok := query["url"]

	if !ok || len(urlLink) != 1 {
		c.BaseController.WriteError(w, ErrBadRequest)
		return
	}
	depth, ok := query["depth"]

	var depthSearch = 4
	if ok && depth != nil && len(depth) == 1 {
		atoi, err := strconv.Atoi(depth[0])
		if err != nil {
			c.BaseController.WriteError(w, ErrBadRequest)
			return
		}
		depthSearch = atoi
	}

	c.crawler.Crawl(urlLink[0], depthSearch, &app.HttpFetcher{})

	urls := make([]string, 0, len(c.crawler.Crawled))
	for k := range  c.crawler.Crawled {
		urlLink = append(urlLink, k)
	}

	responseUrls := struct {
		Crawled []string `json:"crawled"`
	}{Crawled: urls}

	c.BaseController.WriteJsonResponse(w, responseUrls)
}
