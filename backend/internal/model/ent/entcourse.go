// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/internal/model/ent/entcourse"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// EntCourse is the model entity for the EntCourse schema.
type EntCourse struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CourseUrl holds the value of the "courseUrl" field.
	CourseUrl string `json:"courseUrl,omitempty"`
	// PaymentMethod holds the value of the "paymentMethod" field.
	PaymentMethod entcourse.PaymentMethod `json:"paymentMethod,omitempty"`
	// PaymentAmount holds the value of the "paymentAmount" field.
	PaymentAmount float64 `json:"paymentAmount,omitempty"`
	// StartDate holds the value of the "startDate" field.
	StartDate time.Time `json:"startDate,omitempty"`
	// EndDate holds the value of the "endDate" field.
	EndDate time.Time `json:"endDate,omitempty"`
	// Monday holds the value of the "monday" field.
	Monday bool `json:"monday,omitempty"`
	// Tuesday holds the value of the "tuesday" field.
	Tuesday bool `json:"tuesday,omitempty"`
	// Wednesday holds the value of the "wednesday" field.
	Wednesday bool `json:"wednesday,omitempty"`
	// Thursday holds the value of the "thursday" field.
	Thursday bool `json:"thursday,omitempty"`
	// Friday holds the value of the "friday" field.
	Friday bool `json:"friday,omitempty"`
	// Saturday holds the value of the "saturday" field.
	Saturday bool `json:"saturday,omitempty"`
	// Sunday holds the value of the "sunday" field.
	Sunday bool `json:"sunday,omitempty"`
	// MondayStartTime holds the value of the "mondayStartTime" field.
	MondayStartTime time.Time `json:"mondayStartTime,omitempty"`
	// MondayEndTime holds the value of the "mondayEndTime" field.
	MondayEndTime time.Time `json:"mondayEndTime,omitempty"`
	// TuesdayStartTime holds the value of the "tuesdayStartTime" field.
	TuesdayStartTime time.Time `json:"tuesdayStartTime,omitempty"`
	// TuesdayEndTime holds the value of the "tuesdayEndTime" field.
	TuesdayEndTime time.Time `json:"tuesdayEndTime,omitempty"`
	// WednesdayStartTime holds the value of the "wednesdayStartTime" field.
	WednesdayStartTime time.Time `json:"wednesdayStartTime,omitempty"`
	// WednesdayEndTime holds the value of the "wednesdayEndTime" field.
	WednesdayEndTime time.Time `json:"wednesdayEndTime,omitempty"`
	// ThursdayStartTime holds the value of the "thursdayStartTime" field.
	ThursdayStartTime time.Time `json:"thursdayStartTime,omitempty"`
	// ThursdayEndTime holds the value of the "thursdayEndTime" field.
	ThursdayEndTime time.Time `json:"thursdayEndTime,omitempty"`
	// FridayStartTime holds the value of the "fridayStartTime" field.
	FridayStartTime time.Time `json:"fridayStartTime,omitempty"`
	// FridayEndTime holds the value of the "fridayEndTime" field.
	FridayEndTime time.Time `json:"fridayEndTime,omitempty"`
	// SaturdayStartTime holds the value of the "saturdayStartTime" field.
	SaturdayStartTime time.Time `json:"saturdayStartTime,omitempty"`
	// SaturdayEndTime holds the value of the "saturdayEndTime" field.
	SaturdayEndTime time.Time `json:"saturdayEndTime,omitempty"`
	// SundayStartTime holds the value of the "sundayStartTime" field.
	SundayStartTime time.Time `json:"sundayStartTime,omitempty"`
	// SundayEndTime holds the value of the "sundayEndTime" field.
	SundayEndTime time.Time `json:"sundayEndTime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntCourseQuery when eager-loading is set.
	Edges EntCourseEdges `json:"edges"`
}

// EntCourseEdges holds the relations/edges for other nodes in the graph.
type EntCourseEdges struct {
	// Todo holds the value of the todo edge.
	Todo []*EntTodo `json:"todo,omitempty"`
	// Attendance holds the value of the attendance edge.
	Attendance []*EntAttendance `json:"attendance,omitempty"`
	// Post holds the value of the post edge.
	Post []*EntPost `json:"post,omitempty"`
	// OwnedBy holds the value of the ownedBy edge.
	OwnedBy []*EntUser `json:"ownedBy,omitempty"`
	// JoinedBy holds the value of the joinedBy edge.
	JoinedBy []*EntUser `json:"joinedBy,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// TodoOrErr returns the Todo value or an error if the edge
