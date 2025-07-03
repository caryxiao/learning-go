package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	PostTotal int
}

type Post struct {
	gorm.Model
	UserID        uint
	User          User   `gorm:"foreignKey:UserID"`
	Title         string `gorm:"type:varchar(255)"`
	Body          string
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentTotal  uint
	CommentStatus string `gorm:"type:varchar(255);default:'No Comment'"`
}

type Comment struct {
	gorm.Model
	PostID uint
	Post   Post `gorm:"foreignKey:PostID"`
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
	Body   string
}

// AfterCreate 创建文章后增加用户文章总数的计数
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&User{}).Where("id=?", p.UserID).Update("post_total", gorm.Expr("post_total + 1"))
	return
}

// AfterDelete 删除文章后减少用户文章总数的计数
func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	tx.Model(&User{}).Where("id=?", p.UserID).Update("post_total", gorm.Expr("post_total - 1"))
	return
}

// AfterCreate 添加评论后增加文章评论的计数
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	var post Post
	tx.First(&post, c.PostID)
	post.CommentTotal += 1
	if post.CommentStatus == "No Comment" {
		post.CommentStatus = "Has Comment"
	}
	tx.Save(&post)
	return
}

// AfterDelete 删除文章评论后减少文章的评论计数
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	tx.First(&post, c.PostID)
	post.CommentTotal -= 1
	if post.CommentTotal == 1 {
		post.CommentStatus = "No Comment"
	}
	tx.Save(&post)
	return
}

func initData(db *gorm.DB) {
	// 创建相关表
	err := db.AutoMigrate(&User{})
	err = db.AutoMigrate(&Post{})
	err = db.AutoMigrate(&Comment{})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/leaning-go?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	initData(db)
	var user User
	db.FirstOrCreate(&user, User{
		Name: "Cary",
	})

	var commentUser User
	db.FirstOrCreate(&commentUser, User{
		Name: "Tom",
	})

	post1 := Post{
		Title: "Post 1",
		Body:  "Hello World!",
	}
	err = db.Model(&user).Association("Posts").Append(&post1)

	if err != nil {
		log.Fatal(err)
	}

	// 添加多个评论
	comment1 := Comment{
		UserID: commentUser.ID,
		Body:   "Comment 1",
	}
	comment2 := Comment{
		UserID: commentUser.ID,
		Body:   "Comment 2",
	}
	comment3 := Comment{
		UserID: commentUser.ID,
		Body:   "Comment 3",
	}

	err = db.Model(&post1).Association("Comments").Append([]*Comment{&comment1, &comment2, &comment3})
	if err != nil {
		log.Fatal(err)
	}

	db.First(&post1, "id = ?", post1.ID)
	log.Printf("Post ID %d, Comment Total %d, Comment Status: %s", post1.ID, post1.CommentTotal, post1.CommentStatus)

	// 删除第二个评论
	db.Delete(&comment2)
	db.First(&post1, "id = ?", post1.ID)
	log.Printf("Post ID %d, Comment Total %d, Comment Status: %s", post1.ID, post1.CommentTotal, post1.CommentStatus)

	db.Delete(&comment1)
	db.Delete(&comment3)

	db.First(&post1, "id = ?", post1.ID)
	log.Printf("Post ID %d, Comment Total %d, Comment Status: %s", post1.ID, post1.CommentTotal, post1.CommentStatus)
}
