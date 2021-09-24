package server

import (
	"beznet/adullam/auth"
	"beznet/adullam/api"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/api/v1/auth", auth.Authenticate)
	
	http.Handle("/api/v1/student_course_details", auth.IsAuthorized(api.StudentCourseDetails))
}

