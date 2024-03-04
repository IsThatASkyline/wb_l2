package routes

import (
	"net/http"

	"github.com/IsThatASkyline/wb_l2/dev11/calendar"
	"github.com/IsThatASkyline/wb_l2/dev11/middleware"
)

func IntitRoutes(c calendar.Calendar) {
	handler := &Handler{
		c: c,
	}
	http.Handle("/api/create_event", middleware.Logging(http.HandlerFunc(handler.CreateEvent)))
	http.Handle("/api/update_event", middleware.Logging(http.HandlerFunc(handler.UpdateEvent)))
	http.Handle("/api/delete_event", middleware.Logging(http.HandlerFunc(handler.DeleteEvent)))
	http.Handle("/api/events_for_day", middleware.Logging(http.HandlerFunc(handler.ShowEvents)))
	http.Handle("/api/events_for_week", middleware.Logging(http.HandlerFunc(handler.ShowEvents)))
	http.Handle("/api/events_for_month", middleware.Logging(http.HandlerFunc(handler.ShowEvents)))
}
