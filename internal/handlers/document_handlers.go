package handlers

import (
	"net/http"
	"realtime-doc-editor-backend/internal/models"
	"realtime-doc-editor-backend/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateDocument handles the creation of a document
func CreateDocumentHandler(c *gin.Context, documentRepo *repositories.DocumentRepository) {
	var input models.Document
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	doc, err := documentRepo.CreateDocument(input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doc)
}

// GetAllDocuments handles fetching all documents
func GetAllDocumentsHandler(c *gin.Context, documentRepo *repositories.DocumentRepository) {
	docs, err := documentRepo.GetAllDocuments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, docs)
}

// GetDocumentByID handles fetching a single document by its ID
func GetDocumentByIDHandler(c *gin.Context, documentRepo *repositories.DocumentRepository) {
	// Convert the id from string to uint
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	// Since ParseUint returns an unsigned 64-bit integer, we cast it to uint
	doc, err := documentRepo.GetDocumentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}
	c.JSON(http.StatusOK, doc)
}

// UpdateDocument handles updating an existing document by its ID
func UpdateDocumentHandler(c *gin.Context, documentRepo *repositories.DocumentRepository) {
	// Convert the id from string to uint
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	var input models.Document
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	doc, err := documentRepo.UpdateDocument(uint(id), input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doc)
}

// DeleteDocument handles deleting a document by its ID
func DeleteDocumentHandler(c *gin.Context, documentRepo *repositories.DocumentRepository) {
	// Convert the id from string to uint
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	if err := documentRepo.DeleteDocument(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
