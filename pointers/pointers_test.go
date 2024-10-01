package pointers

import "testing"

func TestWallet(t *testing.T) {
	// wallet := Wallet{balance: 20}
	// // wallet.Deposit(BitCoin(10))
	// wallet.Withdraw(BitCoin(10))
	// got := wallet.Balance()
	// expected := BitCoin(10)
	// // got := wallet.Balance()
	// // expected := BitCoin(20)

	// if got != expected {
	// 	t.Errorf("got %s wanted %s", got, expected)
	// }

	// t.Run("Desposit into wallet", func(t *testing.T) {
	// 	wallet := Wallet{}
	// 	wallet.Deposit(BitCoin(10))
	// 	got := wallet.Balance()
	// 	want := BitCoin(10)

	// 	if got != want {
	// 		t.Errorf("got %s, wanted %s", got, want)
	// 	}
	// })

	// t.Run("Withdraw Balance", func(t *testing.T) {
	// 	wallet := Wallet{balance: BitCoin(20)}
	// 	wallet.Withdraw(BitCoin(10))
	// 	got := wallet.Balance()
	// 	want := BitCoin(10)

	// 	if got != want {
	// 		t.Errorf("got %s, wanted %s", got, want)
	// 	}
	// })

	wallet := Wallet{}
	assertBalance := func(t testing.TB, wallet Wallet, want BitCoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, wanted %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Errorf("Didn't get wanted error")
		}
		if got != want {
			t.Errorf("got %s, wanted %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet.Withdraw(BitCoin(10))
		assertBalance(t, wallet, BitCoin(0))
	})
	t.Run("Custom Error", func(t *testing.T) {
		err := wallet.Withdraw(500)
		assertError(t, err, ErrInsufficientFunds)
	})

}
