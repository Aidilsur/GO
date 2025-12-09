package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

type contextKey string


func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, contextKey("b"), "B")
	contextC := context.WithValue(contextA, contextKey("c"), "C")

	contextD := context.WithValue(contextB, contextKey("d"), "D")
	contextE := context.WithValue(contextB, contextKey("e"), "E")

	contextF := context.WithValue(contextC, contextKey("f"), "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value(contextKey("f")))
	fmt.Println(contextF.Value(contextKey("c")))
	fmt.Println(contextF.Value(contextKey("b")))
	fmt.Println(contextA.Value(contextKey("b")))
}

func CreateCounter(ctx context.Context) chan int {
	destionation := make(chan int)

	go func() {
		defer close(destionation)

		counter :=1 
		for {
			select {
			case <- ctx.Done():
				return
			default:
				destionation <- counter
				counter++
			}
		}
	}()

	return destionation
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destionation := CreateCounter(ctx)

	for n := range destionation {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}