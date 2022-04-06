package dto

import "encoding/json"

type ResumeDto struct {
	Person  Person `json:"person"`
	Learn   string `json:"learn,omitempty"`
	Work    string `json:"work,omitempty"`
	Project string `json:"project,omitempty"`
	Dream   string `json:"dream,omitempty"`
}

type Person struct {
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Age     string `json:"age"`
	Stature string `json:"stature"`
	Weight  string `json:"weight"`
}

func NewResumeDto() *ResumeDto {
	return &ResumeDto{}
}

func (r *ResumeDto) ToString() (string, error) {
	bytes, error := json.Marshal(r)
	return string(bytes), error
}
