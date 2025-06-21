package guide

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/**********************
 **********************
 ******** MAIN ********
 **********************
 **********************/

/*****************************
 * MAIN REPOSITORY INTERFACE *
 *****************************/

/* GAME GUIDE */

type gameGuideRepository interface {
	list(limit, offset int) ([]GameGuide, error)
	listByUser(userID uuid.UUID, limit, offset int) ([]GameGuide, error)
	countByUser(userID uuid.UUID) (int64, error)

	create(guide *GameGuide) error
	getByID(id uuid.UUID) (*GameGuide, error)
	getBySlug(slug string) (*GameGuide, error)
	update(guide *GameGuide) error
	delete(id uuid.UUID) error
}

type gameGuideRepoImpl struct {
	db *gorm.DB
}

func newGameGuideRepo(db *gorm.DB) gameGuideRepository {
	return &gameGuideRepoImpl{db: db}
}

func (r *gameGuideRepoImpl) create(guide *GameGuide) error {
	return r.db.Create(guide).Error
}

func (r *gameGuideRepoImpl) getByID(id uuid.UUID) (*GameGuide, error) {
	var guide GameGuide
	err := r.db.Preload("Tags").First(&guide, "id = ?", id).Error
	return &guide, err
}

func (r *gameGuideRepoImpl) getBySlug(slug string) (*GameGuide, error) {
	var guide GameGuide
	err := r.db.Preload("Tags").First(&guide, "slug = ?", slug).Error
	return &guide, err
}

func (r *gameGuideRepoImpl) update(guide *GameGuide) error {
	return r.db.Save(guide).Error
}

func (r *gameGuideRepoImpl) delete(id uuid.UUID) error {
	return r.db.Delete(&GameGuide{}, "id = ?", id).Error
}

func (r *gameGuideRepoImpl) list(limit, offset int) ([]GameGuide, error) {
	var guides []GameGuide
	err := r.db.Preload("Tags").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&guides).Error
	return guides, err
}

func (r *gameGuideRepoImpl) listByUser(authorID uuid.UUID, limit, offset int) ([]GameGuide, error) {
	var guides []GameGuide
	err := r.db.Preload("Tags").
		Where("author_id = ?", authorID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&guides).Error
	return guides, err
}

func (r *gameGuideRepoImpl) countByUser(authorID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&GameGuide{}).
		Where("author_id = ?", authorID).
		Count(&count).Error
	return count, err
}

/* GAME GUIDE */

type gameGuideTagRepository interface {
	create(tag *GameGuideTag) error
	getByName(name string) (*GameGuideTag, error)
	listAll() ([]GameGuideTag, error)
}

type gameGuideTagRepo struct {
	db *gorm.DB
}

func newGameGuideTagRepo(db *gorm.DB) gameGuideTagRepository {
	return &gameGuideTagRepo{db: db}
}

func (r *gameGuideTagRepo) create(tag *GameGuideTag) error {
	return r.db.Create(tag).Error
}

func (r *gameGuideTagRepo) getByName(name string) (*GameGuideTag, error) {
	var tag GameGuideTag
	err := r.db.First(&tag, "name = ?", name).Error
	return &tag, err
}

func (r *gameGuideTagRepo) listAll() ([]GameGuideTag, error) {
	var tags []GameGuideTag
	err := r.db.Order("name ASC").Find(&tags).Error
	return tags, err
}

/******************************
 ******************************
 ******** INTERACTIONS ********
 ******************************
 ******************************/

/***************************
 * INTERACTIONS REPOSITORY *
 ***************************/

/* LIKE */

type gameGuideLikeRepository interface {
	create(like *GameGuideLike) error
	delete(guideID, userID uuid.UUID) error
	exists(guideID, userID uuid.UUID) (bool, error)
	listByUser(userID uuid.UUID) ([]GameGuideLike, error)
}

type gameGuideLikeRepo struct {
	db *gorm.DB
}

func newGameGuideLikeRepo(db *gorm.DB) gameGuideLikeRepository {
	return &gameGuideLikeRepo{db}
}

func (r *gameGuideLikeRepo) create(like *GameGuideLike) error {
	return r.db.Create(like).Error
}

func (r *gameGuideLikeRepo) delete(guideID, userID uuid.UUID) error {
	return r.db.
		Where("guide_id = ? AND user_id = ?", guideID, userID).
		Delete(&GameGuideLike{}).Error
}

func (r *gameGuideLikeRepo) exists(guideID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.
		Model(&GameGuideLike{}).
		Where("guide_id = ? AND user_id = ?", guideID, userID).
		Count(&count).Error
	return count > 0, err
}

func (r *gameGuideLikeRepo) listByUser(userID uuid.UUID) ([]GameGuideLike, error) {
	var likes []GameGuideLike
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&likes).Error
	return likes, err
}

/* VIEW */

type gameGuideViewRepository interface {
	create(view *GameGuideView) error
	exists(guideID, userID uuid.UUID) (bool, error)
}

type gameGuideViewRepo struct {
	db *gorm.DB
}

func newGameGuideViewRepo(db *gorm.DB) gameGuideViewRepository {
	return &gameGuideViewRepo{db}
}

func (r *gameGuideViewRepo) create(view *GameGuideView) error {
	return r.db.Create(view).Error
}

func (r *gameGuideViewRepo) exists(guideID, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.
		Model(&GameGuideView{}).
		Where("guide_id = ? AND user_id = ?", guideID, userID).
		Count(&count).Error
	return count > 0, err
}
