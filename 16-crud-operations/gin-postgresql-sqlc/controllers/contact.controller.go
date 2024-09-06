package controllers

import (
	"context"
	"database/sql"
	db "gin-postgres/db/sqlc"
	"gin-postgres/schemas"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ContactController struct {
	db  *db.Queries
	ctx context.Context
}

// NewContactController returns the new instance of ContactController
func NewContactController(db *db.Queries, ctx context.Context) *ContactController {
	return &ContactController{db, ctx}
}

// CreateContact creates new entry into the contacts table and return data if success
func (cc *ContactController) CreateContact(ctx *gin.Context) {
	var payload *schemas.CreateContact

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed Payload",
			"error":   err.Error(),
		})
		return
	}

	now := time.Now()

	args := &db.CreateContactParams{
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		PhoneNumber: payload.PhoneNumber,
		Street:      payload.Street,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	contact, err := cc.db.CreateContact(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed creating contact",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Contact created successfully",
		"data":    contact,
	})
}

// UpdateContact update the contact and return data if success
func (cc *ContactController) UpdateContact(ctx *gin.Context) {
	var payload *schemas.UpdateContact
	contactId := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed Payload",
			"error":   err.Error(),
		})
		log.Println("update payload error", err)
		return
	}

	now := time.Now()

	args := &db.UpdateContactParams{
		ID:          uuid.MustParse(contactId),
		FirstName:   sql.NullString{String: payload.FirstName, Valid: payload.FirstName != ""},
		LastName:    sql.NullString{String: payload.LastName, Valid: payload.LastName != ""},
		PhoneNumber: sql.NullString{String: payload.PhoneNumber, Valid: payload.PhoneNumber != ""},
		Street:      sql.NullString{String: payload.Street, Valid: payload.Street != ""},
		UpdatedAt:   sql.NullTime{Time: now, Valid: true},
	}

	contact, err := cc.db.UpdateContact(ctx, *args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Failed to retrieve the contact with ID",
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update the contact",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mesage": "Contact update successfully",
		"data":   contact,
	})

}

// GetContactById fetch the contact by ID
func (cc *ContactController) GetContactById(ctx *gin.Context) {
	contactId := ctx.Param("id")

	contact, err := cc.db.GetContactById(ctx, uuid.MustParse(contactId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Failed to retrieve the contact with given ID",
				"error":   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed retrieving contact",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Contact fetched successfully",
		"data":    contact,
	})
}

// GetAllContacts fetch all contacts
func (cc *ContactController) GetAllContacts(ctx *gin.Context) {
	var reqPage = ctx.DefaultQuery("page", "1")
	var reqLimit = ctx.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(reqPage)
	limit, _ := strconv.Atoi(reqLimit)

	offset := (page - 1) * limit

	args := &db.ListContactsManyParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	contacts, err := cc.db.ListContactsMany(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve contacts",
			"error":   err.Error(),
		})
		return
	}

	if contacts == nil {
		contacts = []db.Contact{}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Fetched Successfully",
		"size":    len(contacts),
		"data":    contacts,
	})
}

// DeleteContactById delete the contact by ID from the db
func (cc *ContactController) DeleteContactById(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := cc.db.GetContactById(ctx, uuid.MustParse(id))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Record not found with this ID",
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"error":   err.Error(),
		})
		return
	}

	err = cc.db.DeleteContact(ctx, uuid.MustParse(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete the record",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Contact deleted successfully",
	})
}
