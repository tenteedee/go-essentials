package bank

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/tenteedee/go-essentials/utils"
)

const filename = "bank/balance.txt"

var balance float64

func Main() {
	var err error
	balance, err = utils.ReadFloatFromFile(filename)

	if err != nil {
		fmt.Println(err)
		panic("Cannot continue without balance")
	}

	fmt.Println("====== Welcome to Go Bank ======")
	fmt.Println("Reach us 24/7 on ", randomdata.PhoneNumber())

	for {
		var choice = menu()

		// if choice == 1 {
		// 	checkBalance()
		// } else if choice == 2 {
		// 	deposit()
		// } else if choice == 3 {
		// 	withdraw()
		// } else if choice == 4 {
		// 	fmt.Println("Goodbye!")
		// 	break
		// } else {
		// 	fmt.Println("Invalid choice")
		// }

		switch choice {
		case 1:
			checkBalance()
		case 2:
			deposit()
		case 3:
			withdraw()
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for banking with us!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func checkBalance() {
	fmt.Printf("Your balance is %.2f\n", balance)
}

func deposit() {
	var amount float64
	for {
		fmt.Print("Enter amount to deposit: ")
		fmt.Scan(&amount)

		if amount <= 0 {
			fmt.Println("Invalid amount")
		} else {
			break
		}
	}

	balance += amount
	utils.WriteFloatToFile(filename, balance)
	fmt.Printf("Deposit successful. Your balance is %.2f\n", balance)
}

func withdraw() {
	var amount float64
	for {
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scan(&amount)
		if amount > balance {
			fmt.Println("Insufficient funds. Please enter a valid amount.")
		} else {
			break
		}
	}

	balance -= amount
	utils.WriteFloatToFile(filename, balance)
	fmt.Printf("Withdrawal successful. Your balance is %.2f\n", balance)

}

func menu() int {
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return choice
}
