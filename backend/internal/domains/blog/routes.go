package blog

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	blogGroup := app.Group("/blog")

	/**************************
	 * PUBLIC ROUTES (READ)   *
	 **************************/

	blogGroup.Get("/", h.listPosts)
	blogGroup.Get("/user/:user_id", h.listPostsByUser)
	blogGroup.Get("/search", h.searchPosts)
	blogGroup.Get("/tag/:tag", h.listPostsByTag)
	blogGroup.Get("/recent", h.listRecentPosts)
	blogGroup.Get("/deleted", h.listDeletedPosts)
	blogGroup.Get("/:id", h.getPost)

	/***************************
	 * INTERACTIONS (PUBLIC)   *
	 ***************************/

	blogGroup.Post("/:id/view", h.incrementView)
	blogGroup.Post("/:id/like", h.likePost)
	blogGroup.Delete("/:id/unlike", h.unlikePost)

	/****************************
	 * AUTH REQUIRED FROM HERE  *
	 ****************************/

	blogGroup.Use(middleware.AuthRequired())

	/**************************
	 * BLOG CRUD              *
	 **************************/

	blogGroup.Post("/", h.createPost)
	blogGroup.Put("/:id", h.updatePost)
	blogGroup.Delete("/:id", h.deletePost)

	/**************************
	 * POST STATE MANAGEMENT  *
	 **************************/

	blogGroup.Patch("/:id/soft-delete", h.softDeletePost)
	blogGroup.Patch("/:id/restore", h.restorePost)

	/***************************
	 * FUTURE ROUTES PLACEHOLDER
	 ***************************/

	// blogGroup.Get("/:id/comments", h.getComments)
	// blogGroup.Get("/:id/comments/count", h.getCommentsCount)
	// blogGroup.Get("/:id/views/count", h.getBlogViewCount)
	// blogGroup.Get("/:id/likes/count", h.getBlogLikeCount)
}
