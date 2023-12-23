package repository

import (
	"github.com/caoquy2000/meeting-app/infrastructure"
	"github.com/caoquy2000/meeting-app/models"
)

type PostRepository struct {
	db infrastructure.Database
}

func NewPostRepository(db infrastructure.Database) PostRepository {
	return PostRepository{
		db: db,
	}
}

// Save a post
func (p PostRepository) Save(post models.Post) error {
	return p.db.DB.Create(&post).Error
}

// Find all post
func (p PostRepository) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0

	queryBuilder := p.db.DB.Order("created_at desc").Model(&models.Post{})

	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			p.db.DB.Where("post.title LIKE ? ", queryKeyword))
	}

	err := queryBuilder.
		Where(post).
		Find(&posts).
		Count(&totalRows).Error

	return &posts, totalRows, err
}

// Update post
func (p PostRepository) Update(post models.Post) error {
	return p.db.DB.Save(&post).Error
}

// Find a post
func (p PostRepository) Find(post models.Post) (models.Post, error) {
	var posts models.Post
	err := p.db.DB.
		Debug().
		Model(&models.Post{}).
		Where(&post).
		Take(&posts).Error
	return post, err
}

// Delete post
func (p PostRepository) Delete(post models.Post) error {
	return p.db.DB.Delete(&post).Error
}
