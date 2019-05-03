package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
}

func (this User) ToString() string {
	return "Name:" + this.Name + ";Email" + this.Email + ";Age:" + string(this.Age)
}

type Store struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Customer struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	//1 注册 model
	orm.RegisterModel(new(Store), new(Customer), new(User))
	//2 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//3 注册默认数据库
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")

	fmt.Print(beego.AppConfig.String("db::envRar"))

	dbcon := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	fmt.Print(dbcon)
	orm.RegisterDataBase("default", "mysql", dbcon /*"root:root@tcp(localhost:3306)/test?charset=utf8"*/) //密码为空格式
}

func CreateUsers(users []User) {

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	o := orm.NewOrm()
	if successNums, err := o.InsertMulti(len(users), users); err != nil {
		fmt.Println("insert fail.....")
	} else {
		fmt.Println("success inserted %d datas", successNums)
	}
}

func CreateUser(user *User) {
	o := orm.NewOrm()
	if successNums, err := o.Insert(user); err != nil {
		fmt.Println("insert fail.....")
	} else {
		fmt.Println("success inserted %d datas", successNums)
	}
}

//find all users
func ListUsers() {
	var users []User
	orm.NewOrm().QueryTable("t_user").All(&users)
	for _, user := range users {
		fmt.Println(user.ToString())
	}
}

//get user count
func CountUser() {
	cnt, _ := orm.NewOrm().QueryTable("t_user").Count()
	fmt.Println("All user count:", cnt)
}

//get the only one user
func GetUser() {
	var user User
	err := orm.NewOrm().QueryTable("t_user").Filter("Id", 5).One(&user)
	if err == nil {
		fmt.Println(user.ToString())
	}
}

//get the only one user
func GetUsers() {
	var users []User
	_, err := orm.NewOrm().QueryTable("t_user").Filter("name__contains", "awd").All(&users)
	if err == nil {
		for _, user := range users {
			fmt.Println(user.ToString())
		}
	}
}

//get limit user
func LimitUser() {
	var users []User
	_, err := orm.NewOrm().QueryTable("t_user").Limit(6).OrderBy("-Name").All(&users)
	if err == nil {
		for _, user := range users {
			fmt.Println(user.ToString())
		}
	}
}

//get user limit,offset
func LimitoffsetUser() {
	var users []User
	_, err := orm.NewOrm().QueryTable("t_user").Limit(1, 4).OrderBy("Id").All(&users)
	if err == nil {
		for _, user := range users {
			fmt.Println(user.ToString())
		}
	}
}

//del user
func DelUser(id int) {
	num, err := orm.NewOrm().QueryTable("t_user").Filter("Id", id).Delete()
	fmt.Printf("Affected Num: %s, %s", num, err)
}

//update user
func UpdateUser() {
	num, err := orm.NewOrm().QueryTable("t_user").Filter("name__contains", "awd").Update(orm.Params{
		"name": "#########",
	})
	fmt.Printf("Affected Num: %s, %s", num, err)
}
