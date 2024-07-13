package repo

import (
	"context"

	. "school_council/db"
	"school_council/models"
)

func CheckStudentExistence(stdNumber string, ctx context.Context) (bool, error) {
	var student models.Student
	var count int64
	err := DB.WithContext(ctx).
		Where("std_number = ?", stdNumber).
		First(&student).
		Count(&count).
		Error

	return count > 0, err
}

func CreateStudent(student *models.Student, ctx context.Context) error {
	return DB.WithContext(ctx).Create(student).Error
}

func UpdateStudent(student *models.Student, ctx context.Context) error {
	return DB.WithContext(ctx).Save(student).Error
}

func CreateGrade(grade *models.Grade, ctx context.Context) error {
	return DB.WithContext(ctx).Create(grade).Error
}
