package service

import (
	"testing"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
	"github.com/xuanit/testing/todo/pb"
	"github.com/xuanit/testing/todo/server/repository/mocks"
)

func TestGetToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}
	req := &pb.GetTodoRequest{Id: "123"}
	mockToDoRep.On("Get", req.Id).Return(toDo, nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.GetTodo(nil, req)

	expectedRes := &pb.GetTodoResponse{Item: toDo}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestListToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	listToDo := []*pb.Todo{
		&pb.Todo{
			Id:          "a1",
			Title:       "b",
			Description: "c",
			Completed:   true,
			CreatedAt:   &timestamp.Timestamp{Seconds: 1},
			UpdatedAt:   &timestamp.Timestamp{Seconds: 1},
		},
		&pb.Todo{
			Id:          "a2",
			Title:       "b",
			Description: "c",
			Completed:   true,
			CreatedAt:   &timestamp.Timestamp{Seconds: 1},
			UpdatedAt:   &timestamp.Timestamp{Seconds: 1},
		},
	}
	req := &pb.ListTodoRequest{Limit: 1, NotCompleted: true}

	mockToDoRep.On("List", req.Limit, req.NotCompleted).Return(listToDo[:req.Limit], nil)

	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.ListTodo(nil, req)

	expectedRes := &pb.ListTodoResponse{Items: listToDo[:req.Limit]}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestCreateToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{
		Id:          "a",
		Title:       "b",
		Description: "c",
		Completed:   true,
		CreatedAt:   &timestamp.Timestamp{Seconds: 1},
		UpdatedAt:   &timestamp.Timestamp{Seconds: 1},
	}
	req := &pb.CreateTodoRequest{Item: toDo}

	mockToDoRep.On("Insert", req.Item).Return(nil)

	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.CreateTodo(nil, req)

	expectedRes := &pb.CreateTodoResponse{Id: toDo.Id}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestDeleteToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDoID := "1"
	req := &pb.DeleteTodoRequest{Id: toDoID}

	mockToDoRep.On("Delete", req.Id).Return(nil)

	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.DeleteTodo(nil, req)

	expectedRes := &pb.DeleteTodoResponse{}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}
