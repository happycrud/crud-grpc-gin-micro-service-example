// Code generated by crud. DO NOT EDIT.

package alltypetable

import (
	"time"

	"github.com/cleancrud/crud/xsql"
	"github.com/cleancrud/crud/xsql/dialect"
)

type AllTypeTable struct {
	Id              int64     `json:"id"`               // 自增id
	TInt            int64     `json:"t_int"`            // 小小整型
	SInt            int64     `json:"s_int"`            // 小整数
	MInt            int64     `json:"m_int"`            // 中整数
	BInt            int64     `json:"b_int"`            // 大整数
	F32             float32   `json:"f32"`              // 小浮点
	F64             float64   `json:"f64"`              // 大浮点
	DecimalMysql    float64   `json:"decimal_mysql"`    //
	CharM           string    `json:"char_m"`           //
	VarcharM        string    `json:"varchar_m"`        //
	JsonM           string    `json:"json_m"`           //
	NvarcharM       string    `json:"nvarchar_m"`       //
	NcharM          string    `json:"nchar_m"`          //
	TimeM           string    `json:"time_m"`           //
	DateM           time.Time `json:"date_m"`           //
	DataTimeM       time.Time `json:"data_time_m"`      //
	TimestampM      time.Time `json:"timestamp_m"`      // 创建时间
	TimestampUpdate time.Time `json:"timestamp_update"` // 更新时间
	YearM           string    `json:"year_m"`           // 年
	TText           string    `json:"t_text"`           //
	MText           string    `json:"m_text"`           //
	TextM           string    `json:"text_m"`           //
	LText           string    `json:"l_text"`           //
	BinaryM         []byte    `json:"binary_m"`         //
	BlobM           []byte    `json:"blob_m"`           //
	LBlob           []byte    `json:"l_blob"`           //
	MBlob           []byte    `json:"m_blob"`           //
	TBlob           []byte    `json:"t_blob"`           //
	BitM            []byte    `json:"bit_m"`            //
	EnumM           string    `json:"enum_m"`           //
	SetM            string    `json:"set_m"`            //
	BoolM           int64     `json:"bool_m"`           //
}

const (
	// table tableName is all_type_table
	table = "all_type_table"
	//Id 自增id
	Id = "id"
	//TInt 小小整型
	TInt = "t_int"
	//SInt 小整数
	SInt = "s_int"
	//MInt 中整数
	MInt = "m_int"
	//BInt 大整数
	BInt = "b_int"
	//F32 小浮点
	F32 = "f32"
	//F64 大浮点
	F64 = "f64"
	//DecimalMysql
	DecimalMysql = "decimal_mysql"
	//CharM
	CharM = "char_m"
	//VarcharM
	VarcharM = "varchar_m"
	//JsonM
	JsonM = "json_m"
	//NvarcharM
	NvarcharM = "nvarchar_m"
	//NcharM
	NcharM = "nchar_m"
	//TimeM
	TimeM = "time_m"
	//DateM
	DateM = "date_m"
	//DataTimeM
	DataTimeM = "data_time_m"
	//TimestampM 创建时间
	TimestampM = "timestamp_m"
	//TimestampUpdate 更新时间
	TimestampUpdate = "timestamp_update"
	//YearM 年
	YearM = "year_m"
	//TText
	TText = "t_text"
	//MText
	MText = "m_text"
	//TextM
	TextM = "text_m"
	//LText
	LText = "l_text"
	//BinaryM
	BinaryM = "binary_m"
	//BlobM
	BlobM = "blob_m"
	//LBlob
	LBlob = "l_blob"
	//MBlob
	MBlob = "m_blob"
	//TBlob
	TBlob = "t_blob"
	//BitM
	BitM = "bit_m"
	//EnumM
	EnumM = "enum_m"
	//SetM
	SetM = "set_m"
	//BoolM
	BoolM = "bool_m"
)

