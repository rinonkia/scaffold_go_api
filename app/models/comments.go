// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Comment is an object representing the database table.
type Comment struct {
	CommentID uint      `boil:"comment_id" json:"comment_id" toml:"comment_id" yaml:"comment_id"`
	ArticleID uint      `boil:"article_id" json:"article_id" toml:"article_id" yaml:"article_id"`
	Message   string    `boil:"message" json:"message" toml:"message" yaml:"message"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *commentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L commentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CommentColumns = struct {
	CommentID string
	ArticleID string
	Message   string
	CreatedAt string
}{
	CommentID: "comment_id",
	ArticleID: "article_id",
	Message:   "message",
	CreatedAt: "created_at",
}

var CommentTableColumns = struct {
	CommentID string
	ArticleID string
	Message   string
	CreatedAt string
}{
	CommentID: "comments.comment_id",
	ArticleID: "comments.article_id",
	Message:   "comments.message",
	CreatedAt: "comments.created_at",
}

// Generated where

var CommentWhere = struct {
	CommentID whereHelperuint
	ArticleID whereHelperuint
	Message   whereHelperstring
	CreatedAt whereHelpernull_Time
}{
	CommentID: whereHelperuint{field: "`comments`.`comment_id`"},
	ArticleID: whereHelperuint{field: "`comments`.`article_id`"},
	Message:   whereHelperstring{field: "`comments`.`message`"},
	CreatedAt: whereHelpernull_Time{field: "`comments`.`created_at`"},
}

// CommentRels is where relationship names are stored.
var CommentRels = struct {
	Article string
}{
	Article: "Article",
}

// commentR is where relationships are stored.
type commentR struct {
	Article *Article `boil:"Article" json:"Article" toml:"Article" yaml:"Article"`
}

// NewStruct creates a new relationship struct
func (*commentR) NewStruct() *commentR {
	return &commentR{}
}

func (r *commentR) GetArticle() *Article {
	if r == nil {
		return nil
	}
	return r.Article
}

// commentL is where Load methods for each relationship are stored.
type commentL struct{}

var (
	commentAllColumns            = []string{"comment_id", "article_id", "message", "created_at"}
	commentColumnsWithoutDefault = []string{"article_id", "message", "created_at"}
	commentColumnsWithDefault    = []string{"comment_id"}
	commentPrimaryKeyColumns     = []string{"comment_id"}
	commentGeneratedColumns      = []string{}
)

type (
	// CommentSlice is an alias for a slice of pointers to Comment.
	// This should almost always be used instead of []Comment.
	CommentSlice []*Comment
	// CommentHook is the signature for custom Comment hook methods
	CommentHook func(context.Context, boil.ContextExecutor, *Comment) error

	commentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	commentType                 = reflect.TypeOf(&Comment{})
	commentMapping              = queries.MakeStructMapping(commentType)
	commentPrimaryKeyMapping, _ = queries.BindMapping(commentType, commentMapping, commentPrimaryKeyColumns)
	commentInsertCacheMut       sync.RWMutex
	commentInsertCache          = make(map[string]insertCache)
	commentUpdateCacheMut       sync.RWMutex
	commentUpdateCache          = make(map[string]updateCache)
	commentUpsertCacheMut       sync.RWMutex
	commentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var commentAfterSelectHooks []CommentHook

var commentBeforeInsertHooks []CommentHook
var commentAfterInsertHooks []CommentHook

var commentBeforeUpdateHooks []CommentHook
var commentAfterUpdateHooks []CommentHook

var commentBeforeDeleteHooks []CommentHook
var commentAfterDeleteHooks []CommentHook

var commentBeforeUpsertHooks []CommentHook
var commentAfterUpsertHooks []CommentHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Comment) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Comment) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Comment) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Comment) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Comment) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Comment) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Comment) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Comment) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Comment) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range commentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCommentHook registers your hook function for all future operations.
func AddCommentHook(hookPoint boil.HookPoint, commentHook CommentHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		commentAfterSelectHooks = append(commentAfterSelectHooks, commentHook)
	case boil.BeforeInsertHook:
		commentBeforeInsertHooks = append(commentBeforeInsertHooks, commentHook)
	case boil.AfterInsertHook:
		commentAfterInsertHooks = append(commentAfterInsertHooks, commentHook)
	case boil.BeforeUpdateHook:
		commentBeforeUpdateHooks = append(commentBeforeUpdateHooks, commentHook)
	case boil.AfterUpdateHook:
		commentAfterUpdateHooks = append(commentAfterUpdateHooks, commentHook)
	case boil.BeforeDeleteHook:
		commentBeforeDeleteHooks = append(commentBeforeDeleteHooks, commentHook)
	case boil.AfterDeleteHook:
		commentAfterDeleteHooks = append(commentAfterDeleteHooks, commentHook)
	case boil.BeforeUpsertHook:
		commentBeforeUpsertHooks = append(commentBeforeUpsertHooks, commentHook)
	case boil.AfterUpsertHook:
		commentAfterUpsertHooks = append(commentAfterUpsertHooks, commentHook)
	}
}

// One returns a single comment record from the query.
func (q commentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Comment, error) {
	o := &Comment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for comments")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Comment records from the query.
func (q commentQuery) All(ctx context.Context, exec boil.ContextExecutor) (CommentSlice, error) {
	var o []*Comment

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Comment slice")
	}

	if len(commentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Comment records in the query.
func (q commentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count comments rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q commentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if comments exists")
	}

	return count > 0, nil
}

// Article pointed to by the foreign key.
func (o *Comment) Article(mods ...qm.QueryMod) articleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`article_id` = ?", o.ArticleID),
	}

	queryMods = append(queryMods, mods...)

	return Articles(queryMods...)
}

// LoadArticle allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (commentL) LoadArticle(ctx context.Context, e boil.ContextExecutor, singular bool, maybeComment interface{}, mods queries.Applicator) error {
	var slice []*Comment
	var object *Comment

	if singular {
		var ok bool
		object, ok = maybeComment.(*Comment)
		if !ok {
			object = new(Comment)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeComment)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeComment))
			}
		}
	} else {
		s, ok := maybeComment.(*[]*Comment)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeComment)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeComment))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &commentR{}
		}
		args = append(args, object.ArticleID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &commentR{}
			}

			for _, a := range args {
				if a == obj.ArticleID {
					continue Outer
				}
			}

			args = append(args, obj.ArticleID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`articles`),
		qm.WhereIn(`articles.article_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Article")
	}

	var resultSlice []*Article
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Article")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for articles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for articles")
	}

	if len(commentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Article = foreign
		if foreign.R == nil {
			foreign.R = &articleR{}
		}
		foreign.R.Comments = append(foreign.R.Comments, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ArticleID == foreign.ArticleID {
				local.R.Article = foreign
				if foreign.R == nil {
					foreign.R = &articleR{}
				}
				foreign.R.Comments = append(foreign.R.Comments, local)
				break
			}
		}
	}

	return nil
}

