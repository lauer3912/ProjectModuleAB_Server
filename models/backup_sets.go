package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//备份集
type BackupSets struct {
	Id       string      `orm:"pk;size(36)"`
	Name     string      `orm:"size(32)"`
	Desc     string      `orm:"size(128);null"`
	Oss      *Oss        `orm:"null;rel(fk);on_delete(set_null)"`
	Oas      *Oas        `orm:"null;rel(fk);on_delete(set_null)"`
	Policies []*Policies `orm:"reverse(many)"`
	Hosts    []*Hosts    `orm:"null;rel(m2m);rel_table(backupsets_to_hosts)"`
}

func init() {
	if prefix := beego.AppConfig.String("database::mysqlprefex"); prefix != "" {
		orm.RegisterModelWithPrefix(prefix, new(BackupSets))
	} else {
		orm.RegisterModel(new(BackupSets))
	}
}
