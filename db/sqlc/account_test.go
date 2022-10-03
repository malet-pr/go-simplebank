package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/malet-pr/go-simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	acc, _ := testQueries.CreateAccount(context.Background(), args)
	return acc
}

func findAccount(t *testing.T, id int64) (Account,error) {
	acc, err := testQueries.GetAccount(context.Background(), id)
	return acc, err
}

func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	acc, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, args.Owner, acc.Owner)
	require.Equal(t, args.Balance, acc.Balance)
	require.Equal(t, args.Currency, acc.Currency)
	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.Balance, acc2.Balance)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	args := UpdateAccountBalanceParams{
		ID:      acc1.ID,
		Balance: util.RandomMoney(),
	}
	acc2, err := testQueries.UpdateAccountBalance(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, args.Balance, acc2.Balance)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	args := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.True(t, len(accounts) <= int(args.Limit))
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	acc2, err := findAccount(t, acc1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc2)
}