// was not loaded in eager-loading.
func (e EntCourseEdges) TodoOrErr() ([]*EntTodo, error) {
	if e.loadedTypes[0] {
		return e.Todo, nil
	}
	return nil, &NotLoadedError{edge: "todo"}
}

// AttendanceOrErr returns the Attendance value or an error if the edge
// was not loaded in eager-loading.
func (e EntCourseEdges) AttendanceOrErr() ([]*EntAttendance, error) {
	if e.loadedTypes[1] {
		return e.Attendance, nil
	}
	return nil, &NotLoadedError{edge: "attendance"}
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading.
func (e EntCourseEdges) PostOrErr() ([]*EntPost, error) {
	if e.loadedTypes[2] {
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// OwnedByOrErr returns the OwnedBy value or an error if the edge
// was not loaded in eager-loading.
func (e EntCourseEdges) OwnedByOrErr() ([]*EntUser, error) {
	if e.loadedTypes[3] {
		return e.OwnedBy, nil
	}
	return nil, &NotLoadedError{edge: "ownedBy"}
}

// JoinedByOrErr returns the JoinedBy value or an error if the edge
// was not loaded in eager-loading.
func (e EntCourseEdges) JoinedByOrErr() ([]*EntUser, error) {
	if e.loadedTypes[4] {
		return e.JoinedBy, nil
	}
	return nil, &NotLoadedError{edge: "joinedBy"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntCourse) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case entcourse.FieldMonday, entcourse.FieldTuesday, entcourse.FieldWednesday, entcourse.FieldThursday, entcourse.FieldFriday, entcourse.FieldSaturday, entcourse.FieldSunday:
			values[i] = new(sql.NullBool)
		case entcourse.FieldPaymentAmount:
			values[i] = new(sql.NullFloat64)
		case entcourse.FieldID:
			values[i] = new(sql.NullInt64)
		case entcourse.FieldName, entcourse.FieldCourseUrl, entcourse.FieldPaymentMethod:
			values[i] = new(sql.NullString)
		case entcourse.FieldCreatedAt, entcourse.FieldUpdatedAt, entcourse.FieldStartDate, entcourse.FieldEndDate, entcourse.FieldMondayStartTime, entcourse.FieldMondayEndTime, entcourse.FieldTuesdayStartTime, entcourse.FieldTuesdayEndTime, entcourse.FieldWednesdayStartTime, entcourse.FieldWednesdayEndTime, entcourse.FieldThursdayStartTime, entcourse.FieldThursdayEndTime, entcourse.FieldFridayStartTime, entcourse.FieldFridayEndTime, entcourse.FieldSaturdayStartTime, entcourse.FieldSaturdayEndTime, entcourse.FieldSundayStartTime, entcourse.FieldSundayEndTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EntCourse", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntCourse fields.
func (ec *EntCourse) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entcourse.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = int(value.Int64)
		case entcourse.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				ec.CreatedAt = value.Time
			}
		case entcourse.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				ec.UpdatedAt = value.Time
			}
		case entcourse.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ec.Name = value.String
			}
		case entcourse.FieldCourseUrl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field courseUrl", values[i])
			} else if value.Valid {
				ec.CourseUrl = value.String
			}
		case entcourse.FieldPaymentMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field paymentMethod", values[i])
			} else if value.Valid {
				ec.PaymentMethod = entcourse.PaymentMethod(value.String)
			}
		case entcourse.FieldPaymentAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field paymentAmount", values[i])
			} else if value.Valid {
				ec.PaymentAmount = value.Float64
			}
		case entcourse.FieldStartDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field startDate", values[i])
			} else if value.Valid {
				ec.StartDate = value.Time
			}
		case entcourse.FieldEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field endDate", values[i])
			} else if value.Valid {
				ec.EndDate = value.Time
			}
		case entcourse.FieldMonday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field monday", values[i])
			} else if value.Valid {
				ec.Monday = value.Bool
			}
		case entcourse.FieldTuesday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field tuesday", values[i])
			} else if value.Valid {
				ec.Tuesday = value.Bool
			}
		case entcourse.FieldWednesday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field wednesday", values[i])
			} else if value.Valid {
				ec.Wednesday = value.Bool
			}
		case entcourse.FieldThursday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field thursday", values[i])
			} else if value.Valid {
				ec.Thursday = value.Bool
			}
		case entcourse.FieldFriday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field friday", values[i])
			} else if value.Valid {
				ec.Friday = value.Bool
			}
		case entcourse.FieldSaturday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field saturday", values[i])
			} else if value.Valid {
				ec.Saturday = value.Bool
			}
		case entcourse.FieldSunday:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field sunday", values[i])
			} else if value.Valid {
				ec.Sunday = value.Bool
			}
		case entcourse.FieldMondayStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mondayStartTime", values[i])
			} else if value.Valid {
				ec.MondayStartTime = value.Time
			}
		case entcourse.FieldMondayEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mondayEndTime", values[i])
			} else if value.Valid {
				ec.MondayEndTime = value.Time
			}
		case entcourse.FieldTuesdayStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field tuesdayStartTime", values[i])
			} else if value.Valid {
				ec.TuesdayStartTime = value.Time
			}
		case entcourse.FieldTuesdayEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field tuesdayEndTime", values[i])
			} else if value.Valid {
				ec.TuesdayEndTime = value.Time
			}
		case entcourse.FieldWednesdayStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field wednesdayStartTime", values[i])
			} else if value.Valid {
				ec.WednesdayStartTime = value.Time
			}
		case entcourse.FieldWednesdayEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field wednesdayEndTime", values[i])
			} else if value.Valid {
				ec.WednesdayEndTime = value.Time
			}
		case entcourse.FieldThursdayStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field thursdayStartTime", values[i])
			} else if value.Valid {
				ec.ThursdayStartTime = value.Time
			}
		case entcourse.FieldThursdayEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field thursdayEndTime", values[i])
			} else if value.Valid {
				ec.ThursdayEndTime = value.Time
			}
		case entcourse.FieldFridayStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field fridayStartTime", values[i])
			} else if value.Valid {
				ec.FridayStartTime = value.Time
			}
		case entcourse.FieldFridayEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field fridayEndTime", values[i])
			} else if value.Valid {
				ec.FridayEndTime = value.Time
			}
		case entcourse.FieldSaturdayStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field saturdayStartTime", values[i])
			} else if value.Valid {
				ec.SaturdayStartTime = value.Time
			}
		case entcourse.FieldSaturdayEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field saturdayEndTime", values[i])
			} else if value.Valid {
				ec.SaturdayEndTime = value.Time
			}
		case entcourse.FieldSundayStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sundayStartTime", values[i])
			} else if value.Valid {
				ec.SundayStartTime = value.Time
			}
		case entcourse.FieldSundayEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sundayEndTime", values[i])
			} else if value.Valid {
				ec.SundayEndTime = value.Time
			}
		}
	}
	return nil
}

