// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/internal/ent/entuser"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// EntUser is the model entity for the EntUser schema.
type EntUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// FirstName holds the value of the "firstName" field.
	FirstName string `json:"firstName,omitempty"`
	// LastName holds the value of the "lastName" field.
	LastName string `json:"lastName,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// DateOfBirth holds the value of the "dateOfBirth" field.
	DateOfBirth time.Time `json:"dateOfBirth,omitempty"`
	// ProfilePictureUrl holds the value of the "profilePictureUrl" field.
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
	// IsTutor holds the value of the "isTutor" field.
	IsTutor bool `json:"isTutor,omitempty"`
	// IsStudent holds the value of the "isStudent" field.
	IsStudent bool `json:"isStudent,omitempty"`
	// IsParent holds the value of the "isParent" field.
	IsParent bool `json:"isParent,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntUserQuery when eager-loading is set.
	Edges EntUserEdges `json:"edges"`
}

// EntUserEdges holds the relations/edges for other nodes in the graph.
type EntUserEdges struct {
	// Course holds the value of the course edge.
	Course []*EntCourse `json:"course,omitempty"`
	// Todo holds the value of the todo edge.
	Todo []*EntTodo `json:"todo,omitempty"`
	// Attendance holds the value of the attendance edge.
	Attendance []*EntAttendance `json:"attendance,omitempty"`
	// Post holds the value of the post edge.
	Post []*EntPost `json:"post,omitempty"`
	// Comment holds the value of the comment edge.
	Comment []*EntComment `json:"comment,omitempty"`
	// Join holds the value of the join edge.
	Join []*EntCourse `json:"join,omitempty"`
	// Children holds the value of the children edge.
	Children []*EntUser `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent []*EntUser `json:"parent,omitempty"`
	// Student holds the value of the student edge.
	Student []*EntUser `json:"student,omitempty"`
	// Tutor holds the value of the tutor edge.
	Tutor []*EntUser `json:"tutor,omitempty"`
	// SParent holds the value of the SParent edge.
	SParent []*EntUser `json:"SParent,omitempty"`
	// STutor holds the value of the STutor edge.
	STutor []*EntUser `json:"STutor,omitempty"`
	// CanSee holds the value of the canSee edge.
	CanSee []*EntPost `json:"canSee,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [13]bool
}

// CourseOrErr returns the Course value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) CourseOrErr() ([]*EntCourse, error) {
	if e.loadedTypes[0] {
		return e.Course, nil
	}
	return nil, &NotLoadedError{edge: "course"}
}

// TodoOrErr returns the Todo value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) TodoOrErr() ([]*EntTodo, error) {
	if e.loadedTypes[1] {
		return e.Todo, nil
	}
	return nil, &NotLoadedError{edge: "todo"}
}

// AttendanceOrErr returns the Attendance value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) AttendanceOrErr() ([]*EntAttendance, error) {
	if e.loadedTypes[2] {
		return e.Attendance, nil
	}
	return nil, &NotLoadedError{edge: "attendance"}
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) PostOrErr() ([]*EntPost, error) {
	if e.loadedTypes[3] {
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) CommentOrErr() ([]*EntComment, error) {
	if e.loadedTypes[4] {
		return e.Comment, nil
	}
	return nil, &NotLoadedError{edge: "comment"}
}

// JoinOrErr returns the Join value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) JoinOrErr() ([]*EntCourse, error) {
	if e.loadedTypes[5] {
		return e.Join, nil
	}
	return nil, &NotLoadedError{edge: "join"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) ChildrenOrErr() ([]*EntUser, error) {
	if e.loadedTypes[6] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) ParentOrErr() ([]*EntUser, error) {
	if e.loadedTypes[7] {
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) StudentOrErr() ([]*EntUser, error) {
	if e.loadedTypes[8] {
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// TutorOrErr returns the Tutor value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) TutorOrErr() ([]*EntUser, error) {
	if e.loadedTypes[9] {
		return e.Tutor, nil
	}
	return nil, &NotLoadedError{edge: "tutor"}
}

// SParentOrErr returns the SParent value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) SParentOrErr() ([]*EntUser, error) {
	if e.loadedTypes[10] {
		return e.SParent, nil
	}
	return nil, &NotLoadedError{edge: "SParent"}
}

// STutorOrErr returns the STutor value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) STutorOrErr() ([]*EntUser, error) {
	if e.loadedTypes[11] {
		return e.STutor, nil
	}
	return nil, &NotLoadedError{edge: "STutor"}
}

// CanSeeOrErr returns the CanSee value or an error if the edge
// was not loaded in eager-loading.
func (e EntUserEdges) CanSeeOrErr() ([]*EntPost, error) {
	if e.loadedTypes[12] {
		return e.CanSee, nil
	}
	return nil, &NotLoadedError{edge: "canSee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case entuser.FieldIsTutor, entuser.FieldIsStudent, entuser.FieldIsParent:
			values[i] = new(sql.NullBool)
		case entuser.FieldID:
			values[i] = new(sql.NullInt64)
		case entuser.FieldFirstName, entuser.FieldLastName, entuser.FieldEmail, entuser.FieldPassword, entuser.FieldCountry, entuser.FieldPhone, entuser.FieldProfilePictureUrl:
			values[i] = new(sql.NullString)
		case entuser.FieldCreatedAt, entuser.FieldUpdatedAt, entuser.FieldDateOfBirth:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EntUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntUser fields.
func (eu *EntUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			eu.ID = int(value.Int64)
		case entuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				eu.CreatedAt = value.Time
			}
		case entuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				eu.UpdatedAt = value.Time
			}
		case entuser.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firstName", values[i])
			} else if value.Valid {
				eu.FirstName = value.String
			}
		case entuser.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastName", values[i])
			} else if value.Valid {
				eu.LastName = value.String
			}
		case entuser.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				eu.Email = value.String
			}
		case entuser.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				eu.Password = value.String
			}
		case entuser.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				eu.Country = value.String
			}
		case entuser.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				eu.Phone = value.String
			}
		case entuser.FieldDateOfBirth:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dateOfBirth", values[i])
			} else if value.Valid {
				eu.DateOfBirth = value.Time
			}
		case entuser.FieldProfilePictureUrl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profilePictureUrl", values[i])
			} else if value.Valid {
				eu.ProfilePictureUrl = value.String
			}
		case entuser.FieldIsTutor:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isTutor", values[i])
			} else if value.Valid {
				eu.IsTutor = value.Bool
			}
		case entuser.FieldIsStudent:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isStudent", values[i])
			} else if value.Valid {
				eu.IsStudent = value.Bool
			}
		case entuser.FieldIsParent:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isParent", values[i])
			} else if value.Valid {
				eu.IsParent = value.Bool
			}
		}
	}
	return nil
}