var columns = []string{
	Id,
	TInt,
	SInt,
	MInt,
	BInt,
	F32,
	F64,
	DecimalMysql,
	CharM,
	VarcharM,
	JsonM,
	NvarcharM,
	NcharM,
	TimeM,
	DateM,
	DataTimeM,
	TimestampM,
	TimestampUpdate,
	YearM,
	TText,
	MText,
	TextM,
	LText,
	BinaryM,
	BlobM,
	LBlob,
	MBlob,
	TBlob,
	BitM,
	EnumM,
	SetM,
	BoolM,
}

var columnsSet = map[string]struct{}{
	Id:              {},
	TInt:            {},
	SInt:            {},
	MInt:            {},
	BInt:            {},
	F32:             {},
	F64:             {},
	DecimalMysql:    {},
	CharM:           {},
	VarcharM:        {},
	JsonM:           {},
	NvarcharM:       {},
	NcharM:          {},
	TimeM:           {},
	DateM:           {},
	DataTimeM:       {},
	TimestampM:      {},
	TimestampUpdate: {},
	YearM:           {},
	TText:           {},
	MText:           {},
	TextM:           {},
	LText:           {},
	BinaryM:         {},
	BlobM:           {},
	LBlob:           {},
	MBlob:           {},
	TBlob:           {},
	BitM:            {},
	EnumM:           {},
	SetM:            {},
	BoolM:           {},
}

func Columns() []string {
	return columns
}
func ColumnsSet() map[string]struct{} {
	return columnsSet
}

func (a *AllTypeTable) NewPtr() any {
	return &AllTypeTable{}
}
func (a *AllTypeTable) IsNil() bool {
	return a == nil
}

func (u *AllTypeTable) ScanDst(aa any, columns []string) []any {
	if a, ok := aa.(*AllTypeTable); ok {
		dst := make([]interface{}, 0, len(columns))
		for _, v := range columns {
			switch v {
			case Id:
				dst = append(dst, &a.Id)
			case TInt:
				dst = append(dst, &a.TInt)
			case SInt:
				dst = append(dst, &a.SInt)
			case MInt:
				dst = append(dst, &a.MInt)
			case BInt:
				dst = append(dst, &a.BInt)
			case F32:
				dst = append(dst, &a.F32)
			case F64:
				dst = append(dst, &a.F64)
			case DecimalMysql:
				dst = append(dst, &a.DecimalMysql)
			case CharM:
				dst = append(dst, &a.CharM)
			case VarcharM:
				dst = append(dst, &a.VarcharM)
			case JsonM:
				dst = append(dst, &a.JsonM)
			case NvarcharM:
				dst = append(dst, &a.NvarcharM)
			case NcharM:
				dst = append(dst, &a.NcharM)
			case TimeM:
				dst = append(dst, &a.TimeM)
			case DateM:
				dst = append(dst, &a.DateM)
			case DataTimeM:
				dst = append(dst, &a.DataTimeM)
			case TimestampM:
				dst = append(dst, &a.TimestampM)
			case TimestampUpdate:
				dst = append(dst, &a.TimestampUpdate)
			case YearM:
				dst = append(dst, &a.YearM)
			case TText:
				dst = append(dst, &a.TText)
			case MText:
				dst = append(dst, &a.MText)
			case TextM:
				dst = append(dst, &a.TextM)
			case LText:
				dst = append(dst, &a.LText)
			case BinaryM:
				dst = append(dst, &a.BinaryM)
			case BlobM:
				dst = append(dst, &a.BlobM)
			case LBlob:
				dst = append(dst, &a.LBlob)
			case MBlob:
				dst = append(dst, &a.MBlob)
			case TBlob:
				dst = append(dst, &a.TBlob)
			case BitM:
				dst = append(dst, &a.BitM)
			case EnumM:
				dst = append(dst, &a.EnumM)
			case SetM:
				dst = append(dst, &a.SetM)
			case BoolM:
				dst = append(dst, &a.BoolM)
			}
		}
		return dst
	}
	return []any{}

}

