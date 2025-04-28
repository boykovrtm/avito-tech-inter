package user

import "time"

type User struct {
	username     string
	balance      int32
	boughtItems  []BoughtItems
	transactions []Transaction
}

type Transaction struct {
	id     int64
	from   string
	to     string
	amount int32
	moment time.Time
}

type BoughtItems struct {
	Name  string
	Count int32
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) GetBalance() int32 {
	return u.balance
}

func (u *User) SetBalance(balance int32) {
	u.balance = balance
}

func (u *User) GetTransactions() []Transaction {
	return u.transactions
}

func (u *User) SetTransactions(transactions []Transaction) {
	u.transactions = transactions
}

func (u *User) GetBoughtItems() []BoughtItems {
	return u.boughtItems
}

func (u *User) SetBoughtItems(boughtItems []BoughtItems) {
	u.boughtItems = boughtItems
}

func (t *Transaction) GetID() int64 {
	return t.id
}

func (t *Transaction) SetID(id int64) {
	t.id = id
}

func (t *Transaction) GetFrom() string {
	return t.from
}

func (t *Transaction) SetFrom(from string) {
	t.from = from
}

func (t *Transaction) GetTo() string {
	return t.to
}

func (t *Transaction) SetTo(to string) {
	t.to = to
}

func (t *Transaction) GetAmount() int32 {
	return t.amount
}

func (t *Transaction) SetAmount(amount int32) {
	t.amount = amount
}

func (t *Transaction) GetMoment() time.Time {
	return t.moment
}

func (t *Transaction) SetMoment(moment time.Time) {
	t.moment = moment
}

func NewUser(username string) *User {
	return &User{username: username, balance: 1000, boughtItems: []BoughtItems{}, transactions: []Transaction{}}
}
