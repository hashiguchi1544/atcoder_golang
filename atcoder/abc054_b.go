package atcoder

import "fmt"

func abc054a() {
	var aliceCard, bobCard int
	fmt.Scan(&aliceCard, &bobCard)

	if aliceCard == 1 {
		aliceCard = 14
	}

	if bobCard == 1 {
		bobCard = 14
	}

	if aliceCard > bobCard {
		fmt.Println("Alice")
	} else if bobCard > aliceCard {
		fmt.Println("Bob")
	} else {
		fmt.Println("Draw")
	}
}