// SetArticle of the comment to the related item.
// Sets o.R.Article to related.
// Adds o to related.R.Comments.
func (o *Comment) SetArticle(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Article) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `comments` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"article_id"}),
		strmangle.WhereClause("`", "`", 0, commentPrimaryKeyColumns),
	)
	values := []interface{}{related.ArticleID, o.CommentID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ArticleID = related.ArticleID
	if o.R == nil {
		o.R = &commentR{
			Article: related,
		}
	} else {
		o.R.Article = related
	}

	if related.R == nil {
		related.R = &articleR{
			Comments: CommentSlice{o},
		}
	} else {
		related.R.Comments = append(related.R.Comments, o)
	}

	return nil
}

// Comments retrieves all the records using an executor.
func Comments(mods ...qm.QueryMod) commentQuery {
	mods = append(mods, qm.From("`comments`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`comments`.*"})
	}

	return commentQuery{q}
}

// FindComment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindComment(ctx context.Context, exec boil.ContextExecutor, commentID uint, selectCols ...string) (*Comment, error) {
	commentObj := &Comment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `comments` where `comment_id`=?", sel,
	)

	q := queries.Raw(query, commentID)

	err := q.Bind(ctx, exec, commentObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from comments")
	}

	if err = commentObj.doAfterSelectHooks(ctx, exec); err != nil {
		return commentObj, err
	}

	return commentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Comment) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no comments provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(commentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	commentInsertCacheMut.RLock()
	cache, cached := commentInsertCache[key]
	commentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			commentAllColumns,
			commentColumnsWithDefault,
			commentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(commentType, commentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(commentType, commentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `comments` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `comments` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `comments` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, commentPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into comments")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.CommentID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == commentMapping["comment_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.CommentID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for comments")
	}

CacheNoHooks:
	if !cached {
		commentInsertCacheMut.Lock()
		commentInsertCache[key] = cache
		commentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Comment.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Comment) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	commentUpdateCacheMut.RLock()
	cache, cached := commentUpdateCache[key]
	commentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			commentAllColumns,
			commentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update comments, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `comments` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, commentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(commentType, commentMapping, append(wl, commentPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update comments row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for comments")
	}

	if !cached {
		commentUpdateCacheMut.Lock()
		commentUpdateCache[key] = cache
		commentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q commentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for comments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for comments")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CommentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `comments` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, commentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in comment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all comment")
	}
	return rowsAff, nil
}

var mySQLCommentUniqueColumns = []string{
	"comment_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Comment) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no comments provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(commentColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCommentUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	commentUpsertCacheMut.RLock()
	cache, cached := commentUpsertCache[key]
	commentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			commentAllColumns,
			commentColumnsWithDefault,
			commentColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			commentAllColumns,
			commentPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert comments, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`comments`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `comments` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(commentType, commentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(commentType, commentMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for comments")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.CommentID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == commentMapping["comment_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(commentType, commentMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for comments")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for comments")
	}

CacheNoHooks:
	if !cached {
		commentUpsertCacheMut.Lock()
		commentUpsertCache[key] = cache
		commentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Comment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Comment) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Comment provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), commentPrimaryKeyMapping)
	sql := "DELETE FROM `comments` WHERE `comment_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from comments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for comments")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q commentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no commentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from comments")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for comments")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CommentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(commentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `comments` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, commentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from comment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for comments")
	}

	if len(commentAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Comment) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindComment(ctx, exec, o.CommentID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CommentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `comments`.* FROM `comments` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, commentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CommentSlice")
	}

	*o = slice

	return nil
}

// CommentExists checks if the Comment row exists.
func CommentExists(ctx context.Context, exec boil.ContextExecutor, commentID uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `comments` where `comment_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, commentID)
	}
	row := exec.QueryRowContext(ctx, sql, commentID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if comments exists")
	}

	return exists, nil
}
