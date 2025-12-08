// 进阶gorm
// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
//
//	Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
// 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        int
	UserName  string
	Post      []Post //一对多 一个用户可以发布多篇文章
	PostCount int    `gorm:"default:0"` // 文章数量统计字段
}

type Post struct {
	ID         int
	PostStr    string // post 内容
	PostStatus string // post 状态 无评论 有评论

	// 外键：这篇文章属于哪个用户
	UserID int  // POST 对应的 USER
	User   User // 关联到User模型

	//一对多 一篇文章可以有多个评论
	Comment []Comment
}

type Comment struct {
	ID         int
	CommentStr string //coment 内容

	// 外键：这个评论属于哪篇文章
	PostID int  // Comment 对应的 Post
	Post   Post // 关联到User模型
}

func find_user_post_and_comment(db *gorm.DB, UserName string) ([]Comment, []Post) {
	var user User
	// 	// 1. 结构体 vs 查询条件（你的困惑）
	// type User struct {
	//     UserName string    // 这里用 UserName
	// }
	// 数据库规范：大多数数据库使用 snake_case（蛇形命名）如 user_name
	// // 查询时却要用 user_name
	// db.Where("user_name = ?", "张三")  // 容易写错！
	// // 让 GORM 自动处理转换
	// db.Where(&User{UserName: "张三"}).First(&user)
	db.Where("user_name = ?", UserName).First(&user)

	fmt.Printf("user:%s \n", user.UserName)

	var user_post []Post
	db.Model(&user).Association("Post").Find(&user_post)

	var user_comment []Comment
	db.Model(&user_post).Association("Comment").Find(&user_comment)
	return user_comment, user_post
}

func find_best_post(db *gorm.DB) (Post, error) {
	var post Post
	var postID int

	// 使用子查询找到评论最多的post_id
	subquery := db.Model(&Comment{}).
		Select("post_id").      // Select("post_id") - 只选择post_id字段,我们只需要文章ID，不需要评论的其他信息
		Group("post_id").       // 按文章ID分组,因为我们要统计每篇文章的评论数，所以需要按文章分组 相当于SQL: GROUP BY post_id
		Order("COUNT(*) DESC"). // Order("COUNT(*) DESC") - 按评论数降序排序,COUNT(*) - 统计每组的评论数量,DESC - 降序（从大到小），评论最多的排最前面,相当于SQL: ORDER BY COUNT(*) DESC
		Limit(1)

	// 执行子查询，将结果扫描到 postID
	err := subquery.Scan(&postID).Error
	if err != nil {
		return Post{}, err
	}

	fmt.Printf("找到评论最多的文章ID: %d\n", postID)

	// 2. 用找到的post_id查询完整的Post信息
	err = db.Where("id = ?", postID).First(&post).Error

	return post, err
}

// 要求 ：
// BeforeCreate 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("BeforeCreate post content:", p.PostStr)
	// 可以在这里修改数据
	return nil
}

// AfterCreate - 创建后自动调用
func (p *Post) AfterCreate(tx *gorm.DB) error {
	fmt.Printf("文章创建成功，更新用户文章数量统计...\n")
	fmt.Println("AfterCreate post content:", p.PostStr)

	// 更新用户的文章数量
	result := tx.Model(&User{}).
		Where("id  = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1))
	if result.Error != nil {
		fmt.Printf("更新文章数量失败: %v\n", result.Error)
		return result.Error
	}

	fmt.Printf("用户(ID: %d)的文章数量已更新\n", p.UserID)
	// 可以在这里修改数据
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) BeforeDelete(tx *gorm.DB) error {
	fmt.Printf("评论删除之前，更新文章状态...\n")

	fmt.Printf("----BeforeDelete: postid=%d\n", c.PostID)
	// 保存重要数据到临时字段
	return nil
}

// AfterDelete - 删除后自动调用
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	fmt.Printf("评论删除成功，更新文章状态...\n")
	fmt.Printf("----AfterDelete Comment content: id=%d, commentstr=%s, postid=%d \n", c.ID, c.CommentStr, c.PostID)

	//获取当前评论数量
	var remainingComments int64

	tx.Model(&Post{}).
		Where("id = ?", c.PostID).
		Count(&remainingComments)

	fmt.Println("----remainingComments:", remainingComments)
	remainingComments--
	// 更新用户的评论状态
	if remainingComments == 0 {
		tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("post_status", "无评论")
	}

	fmt.Printf("用户(ID: %d)的文章数量已更新\n", c.PostID)
	// 可以在这里修改数据
	return nil
}

