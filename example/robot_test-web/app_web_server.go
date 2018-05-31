package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认不会解析
	cmdName := r.Form.Get("Cmd")
	testCase := r.Form.Get("Case")
	if cmdName == "" || testCase == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "params error!") //返回
		return
	}
	cmd := exec.Command(cmdName, testCase)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error()) //返回
		return
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error()) //返回
		return
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error()) //返回
		return
	}
	body := string(opBytes)

	if body == "" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "body is empty") //返回
		return
	}
	strs := strings.Split(body, "\n")
	if len(strs) < 6 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "body is empty") //返回
		return
	}
	line := strs[len(strs)-6]
	//15 tests total, 15 passed, 0 failed
	reg := regexp.MustCompile(`(\d+) .+ (\d+) passed, (\d+) failed`)
	arr := reg.FindStringSubmatch(line)
	if len(arr) != 4 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "parse response body error.") //返回
		return
	}
	result := make(map[string]int)
	total, _ := strconv.Atoi(arr[1])
	result["total"] = total
	sucess, _ := strconv.Atoi(arr[2])
	result["sucess"] = sucess
	fail, _ := strconv.Atoi(arr[3])
	result["fail"] = fail
	if result["fail"] > 0 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	res, _ := json.Marshal(result)
	fmt.Fprintf(w, string(res)) //返回
}

func main() {
	http.HandleFunc("/monitor", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Listen error : ", err)
	}
}
