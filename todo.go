package main

type Todo struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Completed bool     `json:"completed"`
	Due       jsonTime `json:"due"`
}

type Todos []Todo
