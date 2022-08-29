package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"backend/ent"
	"backend/ent/entuser"
	"backend/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:root@tcp(0.0.0.0:3307)/tutorkit?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()
	CreateUser(ctx, client)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.EntUser, error) {
	u1, err := client.EntUser.
		Create().
		SetName("Tutor").
		SetEmail("tutor@gmail.com").
		SetPassword("123123").
		SetPhone("111-111-1111").
		SetProfilePictureUrl("").
		SetIsTutor(true).
		SetIsStudent(false).
		SetIsParent(false).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u1)

	u2, err := client.EntUser.
		Create().
		SetName("Student").
		SetEmail("student@gmail.com").
		SetPassword("123123").
		SetPhone("111-111-1111").
		SetProfilePictureUrl("").
		SetIsTutor(false).
		SetIsStudent(true).
		SetIsParent(false).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u2)
	u3, err := client.EntUser.
		Create().
		SetName("Parent").
		SetEmail("parent@gmail.com").
		SetPassword("123123").
		SetPhone("111-111-1111").
		SetProfilePictureUrl("").
		SetIsTutor(false).
		SetIsStudent(false).
		SetIsParent(true).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u3)
	c1, err := client.EntCourse.
		Create().
		SetName("Course1").
		SetCourseUrl("").
		SetPaymentMethod("hours").
		SetPaymentAmount(100).
		SetStartDate(time.Date(2022, 8, 28, 14, 30, 45, 100, time.Local)).
		SetEndDate(time.Date(2022, 9, 28, 14, 30, 45, 100, time.Local)).
		SetMonday(true).
		SetTuesday(false).
		SetWednesday(false).
		SetThursday(false).
		SetFriday(false).
		SetSaturday(false).
		SetSunday(false).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("course was created: ", c1)

	if err := u2.Update().AddParent(u3).Exec(ctx); err != nil {
		log.Fatalf("failed connecting student to its parent: %v", err)
	}
	if err := u2.Update().AddTutor(u1).Exec(ctx); err != nil {
		log.Fatalf("failed connecting student to its tutor: %v", err)
	}
	if err := u3.Update().AddSTutor(u1).Exec(ctx); err != nil {
		log.Fatalf("failed connecting tutor to student's parent: %v", err)
	}
	if err := u1.Update().AddCourse(c1).Exec(ctx); err != nil {
		log.Fatalf("failed connecting tutor to course: %v", err)
	}
	// Query all todo items that depend on other items.
	items, err := client.EntUser.Query().Where(entuser.HasChildren()).All(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Name)
	}

	// Query all todo items that depend on other items.
	items2, err := client.EntUser.Query().Where(entuser.HasStudent()).All(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	for _, t := range items2 {
		fmt.Printf("%d: %q\n", t.ID, t.Name)
	}
	// Query all todo items that depend on other items.
	items3, err := client.EntUser.Query().Where(entuser.HasSParent()).All(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	for _, t := range items3 {
		fmt.Printf("%d: %q\n", t.ID, t.Name)
	}
	items4, err := client.EntUser.Query().Where(entuser.HasCourse()).All(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	for _, t := range items4 {
		fmt.Printf("%d: %q\n", t.ID, t.Name)
	}
	return nil, nil
}
