package structs

import (
	"testing"

	"github.com/pedrogonzalezj/play/custom"
)

func TestWallet(t *testing.T) {

	t.Run("Balance", func(t *testing.T) {
		wallet := Wallet{balance: 10}

		assertBalance(t, wallet, custom.Bitcoin(10))
	})

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(custom.Bitcoin(10))

		assertBalance(t, wallet, custom.Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{custom.Bitcoin(20)}
		err := wallet.Withdraw(custom.Bitcoin(10))

		assertBalance(t, wallet, custom.Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := custom.Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(custom.Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want custom.Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
