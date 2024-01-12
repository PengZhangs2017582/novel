// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"novel/internal/dao/internal"
)

// internalPhoneVerifyDao is internal type for wrapping internal DAO implements.
type internalPhoneVerifyDao = *internal.PhoneVerifyDao

// phoneVerifyDao is the data access object for table nv_phone_verify.
// You can define custom methods on it to extend its functionality as you wish.
type phoneVerifyDao struct {
	internalPhoneVerifyDao
}

var (
	// PhoneVerify is globally public accessible object for table nv_phone_verify operations.
	PhoneVerify = phoneVerifyDao{
		internal.NewPhoneVerifyDao(),
	}
)

// Fill with you ideas below.