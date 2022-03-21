package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Data struct {
	Fdata string
}

type Rt struct {
	Status string
	Url    string
}

func facehandler(w http.ResponseWriter, r *http.Request) {
	h := r.Method
	if h == "GET" {
		fmt.Println("[GET '/face'] id = " + r.URL.Query().Get("id"))
		res := r.URL.Query().Get("id")
		buf, err := ioutil.ReadFile("./face/" + res + ".png")

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(buf)
	} else if h == "POST" {
		fmt.Println("acces '/'dirctry")
		fileInfos, err := ioutil.ReadDir("./face")
		l := len(fileInfos)
		if err != nil {
			fmt.Println(err)
			return
		}
		fp, err := os.Create("./face/" + strconv.Itoa(l) + ".png")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fp.Close()
		mf := r.MultipartForm
		for k, v := range mf.Value {
			fmt.Printf("%v : %v", k, v)
		}
		//err = png.Encode(fp, dst)
		if err != nil {
			stcData := Rt{Status: "err", Url: "http://face/id=" + strconv.Itoa(l)}
			res, _ := json.Marshal(stcData)
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		}
		stcData := Rt{Status: "OK", Url: "http://face/id=" + strconv.Itoa(l)}
		res, err := json.Marshal(stcData)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func eventhandler(w http.ResponseWriter, r *http.Request) {
	h := r.Method
	if h == "GET" {
		fmt.Println("[GET '/event'] id = " + r.URL.Query().Get("id"))
		res := r.URL.Query().Get("id")
		buf, err := ioutil.ReadFile("./event/" + res + ".png")

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(buf)
	} else if h == "POST" {
		fmt.Println("acces '/'dirctry")
		fileInfos, err := ioutil.ReadDir("./event")
		l := len(fileInfos)
		if err != nil {
			fmt.Println(err)
			return
		}
		fp, err := os.Create("./event/" + strconv.Itoa(l) + ".png")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fp.Close()
		mf := r.MultipartForm
		for k, v := range mf.Value {
			fmt.Printf("%v : %v", k, v)
		}
		//err = png.Encode(fp, dst)
		if err != nil {
			stcData := Rt{Status: "err", Url: "http://event/id=" + strconv.Itoa(l)}
			res, _ := json.Marshal(stcData)
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		}
		stcData := Rt{Status: "OK", Url: "http://event/id=" + strconv.Itoa(l)}
		res, err := json.Marshal(stcData)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func main() {
	http.HandleFunc("/face", facehandler)
	http.HandleFunc("/event", eventhandler)
	http.ListenAndServe(":8080", nil)
}
