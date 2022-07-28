package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		received := wallet.Balance()
		if received != expected {
			t.Errorf("received %s expected %s", received, expected)
		}
	}

	assertNoError := func(t testing.TB, received error) {
		t.Helper()
		if received != nil {
			t.Fatal("expected no error but got one")
		}
	}

	assertError := func(t testing.TB, receivedError error, expectedError error) {
		t.Helper()
		if receivedError == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if receivedError != expectedError {
			t.Errorf("received %q expected %q", receivedError, expectedError)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
