package snapdex

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/************************
 * REPOSITORY INTERFACE *
 ************************/

type snapRepository interface {
	create(snap *Snap) error
	getById(id uuid.UUID) (*Snap, error)
	update(snap *Snap) error
	delete(id uuid.UUID) error

	list(limit, offset int) ([]Snap, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]Snap, error)
	countByUser(userID uuid.UUID) (int64, error)
}

type snapRepo struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) snapRepository {
	return &snapRepo{db: db}
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

func (r *snapRepo) create(snap *Snap) error {
	return r.db.Create(snap).Error
}

func (r *snapRepo) update(snap *Snap) error {
	return r.db.Save(snap).Error
}

func (r *snapRepo) delete(id uuid.UUID) error {
	return r.db.Delete(&Snap{}, "id = ?", id).Error
}

func (r *snapRepo) getById(id uuid.UUID) (*Snap, error) {
	var snap Snap
	err := r.db.Where("id = ?", id).First(&snap).Error // 🔧 fix: chain Where before First
	return &snap, err
}

func (r *snapRepo) list(limit, offset int) ([]Snap, error) {
	var snaps []Snap
	err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&snaps).
		Error // 🔧 fix: Order, Limit, Offset must come before Find
	return snaps, err
}

func (r *snapRepo) listByUser(userID uuid.UUID, limit, offset int) ([]Snap, error) {
	var snaps []Snap
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&snaps).
		Error
	return snaps, err
}

func (r *snapRepo) countByUser(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.
		Model(&Snap{}). // 🔧 fix: use Model, not Find
		Where("user_id = ?", userID).
		Count(&count).
		Error
	return count, err
}

/**************************
 * REPOSITORY INTERFACES *
 **************************/

type snapCommentRepository interface {
	create(comment *SnapComment) error
	getByID(id uuid.UUID) (*SnapComment, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]SnapComment, error)
	countByUser(userID uuid.UUID) (int64, error)
	updateStatus(id uuid.UUID, status string) error
	delete(id uuid.UUID) error
	existsByID(id uuid.UUID) (bool, error)
}

type snapLikeRepository interface {
	create(like *SnapLike) error
	delete(snapID, userID uuid.UUID) error
	deleteAllBySnap(snapID uuid.UUID) error
	listUserLikes(userID uuid.UUID) ([]SnapLike, error)
	exists(snapID, userID uuid.UUID) (bool, error)
}

type snapCommentLikeRepository interface {
	create(like *SnapCommentLike) error
	delete(commentID, userID uuid.UUID) error
	listByComment(commentID uuid.UUID) ([]SnapCommentLike, error)
	listUserCommentLikes(userID uuid.UUID) ([]SnapCommentLike, error)
	exists(commentID, userID uuid.UUID) (bool, error)
}

/***************************
 * REPOSITORY STRUCTURES *
 ***************************/

type snapCommentRepo struct {
	db *gorm.DB
}

type snapLikeRepo struct {
	db *gorm.DB
}

type snapCommentLikeRepo struct {
	db *gorm.DB
}

/*******************************
 * CONSTRUCTOR FUNCTIONS *
 *******************************/

func newSnapCommentRepo(db *gorm.DB) snapCommentRepository {
	return &snapCommentRepo{db: db}
}

func newSnapLikeRepo(db *gorm.DB) snapLikeRepository {
	return &snapLikeRepo{db: db}
}

func newSnapCommentLikeRepo(db *gorm.DB) snapCommentLikeRepository {
	return &snapCommentLikeRepo{db: db}
}

/********************************
 * IMPLEMENTATIONS: SnapComment *
 ********************************/

func (r *snapCommentRepo) create(comment *SnapComment) error {
	return r.db.Create(comment).Error
}

func (r *snapCommentRepo) getByID(id uuid.UUID) (*SnapComment, error) {
	var comment SnapComment
	err := r.db.First(&comment, "id = ?", id).Error
	return &comment, err
}

func (r *snapCommentRepo) listByUser(userID uuid.UUID, limit, offset int) ([]SnapComment, error) {
	var comments []SnapComment
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Order("created_at DESC").Find(&comments).Error
	return comments, err
}

func (r *snapCommentRepo) countByUser(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&SnapComment{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *snapCommentRepo) updateStatus(id uuid.UUID, status string) error {
	return r.db.Model(&SnapComment{}).Where("id = ?", id).Update("status", status).Error
}

func (r *snapCommentRepo) delete(id uuid.UUID) error {
	return r.db.Delete(&SnapComment{}, "id = ?", id).Error
}

func (r *snapCommentRepo) existsByID(id uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.Model(&SnapComment{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
	return exists, err
}

/******************************
 * IMPLEMENTATIONS: SnapLike *
 ******************************/

func (r *snapLikeRepo) create(like *SnapLike) error {
	return r.db.Create(like).Error
}

func (r *snapLikeRepo) delete(snapID, userID uuid.UUID) error {
	return r.db.Where("snap_id = ? AND user_id = ?", snapID, userID).Delete(&SnapLike{}).Error
}

func (r *snapLikeRepo) deleteAllBySnap(snapID uuid.UUID) error {
	return r.db.Where("snap_id = ?", snapID).Delete(&SnapLike{}).Error
}

func (r *snapLikeRepo) listUserLikes(userID uuid.UUID) ([]SnapLike, error) {
	var likes []SnapLike
	err := r.db.Where("user_id = ?", userID).Find(&likes).Error
	return likes, err
}

func (r *snapLikeRepo) exists(snapID, userID uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.Model(&SnapLike{}).Select("count(*) > 0").Where("snap_id = ? AND user_id = ?", snapID, userID).Find(&exists).Error
	return exists, err
}

/************************************
 * IMPLEMENTATIONS: SnapCommentLike *
 ************************************/

func (r *snapCommentLikeRepo) create(like *SnapCommentLike) error {
	return r.db.Create(like).Error
}

func (r *snapCommentLikeRepo) delete(commentID, userID uuid.UUID) error {
	return r.db.Where("comment_id = ? AND user_id = ?", commentID, userID).Delete(&SnapCommentLike{}).Error
}

func (r *snapCommentLikeRepo) listByComment(commentID uuid.UUID) ([]SnapCommentLike, error) {
	var likes []SnapCommentLike
	err := r.db.Where("comment_id = ?", commentID).Find(&likes).Error
	return likes, err
}

func (r *snapCommentLikeRepo) listUserCommentLikes(userID uuid.UUID) ([]SnapCommentLike, error) {
	var likes []SnapCommentLike
	err := r.db.Where("user_id = ?", userID).Find(&likes).Error
	return likes, err
}

func (r *snapCommentLikeRepo) exists(commentID, userID uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.Model(&SnapCommentLike{}).Select("count(*) > 0").Where("comment_id = ? AND user_id = ?", commentID, userID).Find(&exists).Error
	return exists, err
}
