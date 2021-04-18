package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/foolin/pagser"
	"github.com/gocolly/colly"
)

var p = pagser.New()

type ProductData struct {
	Title       string `pagser:"h1#title->text()"`
	Price       string `pagser:"div#price span#priceblock_ourprice->text()"`
	ImgUrl      string `pagser:"img#landingImage->attr(src)"`
	Ratings     string `pagser:"div#averageCustomerReviews .a-declarative->first()->text()"`
	Description string `pagser:"div#productDescription p->text()"`
}

// Not needed as of Now
// type PostData struct {
// 	Title       string `json:"Title"`
// 	Price       int    `json:"Price"`
// 	ImgUrl      string `json:"ImgURL"`
// 	Reviews     string `json:"Rating"`
// 	Description string `json:"Description"`
// }

func PostHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Method Not Supported")
		return
	}
	r.ParseMultipartForm(1024)
	URL := r.Form.Get("url")
	var c = colly.NewCollector(
		colly.AllowedDomains("www.amazon.in"),
		colly.MaxDepth(0),
	)
	var proddata ProductData

	// Get the Title
	c.OnHTML("html", func(e *colly.HTMLElement) {
		p.ParseSelection(&proddata, e.DOM)
		fmt.Printf("%v", proddata)
	})

	c.Visit(URL)
	c.Wait()
	json.NewEncoder(w).Encode(proddata)
}

func main() {
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe(":5051", nil)
	// c.Visit("https://www.amazon.in/dp/B076B2BC19/ref=vp_d_pb_TIER3_cmlr_lp_B073NTCT4R_pd?_encoding=UTF8&")
}
