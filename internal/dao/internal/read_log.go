// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReadLogDao is the data access object for table nv_read_log.
type ReadLogDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ReadLogColumns // columns contains all the column names of Table for convenient usage.
}

// ReadLogColumns defines and stores column names for table nv_read_log.
type ReadLogColumns struct {
	Id           string //
	ShortStoryId string //
	OperateType  string // 操作类型：1-阅读，2-加入书架
	OperateTime  string // 操作时间
	ReadDuration string // 阅读时长，秒为单位
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
}

// readLogColumns holds the columns for table nv_read_log.
var readLogColumns = ReadLogColumns{
	Id:           "id",
	ShortStoryId: "short_story_id",
	OperateType:  "operate_type",
	OperateTime:  "operate_time",
	ReadDuration: "read_duration",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewReadLogDao creates and returns a new DAO object for table data access.
func NewReadLogDao() *ReadLogDao {
	return &ReadLogDao{
		group:   "default",
		table:   "nv_read_log",
		columns: readLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReadLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReadLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReadLogDao) Columns() ReadLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReadLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReadLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReadLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
