package main

import (
	"fmt"

	"github.com/ramrodo/tech-assessment-loan-startup/utils"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type Investment struct {
	amount int32
}

type NotCombinationFoundError struct{}

func (e *NotCombinationFoundError) Error() string {
	return "No se logró encontrar una combinación válida"
}

func (Investment) Assign(investment int32) (int32, int32, int32, error) {
	if investment < 100 {
		return 0, 0, 0, &NotCombinationFoundError{}
	}

	credit_type_300_max := int(investment / 300)
	credit_type_500_max := int(investment / 500)
	credit_type_700_max := int(investment / 700)

	// Get combinations until we find the first valid one, otherwise, return an error
	var combination int32

	for i := 0; i <= credit_type_300_max; i++ {
		for j := 0; j <= credit_type_500_max; j++ {
			for k := 0; k <= credit_type_700_max; k++ {
				combination = int32((300 * i) + (500 * j) + (700 * k))

				if combination == investment {
					return int32(i), int32(j), int32(k), nil
				}

				if combination > investment {
					break
				}
			}
		}
	}

	return 0, 0, 0, &NotCombinationFoundError{}
}

func main() {
	investment := Investment{}
	credit_type_300, credit_type_500, credit_type_700, err := investment.Assign(1100)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Puedes asignar:")
		if credit_type_300 > 0 {
			fmt.Printf("\n%d crédito%s de $300", credit_type_300, utils.Plural(credit_type_300))
		}
		if credit_type_500 > 0 {
			fmt.Printf("\n%d crédito%s de $500", credit_type_500, utils.Plural(credit_type_500))
		}
		if credit_type_700 > 0 {
			fmt.Printf("\n%d crédito%s de $700", credit_type_700, utils.Plural(credit_type_700))
		}
		fmt.Println()
	}
}
