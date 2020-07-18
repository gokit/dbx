# gokit/dbx (database records)

forked from `gocraft/dbx` v2

[![GoDoc](https://godoc.org/github.com/gokit/dbx?status.png)](https://godoc.org/github.com/gokit/dbx)

`gokit/dbx` forked from `gocraft/dbr` v2, and Some destructive features have been added, still provides additions to Go's database/sql for super fast performance and convenience.

```
$ go get -u github.com/gokit/dbx
```

```go
import "github.com/gokit/dbx"
```

## Destructive features

- SelectStmt`s Where、Having、GroupBy、OrderBy will first reset and then append condition.
- UpdateStmt`s Where will first reset and then append condition.
- DeleteStmt`s Where will first reset and then append condition.

## New features

- SelectStmt support AndWhere、OrWhere、AndHaving、OrHaving、AddGroupBy、AddOrderBy、ToSql、ToRawSql
- UpdateStmt support AndWhere、OrWhere、ToSql、ToRawSql
- DeleteStmt support AndWhere、OrWhere、ToSql、ToRawSql
- InsertStmt support ToSql、ToRawSql

## Driver support

* MySQL
* PostgreSQL
* SQLite3

## Examples

See [godoc](https://godoc.org/github.com/gokit/dbx) for more examples.

### Open connections

```go
// create a connection (e.g. "postgres", "mysql", or "sqlite3")
conn, _ := Open("postgres", "...", nil)
conn.SetMaxOpenConns(10)

// create a session for each business unit of execution (e.g. a web request or goworkers job)
sess := conn.NewSession(nil)

// create a tx from sessions
sess.Begin()
```

### Create and use Tx

```go
sess := mysqlSession
tx, err := sess.Begin()
if err != nil {
	return
}
defer tx.RollbackUnlessCommitted()

// do stuff...

tx.Commit()
```

### SelectStmt loads data into structs

```go
// columns are mapped by tag then by field
type Suggestion struct {
	ID	int64		// id, will be autoloaded by last insert id
	Title	NullString	`db:"subject"`	// subjects are called titles now
	Url	string		`db:"-"`	// ignored
	secret	string		// ignored
}

// By default gokit/dbx converts CamelCase property names to snake_case column_names.
// You can override this with struct tags, just like with JSON tags.
// This is especially helpful while migrating from legacy systems.
var suggestions []Suggestion
sess := mysqlSession
sess.Select("*").From("suggestions").Load(&suggestions)
```

### SelectStmt with where-value interpolation

```go
// database/sql uses prepared statements, which means each argument
// in an IN clause needs its own question mark.
// gokit/dbx, on the other hand, handles interpolation itself
// so that you can easily use a single question mark paired with a
// dynamically sized slice.

sess := mysqlSession
ids := []int64{1, 2, 3, 4, 5}
sess.Select("*").From("suggestions").Where("id IN ?", ids).OrWhere("id = ?", 0)
```

### SelectStmt with joins

```go
sess := mysqlSession
sess.Select("*").From("suggestions").
	Join("subdomains", "suggestions.subdomain_id = subdomains.id")

sess.Select("*").From("suggestions").
	LeftJoin("subdomains", "suggestions.subdomain_id = subdomains.id")

// join multiple tables
sess.Select("*").From("suggestions").
	Join("subdomains", "suggestions.subdomain_id = subdomains.id").
	Join("accounts", "subdomains.accounts_id = accounts.id")
```

### SelectStmt with raw SQL

```go
SelectBySql("SELECT `title`, `body` FROM `suggestions` ORDER BY `id` ASC LIMIT 10")
```

### InsertStmt adds data from struct

```go
type Suggestion struct {
	ID		int64
	Title		NullString
	CreatedAt	time.Time
}
sugg := &Suggestion{
	Title:		NewNullString("Gopher"),
	CreatedAt:	time.Now(),
}
sess := mysqlSession
sess.InsertInto("suggestions").
	Columns("title").
	Record(&sugg).
	Exec()

// id is set automatically
fmt.Println(sugg.ID)
```

### InsertStmt adds data from value

```go
sess := mysqlSession
sess.InsertInto("suggestions").
	Pair("title", "Gopher").
	Pair("body", "I love go.")
```

## Thanks & Authors

Forked from gocraft/dbr
* [dbr](https://github.com/gocraft/dbr) -provides additions to Go's database/sql for super fast performance and convenience.

Inspiration from these excellent libraries:
* [sqlx](https://github.com/jmoiron/sqlx) - various useful tools and utils for interacting with database/sql.
* [Squirrel](https://github.com/lann/squirrel) - simple fluent query builder.

Authors:
* Jonathan Novak -- [https://github.com/cypriss](https://github.com/cypriss)
* Tai-Lin Chu -- [https://github.com/taylorchu](https://github.com/taylorchu)
* Sponsored by [UserVoice](https://eng.uservoice.com)

Contributors:
* Paul Bergeron -- [https://github.com/dinedal](https://github.com/dinedal) - SQLite dialect
