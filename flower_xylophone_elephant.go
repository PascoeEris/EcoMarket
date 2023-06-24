package main

import "fmt"

// Represents a product or service offered on the marketplace
type Product struct {
	Name string
	Price int
	Category string
	Description string
	EcoRating int
}

// Represents a user profile.
type User struct {
	Name string
	Address string
	Email string
	Phone string
	Favorites []Product
}

// Represents an individual order.
type Order struct {
	Id string
	User User
	Products []Product
	TotalPrice int
	Status string
}

// Represents the marketplace.
type Marketplace struct {
	Name string
	Products []Product
	Users []User
	Orders []Order
}

func (m *Marketplace) AddProduct(p Product) {
	m.Products = append(m.Products, p)
}

func (m *Marketplace) UpdateProduct(p Product) error {
	for i, prod := range m.Products {
		if prod.Name == p.Name {
			m.Products[i] = p
			return nil
		}
	}
	return fmt.Errorf("Product not found")	
}

func (m *Marketplace) RemoveProduct(name string) error {
	for i, prod := range m.Products {
		if prod.Name == name {
			m.Products = append(m.Products[:i], m.Products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Product not found")
}

func (m *Marketplace) AddUser(u User) {
	m.Users = append(m.Users, u)
}

func (m *Marketplace) UpdateUser(u User) error {
	for i, user := range m.Users {
		if user.Name == u.Name {
			m.Users[i] = u
			return nil
		}
	}
	return fmt.Errorf("User not found")
}

func (m *Marketplace) RemoveUser(name string) error {
	for i, user := range m.Users {
		if user.Name == name {
			m.Users = append(m.Users[:i], m.Users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User not found")
}

func (m *Marketplace) AddOrder(o Order) {
	m.Orders = append(m.Orders, o)
}

func (m *Marketplace) UpdateOrder(o Order) error {
	for i, order := range m.Orders {
		if order.Id == o.Id {
			m.Orders[i] = o
			return nil
		}
	}
	return fmt.Errorf("Order not found")	
}

func (m *Marketplace) RemoveOrder(id string) error {
	for i, order := range m.Orders {
		if order.Id == id {
			m.Orders = append(m.Orders[:i], m.Orders[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Order not found")
}

func main() {
	m := Marketplace{
		Name: "Green Marketplace",
		Products: []Product{},
		Users: []User{},
		Orders: []Order{},
	}
	fmt.Println(m)
}