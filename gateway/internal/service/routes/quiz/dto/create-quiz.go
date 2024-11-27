package dto

import "time"

type QuestionType string

const (
	SingleChoice   QuestionType = "single_choice"
	MultipleChoice QuestionType = "multiple_choice"
	FreeText       QuestionType = "free_text"
	Scale          QuestionType = "scale"
)

// QuestionOption представляет вариант ответа для вопросов с выбором
type QuestionOption struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// Dependency представляет зависимость между вопросами
type Dependency struct {
	QuestionID  string `json:"question_id"`
	OptionID    string `json:"option_id"`
	DependentOn string `json:"dependent_on"`
}

// Question представляет вопрос в опроснике
type Question struct {
	ID           string           `json:"id"`
	Text         string           `json:"text"`
	Type         QuestionType     `json:"type"`
	Options      []QuestionOption `json:"options,omitempty"`
	MinValue     *int             `json:"min_value,omitempty"`
	MaxValue     *int             `json:"max_value,omitempty"`
	IsRequired   bool             `json:"is_required"`
	Dependencies []Dependency     `json:"dependencies,omitempty"`
}

// Quiz представляет опросник
type Quiz struct {
	ID           string     `json:"id"`
	Title        string     `json:"title"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	RewardPoints int        `json:"reward_points"`
	IsAnonymous  bool       `json:"is_anonymous"`
	Questions    []Question `json:"questions"`
}
