// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"novel/internal/dao/internal"
)

// internalReadLogDao is internal type for wrapping internal DAO implements.
type internalReadLogDao = *internal.ReadLogDao

// readLogDao is the data access object for table nv_read_log.
// You can define custom methods on it to extend its functionality as you wish.
type readLogDao struct {
	internalReadLogDao
}

var (
	// ReadLog is globally public accessible object for table nv_read_log operations.
	ReadLog = readLogDao{
		internal.NewReadLogDao(),
	}
)

// Fill with you ideas below.
