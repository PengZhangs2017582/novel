// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortStoryDao is the data access object for table nv_short_story.
type ShortStoryDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ShortStoryColumns // columns contains all the column names of Table for convenient usage.
}

// ShortStoryColumns defines and stores column names for table nv_short_story.
type ShortStoryColumns struct {
	Id                string //
	Title             string // 标题
	Content           string // 内容
	RecommendCoverImg string // 推荐封面
	BookCoverImg      string // 书籍封面
	RecommendTitle    string // 推荐标题
	CatId             string // 分类，多个用逗号隔开
	PreReadPercent    string // 试读百分比
	IsAudit           string // 是否审核：1-已审核，2-未审核，3-草稿
	CreatedAt         string //
	UpdatedAt         string //
	DeletedAt         string //
}

// shortStoryColumns holds the columns for table nv_short_story.
var shortStoryColumns = ShortStoryColumns{
	Id:                "id",
	Title:             "title",
	Content:           "content",
	RecommendCoverImg: "recommend_cover_img",
	BookCoverImg:      "book_cover_img",
	RecommendTitle:    "recommend_title",
	CatId:             "cat_id",
	PreReadPercent:    "pre_read_percent",
	IsAudit:           "is_audit",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
}

// NewShortStoryDao creates and returns a new DAO object for table data access.
func NewShortStoryDao() *ShortStoryDao {
	return &ShortStoryDao{
		group:   "default",
		table:   "nv_short_story",
		columns: shortStoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortStoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortStoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortStoryDao) Columns() ShortStoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortStoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortStoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortStoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
