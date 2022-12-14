// Code generated by ent, DO NOT EDIT.

package entuser

import (
	"time"
)

const (
	// Label holds the string label denoting the entuser type in the database.
	Label = "ent_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldFirstName holds the string denoting the firstname field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the lastname field in the database.
	FieldLastName = "last_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldDateOfBirth holds the string denoting the dateofbirth field in the database.
	FieldDateOfBirth = "date_of_birth"
	// FieldProfilePictureUrl holds the string denoting the profilepictureurl field in the database.
	FieldProfilePictureUrl = "profile_picture_url"
	// FieldIsTutor holds the string denoting the istutor field in the database.
	FieldIsTutor = "is_tutor"
	// FieldIsStudent holds the string denoting the isstudent field in the database.
	FieldIsStudent = "is_student"
	// FieldIsParent holds the string denoting the isparent field in the database.
	FieldIsParent = "is_parent"
	// EdgeCourse holds the string denoting the course edge name in mutations.
	EdgeCourse = "course"
	// EdgeTodo holds the string denoting the todo edge name in mutations.
	EdgeTodo = "todo"
	// EdgeAttendance holds the string denoting the attendance edge name in mutations.
	EdgeAttendance = "attendance"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// EdgeComment holds the string denoting the comment edge name in mutations.
	EdgeComment = "comment"
	// EdgeJoin holds the string denoting the join edge name in mutations.
	EdgeJoin = "join"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeStudent holds the string denoting the student edge name in mutations.
	EdgeStudent = "student"
	// EdgeTutor holds the string denoting the tutor edge name in mutations.
	EdgeTutor = "tutor"
	// EdgeSParent holds the string denoting the sparent edge name in mutations.
	EdgeSParent = "SParent"
	// EdgeSTutor holds the string denoting the stutor edge name in mutations.
	EdgeSTutor = "STutor"
	// EdgeCanSee holds the string denoting the cansee edge name in mutations.
	EdgeCanSee = "canSee"
	// Table holds the table name of the entuser in the database.
	Table = "ent_users"
	// CourseTable is the table that holds the course relation/edge. The primary key declared below.
	CourseTable = "ent_user_course"
	// CourseInverseTable is the table name for the EntCourse entity.
	// It exists in this package in order to avoid circular dependency with the "entcourse" package.
	CourseInverseTable = "ent_courses"
	// TodoTable is the table that holds the todo relation/edge.
	TodoTable = "ent_todos"
	// TodoInverseTable is the table name for the EntTodo entity.
	// It exists in this package in order to avoid circular dependency with the "enttodo" package.
	TodoInverseTable = "ent_todos"
	// TodoColumn is the table column denoting the todo relation/edge.
	TodoColumn = "ent_user_todo"
	// AttendanceTable is the table that holds the attendance relation/edge.
	AttendanceTable = "ent_attendances"
	// AttendanceInverseTable is the table name for the EntAttendance entity.
	// It exists in this package in order to avoid circular dependency with the "entattendance" package.
	AttendanceInverseTable = "ent_attendances"
	// AttendanceColumn is the table column denoting the attendance relation/edge.
	AttendanceColumn = "ent_user_attendance"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "ent_posts"
	// PostInverseTable is the table name for the EntPost entity.
	// It exists in this package in order to avoid circular dependency with the "entpost" package.
	PostInverseTable = "ent_posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "ent_user_post"
	// CommentTable is the table that holds the comment relation/edge.
	CommentTable = "ent_comments"
	// CommentInverseTable is the table name for the EntComment entity.
	// It exists in this package in order to avoid circular dependency with the "entcomment" package.
	CommentInverseTable = "ent_comments"
	// CommentColumn is the table column denoting the comment relation/edge.
	CommentColumn = "ent_user_comment"
	// JoinTable is the table that holds the join relation/edge. The primary key declared below.
	JoinTable = "ent_user_join"
	// JoinInverseTable is the table name for the EntCourse entity.
	// It exists in this package in order to avoid circular dependency with the "entcourse" package.
	JoinInverseTable = "ent_courses"
	// ChildrenTable is the table that holds the children relation/edge. The primary key declared below.
	ChildrenTable = "ent_user_parent"
	// ParentTable is the table that holds the parent relation/edge. The primary key declared below.
	ParentTable = "ent_user_parent"
	// StudentTable is the table that holds the student relation/edge. The primary key declared below.
	StudentTable = "ent_user_tutor"
	// TutorTable is the table that holds the tutor relation/edge. The primary key declared below.
	TutorTable = "ent_user_tutor"
	// SParentTable is the table that holds the SParent relation/edge. The primary key declared below.
	SParentTable = "ent_user_STutor"
	// STutorTable is the table that holds the STutor relation/edge. The primary key declared below.
	STutorTable = "ent_user_STutor"
	// CanSeeTable is the table that holds the canSee relation/edge. The primary key declared below.
	CanSeeTable = "ent_post_shareWith"
	// CanSeeInverseTable is the table name for the EntPost entity.
	// It exists in this package in order to avoid circular dependency with the "entpost" package.
	CanSeeInverseTable = "ent_posts"
)

// Columns holds all SQL columns for entuser fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldFirstName,
	FieldLastName,
	FieldEmail,
	FieldPassword,
	FieldCountry,
	FieldPhone,
	FieldDateOfBirth,
	FieldProfilePictureUrl,
	FieldIsTutor,
	FieldIsStudent,
	FieldIsParent,
}

var (
	// CoursePrimaryKey and CourseColumn2 are the table columns denoting the
	// primary key for the course relation (M2M).
	CoursePrimaryKey = []string{"ent_user_id", "ent_course_id"}
	// JoinPrimaryKey and JoinColumn2 are the table columns denoting the
	// primary key for the join relation (M2M).
	JoinPrimaryKey = []string{"ent_user_id", "ent_course_id"}
	// ChildrenPrimaryKey and ChildrenColumn2 are the table columns denoting the
	// primary key for the children relation (M2M).
	ChildrenPrimaryKey = []string{"ent_user_id", "child_id"}
	// ParentPrimaryKey and ParentColumn2 are the table columns denoting the
	// primary key for the parent relation (M2M).
	ParentPrimaryKey = []string{"ent_user_id", "child_id"}
	// StudentPrimaryKey and StudentColumn2 are the table columns denoting the
	// primary key for the student relation (M2M).
	StudentPrimaryKey = []string{"ent_user_id", "student_id"}
	// TutorPrimaryKey and TutorColumn2 are the table columns denoting the
	// primary key for the tutor relation (M2M).
	TutorPrimaryKey = []string{"ent_user_id", "student_id"}
	// SParentPrimaryKey and SParentColumn2 are the table columns denoting the
	// primary key for the SParent relation (M2M).
	SParentPrimaryKey = []string{"ent_user_id", "SParent_id"}
	// STutorPrimaryKey and STutorColumn2 are the table columns denoting the
	// primary key for the STutor relation (M2M).
	STutorPrimaryKey = []string{"ent_user_id", "SParent_id"}
	// CanSeePrimaryKey and CanSeeColumn2 are the table columns denoting the
	// primary key for the canSee relation (M2M).
	CanSeePrimaryKey = []string{"ent_post_id", "ent_user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// FirstNameValidator is a validator for the "firstName" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "lastName" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultIsTutor holds the default value on creation for the "isTutor" field.
	DefaultIsTutor bool
	// DefaultIsStudent holds the default value on creation for the "isStudent" field.
	DefaultIsStudent bool
	// DefaultIsParent holds the default value on creation for the "isParent" field.
	DefaultIsParent bool
)
