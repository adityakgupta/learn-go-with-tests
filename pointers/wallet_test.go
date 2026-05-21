package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted error but got nil")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{10}
		wallet.Withdraw(10)
		assertBalance(t, wallet, 0)
	})

	t.Run("insufficient", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(20)

		assertError(t, err, Err)
		assertBalance(t, wallet, 10)
	})
}
