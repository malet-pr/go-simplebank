package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/malet-pr/go-simplebank/util"
	"github.com/stretchr/testify/require"
)

func createAccountForEntry(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	acc, _ := testQueries.CreateAccount(context.Background(), args)
	//require.NoError(t, err)
	//require.NotEmpty(t, acc)
	require.NotZero(t, acc.ID)
	return acc
}

func createRandomEntry(t *testing.T) Entry {
	acc := createAccountForEntry(t)
	args := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	require.NotZero(t, acc.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestUpdateEntryAmount(t *testing.T) {
	entry1 := createRandomEntry(t)
	args := UpdateEntryAmountParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}
	entry2, err := testQueries.UpdateEntryAmount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, args.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	args := ListEntriesParams{
		Limit:  5,
		Offset: 0,
	}
	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.True(t,len(entries) <= int(args.Limit))
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

