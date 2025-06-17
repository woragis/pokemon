package snapdex

import (
	"pokemon/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterRoutes(app fiber.Router) {
	snaps := app.Group("/snaps")

	// Snap routes
	snaps.Get("/", h.listSnaps)
	snaps.Get("/:id", h.getSnap)
	snaps.Get("/user/:user_id", h.listSnapsByUser)
	snaps.Get("/user/:user_id/count", h.countSnapsByUser)

	snaps.Use(middleware.AuthRequired())
	snaps.Post("/", h.createSnap)
	snaps.Put("/:id", h.updateSnap)
	snaps.Delete("/:id", h.deleteSnap)

	// Snap comment routes
	snapComments := snaps.Group("/comments")
	snapComments.Post("/", h.createComment)
	snapComments.Get("/me", h.listCommentsByUser)
	snapComments.Get("/me/count", h.countCommentsByUser)
	snapComments.Put("/:id/status/:status", h.updateCommentStatus)
	snapComments.Delete("/:id", h.deleteComment)
	snapComments.Get("/exists/:id", h.commentExists)

	// Snap like routes
	snapLikes := snaps.Group("/likes")
	snapLikes.Post("/:snap_id", h.likeSnap)
	snapLikes.Delete("/:snap_id", h.unlikeSnap)
	snapLikes.Delete("/by-snap/:snap_id", h.deleteAllLikesBySnap)
	snapLikes.Get("/me", h.listUserLikes)
	snapLikes.Get("/exists/:snap_id", h.isLikedByUser)

	// Snap comment like routes
	commentLikes := snaps.Group("/comment-likes")
	commentLikes.Post("/:comment_id", h.likeComment)
	commentLikes.Delete("/:comment_id", h.unlikeComment)
	commentLikes.Get("/by-comment/:comment_id", h.listLikesByComment)
	commentLikes.Get("/me", h.listCommentsUserLikes)
	commentLikes.Get("/exists/:comment_id", h.isCommentLikedByUser)
}
