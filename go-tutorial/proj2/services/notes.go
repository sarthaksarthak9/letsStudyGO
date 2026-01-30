package services

import (
	"fmt"
	internal "proj2/internal/model"

	"gorm.io/gorm"
)

// service handles all the business logic
type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(db *gorm.DB) {
	n.db = db
	n.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesService(status bool) ([]*internal.Notes, error) {
	// Pointers are required because GORM needs to MODIFY the struct you pass to it.
	// Without a pointer, Go would give GORM only a copy — and changes would be lost.
	var notes []*internal.Notes

	if err := n.db.Where("status = ?", status).Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

// data flow here from frontend
func (n *NotesService) CreateNotesService(title string, status bool) (*internal.Notes, error) {
	// data := Note{
	// 	Id: 3, Name: "Note 3",
	// }
	note := &internal.Notes{
		Title:  title,
		Status: status,
	}

	if err := n.db.Create(note).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}

	return note, nil
}

func (n *NotesService) UpdateNotesService(title string, status bool, id int) (*internal.Notes, error) {
	// data := Note{
	// 	Id: 3, Name: "Note 3",
	// }
	// note := &internal.Notes{
	// 	Title:  title,
	// 	Status: status,
	// }

	var note *internal.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}

	note.Title = title
	note.Status = status

	if err := n.db.Save(note).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}

	return note, nil
}

func (n *NotesService) DeleteNotesService(id int64) (error){
	
	var note *internal.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err!=nil{
		return err
	}

	if err := n.db.Where("id = ?", id).Delete(&note).Error; err != nil{
		fmt.Print(err)
		return err
	} 

	return nil
}
// GetAllNotes returns all notes without filtering
func (n *NotesService) GetAllNotes() ([]*internal.Notes, error) {
	var notes []*internal.Notes
	
	if err := n.db.Find(&notes).Error; err != nil {
		return nil, err
	}
	
	return notes, nil
}
