package main

import "fmt"

// Размеры дома
type Dimensions struct {
	Width  float64
	Height float64
	Depth  float64
}

// Мебель
type Furniture struct {
	Name   string
	Width  float64
	Height float64
	Depth  float64
}

// Бытовая техника
type Appliance struct {
	Name   string
	Width  float64
	Height float64
	Depth  float64
}

// Состав семьи
type FamilyMember struct {
	Name string
	Age  int
}

// Дом
type House struct {
	Size       Dimensions
	Furnitures []Furniture
	Appliances []Appliance
	Family     []FamilyMember
}

func main() {
	myHouse := House{
		Size: Dimensions{Width: 35, Height: 3, Depth: 25},
		Furnitures: []Furniture{
			{Name: "Диван", Width: 2, Height: 1, Depth: 0.8},
			{Name: "Стол", Width: 1.5, Height: 0.7, Depth: 1},
			{Name: "Стул", Width: 0.5, Height: 0.9, Depth: 0.5},
			{Name: "Кровать", Width: 2, Height: 0.5, Depth: 1.8},
			{Name: "Шкаф", Width: 1, Height: 2.2, Depth: 0.6},
		},
		Appliances: []Appliance{
			{Name: "Телевизор", Width: 1.2, Height: 0.8, Depth: 0.1},
			{Name: "Холодильник", Width: 0.8, Height: 2, Depth: 0.7},
			{Name: "Стиральная машина", Width: 0.6, Height: 0.9, Depth: 0.5},
			{Name: "Микроволновка", Width: 0.5, Height: 0.3, Depth: 0.4},
			{Name: "Кондиционер", Width: 1, Height: 0.3, Depth: 0.2},
		},
		Family: []FamilyMember{
			{Name: "Владислав", Age: 23},
			{Name: "Владислава", Age: 23},
			{Name: "Эрнест (кот)", Age: 2},
			{Name: "Букля (кот)", Age: 10},
		},
	}

	DescribeHouse(myHouse)
}

func DescribeHouse(h House) {
	fmt.Println("Описание квартиры:")
	fmt.Printf("Размеры: Ширина - %v м, Высота - %v м, Глубина - %v м\n", h.Size.Width, h.Size.Height, h.Size.Depth)

	fmt.Println("\nМебель:")
	for _, f := range h.Furnitures {
		fmt.Printf("%v: Ширина - %v м, Высота - %v м, Глубина - %v м\n", f.Name, f.Width, f.Height, f.Depth)
	}

	fmt.Println("\nБытовая техника:")
	for _, a := range h.Appliances {
		fmt.Printf("%v: Ширина - %v м, Высота - %v м, Глубина - %v м\n", a.Name, a.Width, a.Height, a.Depth)
	}

	fmt.Println("\nСостав семьи:")
	for _, fm := range h.Family {
		fmt.Printf("Имя: %v, Возраст: %v лет\n", fm.Name, fm.Age)
	}
}
