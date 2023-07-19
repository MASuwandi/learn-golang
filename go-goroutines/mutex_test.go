package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Mutex
// Sync Mutex
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Counter = ", x)
}

// RW Mutex (Read Write Mutex)
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

// Locking Write
func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

// Locking Read
func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println("After Increment", account.GetBalance())
			}
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Final Balnce: ", account.GetBalance())
}

// Deadlock
// simulation
type UserBalance struct {
	sync.Mutex // Mutex sync.Mutex
	Name       string
	Balance    int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("user1 Lock: ", user1.Name)
	user1.Change(-amount)

	time.Sleep(500 * time.Millisecond)

	user2.Lock()
	fmt.Println("user2 Lock: ", user2.Name)
	user2.Change(amount)

	time.Sleep(500 * time.Millisecond)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Rocket",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Racoon",
		Balance: 1000000,
	}

	fmt.Println("user1: ", user1.Name, ", Balance: ", user1.Balance)
	fmt.Println("user2: ", user2.Name, ", Balance: ", user2.Balance)

	// Scenario 1
	// Transfer(&user1, &user2, 100000)

	// Scenario 2 deadlock
	// lock user 1 and 2 simultanously
	// wait unlock forever
	Transfer(&user1, &user2, 100000)
	Transfer(&user2, &user1, 200000)

	time.Sleep(1250 * time.Millisecond)

	fmt.Println("user1: ", user1.Name, ", Balance: ", user1.Balance)
	fmt.Println("user2: ", user2.Name, ", Balance: ", user2.Balance)
}
