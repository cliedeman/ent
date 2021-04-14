package benchmark

import (
	"context"
	"entgo.io/ent/entc/integration/benchmark/ent"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
)

func BenchmarkQuery1User(b *testing.B) {
	ctx := context.Background()
	driver, err := NewBenchDriver()
	require.NoError(b, err)
	client := ent.NewClient(ent.Driver(driver))

	driver.mock.
		ExpectQuery("SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name` FROM `users`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name"}).
			AddRow(1, 20, "john"))

	users := client.User.Query().AllX(ctx)
	require.NotNil(b, users)
}

func BenchmarkQuery100Users(b *testing.B) {
	ctx := context.Background()
	driver, err := NewBenchDriver()
	require.NoError(b, err)
	client := ent.NewClient(ent.Driver(driver))

	rows := sqlmock.NewRows([]string{"id", "age", "name"})

	for i := 1; i <= 100; i++ {
		rows.AddRow(i, 20, "john")
	}

	driver.mock.
		ExpectQuery("SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name` FROM `users`").
		WillReturnRows(rows)

	users := client.User.Query().AllX(ctx)
	require.NotNil(b, users)
}

func BenchmarkQueryInsertUser(b *testing.B) {
	ctx := context.Background()
	driver, err := NewBenchDriver()
	require.NoError(b, err)
	client := ent.NewClient(ent.Driver(driver))

	driver.mock.
		ExpectExec("INSERT INTO `users` (`age`, `name`) VALUES (?, ?)").
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))

	users := client.User.Create().SetName("john").SetAge(25).SaveX(ctx)
	require.NotNil(b, users)
}

func BenchmarkQueryUsersWithEagerLoad(b *testing.B) {
	ctx := context.Background()
	driver, err := NewBenchDriver()
	require.NoError(b, err)
	client := ent.NewClient(ent.Driver(driver))

	driver.mock.
		ExpectQuery("SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name` FROM `users`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name"}).
			AddRow(1, 20, "john"))

	driver.mock.
		ExpectQuery("SELECT `user_id`, `group_id` FROM `user_groups` WHERE `user_id` IN (?)").
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "group_id"}).
			AddRow(1, 1))

	driver.mock.
		ExpectQuery("SELECT DISTINCT `groups`.`id`, `groups`.`active`, `groups`.`name` FROM `groups` WHERE `groups`.`id` IN (?)").
		WillReturnRows(sqlmock.NewRows([]string{"id", "active"}).
			AddRow(1, true))

	users := client.User.Query().WithGroups().AllX(ctx)
	require.NotNil(b, users)
}