func main() {
	// 1. 连接SQLite数据库（blog.db文件）
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	// 3. 创建一些测试数据
	createBlogData(db)

	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var user_comments []Comment
	var user_posts []Post

	user_comments, user_posts = find_user_post_and_comment(db, "张三")
	for _, post := range user_posts {
		fmt.Printf("ID:%d UserId:%d, PostStr:%s\n", post.ID, post.UserID, post.PostStr)
	}

	for _, comment := range user_comments {
		fmt.Printf("----- ID:%d PostID:%d,CommentStr:%s\n", comment.ID, comment.PostID, comment.CommentStr)
	}

	best_post, err := find_best_post(db)
	if err != nil {
		panic("find_best_post 失败: " + err.Error())
	}
	fmt.Printf("----- find_best_post ID:%d PostID:%d,PostStr: %s\n", best_post.ID, best_post.UserID, best_post.PostStr)

	var users []User
	// BeforeCreate 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	db.Find(&users)
	for _, user := range users {
		fmt.Printf("用户的文章数量统计字段 ID:%d UserName:%s, PostCount:%d\n", user.ID, user.UserName, user.PostCount)
	}

	var before_delete_comments []Comment

	var after_delete_comments []Comment

	var before_delete_post []Post

	var after_delete_post []Post
	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

	db.Find(&before_delete_comments)
	for _, com := range before_delete_comments {
		fmt.Printf(" before_delete_comments ID:%d, PostID:%d\n", com.ID, com.PostID)
	}

	db.Find(&before_delete_post)
	for _, po := range before_delete_post {
		fmt.Printf("before_delete_post ID:%d, PostStatus:%s\n", po.ID, po.PostStatus)
	}

	comment := Comment{
		CommentStr: "写得很详细，对Go新手很有帮助！",
		ID:         1,
		PostID:     1,
	}
	db.Delete(&comment)
	//为什么GORM 批量操作删除无法触发回调？
	// GORM这样设计有几个原因：

	// 性能考虑：批量删除时，如果每条记录都触发回调，性能会很差

	// 一致性：回调中可能有业务逻辑，批量操作时很难保证一致性

	// 明确性：开发者需要明确知道何时会触发回调
	// db.Where("post_id = ?", 1).Delete(&Comment{})

	db.Find(&after_delete_comments)
	for _, com := range after_delete_comments {
		fmt.Printf("after_delete_comments ID:%d, PostID:%d\n", com.ID, com.PostID)
	}

	db.Find(&after_delete_post)
	for _, po := range after_delete_post {
		fmt.Printf("after_delete_post ID:%d, PostStatus:%s\n", po.ID, po.PostStatus)
	}

}

