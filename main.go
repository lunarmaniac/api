package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type Tip struct {
	Tip      string `json:"tip"`
	Category string `json:"category"`
}

var tipsByCategory = make(map[string][]Tip)
var allTips []Tip

// load all JSON files from the folder
func loadTips(folder string) error {
	return filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".json") {
			return nil
		}

		fileBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var tips []Tip
		if err := json.Unmarshal(fileBytes, &tips); err != nil {
			return err
		}

		for _, tip := range tips {
			tipsByCategory[tip.Category] = append(tipsByCategory[tip.Category], tip)
			allTips = append(allTips, tip)
		}

		return nil
	})
}

func randomTipFromSlice(tips []Tip) Tip {
	rand.Seed(time.Now().UnixNano())
	return tips[rand.Intn(len(tips))]
}

func main() {
	if err := loadTips("./tips"); err != nil {
		log.Fatalf("Failed to load tips: %v", err)
	}

	http.HandleFunc("/tip", func(w http.ResponseWriter, r *http.Request) {
		if len(allTips) == 0 {
			http.Error(w, "No tips available", http.StatusInternalServerError)
			return
		}
		tip := randomTipFromSlice(allTips)
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		encoder.Encode(map[string]string{"tip": tip.Tip, "category": tip.Category})
	})
	
	fmt.Println("Starting API server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