// QueryCourse queries the "course" edge of the EntUser entity.
func (eu *EntUser) QueryCourse() *EntCourseQuery {
	return (&EntUserClient{config: eu.config}).QueryCourse(eu)
}

// QueryTodo queries the "todo" edge of the EntUser entity.
func (eu *EntUser) QueryTodo() *EntTodoQuery {
	return (&EntUserClient{config: eu.config}).QueryTodo(eu)
}

// QueryAttendance queries the "attendance" edge of the EntUser entity.
func (eu *EntUser) QueryAttendance() *EntAttendanceQuery {
	return (&EntUserClient{config: eu.config}).QueryAttendance(eu)
}

// QueryPost queries the "post" edge of the EntUser entity.
func (eu *EntUser) QueryPost() *EntPostQuery {
	return (&EntUserClient{config: eu.config}).QueryPost(eu)
}

// QueryComment queries the "comment" edge of the EntUser entity.
func (eu *EntUser) QueryComment() *EntCommentQuery {
	return (&EntUserClient{config: eu.config}).QueryComment(eu)
}

// QueryJoin queries the "join" edge of the EntUser entity.
func (eu *EntUser) QueryJoin() *EntCourseQuery {
	return (&EntUserClient{config: eu.config}).QueryJoin(eu)
}

// QueryChildren queries the "children" edge of the EntUser entity.
func (eu *EntUser) QueryChildren() *EntUserQuery {
	return (&EntUserClient{config: eu.config}).QueryChildren(eu)
}

// QueryParent queries the "parent" edge of the EntUser entity.
func (eu *EntUser) QueryParent() *EntUserQuery {
	return (&EntUserClient{config: eu.config}).QueryParent(eu)
}

// QueryStudent queries the "student" edge of the EntUser entity.
func (eu *EntUser) QueryStudent() *EntUserQuery {
	return (&EntUserClient{config: eu.config}).QueryStudent(eu)
}

// QueryTutor queries the "tutor" edge of the EntUser entity.
func (eu *EntUser) QueryTutor() *EntUserQuery {
	return (&EntUserClient{config: eu.config}).QueryTutor(eu)
}

// QuerySParent queries the "SParent" edge of the EntUser entity.
func (eu *EntUser) QuerySParent() *EntUserQuery {
	return (&EntUserClient{config: eu.config}).QuerySParent(eu)
}

// QuerySTutor queries the "STutor" edge of the EntUser entity.
func (eu *EntUser) QuerySTutor() *EntUserQuery {
	return (&EntUserClient{config: eu.config}).QuerySTutor(eu)
}

// QueryCanSee queries the "canSee" edge of the EntUser entity.
func (eu *EntUser) QueryCanSee() *EntPostQuery {
	return (&EntUserClient{config: eu.config}).QueryCanSee(eu)
}

// Update returns a builder for updating this EntUser.
// Note that you need to call EntUser.Unwrap() before calling this method if this EntUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (eu *EntUser) Update() *EntUserUpdateOne {
	return (&EntUserClient{config: eu.config}).UpdateOne(eu)
}

// Unwrap unwraps the EntUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eu *EntUser) Unwrap() *EntUser {
	_tx, ok := eu.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntUser is not a transactional entity")
	}
	eu.config.driver = _tx.drv
	return eu
}

// String implements the fmt.Stringer.
func (eu *EntUser) String() string {
	var builder strings.Builder
	builder.WriteString("EntUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", eu.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(eu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(eu.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("firstName=")
	builder.WriteString(eu.FirstName)
	builder.WriteString(", ")
	builder.WriteString("lastName=")
	builder.WriteString(eu.LastName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(eu.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(eu.Password)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(eu.Country)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(eu.Phone)
	builder.WriteString(", ")
	builder.WriteString("dateOfBirth=")
	builder.WriteString(eu.DateOfBirth.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("profilePictureUrl=")
	builder.WriteString(eu.ProfilePictureUrl)
	builder.WriteString(", ")
	builder.WriteString("isTutor=")
	builder.WriteString(fmt.Sprintf("%v", eu.IsTutor))
	builder.WriteString(", ")
	builder.WriteString("isStudent=")
	builder.WriteString(fmt.Sprintf("%v", eu.IsStudent))
	builder.WriteString(", ")
	builder.WriteString("isParent=")
	builder.WriteString(fmt.Sprintf("%v", eu.IsParent))
	builder.WriteByte(')')
	return builder.String()
}

// EntUsers is a parsable slice of EntUser.
type EntUsers []*EntUser

func (eu EntUsers) config(cfg config) {
	for _i := range eu {
		eu[_i].config = cfg
	}
}
