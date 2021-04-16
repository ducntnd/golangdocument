package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	//"github.com/kataras/iris/v12"
)

func main()  {
	//app := iris.New()

	ExampleDB_Model()
}

type User struct {
	//tableName struct{} `pg:",discard_unknown_columns"`

	Id     int64
	Name   string
	Emails []string
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User `pg:"rel:has-one"`
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func ExampleDB_Model() {
	db := pg.Connect(&pg.Options{
		Addr: ":5432",
		User: "",
		Password: "1234",
		Database: "postgres",
	})

	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	user1 := &User{
		Name: "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	_, err = db.Model(user1).Insert()
	if err != nil{
		panic(err)
	}

	_,err = db.Model(&User{
		Name:   "root",
		Emails: []string{"root1@root", "root2@root"},
	}).Insert()
	if err != nil{
		panic(err)
	}

	story1 := &Story{
		Title:    "Cool story",
		AuthorId: user1.Id,
	}
	_, err = db.Model(story1).Insert()
	if err != nil {
		panic(err)
	}

	// Select user by primary key.
	user := &User{Id: user1.Id}
	err = db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}

	// Select all users.
	var users []User
	err = db.Model(&users).Select()
	if err != nil {
		panic(err)
	}

	// Select story and associated author in one query.
	story := new(Story)
	err = db.Model(story).
		Relation("Author").
		Where("story.id = ?", story1.Id).
		Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
	fmt.Println(users)
	fmt.Println(story)

}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Story)(nil),
	}


	for _,model := range models{
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil{
			return err
		}
	}
	return nil
}
