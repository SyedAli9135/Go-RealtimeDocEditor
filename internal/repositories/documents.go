package repositories

import (
	"realtime-doc-editor-backend/internal/models"

	"gorm.io/gorm"
)

// DocumentRepository provides CRUD operations for documents
type DocumentRepository struct {
	DB *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{DB: db}
}

// CreateDocument creates a new document in the database
func (repo *DocumentRepository) CreateDocument(title, content string) (*models.Document, error) {
	doc := &models.Document{Title: title, Content: content}
	if err := repo.DB.Create(doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

// GetDocumentByID fetches a document by its ID
func (repo *DocumentRepository) GetDocumentByID(id uint) (*models.Document, error) {
	var doc models.Document
	if err := repo.DB.First(&doc, id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

// GetAllDocuments retrieves all documents
func (repo *DocumentRepository) GetAllDocuments() ([]models.Document, error) {
	var docs []models.Document
	if err := repo.DB.Find(&docs).Error; err != nil {
		return nil, err
	}
	return docs, nil
}

// UpdateDocument updates an existing document's content
func (repo *DocumentRepository) UpdateDocument(id uint, title, content string) (*models.Document, error) {
	var doc models.Document
	if err := repo.DB.First(&doc, id).Error; err != nil {
		return nil, err
	}
	doc.Title = title
	doc.Content = content
	if err := repo.DB.Save(&doc).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

// DeleteDocument deletes a document by its ID
func (repo *DocumentRepository) DeleteDocument(id uint) error {
	if err := repo.DB.Delete(&models.Document{}, id).Error; err != nil {
		return err
	}
	return nil
}
