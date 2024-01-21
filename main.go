package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func main() {
	searcher := Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch(searcher))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("shakesearch available at http://localhost:%s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Searcher struct {
	CompleteWorks string
	SuffixArray   *suffixarray.Index
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		query := params["q"]
		if len(query) < 1 { 
    		w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		// pass through the "page" param to Search to return the next page
		page := params["page"]
		page_number := 0
		if len(page) < 1 { 
    		fmt.Println("page param missing")
		} else {
			page_number, _ = strconv.Atoi(page[0])
		}
		results := searcher.Search(query[0], page_number)
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	// store the complete works as a lowercase array so we can make case insensitive lookups against it.
	s.SuffixArray = suffixarray.New(bytes.ToLower(dat))
	return nil
}

func (s *Searcher) Search(query string, page_number int) []string {
	// also make the query string lowercase so the lookup against the character array is case insensitive
	idxs := s.SuffixArray.Lookup(bytes.ToLower([]byte(query)), -1)
	// sort idxs so the next page returns a new a set of results
	sort.Ints(idxs)
	results := []string{}

	// limit result to MAX 20 entries per page
	start_idx := page_number*20
	// cap the end index to be the len of the idxs array (we can only return what's remaining)
	end_idx := int(math.Min(float64((page_number + 1)*20), float64(len(idxs))))
	for _, idx := range idxs[start_idx: end_idx] {
		results = append(results, s.CompleteWorks[idx-250:idx+250])
	}
	return results
}
