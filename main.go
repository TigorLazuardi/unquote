package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var verbose bool

func init() {
	flag.BoolVar(&verbose, "v", false, "verbose logging")
	flag.Parse()
	level := "info"

	if verbose {
		level = "debug"
	}

	ll, err := log.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	log.SetFormatter(&log.TextFormatter{
		QuoteEmptyFields: true,
	})
	log.SetLevel(ll)
}

func main() {
	stdin := readStdIn()
	argInput := readArgs()
	if stdin == nil && argInput == "" {
		log.Fatal("empty input")
	}

	if len(stdin) > 0 {
		work(string(stdin))
	} else {
		work(argInput)
	}
}

func work(s string) {
	val := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &val)
	if err != nil {
		c := handleString(s)
		if m, ok := c.(map[string]interface{}); ok {
			fmt.Println(marshaller(handleJSONInput(m)))
		} else {
			fmt.Println(marshaller(c))
		}
	} else {
		fmt.Println(marshaller(handleJSONInput(val)))
	}
}

func marshaller(v interface{}) interface{} {
	log.Debug(v)
	if s, ok := v.(map[string]interface{}); ok {
		e, _ := json.Marshal(s)
		return string(e)
	}
	return v
}

func readStdIn() []byte {
	info, err := os.Stdin.Stat()
	if err != nil {
		log.Panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		log.Trace("empty pipe input")
		return nil
	}
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err.Error())
	}
	return input
}

func readArgs() string {
	args := flag.Args()
	if len(args) < 1 {
		log.Debug("empty input from arguments")
		return ""
	}
	return args[0]
}

// handleJSONInput function
func handleJSONInput(payload map[string]interface{}) interface{} {
	y := make(map[string]interface{})
	for k, v := range payload {
		if m, ok := v.(map[string]interface{}); ok {
			v = handleJSONInput(m)
		}
		if s, ok := v.(string); ok {
			v = handleString(s)
		}
		y[k] = v
	}
	return y
}

func handleString(s string) interface{} {
	x, err := strconv.Unquote(s)
	// Not Escaped Values
	if err != nil {
		log.WithFields(map[string]interface{}{
			"value": s,
		}).Debug(err.Error())
		x = s
	}

	m := make(map[string]interface{})
	err = json.Unmarshal([]byte(x), &m)
	// Not JSON Object
	if err != nil {
		log.Trace(err)
		return x
	}
	return handleJSONInput(m)
}
