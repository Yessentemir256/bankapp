package card

import (
	"bank/pkg/bank/types"
)

// Total вычисляет общую сумму средств на всех картах.
// Отрицательные суммы и суммы на заблокированных картах игнорируются
func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if card.Balance <= 0 {
			continue
		}

		if !card.Active {
			continue
		}
		sum += card.Balance
	}
	return sum
}
