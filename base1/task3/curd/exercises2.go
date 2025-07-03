package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Balance float64 `gorm:"type:decimal(10,2)"`
}

type Transaction struct {
	FromAccountID uint
	FromAccount   Account `gorm:"foreignKey:FromAccountID"`
	ToAccountID   uint
	ToAccount     Account `gorm:"foreignKey:ToAccountID"`
	Amount        float64 `gorm:"type:decimal(10,2)"`
}

func (a *Account) Transfer(tx *gorm.DB, toAccount *Account, amount float64) bool {

	if toAccount == nil {
		fmt.Println("ToAccount is nil")
		return false
	}

	if amount <= 0 {
		fmt.Println("转账金额错误: ", amount)
		return false
	}

	if amount > toAccount.Balance {
		fmt.Println("客户余额不足, 当前账户余额为:", toAccount.Balance, ",你要转账的金额为:", amount)
		return false
	}

	err := tx.Transaction(func(tx *gorm.DB) error {
		tx.Create(&Transaction{
			FromAccountID: a.ID,
			ToAccountID:   toAccount.ID,
			Amount:        amount,
		})
		tx.Debug().Model(a).Where("balance = ?", a.Balance).Update("balance", a.Balance-amount)
		tx.Debug().Model(toAccount).Where("balance = ?", toAccount.Balance).Update("balance", toAccount.Balance+amount)
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/leaning-go?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Account{})
	err = db.AutoMigrate(&Transaction{})
	if err != nil {
		panic(err)
	}

	// 先清空数据
	db.Exec("SET FOREIGN_KEY_CHECKS=0")
	db.Exec("TRUNCATE TABLE transactions")
	db.Exec("TRUNCATE TABLE accounts")
	db.Exec("SET FOREIGN_KEY_CHECKS=1")

	db.Create(&Account{
		Balance: 1000,
	})

	db.Create(&Account{
		Balance: 1000,
	})

	var fromAccount1 Account
	var toAccount1 Account
	res1 := db.First(&fromAccount1, "id = ?", 1)
	if res1.Error != nil {
		fmt.Println(res1.Error)
	}
	res2 := db.First(&toAccount1, "id = ?", 2)
	if res2.Error != nil {
		fmt.Println(res2.Error)
	}

	fromAccount1.Transfer(db, &toAccount1, 100)
}
