package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,
		},
	}))

	// Output:
	// 100000
	// 300000
	// 0
	// 0

}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: -10_000_00,
			Active:  true,
			PAN:     "4400 4301 2671 3749",
		},
		{
			Balance: 10_000_00,
			Active:  true,
			PAN:     "4405 6398 6834 0132",
		},
		{
			Balance: 10_000_00,
			Active:  false,
			PAN:     "5395 4524 1605 2228",
		},
	}

	sources := PaymentSources(cards)

	for _, source := range sources {
		fmt.Println(source.Number)
	}

	// Output:
	// 4405 6398 6834 0132

}
