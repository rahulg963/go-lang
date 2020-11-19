package main

type Query struct {
	Term     string `json:"term"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

type Result struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// func helper(w http.ResponseWriter, r *http.Request) {
// 	// get url
// 	// url := r.URL
// 	// query := url.Query() // map[string][]string
// 	// q := query["q"]	//[]string{"products"}
// 	// page := query.Get("page") // returns first element of list-> "1"

// 	// process form data
// 	// err := r.ParseForm()
// 	// f := r.Form
// 	// username := f["username"] // slice of string
// 	// pass := f.Get("password") // single value

// 	// reading from json
// 	dec := json.NewDecoder(r.Body)
// 	var query Query
// 	err := dec.Decode(&query)

// 	// writing to json
// 	var results []Result = models.GetUsers()
// 	enc := json.NewEncoder(w)
// 	err := enc.Encode(results)
// }
