package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"school_council/dto"
	"school_council/models"
	"school_council/repo"
	"school_council/utils"

	v "github.com/cohesivestack/valgo"
	"gorm.io/gorm"
)

type m map[string]any

func Create(w http.ResponseWriter, r *http.Request) {
	body := dto.CreateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		slog.Error(err.Error())
		utils.ErrResp(w, http.StatusBadRequest, err.Error())
		return
	}

	val := v.
		Is(v.String(body.Name, "name").Not().Blank()).
		Is(v.String(body.Surname, "surname").Not().Blank()).
		Is(v.String(body.StdNumber, "stdNumber").Not().Blank())
		// TODO add validation for grades

	if !val.Valid() {
		utils.JsonResp(w, m{"error": val.Errors()}, http.StatusBadRequest)
		return
	}

	log.Printf("body: %+v", body)

	studentExists, err := repo.CheckStudentExistence(body.StdNumber, r.Context())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error(err.Error())
		utils.InternalErrResp(w, err)
		return
	}
	if studentExists {
		utils.ErrResp(w, http.StatusConflict, "student already exists")
		return
	}

	grades := reduceGrades(body.Grades)

	student := &models.Student{
		Name:      body.Name,
		Surname:   body.Surname,
		StdNumber: body.StdNumber,
		Grades:    grades,
	}

	if err := repo.CreateStudent(student, r.Context()); err != nil {
		slog.Error("create student", "err", err.Error())
		utils.InternalErrResp(w, err)
		return
	}

	respGrades := make([]dto.GradeResponse, 0, len(student.Grades))
	for _, grade := range student.Grades {
		respGrades = append(respGrades, dto.GradeResponse{
			Code:  grade.Code,
			Value: grade.Value,
		})
	}

	resp := dto.CreateResponse{
		Student: dto.Student{
			Name:      student.Name,
			Surname:   student.Surname,
			StdNumber: student.StdNumber,
		},
		ID:     student.ID,
		Grades: respGrades,
	}

	// student GET endpoint is not implemented yet
	w.Header().Set("Location", fmt.Sprintf("/students/%d", student.ID))
	utils.JsonResp(w, m{"status": "success", "data": resp}, http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	body := dto.UpdateRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		slog.Error(err.Error())
		utils.ErrResp(w, http.StatusBadRequest, err.Error())
		return
	}

	val := v.
		Is(v.String(body.Name, "name").Not().Blank()).
		Is(v.String(body.Surname, "surname").Not().Blank())
		// TODO add validation for grades

	if !val.Valid() {
		utils.JsonResp(w, m{"error": val.Errors()}, http.StatusBadRequest)
		return
	}

	stdNumber := r.PathValue("stdNumber")
	if stdNumber == "" {
		utils.ErrResp(w, http.StatusBadRequest, "stdNumber param is required")
		return
	}

	studentExists, err := repo.CheckStudentExistence(stdNumber, r.Context())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error(err.Error())
		utils.InternalErrResp(w, err)
		return
	}
	if !studentExists {
		utils.ErrResp(w, http.StatusNotFound, "student not found")
		return
	}

	grades := reduceGrades(body.Grades)

	student := &models.Student{
		Name:      body.Name,
		Surname:   body.Surname,
		Grades:    grades,
		StdNumber: stdNumber,
	}

	if err := repo.UpdateStudent(student, r.Context()); err != nil {
		slog.Error("update student", "err", err.Error())
		utils.InternalErrResp(w, err)
		return
	}

	respGrades := make([]dto.GradeResponse, 0, len(student.Grades))
	for _, grade := range student.Grades {
		respGrades = append(respGrades, dto.GradeResponse{
			Code:  grade.Code,
			Value: grade.Value,
		})
	}

	resp := dto.CreateResponse{
		Student: dto.Student{
			Name:      student.Name,
			Surname:   student.Surname,
			StdNumber: student.StdNumber,
		},
		ID:     student.ID,
		Grades: respGrades,
	}

	fmt.Printf("resp: %vn", resp)

	// student GET endpoint is not implemented yet
	w.Header().Set("Location", fmt.Sprintf("/students/%d", student.ID))
	utils.JsonResp(w, m{"status": "success", "data": resp})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func reduceGrades(grades []dto.GradeRequest) []models.Grade {
	gradeMap := make(map[string][]float64)
	for _, grade := range grades {
		gradeMap[grade.Code] = append(gradeMap[grade.Code], float64(grade.Value))
	}

	reducedGrades := make([]models.Grade, 0, len(gradeMap))
	for code, values := range gradeMap {
		sum := 0.0
		for _, value := range values {
			sum += value
		}
		avg := sum / float64(len(values))
		reducedGrades = append(reducedGrades, models.Grade{
			Code:  code,
			Value: avg,
		})
	}

	return reducedGrades
}
