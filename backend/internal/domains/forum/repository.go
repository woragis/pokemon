package forum

import (
	"errors"

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

type topicRepository interface {
	create(topic *Topic) error
	getByID(id uuid.UUID) (*Topic, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]Topic, error)
	list(limit, offset int) ([]Topic, error)
	update(topic *Topic) error
	delete(id uuid.UUID) error
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

type repository struct {
	db *gorm.DB
}

func newTopicRepository(db *gorm.DB) topicRepository {
	return &repository{db}
}

func (r *repository) create(topic *Topic) error {
	return r.db.Create(topic).Error
}

func (r *repository) getByID(id uuid.UUID) (*Topic, error) {
	var topic Topic
	err := r.db.Preload("User").First(&topic, "id = ?", id).Error
	return &topic, err
}

func (r *repository) listByUser(userID uuid.UUID, limit, offset int) ([]Topic, error) {
	var topics []Topic
	err := r.db.
		Preload("User").
		Order("created_at DESC").
		Where("user_id= ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&topics).Error
	return topics, err
}

func (r *repository) list(limit, offset int) ([]Topic, error) {
	var topics []Topic
	err := r.db.
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&topics).Error
	return topics, err
}

func (r *repository) update(topic *Topic) error {
	return r.db.Save(topic).Error
}

func (r *repository) delete(id uuid.UUID) error {
	return r.db.Delete(&Topic{}, "id = ?", id).Error
}

/**************************************
 **************************************
 ************ INTERACTIONS ************
 **************************************
 **************************************/

/********************************
 * COMMENT REPOSITORY INTERFACE *
 ********************************/

type topicCommentRepository interface {
	create(c *TopicComment) error
	getByID(id uuid.UUID) (*TopicComment, error)
	update(c *TopicComment) error
	delete(id uuid.UUID) error
	listByTopic(topicID uuid.UUID) ([]TopicComment, error)
	count(topicID uuid.UUID) (int64, error)
}

type topicCommentRepositoryImpl struct {
	db *gorm.DB
}

func newTopicCommentRepository(db *gorm.DB) topicCommentRepository {
	return &topicCommentRepositoryImpl{db: db}
}

/*************************************
 * COMMENT REPOSITORY IMPLEMENTATION *
 *************************************/

func (r *topicCommentRepositoryImpl) create(c *TopicComment) error {
	return r.db.Create(c).Error
}

func (r *topicCommentRepositoryImpl) getByID(id uuid.UUID) (*TopicComment, error) {
	var comment TopicComment
	if err := r.db.Preload("User").Preload("Replies").First(&comment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *topicCommentRepositoryImpl) update(c *TopicComment) error {
	return r.db.Save(c).Error
}

func (r *topicCommentRepositoryImpl) delete(id uuid.UUID) error {
	return r.db.Delete(&TopicComment{}, "id = ?", id).Error
}

func (r *topicCommentRepositoryImpl) listByTopic(topicID uuid.UUID) ([]TopicComment, error) {
	var comments []TopicComment
	err := r.db.
		Preload("User").
		Preload("Replies").
		Where("topic_id = ? AND parent_id IS NULL", topicID).
		Order("created_at ASC").
		Find(&comments).Error
	return comments, err
}

func (r *topicCommentRepositoryImpl) count(topicID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.
		Where("topic_id = ?", topicID).
		Count(&count).Error
	return count, err
}

/*********************************
 * CATEGORY REPOSITORY INTERFACE *
 *********************************/

type topicCategoryRepository interface {
	create(c *TopicCategory) error
	getByID(id uuid.UUID) (*TopicCategory, error)
	update(c *TopicCategory) error
	delete(id uuid.UUID) error
	list(limit, offset int) ([]TopicCategory, error)
}

type topicCategoryRepositoryImpl struct {
	db *gorm.DB
}

func newTopicCategoryRepository(db *gorm.DB) topicCategoryRepository {
	return &topicCategoryRepositoryImpl{db: db}
}

/**************************************
 * CATEGORY REPOSITORY IMPLEMENTATION *
 **************************************/

func (r *topicCategoryRepositoryImpl) create(c *TopicCategory) error {
	return r.db.Create(c).Error
}

func (r *topicCategoryRepositoryImpl) getByID(id uuid.UUID) (*TopicCategory, error) {
	var category TopicCategory
	err := r.db.First(&category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *topicCategoryRepositoryImpl) update(c *TopicCategory) error {
	return r.db.Save(c).Error
}

func (r *topicCategoryRepositoryImpl) delete(id uuid.UUID) error {
	return r.db.Delete(&TopicCategory{}, "id = ?", id).Error
}

func (r *topicCategoryRepositoryImpl) list(limit, offset int) ([]TopicCategory, error) {
	var categories []TopicCategory
	err := r.db.
		Order("name ASC").
		Limit(limit).
		Offset(offset).
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

/*****************************
 * LIKE REPOSITORY INTERFACE *
 *****************************/

type topicLikeRepository interface {
	create(like *TopicLike) error
	get(topicID, userID uuid.UUID) (*TopicLike, error)
	delete(topicID, userID uuid.UUID) error
	count(topicID uuid.UUID, likeValue bool) (int64, error)
}

type topicLikeRepo struct {
	db *gorm.DB
}

func newTopicLikeRepository(db *gorm.DB) topicLikeRepository {
	return &topicLikeRepo{db: db}
}

/**********************************
 * LIKE REPOSITORY IMPLEMENTATION *
 **********************************/

func (r *topicLikeRepo) count(topicID uuid.UUID, likeValue bool) (int64, error) {
	var count int64
	err := r.db.Model(&TopicLike{}).
		Where("topic_id = ? AND like = ?", topicID, likeValue).
		Count(&count).Error
	return count, err
}

func (r *topicLikeRepo) create(like *TopicLike) error {
	// Delete any existing like for the user-topic pair (if exists)
	if err := r.delete(like.TopicID, like.UserID); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Create the new like
	return r.db.Create(like).Error
}

func (r *topicLikeRepo) get(topicID, userID uuid.UUID) (*TopicLike, error) {
	var like TopicLike
	err := r.db.Where("topic_id = ? AND user_id = ?", topicID, userID).First(&like).Error
	if err != nil {
		return nil, err
	}
	return &like, nil
}

func (r *topicLikeRepo) delete(topicID, userID uuid.UUID) error {
	return r.db.Where("topic_id = ? AND user_id = ?", topicID, userID).Delete(&TopicLike{}).Error
}

/*****************************
 * VIEW REPOSITORY INTERFACE *
 *****************************/

type topicViewRepository interface {
	create(view *TopicView) error
	listByUser(userID uuid.UUID, limit, offset int) ([]TopicView, error)
}

type topicViewRepoImpl struct {
	db *gorm.DB
}

func newTopicViewRepository(db *gorm.DB) topicViewRepository {
	return &topicViewRepoImpl{db: db}
}

/**********************************
 * VIEW REPOSITORY IMPLEMENTATION *
 **********************************/

func (r *topicViewRepoImpl) create(view *TopicView) error {
	return r.db.Create(view).Error
}

func (r *topicViewRepoImpl) listByUser(userID uuid.UUID, limit, offset int) ([]TopicView, error) {
	var views []TopicView
	err := r.db.
		Preload("Topic").
		Preload("User").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&views).Error

	return views, err
}

/*************************************
 * COMMENT LIKE REPOSITORY INTERFACE *
 *************************************/

type commentLikeRepository interface {
	create(like *TopicCommentLike) error
	update(like *TopicCommentLike) error
	get(commentID, userID uuid.UUID) (*TopicCommentLike, error)
	delete(commentID, userID uuid.UUID) error
	count(commentID uuid.UUID) (likes, dislikes int64, err error)
}

type commentLikeRepo struct {
	db *gorm.DB
}

// Constructor
func newCommentLikeRepository(db *gorm.DB) commentLikeRepository {
	return &commentLikeRepo{db: db}
}

/******************************************
 * COMMENT LIKE REPOSITORY IMPLEMENTATION *
 ******************************************/

// Create a new CommentLike
func (r *commentLikeRepo) create(like *TopicCommentLike) error {
	return r.db.Create(like).Error
}

// Update an existing CommentLike
func (r *commentLikeRepo) update(like *TopicCommentLike) error {
	return r.db.Save(like).Error
}

// Get retrieves a CommentLike by commentID and userID
func (r *commentLikeRepo) get(commentID, userID uuid.UUID) (*TopicCommentLike, error) {
	var like TopicCommentLike
	err := r.db.
		Where("comment_id = ? AND user_id = ?", commentID, userID).
		First(&like).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &like, err
}

// Delete removes a CommentLike by commentID and userID
func (r *commentLikeRepo) delete(commentID, userID uuid.UUID) error {
	return r.db.
		Where("comment_id = ? AND user_id = ?", commentID, userID).
		Delete(&TopicCommentLike{}).Error
}

func (r *commentLikeRepo) count(commentID uuid.UUID) (int64, int64, error) {
	var likes int64
	var dislikes int64

	err := r.db.Model(&TopicCommentLike{}).
		Where("comment_id = ? AND like = TRUE", commentID).
		Count(&likes).Error
	if err != nil {
		return 0, 0, err
	}

	err = r.db.Model(&TopicCommentLike{}).
		Where("comment_id = ? AND like = FALSE", commentID).
		Count(&dislikes).Error
	if err != nil {
		return 0, 0, err
	}

	return likes, dislikes, nil
}
