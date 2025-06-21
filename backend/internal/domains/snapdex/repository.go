package snapdex

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/**********************
 **********************
 ******** MAIN ********
 **********************
 **********************/

/************************
 * REPOSITORY INTERFACE *
 ************************/

type snapRepository interface {
	list(limit, offset int) ([]Snap, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]Snap, error)
	countByUser(userID uuid.UUID) (int64, error)

	create(snap *Snap) error
	getById(id uuid.UUID) (*Snap, error)
	update(snap *Snap) error
	delete(id uuid.UUID) error
}

type snapRepoImpl struct {
	db *gorm.DB
}

func newRepo(db *gorm.DB) snapRepository {
	return &snapRepoImpl{db: db}
}

/*****************************
 * REPOSITORY IMPLEMENTATION *
 *****************************/

func (r *snapRepoImpl) create(snap *Snap) error {
	return r.db.Create(snap).Error
}

func (r *snapRepoImpl) update(snap *Snap) error {
	return r.db.Save(snap).Error
}

func (r *snapRepoImpl) delete(id uuid.UUID) error {
	return r.db.Delete(&Snap{}, "id = ?", id).Error
}

func (r *snapRepoImpl) getById(id uuid.UUID) (*Snap, error) {
	var snap Snap
	err := r.db.Where("id = ?", id).First(&snap).Error // 🔧 fix: chain Where before First
	return &snap, err
}

func (r *snapRepoImpl) list(limit, offset int) ([]Snap, error) {
	var snaps []Snap
	err := r.db.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&snaps).
		Error // 🔧 fix: Order, Limit, Offset must come before Find
	return snaps, err
}

func (r *snapRepoImpl) listByUser(userID uuid.UUID, limit, offset int) ([]Snap, error) {
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

func (r *snapRepoImpl) countByUser(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.
		Model(&Snap{}). // 🔧 fix: use Model, not Find
		Where("user_id = ?", userID).
		Count(&count).
		Error
	return count, err
}

/******************************
 ******************************
 ******** INTERACTIONS ********
 ******************************
 ******************************/

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

type snapCommentRepoImpl struct {
	db *gorm.DB
}

type snapLikeRepoImpl struct {
	db *gorm.DB
}

type snapCommentLikeRepoImpl struct {
	db *gorm.DB
}

/*******************************
 * CONSTRUCTOR FUNCTIONS *
 *******************************/

func newSnapCommentRepo(db *gorm.DB) snapCommentRepository {
	return &snapCommentRepoImpl{db: db}
}

func newSnapLikeRepo(db *gorm.DB) snapLikeRepository {
	return &snapLikeRepoImpl{db: db}
}

func newSnapCommentLikeRepo(db *gorm.DB) snapCommentLikeRepository {
	return &snapCommentLikeRepoImpl{db: db}
}

/********************************
 * IMPLEMENTATIONS: SnapComment *
 ********************************/

func (r *snapCommentRepoImpl) create(comment *SnapComment) error {
	return r.db.Create(comment).Error
}

func (r *snapCommentRepoImpl) getByID(id uuid.UUID) (*SnapComment, error) {
	var comment SnapComment
	err := r.db.First(&comment, "id = ?", id).Error
	return &comment, err
}

func (r *snapCommentRepoImpl) listByUser(userID uuid.UUID, limit, offset int) ([]SnapComment, error) {
	var comments []SnapComment
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Order("created_at DESC").Find(&comments).Error
	return comments, err
}

func (r *snapCommentRepoImpl) countByUser(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&SnapComment{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *snapCommentRepoImpl) updateStatus(id uuid.UUID, status string) error {
	return r.db.Model(&SnapComment{}).Where("id = ?", id).Update("status", status).Error
}

func (r *snapCommentRepoImpl) delete(id uuid.UUID) error {
	return r.db.Delete(&SnapComment{}, "id = ?", id).Error
}

func (r *snapCommentRepoImpl) existsByID(id uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.Model(&SnapComment{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Error
	return exists, err
}

/******************************
 * IMPLEMENTATIONS: SnapLike *
 ******************************/

func (r *snapLikeRepoImpl) create(like *SnapLike) error {
	return r.db.Create(like).Error
}

func (r *snapLikeRepoImpl) delete(snapID, userID uuid.UUID) error {
	return r.db.Where("snap_id = ? AND user_id = ?", snapID, userID).Delete(&SnapLike{}).Error
}

func (r *snapLikeRepoImpl) deleteAllBySnap(snapID uuid.UUID) error {
	return r.db.Where("snap_id = ?", snapID).Delete(&SnapLike{}).Error
}

func (r *snapLikeRepoImpl) listUserLikes(userID uuid.UUID) ([]SnapLike, error) {
	var likes []SnapLike
	err := r.db.Where("user_id = ?", userID).Find(&likes).Error
	return likes, err
}

func (r *snapLikeRepoImpl) exists(snapID, userID uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.Model(&SnapLike{}).Select("count(*) > 0").Where("snap_id = ? AND user_id = ?", snapID, userID).Find(&exists).Error
	return exists, err
}

/************************************
 * IMPLEMENTATIONS: SnapCommentLike *
 ************************************/

func (r *snapCommentLikeRepoImpl) create(like *SnapCommentLike) error {
	return r.db.Create(like).Error
}

func (r *snapCommentLikeRepoImpl) delete(commentID, userID uuid.UUID) error {
	return r.db.Where("comment_id = ? AND user_id = ?", commentID, userID).Delete(&SnapCommentLike{}).Error
}

func (r *snapCommentLikeRepoImpl) listByComment(commentID uuid.UUID) ([]SnapCommentLike, error) {
	var likes []SnapCommentLike
	err := r.db.Where("comment_id = ?", commentID).Find(&likes).Error
	return likes, err
}

func (r *snapCommentLikeRepoImpl) listUserCommentLikes(userID uuid.UUID) ([]SnapCommentLike, error) {
	var likes []SnapCommentLike
	err := r.db.Where("user_id = ?", userID).Find(&likes).Error
	return likes, err
}

func (r *snapCommentLikeRepoImpl) exists(commentID, userID uuid.UUID) (bool, error) {
	var exists bool
	err := r.db.Model(&SnapCommentLike{}).Select("count(*) > 0").Where("comment_id = ? AND user_id = ?", commentID, userID).Find(&exists).Error
	return exists, err
}
