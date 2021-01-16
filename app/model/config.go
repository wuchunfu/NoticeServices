// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"NoticeServices/app/model/internal"
	"github.com/gogf/gf/database/gdb"
)

// Config is the golang structure for table config.
type Config internal.Config

// Fill with you ideas below.

type ConfigData struct {
	AppId       string `orm:"app_id"       json:"app_id"`       //
	Name        string `orm:"name"         json:"name"`         //
	SendGateway string `orm:"send_gateway" json:"send_gateway"` //
	Type        string `orm:"type"         json:"type"`         //
}

// Service修改内容
type ConfigUpData struct {
	ConfigData
	Id string
}

type ConfigServiceGetListReq struct {
	Key         string // 关键字
	Name        string //
	SendGateway string //
	Type        string //
	Page        int    `d:"1"  v:"min:0#分页号码错误"`     // 分页号码
	Size        int    `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
}

type ConfigServiceGetListRes struct {
	List  gdb.Result `json:"list"`  // 列表
	Page  int        `json:"page"`  // 分页码
	Size  int        `json:"size"`  // 分页数量
	Total int        `json:"total"` // 数据总数
}