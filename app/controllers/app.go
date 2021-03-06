package controllers

import (
	"github.com/revel/revel"
	"GBlog/app/models"
	"time"
	"fmt"
)
type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	//dao := models.NewDao(c.MongoSession)
	blogs := dao.FindBlogs()
	now := time.Now().Add(-1 * time.Hour)
	fmt.Println(now)
	recentCnt :=dao.FindBlogsByDate(now);
	fmt.Println(recentCnt)
	return c.Render(blogs,recentCnt)
}
func (c App) BlogInfor(id string,rcnt int) revel.Result {
	fmt.Println(id);
	fmt.Println(rcnt);
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	blog := dao.FindBlogById(id)
	fmt.Println(blog.ReadCnt);
	if(blog.ReadCnt==rcnt){
		blog.ReadCnt = rcnt+1
		dao.UpdateBlogById(id,blog)
	}
	comments := dao.FindCommentsByBlogId(blog.Id);
	if len(comments)==0&&blog.CommentCnt!=0{
		blog.CommentCnt=0;
		dao.UpdateBlogById(id,blog)
	}else if len(comments)!=blog.CommentCnt{
		blog.CommentCnt=len(comments);
		dao.UpdateBlogById(id,blog)
	}
	fmt.Println(comments);
	return c.Render(blog,rcnt,comments)
}
func (c App) WBlog() revel.Result {
	return c.Render()
}
func (c App) Message() revel.Result {
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	//dao := models.NewDao(c.MongoSession)
	messages := dao.FindAllMessages()
	return c.Render(messages)
}
func (c App) History() revel.Result {
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	dao.CreateAllHistory();
	historys := dao.FindHistory();
	for i,_ := range historys{
		historys[i].Blogs =dao.FindBlogsByYear(historys[i].Year);	
	}
	fmt.Println(historys)
	return c.Render(historys)
}
func (c App) About() revel.Result {
	return c.Render()
}
func (c App) Emails() revel.Result {
	dao, err := models.NewDao()
	if err != nil {
		c.Response.Status = 500
		return c.RenderError(err)
	}
	defer dao.Close()
	emails :=dao.FindAllEmails();
	return c.Render(emails)
}