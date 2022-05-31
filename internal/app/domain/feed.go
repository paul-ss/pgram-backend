package domain

// FeedGet is api model
type FeedGet struct {
	Limit int64  `json:"Limit"`
	Since int64  `json:"Since"`
	Sort  string `json:"Sort"`
	Desc  bool   `json:"Desc"`
}

type FeedGetRepo struct {
	Limit int64 `json:"Limit"`
	Since int64 `json:"Since"`
	Desc  bool  `json:"Desc"`
}

type FeedResponse struct {
	PostList []Post
}
