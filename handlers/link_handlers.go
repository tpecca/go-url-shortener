package handlers

import (
	"context"
	"go-url-shortener/db"
	"go-url-shortener/models"
	"go-url-shortener/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateShortLink(c *gin.Context) {
	entry := models.LinkEntry{
		OriginalURL: c.PostForm("url"),
		Hash:        utils.CreateHash(6),
	}

	result, err := db.LinksDB.InsertOne(context.Background(), entry)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			var existing models.LinkEntry
			if findErr := db.LinksDB.FindOne(context.Background(), bson.M{"original_url": entry.OriginalURL}).Decode(&existing); findErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "Database lookup failed"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:4000/" + existing.Hash})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to store link"})
		return
	}

	if result != nil {
		c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:4000/" + entry.Hash})
	}
}

func RedirectToOriginal(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Missing hash parameter"})
		return
	}

	var entry models.LinkEntry
	if err := db.LinksDB.FindOne(context.Background(), bson.M{"hash": hash}).Decode(&entry); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Link not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, entry.OriginalURL)
}
