package service

import (
	"backend/go/dao"
	"backend/go/entity"
	"backend/go/request"
)

/*
新建问卷信息
*/
func CreateQuestionnaire(createQuestionnaireRequest *request.CreateQuestionnaireRequest) (err error) {
	var questionnaire entity.Questionnaire
	questionnaire.StudentId = createQuestionnaireRequest.StudentId
	questionnaire.CourseId = createQuestionnaireRequest.CourseId
	questionnaire.Satisfaction = createQuestionnaireRequest.Satisfaction
	questionnaire.Comment = createQuestionnaireRequest.Comment
	questionnaire.Suggestion = createQuestionnaireRequest.Suggestion
	student, err := GetStudentById(createQuestionnaireRequest.StudentId)
	questionnaire.StudentName = student.Name
	course, err := GetCourseById(createQuestionnaireRequest.CourseId)
	questionnaire.CourseName = course.Name
	if err = dao.SqlSession.Create(questionnaire).Error; err != nil {
		return err
	}
	return
}

/*
获取问卷集合
*/
func GetAllQuestionnaire() (questionnaireList []*entity.Questionnaire, err error) {
	if err = dao.SqlSession.Find(&questionnaireList).Error; err != nil {
		return nil, err
	}
	return
}

/*
根据id删除问卷
*/
func DeleteQuestionnaireById(id uint) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.Questionnaire{}).Error
	return
}

/*
根据id查询问卷
*/
func GetQuestionnaireById(id uint) (questionnaire *entity.Questionnaire, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(questionnaire).Error; err != nil {
		return nil, err
	}
	return
}

/*
更新问卷信息
*/
func UpdateQuestionnaire(questionnaire *entity.Questionnaire) (err error) {
	err = dao.SqlSession.Save(questionnaire).Error
	return
}
