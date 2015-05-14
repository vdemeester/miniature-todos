package main

import (
	"fmt"
	"time"
)

var currentId int

var todos Todos

// Give us some seed data
func init() {
	due, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	// Format is 2015-05-14T14:14:00
	defaultFormat := "2006-01-02T15:04:00"
	RepoCreateTodo(Todo{Name: "Write presentation", Completed: true})
	RepoCreateTodo(Todo{Name: "Host meetup", Due: jsonTime{time.Now(), defaultFormat}})
	RepoCreateTodo(Todo{Name: "Initiate project Y", Due: jsonTime{due, defaultFormat}})
	RepoCreateTodo(Todo{Name: "Prepare retro. n.100"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
