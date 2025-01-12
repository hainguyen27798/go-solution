package examples

import (
	"fmt"
	"github.com/hainguyen27798/go-solution/internal/utils"
	"time"
)

type EmailTask struct {
	Email   string
	Subject string
}

func (e *EmailTask) Process(workerId int) {
	fmt.Printf("Worker #%d - Sending email task: %s\n", workerId, e.Subject)
	time.Sleep(time.Second * 1)
	fmt.Printf("Worker #%d - Sending email task: %s - Done\n", workerId, e.Subject)
}

type ImageTask struct {
	ImageURL string
}

func (i *ImageTask) Process(workerId int) {
	fmt.Printf("Worker #%d - Processing image task: %s\n", workerId, i.ImageURL)
	time.Sleep(time.Second * 10)
	fmt.Printf("Worker #%d - Processing image task: %s - Done\n", workerId, i.ImageURL)
}

func RunWorkerPoolExample() {
	tasks := []utils.Task{
		&EmailTask{Subject: "Hello", Email: "test@gmail.com"},
		&EmailTask{Subject: "Hi", Email: "test2@gmail.com"},
		&ImageTask{ImageURL: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"},
		&EmailTask{Subject: "Goodbye", Email: "test3@gmail.com"},
		&ImageTask{ImageURL: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"},
		&EmailTask{Subject: "Bye", Email: "test4@gmail.com"},
		&EmailTask{Subject: "Bye2", Email: "test5@gmail.com"},
		&ImageTask{ImageURL: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"},
		&EmailTask{Subject: "Bye3", Email: "test6@gmail.com"},
	}

	wp := utils.WorkerPool{
		Tasks:       tasks,
		Concurrency: 3,
	}

	wp.Run()
}
