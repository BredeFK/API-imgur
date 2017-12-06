package main

// ImagePost struct for posting image
type ImagePost struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

// CommentPost struct for posting comment
type CommentPost struct {
	ImageID  string `json:"imageID"`
	Comment  string `json:"comment"`
	ParentID string `json:"parentID"`
}