func (a *AllTypeTable) Values() []any {
	if a.IsNil() {
		return []any{}
	}
	return []interface{}{a.Id, a.TInt, a.SInt, a.MInt, a.BInt, a.F32, a.F64, a.DecimalMysql, a.CharM, a.VarcharM, a.JsonM, a.NvarcharM, a.NcharM, a.TimeM, a.DateM, a.DataTimeM, a.TimestampM, a.TimestampUpdate, a.YearM, a.TText, a.MText, a.TextM, a.LText, a.BinaryM, a.BlobM, a.LBlob, a.MBlob, a.TBlob, a.BitM, a.EnumM, a.SetM, a.BoolM}
}
func (a *AllTypeTable) GetAutoIncrPk() (int64, string) {

	if a.IsNil() {
		return 0, Id
	}
	return a.Id, Id

}
func (a *AllTypeTable) SetAutoIncrPk(id int64) {

	if a.IsNil() {
		return
	}
	a.Id = int64(id)

}

func (a *AllTypeTable) Columns() []string {
	return columns
}
func (a *AllTypeTable) ColumnsSet() map[string]struct{} {
	return columnsSet
}
func (a *AllTypeTable) Schema() string {
	return ""
}

func (a *AllTypeTable) Dialect() string {

	return dialect.MySQL
}
func (a *AllTypeTable) Table() string {
	return table
}

const (
	IdOp = xsql.FieldOp[int64](Id)

	TIntOp = xsql.FieldOp[int64](TInt)

	SIntOp = xsql.FieldOp[int64](SInt)

	MIntOp = xsql.FieldOp[int64](MInt)

	BIntOp = xsql.FieldOp[int64](BInt)

	F32Op = xsql.FieldOp[float32](F32)

	F64Op = xsql.FieldOp[float64](F64)

	DecimalMysqlOp = xsql.FieldOp[float64](DecimalMysql)

	CharMOp = xsql.StrFieldOp(CharM)

	VarcharMOp = xsql.StrFieldOp(VarcharM)

	NvarcharMOp = xsql.StrFieldOp(NvarcharM)

	NcharMOp = xsql.StrFieldOp(NcharM)

	TimeMOp = xsql.StrFieldOp(TimeM)

	DateMOp = xsql.FieldOp[string](DateM)

	DataTimeMOp = xsql.FieldOp[string](DataTimeM)

	TimestampMOp = xsql.FieldOp[string](TimestampM)

	TimestampUpdateOp = xsql.FieldOp[string](TimestampUpdate)

	YearMOp = xsql.StrFieldOp(YearM)

	BinaryMOp = xsql.FieldOp[[]byte](BinaryM)

	BitMOp = xsql.FieldOp[[]byte](BitM)

	EnumMOp = xsql.StrFieldOp(EnumM)

	SetMOp = xsql.StrFieldOp(SetM)

	BoolMOp = xsql.FieldOp[int64](BoolM)
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
	return &Creater{xsql.NewInsertExecutor[*AllTypeTable](db)}
}

type Creater struct {
	*xsql.InsertExecutor[*AllTypeTable]
}

func (c *Creater) SetAllTypeTable(a ...*AllTypeTable) *Creater {
	c.SetItems(a...)
	return c
}

func Find(db xsql.ExecQuerier) *xsql.SelectExecutor[*AllTypeTable] {
	return xsql.NewSelectExecutor[*AllTypeTable](db)
}
func Delete(db xsql.ExecQuerier) *xsql.DeleteExecutor[*AllTypeTable] {
	return xsql.NewDeleteExecutor[*AllTypeTable](db)
}
func Update(db xsql.ExecQuerier) *Updater {
	return &Updater{xsql.NewUpdateExecutor[*AllTypeTable](db)}
}

type Updater struct {
	*xsql.UpdateExecutor[*AllTypeTable]
}

func (u *Updater) SetId(arg int64) *Updater {
	u.Set(Id, arg)
	return u
}

func (u *Updater) SetTInt(arg int64) *Updater {
	u.Set(TInt, arg)
	return u
}

func (u *Updater) AddTInt(arg interface{}) *Updater {
	u.Add(TInt, arg)
	return u
}

func (u *Updater) SetSInt(arg int64) *Updater {
	u.Set(SInt, arg)
	return u
}

func (u *Updater) AddSInt(arg interface{}) *Updater {
	u.Add(SInt, arg)
	return u
}

