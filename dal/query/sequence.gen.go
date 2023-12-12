// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"shorturl/dal/model"
)

func newSequence(db *gorm.DB, opts ...gen.DOOption) sequence {
	_sequence := sequence{}

	_sequence.sequenceDo.UseDB(db, opts...)
	_sequence.sequenceDo.UseModel(&model.Sequence{})

	tableName := _sequence.sequenceDo.TableName()
	_sequence.ALL = field.NewAsterisk(tableName)
	_sequence.ID = field.NewInt64(tableName, "id")
	_sequence.Stub = field.NewString(tableName, "stub")
	_sequence.Timestamp = field.NewTime(tableName, "timestamp")

	_sequence.fillFieldMap()

	return _sequence
}

// sequence 序号表
type sequence struct {
	sequenceDo sequenceDo

	ALL       field.Asterisk
	ID        field.Int64
	Stub      field.String
	Timestamp field.Time

	fieldMap map[string]field.Expr
}

func (s sequence) Table(newTableName string) *sequence {
	s.sequenceDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sequence) As(alias string) *sequence {
	s.sequenceDo.DO = *(s.sequenceDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sequence) updateTableName(table string) *sequence {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Stub = field.NewString(table, "stub")
	s.Timestamp = field.NewTime(table, "timestamp")

	s.fillFieldMap()

	return s
}

func (s *sequence) WithContext(ctx context.Context) ISequenceDo { return s.sequenceDo.WithContext(ctx) }

func (s sequence) TableName() string { return s.sequenceDo.TableName() }

func (s sequence) Alias() string { return s.sequenceDo.Alias() }

func (s sequence) Columns(cols ...field.Expr) gen.Columns { return s.sequenceDo.Columns(cols...) }

func (s *sequence) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sequence) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["id"] = s.ID
	s.fieldMap["stub"] = s.Stub
	s.fieldMap["timestamp"] = s.Timestamp
}

func (s sequence) clone(db *gorm.DB) sequence {
	s.sequenceDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sequence) replaceDB(db *gorm.DB) sequence {
	s.sequenceDo.ReplaceDB(db)
	return s
}

type sequenceDo struct{ gen.DO }

type ISequenceDo interface {
	gen.SubQuery
	Debug() ISequenceDo
	WithContext(ctx context.Context) ISequenceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISequenceDo
	WriteDB() ISequenceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISequenceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISequenceDo
	Not(conds ...gen.Condition) ISequenceDo
	Or(conds ...gen.Condition) ISequenceDo
	Select(conds ...field.Expr) ISequenceDo
	Where(conds ...gen.Condition) ISequenceDo
	Order(conds ...field.Expr) ISequenceDo
	Distinct(cols ...field.Expr) ISequenceDo
	Omit(cols ...field.Expr) ISequenceDo
	Join(table schema.Tabler, on ...field.Expr) ISequenceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISequenceDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISequenceDo
	Group(cols ...field.Expr) ISequenceDo
	Having(conds ...gen.Condition) ISequenceDo
	Limit(limit int) ISequenceDo
	Offset(offset int) ISequenceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISequenceDo
	Unscoped() ISequenceDo
	Create(values ...*model.Sequence) error
	CreateInBatches(values []*model.Sequence, batchSize int) error
	Save(values ...*model.Sequence) error
	First() (*model.Sequence, error)
	Take() (*model.Sequence, error)
	Last() (*model.Sequence, error)
	Find() ([]*model.Sequence, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sequence, err error)
	FindInBatches(result *[]*model.Sequence, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Sequence) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISequenceDo
	Assign(attrs ...field.AssignExpr) ISequenceDo
	Joins(fields ...field.RelationField) ISequenceDo
	Preload(fields ...field.RelationField) ISequenceDo
	FirstOrInit() (*model.Sequence, error)
	FirstOrCreate() (*model.Sequence, error)
	FindByPage(offset int, limit int) (result []*model.Sequence, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISequenceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	ReplaceStub() (err error)
	LastInsertID() (result uint64, err error)
}

// REPLACE INTO @@table (stub) VALUES ('a')
func (s sequenceDo) ReplaceStub() (err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("REPLACE INTO sequence (stub) VALUES ('a') ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Exec(generateSQL.String()) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT LAST_INSERT_ID()
func (s sequenceDo) LastInsertID() (result uint64, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT LAST_INSERT_ID() ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (s sequenceDo) Debug() ISequenceDo {
	return s.withDO(s.DO.Debug())
}

func (s sequenceDo) WithContext(ctx context.Context) ISequenceDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sequenceDo) ReadDB() ISequenceDo {
	return s.Clauses(dbresolver.Read)
}

func (s sequenceDo) WriteDB() ISequenceDo {
	return s.Clauses(dbresolver.Write)
}

func (s sequenceDo) Session(config *gorm.Session) ISequenceDo {
	return s.withDO(s.DO.Session(config))
}

func (s sequenceDo) Clauses(conds ...clause.Expression) ISequenceDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sequenceDo) Returning(value interface{}, columns ...string) ISequenceDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sequenceDo) Not(conds ...gen.Condition) ISequenceDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sequenceDo) Or(conds ...gen.Condition) ISequenceDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sequenceDo) Select(conds ...field.Expr) ISequenceDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sequenceDo) Where(conds ...gen.Condition) ISequenceDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sequenceDo) Order(conds ...field.Expr) ISequenceDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sequenceDo) Distinct(cols ...field.Expr) ISequenceDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sequenceDo) Omit(cols ...field.Expr) ISequenceDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sequenceDo) Join(table schema.Tabler, on ...field.Expr) ISequenceDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sequenceDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISequenceDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sequenceDo) RightJoin(table schema.Tabler, on ...field.Expr) ISequenceDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sequenceDo) Group(cols ...field.Expr) ISequenceDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sequenceDo) Having(conds ...gen.Condition) ISequenceDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sequenceDo) Limit(limit int) ISequenceDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sequenceDo) Offset(offset int) ISequenceDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sequenceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISequenceDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sequenceDo) Unscoped() ISequenceDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sequenceDo) Create(values ...*model.Sequence) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sequenceDo) CreateInBatches(values []*model.Sequence, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sequenceDo) Save(values ...*model.Sequence) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sequenceDo) First() (*model.Sequence, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sequence), nil
	}
}

func (s sequenceDo) Take() (*model.Sequence, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sequence), nil
	}
}

func (s sequenceDo) Last() (*model.Sequence, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sequence), nil
	}
}

func (s sequenceDo) Find() ([]*model.Sequence, error) {
	result, err := s.DO.Find()
	return result.([]*model.Sequence), err
}

func (s sequenceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sequence, err error) {
	buf := make([]*model.Sequence, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sequenceDo) FindInBatches(result *[]*model.Sequence, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sequenceDo) Attrs(attrs ...field.AssignExpr) ISequenceDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sequenceDo) Assign(attrs ...field.AssignExpr) ISequenceDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sequenceDo) Joins(fields ...field.RelationField) ISequenceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sequenceDo) Preload(fields ...field.RelationField) ISequenceDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sequenceDo) FirstOrInit() (*model.Sequence, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sequence), nil
	}
}

func (s sequenceDo) FirstOrCreate() (*model.Sequence, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sequence), nil
	}
}

func (s sequenceDo) FindByPage(offset int, limit int) (result []*model.Sequence, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sequenceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sequenceDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sequenceDo) Delete(models ...*model.Sequence) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sequenceDo) withDO(do gen.Dao) *sequenceDo {
	s.DO = *do.(*gen.DO)
	return s
}
