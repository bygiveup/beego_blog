package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	
)
//category 分类
//Title 名称
//created创建时间
//views  浏览次数
//`orm:"index"   索引
//TopicTime  最后文章发表时间
//Topictime  多少篇文章
//TopicLastUserId 最后操作的用户id
type Category struct{
	Id		int64
	Title	string
	Created	time.Time `orm:"index"`
	Views	int64 `orm:"index"`
	TopicTime	time.Time  `orm:"index"`
	TopicCount	int64
	TopicLastUserId int64
}

//Uid  谁写的
//Content  内容
//Attachment  附件
//Updated 最后更新
//Author  作者信息
//ReplyTime  最后回复时间
//ReplyCount 最后评论个数
//ReplyLastUserId  评论id
type Topic struct{
	Id  int64
	Uid int64
	Title	string
	Content  string  `orm:"size(5000)"`
	Attachment	string
	Created	time.Time `orm:"index"`
	Updated	time.Time `orm:"index"`
	Views	int64 `orm:"index"`
	Author	string
	ReplyTime	time.Time  `orm:"index"`
	ReplyCount	int64
	ReplyLastUserId	int64
}

func init() {
	maxIdle := 30   //设置最大空闲连接
	maxConn := 30  //设置最大数据库连接 
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)
	
    // 需要在init中注册定义的model
    orm.RegisterModel(new(Category), new(Topic))
}

//分类
func AddCategory(name string) error{
	o := orm.NewOrm()
	cate := &Category{Title: name}
	//查询
	qs := o.QueryTable("category") 
	err := qs.Filter("title",name).One(cate)
	if err == nil{
		return err
	}
	_,err = o.Insert(cate)
	if err != nil{
		return err
	}
	return nil
}

//分类获取所有分类目录
func GetAllCategories()([]*Category,error){
	//获取orm对象
	o := orm.NewOrm()

	cates := make([]*Category,0)
	qs := o.QueryTable("category")
	_,err := qs.All(&cates)
	return cates,err
}
