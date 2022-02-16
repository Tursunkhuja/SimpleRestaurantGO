package employee

import (
	"SimpleRestaurantGo/pkg/menu"
	"reflect"
	"strconv"
)

type Employee struct {
	Name       string
	Age        int
	Order      interface{} //The empty interface denoted by interface{} can hold values of any type
	orderCount int
}

func (e *Employee) NewRequest(menuItem string, quantity int) interface{} {
	e.orderCount++

	if e.orderCount%3 == 0 {
		if menuItem == "chicken" {
			egg := menu.Egg{}
			egg.Init(quantity)
			e.Order = egg
		}
		if menuItem == "egg" {
			chicken := menu.Chicken{}
			chicken.Init(quantity)
			e.Order = chicken
		}
	} else {
		if menuItem == "chicken" {
			chicken := menu.Chicken{}
			chicken.Init(quantity)
			e.Order = chicken
		}
		if menuItem == "egg" {
			egg := menu.Egg{}
			egg.Init(quantity)
			e.Order = egg
		}
	}

	return e.Order
}

func (e *Employee) CopyRequest() interface{} {
	if typeofObject(e.Order) == "Egg" {
		egg := reflect.ValueOf(e.Order).Interface().(menu.Egg)

		newEgg := menu.Egg{}
		newEgg.Init(egg.GetQuantity())
		newEgg.Quality = egg.GetQuality()
		return newEgg
	}
	if typeofObject(e.Order) == "Chicken" {
		chicken := reflect.ValueOf(e.Order).Interface().(menu.Chicken)

		newChicken := menu.Chicken{}
		newChicken.Init(chicken.Quantity)
		return newChicken
	}

	return nil
}

func (e *Employee) Prepare() bool {
	if typeofObject(e.Order) == "Chicken" {
		chicken := reflect.ValueOf(e.Order).Interface().(menu.Chicken)
		for i := 0; i < chicken.GetQuantity(); i++ {
			chicken.CutUp()
		}
		chicken.Cook()
		return true
	}
	if typeofObject(e.Order) == "Egg" {
		egg := reflect.ValueOf(e.Order).Interface().(menu.Egg)
		for i := 0; i < egg.GetQuantity(); i++ {
			egg.Crack()
			egg.DiscardShell()
		}
		egg.Cook()
		return true
	}
	return false
}

func (e *Employee) Inspect() string {
	result := "no inspection is required"

	if typeofObject(e.Order) == "Egg" {
		egg := reflect.ValueOf(e.Order).Interface().(menu.Egg)
		result = strconv.Itoa(egg.GetQuality())
	}

	return result
}

// Using reflect.TypeOf()
func typeofObject(variable interface{}) string {
	return reflect.TypeOf(variable).String()
}
