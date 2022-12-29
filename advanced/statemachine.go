package advanced

import "fmt"

func LearnStatemachine() {

	// Example for a bank account
	//startAccountBalance()

	//
	sm := New()
	fmt.Println(sm.foo())
	sm.state = "A"
	fmt.Println(sm.foo())
	close(sm.quitc)

}

func startAccountBalance() {
	accountBalance := 0
	// incCh to increase our account balance
	incCh := make(chan int)

	// outCh to decrease our account balance
	outCh := make(chan int)

	// readCh to read our account balance;
	readCh := make(chan chan int)

	for {
		select {
		case i := <-incCh:
			accountBalance = accountBalance + i
		case i := <-outCh:
			accountBalance = accountBalance - i
		case ch := <-readCh:
			ch <- accountBalance
		}
	}
}

type stateMachine struct {
	state   string
	actionc chan func()
	quitc   chan struct{}
}

func (sm *stateMachine) loop() {
	for {
		select {
		case f := <-sm.actionc:
			f()
		case <-sm.quitc:
			return
		}
	}
}

func (sm *stateMachine) foo() string {
	result := make(chan string)
	sm.actionc <- func() {
		if sm.state == "A" {
			sm.state = "B"
		}
		result <- sm.state
	}
	return <-result
}

func New() *stateMachine {
	sm := &stateMachine{
		state:   "initial",
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go sm.loop()
	return sm
}
