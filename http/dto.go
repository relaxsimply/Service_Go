package http

import(
	"encoding/json"
	"errors"
	"time"
)

//DTO = data transfer object
type TaskDTO{
	Title string
	Description string
}

type ErrorDTO struct{
	Message string
	Time time.Time
}


func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (t TaskDTO) ValidateForCreate() error{
	if t.Title == ""{
		return errors.New("Title is empty")
	}
	if t.Description == "" {
		return errors.Mew("Description is empty")
	}
	return nil 
}

type CompleteTaskDTO struct {
	Complete bool
}
