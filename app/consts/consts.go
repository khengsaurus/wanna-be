package consts

type ContextKey string

var (
	PgConnPoolKey = ContextKey("pg_pool")
	UsersTable    = "users"
	ExpensesTable = "expenses"
)
