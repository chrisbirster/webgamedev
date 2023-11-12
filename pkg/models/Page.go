package models

type (
	Link struct {
		Name          string
		Path          string
		IsCurrentPage bool
	}

	Arthor struct {
		Name string
		Link string
	}

	MetaData struct {
		Arthor     Arthor
		Date       string
		Categories []string
		Tags       []string
	}
	Post struct {
		Title       string
		Description string
		MetaData    MetaData
		Classes     string
	}

	Tutorial struct {
		Name  string
		Link  string
		Image string
	}

	Community struct {
		Name  string
		Link  string
		Image string
	}

	Course struct {
		Name  string
		Link  string
		Image string
	}

	BlogIndex struct {
		Name        string
		MainHeading string
		SubHeading  string
		Title       string
		Posts       []Post
		MetaData    MetaData
		NavLinks    []Link
	}

	Page struct {
		Index       int8
		Name        string
		MainHeading string
		SubHeading  string
		Title       string
		Posts       []Post
		MetaData    MetaData
		NavLinks    []Link
		Tutorials   []Tutorial
		Community   []Community
		Course      Course
		RecentPosts []Post
	}
)
