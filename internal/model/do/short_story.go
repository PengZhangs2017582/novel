// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortStory is the golang structure of table nv_short_story for DAO operations like Where/Data.
type ShortStory struct {
	g.Meta            `orm:"table:nv_short_story, do:true"`
	Id                interface{} //
	Title             interface{} // 标题
	Content           interface{} // 内容
	RecommendCoverImg interface{} // 推荐封面
	BookCoverImg      interface{} // 书籍封面
	RecommendTitle    interface{} // 推荐标题
	CatId             interface{} // 分类，多个用逗号隔开
	PreReadPercent    interface{} // 试读百分比
	IsAudit           interface{} // 是否审核：1-已审核，2-未审核，3-草稿
	CreatedAt         *gtime.Time //
	UpdatedAt         *gtime.Time //
	DeletedAt         *gtime.Time //
}
