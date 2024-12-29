package unit_tests

import (
	"realtime-doc-editor-backend/internal/models"
	"realtime-doc-editor-backend/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Initialize an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Document{})
	return db
}

func TestCreateDocument(t *testing.T) {
	db := setupTestDB()
	repo := repositories.NewDocumentRepository(db)

	// Test CreateDocument
	doc, err := repo.CreateDocument("Test Title", "Test Content")
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	assert.Equal(t, "Test Title", doc.Title)
	assert.Equal(t, "Test Content", doc.Content)
}

func TestGetDocumentByID(t *testing.T) {
	db := setupTestDB()
	repo := repositories.NewDocumentRepository(db)

	// Create a document to test retrieval
	createdDoc, _ := repo.CreateDocument("Test Title", "Test Content")

	// Test GetDocumentByID
	doc, err := repo.GetDocumentByID(createdDoc.ID)
	assert.NoError(t, err)
	assert.NotNil(t, doc)
	assert.Equal(t, createdDoc.ID, doc.ID)
	assert.Equal(t, "Test Title", doc.Title)
}

func TestGetAllDocuments(t *testing.T) {
	db := setupTestDB()
	repo := repositories.NewDocumentRepository(db)

	// Create documents to test retrieval
	_, _ = repo.CreateDocument("Test Title 1", "Content 1")
	_, _ = repo.CreateDocument("Test Title 2", "Content 2")

	// Test GetAllDocuments
	docs, err := repo.GetAllDocuments()
	assert.NoError(t, err)
	assert.Len(t, docs, 2)
}

func TestUpdateDocument(t *testing.T) {
	db := setupTestDB()
	repo := repositories.NewDocumentRepository(db)

	// Create a document to test updating
	createdDoc, _ := repo.CreateDocument("Test Title", "Test Content")

	// Test UpdateDocument
	updatedDoc, err := repo.UpdateDocument(createdDoc.ID, "Updated Title", "Updated Content")
	assert.NoError(t, err)
	assert.NotNil(t, updatedDoc)
	assert.Equal(t, "Updated Title", updatedDoc.Title)
	assert.Equal(t, "Updated Content", updatedDoc.Content)
}

func TestDeleteDocument(t *testing.T) {
	db := setupTestDB()
	repo := repositories.NewDocumentRepository(db)

	// Create a document to test deletion
	createdDoc, _ := repo.CreateDocument("Test Title", "Test Content")

	// Test DeleteDocument
	err := repo.DeleteDocument(createdDoc.ID)
	assert.NoError(t, err)

	// Try to get the deleted document
	doc, err := repo.GetDocumentByID(createdDoc.ID)
	assert.Error(t, err)
	assert.Nil(t, doc)
}
