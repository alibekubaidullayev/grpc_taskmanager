package server

import (
	"context"
	"fmt"
	"tm/src/models"
	"tm/src/pb"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedTaskManagerServer
	Db *gorm.DB
}

func (s *Server) Create(
	ctx context.Context, in *pb.CreateTaskRequest,
) (*pb.GetTaskResponse, error) {
	var newTask models.Task
	newTask.FromCreate(in)

	if err := s.Db.Create(&newTask).Error; err != nil {
		return nil, err
	}

	return newTask.GetResponse(), nil
}

func (s *Server) Get(
	ctx context.Context, in *pb.IdRequest,
) (*pb.GetTaskResponse, error) {
	id := uint(in.Id)
	if id == 0 {
		return nil, fmt.Errorf("Invalid id (must be uint and > 0)")
	}

	var task models.Task
	if err := s.Db.First(&task, id).Error; err != nil {
		return nil, err
	}

	return task.GetResponse(), nil
}

func (s *Server) List(
	ctx context.Context, _ *emptypb.Empty,
) (*pb.ListTasksResponse, error) {
	var tasks []models.Task
	if err := s.Db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	var res []*pb.GetTaskResponse
	for _, task := range tasks {
		t := task
		res = append(res, t.GetResponse())
	}

	return &pb.ListTasksResponse{
		Tasks: res,
	}, nil
}

func (s *Server) Update(
	ctx context.Context, in *pb.UpdateTaskRequest,
) (*pb.UpdateTaskResponse, error) {
	id := uint(in.Id)
	if id == 0 {
		return nil, fmt.Errorf("Invalid id (must be uint and > 0)")
	}

	var updateTask models.Task
	err := updateTask.FromUpdate(in)
	if err != nil {
		return nil, err
	}

	var existingTask models.Task
	if err := s.Db.First(&existingTask, id).Error; err != nil {
		return nil, err
	}

	if err := s.Db.Model(&existingTask).Updates(&updateTask).Error; err != nil {
		return nil, err
	}
	updateTask.ID = existingTask.ID
	updateTask.UpdatedAt = existingTask.UpdatedAt

	return updateTask.GetUpdatedResponse(), nil
}

func (s *Server) Delete(
	ctx context.Context, in *pb.IdRequest,
) (*pb.GetTaskResponse, error) {
	id := uint(in.Id)
	if id == 0 {
		return nil, fmt.Errorf("invalid id (must be uint and > 0)")
	}

	var task models.Task
	if err := s.Db.First(&task, id).Error; err != nil {
		return nil, err
	}

	if err := s.Db.Delete(&task).Error; err != nil {
		return nil, err
	}

	return task.GetResponse(), nil
}
