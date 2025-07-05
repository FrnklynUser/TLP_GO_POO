package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Table    int
	Priority string
}

type Restaurant struct {
	urgentOrders chan Order
	normalOrders chan Order
	kitchenReady chan Order
}

func (r *Restaurant) processOrder(order Order, cookID int) {
	fmt.Printf("👨‍🍳 Cocinero %d procesando #%d...\n", cookID, order.ID)
	time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
	fmt.Printf("✅ Pedido #%d completado por cocinero %d\n", order.ID, cookID)
	r.kitchenReady <- order
}

func (r *Restaurant) cookWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case order := <-r.urgentOrders:
			r.processOrder(order, id)
		case order := <-r.normalOrders:
			r.processOrder(order, id)
		case <-time.After(1 * time.Second):
			fmt.Printf("💤 Cocinero %d sin trabajo, termina\n", id)
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	restaurant := &Restaurant{
		urgentOrders: make(chan Order),
		normalOrders: make(chan Order, 3),
		kitchenReady: make(chan Order),
	}

	const numWorkers = 2
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go restaurant.cookWorker(i, &wg)
	}

	orders := []Order{
		{ID: 1, Table: 1, Priority: "normal"},
		{ID: 2, Table: 2, Priority: "urgent"},
		{ID: 3, Table: 3, Priority: "normal"},
		{ID: 4, Table: 4, Priority: "urgent"},
	}

	for _, order := range orders {
		if order.Priority == "urgent" {
			restaurant.urgentOrders <- order
		} else {
			restaurant.normalOrders <- order
		}
	}

	close(restaurant.urgentOrders)
	close(restaurant.normalOrders)

	wg.Wait()
	close(restaurant.kitchenReady)

	fmt.Println("\n🎉 Todos los pedidos procesados:")
	for ready := range restaurant.kitchenReady {
		fmt.Printf("🍽 Pedido #%d servido\n", ready.ID)
	}
}
