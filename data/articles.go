package data

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles = &[]Article{
	{
		ID:      "1",
		Title:   "Title 1",
		Desc:    "Desc 1",
		Content: "Content 1",
	},
	{
		ID:      "2",
		Title:   "Title 2",
		Desc:    "Desc 2",
		Content: "Content 2",
	},
}