// 创建Blog测试数据
func createBlogData(db *gorm.DB) {
	var userCount int64
	db.Model(&User{}).Count(&userCount)

	// 如果已有用户数据，说明已初始化过
	if userCount > 0 {
		fmt.Println("数据库已有数据，跳过初始化")
		return
	}

	// 2. 自动创建表
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("创建表失败: " + err.Error())
	}

	fmt.Println("数据库表创建成功！")

	// 创建4个用户
	users := []User{
		{UserName: "张三"},
		{UserName: "李四"},
		{UserName: "王五"},
		{UserName: "赵六"},
	}
	db.Create(&users)

	fmt.Printf("创建了 %d 个用户\n", len(users))

	// 为每个用户创建文章
	posts := []Post{
		// 张三的文章（技术类）
		{PostStr: "Go语言入门：从Hello World到并发编程", UserID: users[0].ID},
		{PostStr: "GORM使用心得：如何高效操作数据库", UserID: users[0].ID},
		{PostStr: "Web开发实战：用Go构建RESTful API", UserID: users[0].ID},

		// 李四的文章（生活类）
		{PostStr: "周末爬山记：登顶凤凰山的快乐", UserID: users[1].ID},
		{PostStr: "美食分享：在家做正宗重庆火锅", UserID: users[1].ID},
		{PostStr: "电影推荐：今年最值得看的5部电影", UserID: users[1].ID},

		// 王五的文章（学习类）
		{PostStr: "英语学习技巧：如何快速提升词汇量", UserID: users[2].ID},
		{PostStr: "时间管理：高效工作者的7个习惯", UserID: users[2].ID},
		{PostStr: "读书笔记：《如何阅读一本书》精华总结", UserID: users[3].ID},

		// 赵六的文章（科技类）
		{PostStr: "人工智能现状与未来发展趋势", UserID: users[3].ID},
		{PostStr: "区块链技术：从原理到应用", UserID: users[3].ID},
		{PostStr: "云计算入门：AWS vs Azure vs GCP比较", UserID: users[3].ID},
	}
	db.Create(&posts)

	fmt.Printf("创建了 %d 篇文章\n", len(posts))

	// 为每篇文章创建评论
	comments := []Comment{
		// 对张三第1篇文章的评论
		{CommentStr: "写得很详细，对Go新手很有帮助！", PostID: posts[0].ID},
		{CommentStr: "期待更多的Go语言教程", PostID: posts[1].ID},
		{CommentStr: "代码示例很清晰，感谢分享", PostID: posts[1].ID},

		// 对张三第2篇文章的评论
		{CommentStr: "GORM确实比原生SQL方便多了", PostID: posts[1].ID},
		{CommentStr: "学到了关联查询的技巧，很有用", PostID: posts[1].ID},

		// 对张三第3篇文章的评论
		{CommentStr: "API设计得很规范，学习了", PostID: posts[2].ID},
		{CommentStr: "可以分享一下完整的项目代码吗？", PostID: posts[2].ID},

		// 对李四第1篇文章的评论
		{CommentStr: "凤凰山风景确实很美，我也去过", PostID: posts[3].ID},
		{CommentStr: "爬山要注意安全，带足水和食物", PostID: posts[3].ID},

		// 对李四第2篇文章的评论
		{CommentStr: "看着就流口水了，今晚就试试", PostID: posts[4].ID},
		{CommentStr: "火锅底料是自己做的吗？", PostID: posts[4].ID},
		{CommentStr: "配菜清单很详细，谢谢分享", PostID: posts[4].ID},

		// 对李四第3篇文章的评论
		{CommentStr: "第3部电影确实很感人", PostID: posts[5].ID},
		{CommentStr: "还有什么其他类型的电影推荐吗？", PostID: posts[5].ID},

		// 对王五第1篇文章的评论
		{CommentStr: "背单词的方法很实用", PostID: posts[6].ID},
		{CommentStr: "每天坚持学习很重要", PostID: posts[6].ID},

		// 对王五第2篇文章的评论
		{CommentStr: "时间管理对我帮助很大", PostID: posts[7].ID},
		{CommentStr: "番茄工作法确实有效", PostID: posts[7].ID},

		// 对王五第3篇文章的评论
		{CommentStr: "这本书我也读过，总结得很到位", PostID: posts[8].ID},

		// 对赵六第1篇文章的评论
		{CommentStr: "AI发展真的太快了", PostID: posts[9].ID},
		{CommentStr: "未来会有更多AI应用场景", PostID: posts[9].ID},

		// 对赵六第2篇文章的评论
		{CommentStr: "区块链不只是比特币，应用很广", PostID: posts[10].ID},
		{CommentStr: "希望能讲讲智能合约", PostID: posts[11].ID},

		// 对赵六第3篇文章的评论
		{CommentStr: "对比分析很全面", PostID: posts[11].ID},
		{CommentStr: "我们公司用的是AWS", PostID: posts[11].ID},
		{CommentStr: "GCP在某些方面确实有优势", PostID: posts[11].ID},
	}
	db.Create(&comments)

	fmt.Printf("创建了 %d 条评论\n", len(comments))
	fmt.Println("所有测试数据创建完成！")
}
