package Structures

import (
	"encoding/json"
	"fmt"

	"strconv"
)

type Account struct {
	Dpi      float64 `json:"Dpi"`
	Email    string  `json:"Correo"`
	Password string  `json:"Password"`
	Name     string  `json:"Nombre"`
	User     string  `json:"Cuenta"`
}

type NodeAccounts struct {
	Accounts [5]*Account
	Nodes    [5]*NodeAccounts
	Father   *NodeAccounts
	Leaf     bool
	N        int
}

func (this *NodeAccounts) Init(father *NodeAccounts) {
	//this.Accounts = make([]*Account, 5)
	//this.Nodes = make([]*NodeAccounts, 5)
	this.Father = father
	this.Leaf = true
}

func (this *NodeAccounts) SearchAccount(dpi float64, password string) *Account {
	found := &Account{}

	if this == nil {
		return found
	}

	if this.Leaf {
		for i := 0; i < this.N; i++ {
			if this.Accounts[i].Dpi == dpi && this.Accounts[i].Password == password {
				found = this.Accounts[i]
				break
			}
		}
	} else {
		encontrado := false
		for i := 0; i < this.N; i++ {
			if dpi == this.Accounts[i].Dpi && password == this.Accounts[i].Password {
				encontrado = true
				found = this.Accounts[i]
				break
			}
		}

		if !encontrado {
			for i := 0; i < this.N; i++ {
				if dpi < this.Accounts[i].Dpi {
					encontrado = true
					found = this.Nodes[i].SearchAccount(dpi, password)
					break
				}
			}
		}

		if !encontrado {
			found = this.Nodes[this.N].SearchAccount(dpi, password)
		}
	}

	return found
}

func (this *Account) AccountToJson() string {
	json, _ := json.MarshalIndent(this, "", "\t")

	return string(json)
}

func (this *NodeAccounts) Add(father *NodeAccounts, newAccount *Account) *NodeAccounts {
	if this.Leaf {
		fmt.Println("Se está añadiendo: ")
		fmt.Println(&newAccount)
		this = this.InsertAccount(newAccount)
	} else {
		fmt.Println("Se esta buscando hoja")
		found := false

		for i := 0; i < this.N; /*-1*/ i++ {
			if newAccount.Dpi < this.Accounts[i].Dpi {
				found = true
				this.Nodes[i].Add(this, newAccount)
				break
			}
		}

		if !found {
			this.Nodes[this.N].InsertAccount(newAccount)
		}
	}

	if this.N == 5 {
		fmt.Println("Se está balanceando")
		if this.Father == nil {
			c := this
			this = nil
			this = &NodeAccounts{}
			this.Init(nil)

			fmt.Println("*/*/*/*/*/*/*/*/*/*/")
			fmt.Println(c)
			fmt.Println(this)
			fmt.Println("*/*/*/*/*/*/*/*/*/*/")

			this = this.InsertAccount(c.Accounts[2])
			this.Nodes[0] = &NodeAccounts{
				//Accounts: make([]*Account, 5),
				//Nodes:    make([]*NodeAccounts, 5),
				Father: this,
				Leaf:   true,
			}

			this.Nodes[1] = &NodeAccounts{
				//Accounts: make([]*Account, 5),
				//Nodes:    make([]*NodeAccounts, 5),
				Father: this,
				Leaf:   true,
			}

			this.Nodes[0] = this.Nodes[0].InsertAccount(c.Accounts[0])
			this.Nodes[0] = this.Nodes[0].InsertAccount(c.Accounts[1])
			this.Nodes[1] = this.Nodes[1].InsertAccount(c.Accounts[3])
			this.Nodes[1] = this.Nodes[1].InsertAccount(c.Accounts[4])

			this.Leaf = false
			//this.N = 2
		} else {
			mkey := this.Accounts[2]
			this.Father = this.Father.InsertAccount(mkey)
			index := 0

			for index = 0; index < this.Father.N; index++ {
				if this.Father.Accounts[index] == mkey {
					break
				}
			}

			for i := this.Father.N; i > index+1; i-- {
				this.Father.Nodes[i] = this.Father.Nodes[i-1]
			}

			this.Father.Nodes[index+1] = &NodeAccounts{
				//Accounts: make([]*Account, 5),
				//Nodes:    make([]*NodeAccounts, 5),
				Father: this.Father,
				Leaf:   true,
			}

			this.Father.Nodes[index+1] = this.Father.Nodes[index+1].InsertAccount(this.Accounts[3])
			this.Father.Nodes[index+1] = this.Father.Nodes[index+1].InsertAccount(this.Accounts[4])

			aux := this

			this.Father.Nodes[index] = &NodeAccounts{
				//Accounts: make([]*Account, 5),
				//Nodes:    make([]*NodeAccounts, 5),
				Father: this.Father,
				Leaf:   true,
			}

			this.Father.Nodes[index] = this.Father.Nodes[index].InsertAccount(aux.Accounts[0])
			this.Father.Nodes[index] = this.Father.Nodes[index].InsertAccount(aux.Accounts[1])
		}
	}

	this.N = this.NumberOfAccounts()

	return this
}

func (this *NodeAccounts) Show(h int) {
	fmt.Println("Level " + strconv.Itoa(h) + ": ")

	fmt.Print("[ ")
	for i := 0; i < this.N; i++ {
		fmt.Print(this.Accounts[i].Dpi)
		fmt.Print(" ")
	}

	fmt.Println(" ]")

	for i := 0; i < 5; i++ {
		if this.Nodes[i] != nil {
			this.Nodes[i].Show(h + 1)
		}
	}

}

func (this *NodeAccounts) InsertAccount(new *Account) *NodeAccounts {
	this.Accounts[this.N] = new
	this.N++
	if this.N > 1 {
		this.Sort()
	}

	fmt.Println("==========================================")
	for i := 0; i < this.N; i++ {
		fmt.Println(this.Accounts[i])
		fmt.Println(this.N)
	}
	fmt.Println("==========================================")

	return this
}

func (this *NodeAccounts) Sort() {

	var i, j int
	var temp *Account

	for i = 0; i < this.N; i++ {
		for j = 0; j < this.N-i-1; j++ {
			if this.Accounts[j].Dpi > this.Accounts[j+1].Dpi {
				temp = this.Accounts[j]
				this.Accounts[j] = this.Accounts[j+1]
				this.Accounts[j+1] = temp
			}
		}
	}

}

/*func (this *NodeAccounts) Init() {
	this.Accounts = make([]Account, 5)
	this.Nodes = make([]NodeAccounts, 5)
}*/

func (this *NodeAccounts) NumberOfNodes() int {
	j := 0
	for i := 0; i < 5; i++ {
		if this.Nodes[i] != nil {
			j++
		}
	}

	return j
}

func (this *NodeAccounts) NumberOfAccounts() int {
	j := 0
	for i := 0; i < 5; i++ {
		if this.Accounts[i] != nil {
			j++
		}
	}

	return j
}
