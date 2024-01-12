// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortStory is the golang structure for table short_story.
type ShortStory struct {
	Id                uint        `json:"id"                description:""`
	Title             string      `json:"title"             description:"标题"`
	Content           string      `json:"content"           description:"内容"`
	RecommendCoverImg string      `json:"recommendCoverImg" description:"推荐封面"`
	BookCoverImg      string      `json:"bookCoverImg"      description:"书籍封面"`
	RecommendTitle    string      `json:"recommendTitle"    description:"推荐标题"`
	CatId             string      `json:"catId"             description:"分类，多个用逗号隔开"`
	PreReadPercent    uint        `json:"preReadPercent"    description:"试读百分比"`
	IsAudit           uint        `json:"isAudit"           description:"是否审核：1-已审核，2-未审核，3-草稿"`
	CreatedAt         *gtime.Time `json:"createdAt"         description:""`
	UpdatedAt         *gtime.Time `json:"updatedAt"         description:""`
	DeletedAt         *gtime.Time `json:"deletedAt"         description:""`
}
