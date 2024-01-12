// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PhoneVerifyDao is the data access object for table nv_phone_verify.
type PhoneVerifyDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns PhoneVerifyColumns // columns contains all the column names of Table for convenient usage.
}

// PhoneVerifyColumns defines and stores column names for table nv_phone_verify.
type PhoneVerifyColumns struct {
	Id        string //
	Phone     string // 手机号
	Code      string // 验证码
	Type      string // 类型：1-注册/登录，2-找回密码
	CreatedAt string //
	UpdatedAt string //
}

// phoneVerifyColumns holds the columns for table nv_phone_verify.
var phoneVerifyColumns = PhoneVerifyColumns{
	Id:        "id",
	Phone:     "phone",
	Code:      "code",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPhoneVerifyDao creates and returns a new DAO object for table data access.
func NewPhoneVerifyDao() *PhoneVerifyDao {
	return &PhoneVerifyDao{
		group:   "default",
		table:   "nv_phone_verify",
		columns: phoneVerifyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PhoneVerifyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PhoneVerifyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PhoneVerifyDao) Columns() PhoneVerifyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PhoneVerifyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PhoneVerifyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PhoneVerifyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
