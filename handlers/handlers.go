package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "example.com/go-ci-sample/calc"
)

func Health(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// /add?a=1&b=2 -> {"result":3}
func AddHandler(w http.ResponseWriter, r *http.Request) {
    aStr := r.URL.Query().Get("a")
    bStr := r.URL.Query().Get("b")

    a, err := strconv.Atoi(aStr)
    if err != nil {
        http.Error(w, "invalid 'a'", http.StatusBadRequest)
        return
    }
    b, err := strconv.Atoi(bStr)
    if err != nil {
        http.Error(w, "invalid 'b'", http.StatusBadRequest)
        return
    }

    res := calc.Add(a, b)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]int{"result": res})
}
