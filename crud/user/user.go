// Code generated by crud. DO NOT EDIT.

package user

import (
	"time"

	"github.com/cleancrud/crud/xsql"
	"github.com/cleancrud/crud/xsql/dialect"
)

type User struct {
	Id    int64     `json:"id"`    // id字段
	Name  string    `json:"name"`  // 名称
	Age   int64     `json:"age"`   // 年龄
	Ctime time.Time `json:"ctime"` // 创建时间
	Mtime time.Time `json:"mtime"` // 更新时间
}

const (
	// table tableName is user
	table = "user"
	//Id id字段
	Id = "id"
	//Name 名称
	Name = "name"
	//Age 年龄
	Age = "age"
	//Ctime 创建时间
	Ctime = "ctime"
	//Mtime 更新时间
	Mtime = "mtime"
)

var columns = []string{
	Id,
	Name,
	Age,
	Ctime,
	Mtime,
}

var columnsSet = map[string]struct{}{
	Id:    {},
	Name:  {},
	Age:   {},
	Ctime: {},
	Mtime: {},
}

func Columns() []string {
	return columns
}
func ColumnsSet() map[string]struct{} {
	return columnsSet
}

func (a *User) NewPtr() any {
	return &User{}
}
func (a *User) IsNil() bool {
	return a == nil
}

func (u *User) ScanDst(aa any, columns []string) []any {
	if a, ok := aa.(*User); ok {
		dst := make([]interface{}, 0, len(columns))
		for _, v := range columns {
			switch v {
			case Id:
				dst = append(dst, &a.Id)
			case Name:
				dst = append(dst, &a.Name)
			case Age:
				dst = append(dst, &a.Age)
			case Ctime:
				dst = append(dst, &a.Ctime)
			case Mtime:
				dst = append(dst, &a.Mtime)
			}
		}
		return dst
	}
	return []any{}

}

func (a *User) Values() []any {
	if a.IsNil() {
		return []any{}
	}
	return []interface{}{a.Id, a.Name, a.Age, a.Ctime, a.Mtime}
}
func (a *User) GetAutoIncrPk() (int64, string) {

	if a.IsNil() {
		return 0, Id
	}
	return a.Id, Id

}
func (a *User) SetAutoIncrPk(id int64) {

	if a.IsNil() {
		return
	}
	a.Id = int64(id)

}

func (a *User) Columns() []string {
	return columns
}
func (a *User) ColumnsSet() map[string]struct{} {
	return columnsSet
}
func (a *User) Schema() string {
	return ""
}

func (a *User) Dialect() string {

	return dialect.MySQL
}
func (a *User) Table() string {
	return table
}

const (
	IdOp = xsql.FieldOp[int64](Id)

	NameOp = xsql.StrFieldOp(Name)

	AgeOp = xsql.FieldOp[int64](Age)

	CtimeOp = xsql.FieldOp[string](Ctime)

	MtimeOp = xsql.FieldOp[string](Mtime)
)

func And(predicates ...xsql.WhereFunc) xsql.WhereFunc {
	return xsql.AndOp(predicates...)
}
func Or(predicates ...xsql.WhereFunc) xsql.WhereFunc {
	return xsql.OrOp(predicates...)
}
func Not(predicate xsql.WhereFunc) xsql.WhereFunc {
	return xsql.NotOp(predicate)
}

func Create(db xsql.ExecQuerier) *Creater {
	return &Creater{xsql.NewInsertExecutor[*User](db)}
}

type Creater struct {
	*xsql.InsertExecutor[*User]
}

func (c *Creater) SetUser(a ...*User) *Creater {
	c.SetItems(a...)
	return c
}

func Find(db xsql.ExecQuerier) *xsql.SelectExecutor[*User] {
	return xsql.NewSelectExecutor[*User](db)
}
func Delete(db xsql.ExecQuerier) *xsql.DeleteExecutor[*User] {
	return xsql.NewDeleteExecutor[*User](db)
}
func Update(db xsql.ExecQuerier) *Updater {
	return &Updater{xsql.NewUpdateExecutor[*User](db)}
}

type Updater struct {
	*xsql.UpdateExecutor[*User]
}

func (u *Updater) SetId(arg int64) *Updater {
	u.Set(Id, arg)
	return u
}

func (u *Updater) SetName(arg string) *Updater {
	u.Set(Name, arg)
	return u
}

func (u *Updater) SetAge(arg int64) *Updater {
	u.Set(Age, arg)
	return u
}

func (u *Updater) AddAge(arg interface{}) *Updater {
	u.Add(Age, arg)
	return u
}

func (u *Updater) SetCtime(arg time.Time) *Updater {
	u.Set(Ctime, arg)
	return u
}

func (u *Updater) SetMtime(arg time.Time) *Updater {
	u.Set(Mtime, arg)
	return u
}