// QueryTodo queries the "todo" edge of the EntCourse entity.
func (ec *EntCourse) QueryTodo() *EntTodoQuery {
	return (&EntCourseClient{config: ec.config}).QueryTodo(ec)
}

// QueryAttendance queries the "attendance" edge of the EntCourse entity.
func (ec *EntCourse) QueryAttendance() *EntAttendanceQuery {
	return (&EntCourseClient{config: ec.config}).QueryAttendance(ec)
}

// QueryPost queries the "post" edge of the EntCourse entity.
func (ec *EntCourse) QueryPost() *EntPostQuery {
	return (&EntCourseClient{config: ec.config}).QueryPost(ec)
}

// QueryOwnedBy queries the "ownedBy" edge of the EntCourse entity.
func (ec *EntCourse) QueryOwnedBy() *EntUserQuery {
	return (&EntCourseClient{config: ec.config}).QueryOwnedBy(ec)
}

// QueryJoinedBy queries the "joinedBy" edge of the EntCourse entity.
func (ec *EntCourse) QueryJoinedBy() *EntUserQuery {
	return (&EntCourseClient{config: ec.config}).QueryJoinedBy(ec)
}

// Update returns a builder for updating this EntCourse.
// Note that you need to call EntCourse.Unwrap() before calling this method if this EntCourse
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *EntCourse) Update() *EntCourseUpdateOne {
	return (&EntCourseClient{config: ec.config}).UpdateOne(ec)
}

