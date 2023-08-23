package module

import "github.com/jinzhu/gorm"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	gorm.Model    `json:"-"`
	VideoId       int64  `json:"id,omitempty" gorm:"unique;not null;index:idx_video_id"`
	AuthorId      int64  `json:"-" gorm:"not null"`
	Title         string `json:"title" bind:"required"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type VideoList struct {
	AuthorId   int64 `json:"-"`
	AllVideoes []Video
}

type Comment struct {
	gorm.Model `json:"-"`
	CommentId  int64  `json:"comment_id,omitempty" gorm:"unique"`
	UserId     int64  `json:"user_id,omitempty"`
	VideoId    int64  `json:"video_id" bind:"required"`
	Action     int8   `json:"action_type" binding:"required;oneof=1 2" gorm:"-"`
	Content    string `json:"content,omitempty"`
}

type CommentList struct {
	VideoId     int64 `json:"-"`
	AllComments []Comment
}

type User struct {
	gorm.Model      `json:"-"`
	UserId          int64  `json:"id,omitempty" gorm:"index:idx_user_id"`                                   // 用户id
	Name            string `json:"name,omitempty" binding:"required" gorm:"unique;not null;index:idx_name"` // 用户名称
	Password        string `json:"-" binding:"required" gorm:"not null"`
	FollowCount     int64  `json:"follow_count,omitempty" gorm:"default:0"` // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty"`                // 粉丝总数
	IsFollow        bool   `json:"is_follow,omitempty"`                     // true-已关注，false-未关注
	Avatar          string `json:"avatar,omitempty"`                        //用户头像
	BackgroundImage string `json:"background_image,omitempty"`              //用户个人页顶部大图
	Signature       string `json:"signature,omitempty"`                     //个人简介
	TotalFavorited  string `json:"total_favorited,omitempty"`               //获赞数量
	WorkCount       int    `json:"work_count,omitempty"`                    //作品数量
	FavoriteCount   int    `json:"favorite_count,omitempty"`                //点赞数量
	Token           string `json:"-" gorm:"-"`
}

type FavoriteVideo struct {
	gorm.Model `json:"-"`
	UserId     int64 `json:"user_id"`
	VideoId    int64 `json:"video_id"`
	Action     int8  `json:"action_type" binding:"oneof=1 2" gorm:"-"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
