package convert

import (
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

func PersonalV12Pb(v1Personal *v1.Personal) (pbPersonal *pb.Personal) {

	files := []*pb.FileInfo{}
	if len(v1Personal.Files) > 0 {
		for _, v := range v1Personal.Files {
			files = append(files, FileInfoV12Pb(v))
		}
	}
	pbPersonal = &pb.Personal{
		Id:       v1Personal.ID,
		Files:    files,
		Type:     v1Personal.Type,
		Name:     v1Personal.Name,
		Userid:   v1Personal.Userid,
		Address:  v1Personal.Address,
		ClubName: v1Personal.ClubName,
	}
	return
}

func UserV12Pb(v1User *v1.User) (pbUser *pb.User) {
	pbUser = &pb.User{
		Id:       v1User.ID,
		Name:     v1User.Name,
		Ps:       v1User.Ps,
		Role:     v1User.Role,
		Phone:    v1User.Phone,
		Email:    v1User.Email,
		Avatar:   v1User.Avatar,
		NickName: v1User.NickName,
	}
	return
}

func CaseDataV12Pb(v1CaseData *v1.CaseData) (pbCaseData *pb.CaseData) {
	pbCaseData = &pb.CaseData{
		Id:     v1CaseData.ID,
		Userid: v1CaseData.Userid,
	}
	return
}

func ClassV12Pb(v1Class *v1.Class) (pbClass *pb.Class) {
	pbClass = &pb.Class{
		Id:      v1Class.ID,
		Grade:   v1Class.Grade,
		Faculty: v1Class.Faculty,
		Number:  v1Class.Number,
		Subject: v1Class.Subject,
	}
	return
}

func MeasureDetailV12Pb(v1MeasureDetail *v1.MeasureDetail) (pbMeasureDetail *pb.MeasureDetail) {
	pbMeasureDetail = &pb.MeasureDetail{
		Id:     v1MeasureDetail.ID,
		Userid: v1MeasureDetail.Userid,
	}
	return
}

func MeasureResultV12Pb(v1MeasureResult *v1.MeasureResult) (pbMeasureResult *pb.MeasureResult) {
	pbMeasureResult = &pb.MeasureResult{
		Id:     v1MeasureResult.ID,
		Userid: v1MeasureResult.Userid,
	}
	return
}

func MoveDataV12Pb(v1MoveData *v1.MoveData) (pbMoveData *pb.MoveData) {
	pbMoveData = &pb.MoveData{
		Id:     v1MoveData.ID,
		Userid: v1MoveData.Userid,
	}
	return
}

func MovePrescriptionV12Pb(v1MovePrescription *v1.MovePrescription) (pbMovePrescription *pb.MovePrescription) {
	pbMovePrescription = &pb.MovePrescription{
		Id:     v1MovePrescription.ID,
		Userid: v1MovePrescription.Userid,
	}
	return
}

func FileInfoV12Pb(v1FileInfo *v1.FileInfo) (pbFileInfo *pb.FileInfo) {
	pbFileInfo = &pb.FileInfo{
		Name: v1FileInfo.Name,
		Url:  v1FileInfo.URL,
	}
	return
}

func QuestionOptionV12Pb(v1FileInfo *v1.QuestionOptions) (pbFileInfo *pb.QuestionOptions) {
	pbFileInfo = &pb.QuestionOptions{
		Score: v1FileInfo.Score,
		Desc:  v1FileInfo.Desc,
		Id: v1FileInfo.ID,
	}
	return
}

func NoteInfoV12Pb(v1NoteInfo *v1.NoteInfo) (pbNoteInfo *pb.NoteInfo) {
	pbNoteInfo = &pb.NoteInfo{
		Userid:  v1NoteInfo.Userid,
		Score:   v1NoteInfo.Score,
		Content: v1NoteInfo.Content,
		Status:  v1NoteInfo.Status,
		Value:   v1NoteInfo.Value,
	}
	return
}

func CourseV12Pb(v1Course *v1.Course) (pbCourse *pb.Course) {
	files := []*pb.FileInfo{}
	if len(v1Course.Files) > 0 {
		for _, v := range v1Course.Files {
			files = append(files, FileInfoV12Pb(v))
		}
	}
	pbCourse = &pb.Course{
		Id:       v1Course.ID,
		Userid:   v1Course.Userid,
		Title:    v1Course.Title,
		Files:    files,
		Content:  v1Course.Content,
		Desc:     v1Course.Desc,
		Type:     v1Course.Type,
		Create:   v1Course.Create,
		Update:   v1Course.Update,
		Username: v1Course.Username,
	}
	return
}

func WorkV12Pb(v1Work *v1.Work) (pbWork *pb.Work) {
	files := []*pb.FileInfo{}
	if len(v1Work.Files) > 0 {
		for _, v := range v1Work.Files {
			files = append(files, FileInfoV12Pb(v))
		}
	}
	noteInfos := []*pb.NoteInfo{}
	if len(v1Work.NoteInfo) > 0 {
		for _, v := range v1Work.NoteInfo {
			noteInfos = append(noteInfos, NoteInfoV12Pb(v))
		}
	}
	noters := []string{}
	if len(v1Work.Noter) > 0 {
		noters = append(noters, v1Work.Noter...)
	}
	pbWork = &pb.Work{
		Id:            v1Work.ID,
		Userid:        v1Work.Userid,
		Course:        v1Work.Course,
		CourseTitle:   v1Work.CourseTitle,
		Files:         files,
		Content:       v1Work.Content,
		Note:          v1Work.Note,
		Score:         v1Work.Score,
		Type:          v1Work.Type,
		Create:        v1Work.Create,
		Update:        v1Work.Update,
		Username:      v1Work.Username,
		CourseCreater: v1Work.CourseCreater,
		Teacherid:     v1Work.Teacherid,
		NoteInfo:      noteInfos,
		Noter:         noters,
	}
	return
}

func WorkSubmitV12Pb(v1WorkSubmit *v1.WorkSubmit) (pbWorkSubmit *pb.WorkSubmit) {
	pbWorkSubmit = &pb.WorkSubmit{
		Id:        v1WorkSubmit.ID,
		Studentid: v1WorkSubmit.Studentid,
		Taskid:    v1WorkSubmit.Taskid,
		Content:   v1WorkSubmit.Content,
	}
	return
}

func CommentV12Pb(v1Comment *v1.Comment) (pbComment *pb.Comment) {
	pbComment = &pb.Comment{
		Id:          v1Comment.ID,
		Courseid:    v1Comment.Courseid,
		CommentUser: v1Comment.CommentUser,
		Comment:     v1Comment.Comment,
		Score:       v1Comment.Score,
	}
	return
}

func SportTypeV12Pb(v1SportType *v1.SportType) (pbSportType *pb.SportType) {
	pbSportType = &pb.SportType{
		Id:       v1SportType.ID,
		Label:    v1SportType.Label,
		Code:     v1SportType.Code,
		Level:    v1SportType.Level,
		ParentId: v1SportType.ParentID,
	}
	return
}

func EvaluateV12Pb(v1Evaluate *v1.Evaluate) (pbEvaluate *pb.Evaluate) {
	pbEvaluate = &pb.Evaluate{
		Id:        v1Evaluate.ID,
		Name:      v1Evaluate.Name,
		Excellent: v1Evaluate.Excellent,
		Good:      v1Evaluate.Good,
		Pass:      v1Evaluate.Pass,
		Flunk:     v1Evaluate.Flunk,
	}
	return
}

func SportItemV12Pb(v1SportItem *v1.SportItem) (pbSportItem *pb.SportItem) {
	options := []*pb.QuestionOptions{}
	if len(v1SportItem.QuestionOptions) > 0 {
		for _, v := range v1SportItem.QuestionOptions {
			options = append(options, QuestionOptionV12Pb(v))
		}
	}

	pbSportItem = &pb.SportItem{
		Id:              v1SportItem.ID,
		First:           v1SportItem.First,
		Second:          v1SportItem.Second,
		Label:           v1SportItem.Label,
		UploadFile:      v1SportItem.UploadFile,
		FileName:        v1SportItem.FileName,
		FileType:        v1SportItem.FileType,
		CreateTime:      v1SportItem.CreateTime,
		UpdateTime:      v1SportItem.UpdateTime,
		QuestionOptions: options,
		Approver: v1SportItem.Approver,
	}
	return
}

func HomeWorkV12Pb(v1HomeWork *v1.HomeWork) (pbHomeWork *pb.HomeWork) {
	solutions := []*pb.ApplyItem{}
	if len(v1HomeWork.ApplyItems) > 0 {
		for _, v := range v1HomeWork.ApplyItems {
			solutions = append(solutions, ApplyItemV12Pb(v))
		}
	}

	pbHomeWork = &pb.HomeWork{
		Id:          v1HomeWork.ID,
		ApplyUserId: v1HomeWork.ApplyUserID,
		Status:       v1HomeWork.Status,
		ClubName:   v1HomeWork.ClubName,
		Create:     v1HomeWork.Create,
		Update:     v1HomeWork.Update,
		ApplyItems: solutions,
	}
	return
}

func EvaluateInfoV12Pb(v1EvaluateInfo *v1.EvaluateInfo) (pbEvaluateInfo *pb.EvaluateInfo) {
	evaluates := []*pb.Evaluate{}
	if len(v1EvaluateInfo.Evaluates) > 0 {
		for _, v := range v1EvaluateInfo.Evaluates {
			evaluates = append(evaluates, EvaluateV12Pb(v))
		}
	}
	pbEvaluateInfo = &pb.EvaluateInfo{
		Id:         v1EvaluateInfo.ID,
		HomeWorkId: v1EvaluateInfo.HomeWorkID,
		SolutionId: v1EvaluateInfo.SolutionID,
		UserId:     v1EvaluateInfo.UserID,
		Name:       v1EvaluateInfo.Name,
		Evaluates:  evaluates,
		Note:       v1EvaluateInfo.Note,
	}
	return
}

func SolutionV12Pb(v1Solution *v1.Solution) (pbSolution *pb.Solution) {
	evaluates := []*pb.EvaluateInfo{}
	if len(v1Solution.EvaluatesInfo) > 0 {
		for _, v := range v1Solution.EvaluatesInfo {
			evaluates = append(evaluates, EvaluateInfoV12Pb(v))
		}
	}
	teachers := []string{}
	if len(v1Solution.Teachers) > 0 {
		teachers = append(teachers, v1Solution.Teachers...)
	}
	pbSolution = &pb.Solution{
		Id:            v1Solution.ID,
		Title:         v1Solution.Title,
		WorkId:        v1Solution.WorkID,
		StudentId:     v1Solution.StudentID,
		StudentName:   v1Solution.StudentName,
		Content:       v1Solution.Content,
		EvaluatesInfo: evaluates,
		CommitTime:    v1Solution.CommitTime,
		Desc:          v1Solution.Desc,
		Teachers:      teachers,
	}
	return
}

func ApplyItemV12Pb(v1Solution *v1.ApplyItem) (pbSolution *pb.ApplyItem) {
	files := []*pb.FileInfo{}
	if len(v1Solution.Files) > 0 {
		for _, v := range v1Solution.Files {
			files = append(files, FileInfoV12Pb(v))
		}
	}
	pbSolution = &pb.ApplyItem{
		TemplateId:  v1Solution.Templateid,
		Value:       v1Solution.Value,
		ReviewValue: v1Solution.ReviewValue,
		ReviewUser:  v1Solution.ReviewUser,
		Files:       files,
	}
	return
}
