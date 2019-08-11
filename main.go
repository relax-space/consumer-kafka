package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	notifyKill()
	fmt.Printf(">> env  addr:%v,port:%v,topic:%v \n", os.Getenv("KAFKA_HOST"), os.Getenv("PORT"), os.Getenv("TOPIC"))
	config := Config{
		Brokers: []string{os.Getenv("KAFKA_HOST") + ":" + os.Getenv("PORT")},
		Topic:   os.Getenv("TOPIC"),
	}
	if err := Consume("consumer-kafka", config,
		EventFruit,
	); err != nil {
		panic(err)
	}

	fmt.Println("consumer is ready!")

	time.Sleep(100 * time.Hour)
}

func notifyKill() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt)
	go func() {
		for s := range signals {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				os.Exit(0)
			}
		}
	}()

}
