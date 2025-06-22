package blog

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/******************************
 ******************************
 ************ MAIN ************
 ******************************
 ******************************/

/************************
 * REPOSITORY INTERFACE *
 ************************/

type blogRepository interface {
	list(limit int, offset int) ([]Post, error)
	listByUser(userID uuid.UUID, limit int, offset int) ([]Post, error)

	create(post *Post) error
	getByID(id uuid.UUID) (*Post, error)
	update(post *Post) error
	delete(id uuid.UUID) error

	// Extended Services
	search(query string, limit int, offset int) ([]Post, error)
	listByTag(tag string, limit int, offset int) ([]Post, error)
	listRecent(limit int) ([]Post, error)
	softDelete(id uuid.UUID) error
	restore(id uuid.UUID) error
	listDeleted(limit int, offset int) ([]Post, error)
	isAuthor(postID, userID uuid.UUID) (bool, error)
	countByUser(userID uuid.UUID) (int, error)
	countTotal() (int, error)
	archive(id uuid.UUID) error
	exists(id uuid.UUID) (bool, error)
	incrementViewCount(id uuid.UUID) error
	likePost(userID, postID uuid.UUID) error
	unlikePost(userID, postID uuid.UUID) error
	listReported(limit int, offset int) ([]Post, error)
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

type repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) blogRepository {
	return &repository{db}
}

func (r *repository) create(post *Post) error {
	return r.db.Create(post).Error
}

func (r *repository) getByID(id uuid.UUID) (*Post, error) {
	var post Post
	err := r.db.Preload("Userr").First(&post, "id = ?", id).Error
	return &post, err
}

func (r *repository) listByUser(userID uuid.UUID, limit int, offset int) ([]Post, error) {
	var posts []Post
	err := r.db.Preload("Userr").
		Where("user_id = ?", userID).
		Limit(limit).Offset(offset).
		Find(&posts).Error
	return posts, err
}

func (r *repository) list(limit int, offset int) ([]Post, error) {
	var posts []Post
	err := r.db.Preload("User").
		Order("created_at ASC").
		Limit(limit).Offset(offset).
		Find(&posts).Error
	return posts, err
}

func (r *repository) update(post *Post) error {
	return r.db.Save(post).Error
}

func (r *repository) delete(id uuid.UUID) error {
	return r.db.Delete(&Post{}, "id = ?", id).Error
}

/**************************************
 **************************************
 ************ EXTENSIONS **************
 **************************************
 **************************************/

/***************************
 * SEARCH AND FILTERING   *
 ***************************/

func (r *repository) search(query string, limit int, offset int) ([]Post, error) {
	var posts []Post
	err := r.db.Preload("User").
		Where("title ILIKE ? OR content ILIKE ?", "%"+query+"%", "%"+query+"%").
		Limit(limit).Offset(offset).
		Find(&posts).Error
	return posts, err
}

func (r *repository) listByTag(tag string, limit int, offset int) ([]Post, error) {
	var posts []Post
	err := r.db.
		Joins("JOIN post_tags ON post_tags.post_id = posts.id").
		Joins("JOIN tags ON tags.id = post_tags.tag_id").
		Where("tags.name = ?", tag).
		Limit(limit).Offset(offset).
		Find(&posts).Error
	return posts, err
}

func (r *repository) listRecent(limit int) ([]Post, error) {
	var posts []Post
	err := r.db.Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Find(&posts).Error
	return posts, err
}

/***************************
 * POST STATE MANAGEMENT   *
 ***************************/

func (r *repository) softDelete(id uuid.UUID) error {
	return r.db.Model(&Post{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).
		Error
}

func (r *repository) restore(id uuid.UUID) error {
	return r.db.Unscoped().Model(&Post{}).
		Where("id = ?", id).
		Update("deleted_at", nil).
		Error
}

func (r *repository) listDeleted(limit int, offset int) ([]Post, error) {
	var posts []Post
	err := r.db.Unscoped().
		Where("deleted_at IS NOT NULL").
		Limit(limit).Offset(offset).
		Find(&posts).Error
	return posts, err
}

func (r *repository) archive(id uuid.UUID) error {
	return r.db.Model(&Post{}).
		Where("id = ?", id).
		Update("archived", true).
		Error
}

/***************************
 * METRICS AND VALIDATION  *
 ***************************/

func (r *repository) isAuthor(postID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&Post{}).
		Where("id = ? AND user_id = ?", postID, userID).
		Count(&count).Error
	return count > 0, err
}

func (r *repository) countByUser(userID uuid.UUID) (int, error) {
	var count int64
	err := r.db.Model(&Post{}).
		Where("user_id = ?", userID).
		Count(&count).Error
	return int(count), err
}

func (r *repository) countTotal() (int, error) {
	var count int64
	err := r.db.Model(&Post{}).
		Count(&count).Error
	return int(count), err
}

func (r *repository) exists(id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&Post{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

/**************************
 * INTERACTION FEATURES   *
 **************************/

func (r *repository) incrementViewCount(id uuid.UUID) error {
	return r.db.Model(&Post{}).
		Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).
		Error
}

func (r *repository) likePost(userID, postID uuid.UUID) error {
	like := PostLike{UserID: userID, PostID: postID}
	return r.db.Create(&like).Error
}

func (r *repository) unlikePost(userID, postID uuid.UUID) error {
	return r.db.
		Where("user_id = ? AND post_id = ?", userID, postID).
		Delete(&PostLike{}).Error
}

func (r *repository) listReported(limit int, offset int) ([]Post, error) {
	var posts []Post
	err := r.db.
		Joins("JOIN post_reports ON post_reports.post_id = posts.id").
		Group("posts.id").
		Limit(limit).Offset(offset).
		Find(&posts).Error
	return posts, err
}
