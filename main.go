package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// Timer structure
type Timer struct {
	isFinished bool
}

func main() {

	myTimer := &Timer{isFinished: false}
	go myTimer.timer()

	// for !myTimer.isFinished {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println(myTimer.isFinished)
	// }

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/envs", environmentVariables)
	http.HandleFunc("/readdir", readDir)
	http.HandleFunc("/readfile", readFile)
	http.HandleFunc("/healthcheckok", healthcheckok)
	http.HandleFunc("/healthcheckfail", myTimer.healthcheckfail)

	http.ListenAndServe(":8000", nil)
}

// just check if it works
// localhost/
func helloWorld(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello world!\n")
}

// readDir lists all items in a dir from get request
// localhost/readdir?directory=dirname
func readDir(writer http.ResponseWriter, request *http.Request) {
	directory := request.FormValue("q")
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
		io.WriteString(writer, err.Error())
	} else {
		filenames := make([]string, len(files))
		for i, f := range files {
			filenames[i] = f.Name()
		}
		io.WriteString(writer, "content of "+directory+":\n"+strings.Join(filenames, "\n"))
	}

}

func readFile(writer http.ResponseWriter, request *http.Request) {
	file := request.FormValue("q")
	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		io.WriteString(writer, err.Error())
	} else {
		io.WriteString(writer, "content of "+file+":\n"+string(fileContent))
	}

}

// environmentVariables lists all environmental variables in system
// localhost/envs
func environmentVariables(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, strings.Join(os.Environ(), "\n"))
}

// healthcheckok returns 200
// localhost/envs
func healthcheckok(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
}

// healthcheckfail returns 200 and after 3 minutes switch to returning 400
// localhost/envs
func (t *Timer) healthcheckfail(writer http.ResponseWriter, request *http.Request) {
	if t.isFinished {
		writer.WriteHeader(400)
	} else {
		writer.WriteHeader(200)
	}
}

func (t *Timer) timer() {
	time.Sleep(3 * time.Minute)
	t.isFinished = true
}
