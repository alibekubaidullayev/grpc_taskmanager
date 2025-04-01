package models

import (
	"errors"
	"fmt"
	"tm/src/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string  `gorm:"size:64"`
	Description *string `gorm:"size:512"`
}

func (t *Task) SetTitle(title string) error {
	if len(title) == 0 {
		return errors.New("can't have empty title")
	}

	if len(title) > 64 {
		return errors.New("length of title of the task can't be longer than 64 characters")
	}

	t.Title = title
	return nil
}

func (t *Task) SetDescription(description *string) error {
	if description != nil && len(*description) > 512 {
		return fmt.Errorf("Length of description can't be longer than 512 characters")
	}

	t.Description = description
	return nil
}

func (t *Task) FromCreate(in *pb.CreateTaskRequest) error {
	if err := t.SetTitle(in.Title); err != nil {
		return err
	}
	if err := t.SetDescription(in.Description); err != nil {
		return err
	}

	return nil
}

func (t *Task) FromUpdate(in *pb.UpdateTaskRequest) error {
	if in.Title == nil && in.Description == nil {
		return errors.New("no update field provided")
	}

	if title := in.Title; title != nil {
		if err := t.SetTitle(*title); err != nil {
			return err
		}
	}

	if desc := in.Description; desc != nil {
		if err := t.SetDescription(desc); err != nil {
			return err
		}
	}

	return nil
}

func (t *Task) GetResponse() *pb.GetTaskResponse {
	var description string
	if t.Description != nil {
		description = *t.Description
	}

	var deletedAt *timestamppb.Timestamp
	if t.DeletedAt.Valid {
		deletedAt = timestamppb.New(t.DeletedAt.Time)
	}

	return &pb.GetTaskResponse{
		Id:          uint64(t.ID),
		Title:       t.Title,
		Description: description,
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   timestamppb.New(t.UpdatedAt),
		DeletedAt:   deletedAt,
	}
}

func (t *Task) GetUpdatedResponse() *pb.UpdateTaskResponse {
	var title *string
	if t.Title != "" {
		title = &t.Title
	}

	return &pb.UpdateTaskResponse{
		Id:          uint64(t.ID),
		Title:       title,
		Description: t.Description,
		UpdatedAt:   timestamppb.New(t.UpdatedAt),
	}
}
