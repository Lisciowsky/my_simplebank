package db

import (
	"context"
	"testing"
)

func TestCreateEntry(t *testing.T) {
	ctx := context.Background()

	// Define test parameters
	arg := CreateEntryParams{
		AccountID: 1,   // Example account ID
		Amount:    100, // Example amount
	}

	// Call the function
	entry, err := testQueries.CreateEntry(ctx, arg)
	if err != nil {
		t.Fatal(err)
	}

	// Check the results
	if entry.AccountID != arg.AccountID || entry.Amount != arg.Amount {
		t.Errorf("unexpected result: got %v, want %v", entry, arg)
	}
}

func TestCreateTransfer(t *testing.T) {
	ctx := context.Background()

	// Define test parameters
	arg := CreateTransferParams{
		FromAccountID: 1, // Example account IDs
		ToAccountID:   2,
		Amount:        100,
	}

	// Call the function
	transfer, err := testQueries.CreateTransfer(ctx, arg)
	if err != nil {
		t.Fatal(err)
	}

	// Check the results
	if transfer.FromAccountID != arg.FromAccountID || transfer.ToAccountID != arg.ToAccountID || transfer.Amount != arg.Amount {
		t.Errorf("unexpected result: got %v, want %v", transfer, arg)
	}
}
