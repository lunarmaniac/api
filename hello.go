package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Tips struct {
	Tips []string `json:"tips"`
}

func main() {
	jsonFile, err := os.Open("tips.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)
	var tips Tips
	json.Unmarshal(byteValue, &tips)

	fmt.Println("Starting API server")

	s := rand.NewSource(time.Now().Unix())
	R := rand.New(s)

	http.HandleFunc("/tip", func(w http.ResponseWriter, r *http.Request) {
		i := R.Intn(len(tips.Tips))
		randomTip := tips.Tips[i]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"tip": randomTip})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
