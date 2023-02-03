package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Rota{
	{
		URI:         "/posts",
		Method:      http.MethodPost,
		Function:    controllers.CreatePost,
		RequireAuth: true,
	},
	{
		URI:         "/posts",
		Method:      http.MethodGet,
		Function:    controllers.GetPosts,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{id}",
		Method:      http.MethodGet,
		Function:    controllers.GetPost,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{id}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePost,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{id}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePost,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}/posts",
		Method:      http.MethodGet,
		Function:    controllers.GetPostsByUser,
		RequireAuth: true,
	},
	{
		URI:         "/posts/{id}/like",
		Method:      http.MethodPost,
		Function:    controllers.LikePost,
		RequireAuth: true,
	},
}