// Unwrap unwraps the EntCourse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *EntCourse) Unwrap() *EntCourse {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntCourse is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *EntCourse) String() string {
	var builder strings.Builder
	builder.WriteString("EntCourse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ec.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(ec.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(ec.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ec.Name)
	builder.WriteString(", ")
	builder.WriteString("courseUrl=")
	builder.WriteString(ec.CourseUrl)
	builder.WriteString(", ")
	builder.WriteString("paymentMethod=")
	builder.WriteString(fmt.Sprintf("%v", ec.PaymentMethod))
	builder.WriteString(", ")
	builder.WriteString("paymentAmount=")
	builder.WriteString(fmt.Sprintf("%v", ec.PaymentAmount))
	builder.WriteString(", ")
	builder.WriteString("startDate=")
	builder.WriteString(ec.StartDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("endDate=")
	builder.WriteString(ec.EndDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("monday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Monday))
	builder.WriteString(", ")
	builder.WriteString("tuesday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Tuesday))
	builder.WriteString(", ")
	builder.WriteString("wednesday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Wednesday))
	builder.WriteString(", ")
	builder.WriteString("thursday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Thursday))
	builder.WriteString(", ")
	builder.WriteString("friday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Friday))
	builder.WriteString(", ")
	builder.WriteString("saturday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Saturday))
	builder.WriteString(", ")
	builder.WriteString("sunday=")
	builder.WriteString(fmt.Sprintf("%v", ec.Sunday))
	builder.WriteString(", ")
	builder.WriteString("mondayStartTime=")
	builder.WriteString(ec.MondayStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mondayEndTime=")
	builder.WriteString(ec.MondayEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tuesdayStartTime=")
	builder.WriteString(ec.TuesdayStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tuesdayEndTime=")
	builder.WriteString(ec.TuesdayEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("wednesdayStartTime=")
	builder.WriteString(ec.WednesdayStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("wednesdayEndTime=")
	builder.WriteString(ec.WednesdayEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("thursdayStartTime=")
	builder.WriteString(ec.ThursdayStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("thursdayEndTime=")
	builder.WriteString(ec.ThursdayEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("fridayStartTime=")
	builder.WriteString(ec.FridayStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("fridayEndTime=")
	builder.WriteString(ec.FridayEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("saturdayStartTime=")
	builder.WriteString(ec.SaturdayStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("saturdayEndTime=")
	builder.WriteString(ec.SaturdayEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sundayStartTime=")
	builder.WriteString(ec.SundayStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sundayEndTime=")
	builder.WriteString(ec.SundayEndTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// EntCourses is a parsable slice of EntCourse.
type EntCourses []*EntCourse

func (ec EntCourses) config(cfg config) {
	for _i := range ec {
		ec[_i].config = cfg
	}
}
