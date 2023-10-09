//problem description
// a consumer will order a pizza, the order will be taken and there will be certain conditions that will decide
//if the pizza can be made or not, if it is being made the we will have to make the consumer wait andf channels will be used to communicate in between different go routines

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10 // no. of pizzas will be made in a day

var pizzasMade, pizzasFailed, total int // no. of pizzas made in a day and no. of pizzas failed in a day and the total no. of pizzas made in a day

type Producer struct {
	data chan pizzaOrder //data channel of type pizzaOrder
	quit chan chan error //quit channel holding a channel or errors(meaning this will quit making pizza due to any reason and that reason will be  accessed by error channel)
}

type pizzaOrder struct {
	pizzaNumber int    //to keep track which pizza is being made or failed
	message     string //to print the appropriate message whether the pizza was made or failed
	success     bool
}

//closing channel function i will be able to access the channels present in type producer with p and with the help of close channel i will be able to access these througout code

func (p *Producer) close() error {
	ch := make(chan error) //channel of type error
	p.quit <- ch           // storing the error message in the quit channel
	return <-ch            //then returnin whatever the message is

}

func makePizza(pizzaNumber int) *pizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recived order #%d\n", pizzaNumber)

		//generating a random number to ensure pizza's success or failure
		//if the random no. is less than 5 pizza fails else it is made
		rmd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rmd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		color.Cyan("Making pizza #%d. it will take %d seconds....\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		//if the pizza making failed, now we set the reasongs why it failed
		if rmd <= 2 {
			msg = fmt.Sprintf("*** we ran out of ingredients for pizza #%d", pizzaNumber)
		} else if rmd <= 4 {
			msg = fmt.Sprintf("*** The cook quit making pizza #%d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("*** pizza #%d is ready", pizzaNumber)
		}

		p := pizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &pizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	//keep track of which pizza we are making
	var i = 0
	//run forever or until we recieve a quit notification
	//try to make pizzas

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber

			select {
			//we tried to make pizza by sending something to data channel
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}

		//decision
	}
}

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//print a mesage
	color.Green("The Pizzeria is open!!!!")
	color.Yellow("===============================")

	//create a producer(pizza maker)
	pizzaJob := &Producer{
		data: make(chan pizzaOrder), // in pizzaJob the date(order will be sent and the message will be returned through quit channel whether the pizza making was quit by completing the order or failing the order)
		quit: make(chan chan error),
	}

	//run the producer in the background(meaning producer will be a go routine)
	go pizzeria(pizzaJob)

	//create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!!")
			}
		} else {
			color.Cyan("Done making pizzas.....")
			err := pizzaJob.close()

			if err != nil {
				color.Red("***Error closing channel", err)
			}
		}
	}

	//print out the ending message
	color.Yellow("========================================")
	color.Yellow("Done for the day!!!")

	color.Cyan("We made %d pizzas, but failed to make %d , with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Yellow("it was not a very good day....")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day....")
	case pizzasFailed >= 2:
		color.Green("It was a pretty good day!!!")
	default:
		color.Green("It was a great day!!!")
	}
}
