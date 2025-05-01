package models

type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  // 子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` // 父评论列表
	ParentCommentID    *uint           `json:"parent_comment_id"`                               // 父评论id，移除 size:10
	Content            string          `gorm:"size:256" json:"content"`                         // 评论内容
	DiggCount          int             `gorm:"type:tinyint;default:0" json:"digg_count"`        // 点赞数，修正为 type:tinyint
	CommentCount       int             `gorm:"type:tinyint;default:0" json:"comment_count"`     // 评论数，修正为 type:tinyint
	Article            ArticleModel    `gorm:"foreignKey:ArticleID" json:"-"`                   // 关联的文章
	ArticleID          uint            `json:"article_id"`                                      // 文章id
	User               UserModel       `gorm:"foreignKey:UserID" json:"user"`                   // 关联的用户，添加 foreignKey
	UserID             uint            `json:"user_id"`                                         // 评论的用户，移除 size:10
}
