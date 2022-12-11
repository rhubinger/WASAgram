package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// Login
	rt.router.POST("/login", rt.Login)

	// Users
	rt.router.GET("/searchresult", rt.SearchUser)
	rt.router.GET("/users/:uid", rt.GetUser)
	rt.router.GET("/users/:uid/posts", rt.GetPosts)
	rt.router.GET("/users/:uid/posts/count", rt.GetPostCount)
	rt.router.GET("/users/:uid/stream", rt.GetStream)
	rt.router.PUT("/users/:uid/username", rt.ChangeUsername)

	// Followers
	rt.router.GET("/users/:uid/followed", rt.GetFollowed)
	rt.router.GET("/users/:uid/followed/count", rt.GetFollowedCount)
	rt.router.GET("/users/:uid/followers", rt.GetFollowers)
	rt.router.GET("/users/:uid/followers/count", rt.GetFollowerCount)
	rt.router.POST("/users/:uid/followers/:fid", rt.Follow)
	rt.router.DELETE("/users/:uid/followers/:fid", rt.Unfollow)

	// Banning
	rt.router.GET("/users/:uid/banned", rt.GetBanned)
	rt.router.GET("/users/:uid/banned/count", rt.GetBannedCount)
	rt.router.POST("/users/:uid/banned/:bid", rt.Ban)
	rt.router.DELETE("/users/:uid/banned/:bid", rt.Unban)

	// Posts
	rt.router.PUT("/posts", rt.CreatePost)
	rt.router.GET("/posts/:pid", rt.GetPost)
	rt.router.DELETE("/posts/:pid", rt.DeletePost)

	// Likes
	rt.router.GET("/posts/:pid/likes", rt.GetLikes)
	rt.router.GET("/posts/:pid/likes/count", rt.GetLikeCount)
	rt.router.POST("/posts/:pid/likes/:lid", rt.LikePost)
	rt.router.DELETE("/posts/:pid/likes/:lid", rt.DeleteLike)

	// Comments
	rt.router.PUT("/posts/:pid/comments", rt.CreateComment)
	rt.router.GET("/posts/:pid/comments", rt.GetComments)
	rt.router.GET("/posts/:pid/comments/count", rt.GetCommentCount)
	rt.router.DELETE("/posts/:pid/comments/:cid", rt.DeleteComment)

	// Picture
	rt.router.GET("/pictures/:pid", rt.GetPicture)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
