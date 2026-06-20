Just reading about interfaces and implementing them once or twice does not really help..

Here is the idiomatic way of implementing the interfaces **WHEN THE TIME IS RIGHT!!**

### 1. Create structs according to the requirement

```go
type DebitCard struct {
    cardNumber string
    isActive   bool
    balance    int
}

type UpiId struct {
    balance int
}
```

### 2. Implementing the methods

When the struct contains business data that is better to be not exposed directly, we may opt to go for methods.

```go
func (debitCard *DebitCard) canPurchase(price int) bool {
    if debitCard.balance >= price && debitCard.isActive {
        return true
    }
    return false
}

func (upiId *UpiId) canPurchase(price int) bool {
    if upiId.balance >= price {
        return true
    }
    return false
}
```

### 3. REALISATION...

Two structs have different data, but their methods have similar functionality. Would it not be better to have them grouped in some way?

```go
type financeOperations interface {
    canPurchase(amount int) bool
}
```

---

Now we have the interface that links the method and the structs have already implemented those functions.

How about we create a common function that will call the interface? The functionality behind it can differ depending on the struct we pass.

---

### 4. Interface level function definition

```go
func purchasePrecheck(f financeOperations, amount int) {
    if f.canPurchase(amount) {
        fmt.Println("MONEY MONEY MONEY!")
        return
    }

    fmt.Println("BROKE")
}
```

```go
// main.go

func main() {
    newCard := &DebitCard{
        cardNumber: "fub23nvsie",
        isActive:   true,
        balance:    21341,
    }

    purchasePrecheck(newCard, 1212)
}
```
