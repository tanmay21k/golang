package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - in file (not atm)
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var Courses []Course

// fake db
func main() {

	Courses = append(Courses, Course{
		CourseId:    "123",
		CourseName:  "exmaple",
		CoursePrice: "231",
		Author: &Author{
			FullName: "Writer1",
			Website:  "www.youtube.com",
		},
	})
	fmt.Printf("%+v\n", Courses)
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", CreateOneCourse).Methods("POST")
	router.HandleFunc("/update", UpdateCourse).Methods("PUT")
	router.HandleFunc("/delete", deleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}

// creating router

// middlewares
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" //&& c.CourseName == ""
}

// controlers - in seperate file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> API CREATED</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grabbing id as per the params
	params := mux.Vars(r)

	// loop through the courses and find matching id as param and return the response
	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("NO course of given id")

	}
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	// check if the body is empty
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the json")
		return
	}

	// generate unique id, id conversion to string
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	randomNum := rand.New(source)

	// Now you can use `randomNum` for generating random numbers
	course.CourseId = strconv.Itoa((randomNum.Intn(100)))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)

	// append this new course into courses
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// loop, id , remove and add with myId
	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: send a response when Id is not found
	json.NewEncoder(w).Encode("The courseID was not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			json.NewEncoder(w).Encode("Record deleted successfully")
			return
		}
	}
}
