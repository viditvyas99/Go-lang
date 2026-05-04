package models

import (
	"example/api/models"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	ID          int
	Name        string `binfding:"required"`
	Description string
	DateTime    time.Time
	userId      int
}

var events = []Event{}

func (e Event) save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}

func CreateEvenet(context *gin.Context) {
	var event models.Event
	context.ShouldBindBodyWithJSON(&event)
}
