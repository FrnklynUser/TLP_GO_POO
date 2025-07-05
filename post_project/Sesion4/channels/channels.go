package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID       int
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

	// Simular la llegada de pedidos
	go func() {
		for i := 1; i <= 5; i++ {
			priority := "normal"
			if i%2 == 0 {
				priority = "urgent"
			}
			restaurant.urgentOrders <- Order{ID: i, Priority: priority}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Esperar a que todos los trabajadores terminen
	wg.Wait()
}
