package controller

import (
	internal "proj2/internal/model"
	"proj2/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	// defining object of struct that we have in noteservice
	noteService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, notesService services.NotesService) {
	// All routes in this group will start with /notes.
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	notes.PUT("/", n.UpdateNotes())
	notes.DELETE("/:id", n.DeleteNotes())
	n.noteService = notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Query("status")

		// Pointers are required because GORM needs to MODIFY the struct you pass to it.
		// Without a pointer, Go would give GORM only a copy — and changes would be lost.
		var notes []*internal.Notes
		var err error

		if status == "" {
			notes, err = n.noteService.GetAllNotes()
		} else {
			actualStatus, parseErr := strconv.ParseBool(status)
			if parseErr != nil {
				c.JSON(400, gin.H{
					"message": "Invalid status parameter. Use true or false.",
				})
				return
			}
			notes, err = n.noteService.GetNotesService(actualStatus)
		}

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"notes": notes,
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title"`
		Status bool   `json:"status"`
	}

	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		}
		note, err := n.noteService.CreateNotesService(noteBody.Title, noteBody.Status)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"note": note,
		})

		// c.JSON(200, gin.H{
		// 	"notes": n.noteService.CreateNotesService(),
		// })
	}
}

func (n *NotesController) UpdateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title"`
		Status bool   `json:"status"`
		Id     int    `json:"id" binding:"required"`
	}

	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		}
		note, err := n.noteService.UpdateNotesService(noteBody.Title, noteBody.Status, noteBody.Id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"note": note,
		})

		// c.JSON(200, gin.H{
		// 	"notes": n.noteService.CreateNotesService(),
		// })
	}
}

func (n *NotesController) DeleteNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// we need int64 ?? understand
		noteId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		// understand = here ??
		err = n.noteService.DeleteNotesService(noteId)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"note": "Note deleted succesfully",
		})

	}
}

// we have to register this controller in main.go
