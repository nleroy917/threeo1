package main

type shortenUri struct {
	Url string `form:"url"`
}

type redirectUri struct {
	Id string `uri:"id" binding:"required"`
}