func (u *Updater) SetMInt(arg int64) *Updater {
	u.Set(MInt, arg)
	return u
}

func (u *Updater) AddMInt(arg interface{}) *Updater {
	u.Add(MInt, arg)
	return u
}

func (u *Updater) SetBInt(arg int64) *Updater {
	u.Set(BInt, arg)
	return u
}

func (u *Updater) AddBInt(arg interface{}) *Updater {
	u.Add(BInt, arg)
	return u
}

func (u *Updater) SetF32(arg float32) *Updater {
	u.Set(F32, arg)
	return u
}

func (u *Updater) AddF32(arg interface{}) *Updater {
	u.Add(F32, arg)
	return u
}

func (u *Updater) SetF64(arg float64) *Updater {
	u.Set(F64, arg)
	return u
}

func (u *Updater) AddF64(arg interface{}) *Updater {
	u.Add(F64, arg)
	return u
}

func (u *Updater) SetDecimalMysql(arg float64) *Updater {
	u.Set(DecimalMysql, arg)
	return u
}

func (u *Updater) AddDecimalMysql(arg interface{}) *Updater {
	u.Add(DecimalMysql, arg)
	return u
}

func (u *Updater) SetCharM(arg string) *Updater {
	u.Set(CharM, arg)
	return u
}

func (u *Updater) SetVarcharM(arg string) *Updater {
	u.Set(VarcharM, arg)
	return u
}

func (u *Updater) SetJsonM(arg string) *Updater {
	u.Set(JsonM, arg)
	return u
}

func (u *Updater) SetNvarcharM(arg string) *Updater {
	u.Set(NvarcharM, arg)
	return u
}

func (u *Updater) SetNcharM(arg string) *Updater {
	u.Set(NcharM, arg)
	return u
}

func (u *Updater) SetTimeM(arg string) *Updater {
	u.Set(TimeM, arg)
	return u
}

func (u *Updater) SetDateM(arg time.Time) *Updater {
	u.Set(DateM, arg)
	return u
}

func (u *Updater) SetDataTimeM(arg time.Time) *Updater {
	u.Set(DataTimeM, arg)
	return u
}

func (u *Updater) SetTimestampM(arg time.Time) *Updater {
	u.Set(TimestampM, arg)
	return u
}

func (u *Updater) SetTimestampUpdate(arg time.Time) *Updater {
	u.Set(TimestampUpdate, arg)
	return u
}

func (u *Updater) SetYearM(arg string) *Updater {
	u.Set(YearM, arg)
	return u
}

func (u *Updater) SetTText(arg string) *Updater {
	u.Set(TText, arg)
	return u
}

func (u *Updater) SetMText(arg string) *Updater {
	u.Set(MText, arg)
	return u
}

func (u *Updater) SetTextM(arg string) *Updater {
	u.Set(TextM, arg)
	return u
}

func (u *Updater) SetLText(arg string) *Updater {
	u.Set(LText, arg)
	return u
}

func (u *Updater) SetBinaryM(arg []byte) *Updater {
	u.Set(BinaryM, arg)
	return u
}

func (u *Updater) SetBlobM(arg []byte) *Updater {
	u.Set(BlobM, arg)
	return u
}

func (u *Updater) SetLBlob(arg []byte) *Updater {
	u.Set(LBlob, arg)
	return u
}

func (u *Updater) SetMBlob(arg []byte) *Updater {
	u.Set(MBlob, arg)
	return u
}

func (u *Updater) SetTBlob(arg []byte) *Updater {
	u.Set(TBlob, arg)
	return u
}

func (u *Updater) SetBitM(arg []byte) *Updater {
	u.Set(BitM, arg)
	return u
}

func (u *Updater) SetEnumM(arg string) *Updater {
	u.Set(EnumM, arg)
	return u
}

func (u *Updater) SetSetM(arg string) *Updater {
	u.Set(SetM, arg)
	return u
}

func (u *Updater) SetBoolM(arg int64) *Updater {
	u.Set(BoolM, arg)
	return u
}

func (u *Updater) AddBoolM(arg interface{}) *Updater {
	u.Add(BoolM, arg)
	return u
}
