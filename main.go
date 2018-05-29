package main

import (
	"fmt"
	"time"
	"os/signal"
	"os"
	"syscall"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"./task_manager"
)

func main() {
	fmt.Println("start...")
	tm:=task_manager.NewTaskManager()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		buf,err:=ioutil.ReadAll( request.Body)
		if err!=nil{
			fmt.Println(err)
			return
		}

		var req task_manager.Request
		json.Unmarshal(buf,&req)
		if err!=nil{
			fmt.Println(err)
			return
		}

		var respon task_manager.Request
		if request.Method==http.MethodDelete{
			err:=tm.DelTask(req.ID)
			if err!=nil{
				respon.Ok=false
				respon.Error=err.Error()
			}else{
				respon.Ok=true
				respon.ID=req.ID
			}
		}else if request.Method==http.MethodPost{
			err:=tm.AddTask(&task_manager.Task{
				ID:req.ID,
				NextTime:time.Now().UnixNano(),
				Interval:1000000*req.Interval,
				Cmd:req.Cmd,
				Args:req.Args,
				IsActive:true,
			})
			if err!=nil{
				respon.Ok=false
				respon.Error=err.Error()
			}else{
				respon.Ok=true
				respon.ID=req.ID
			}
		}else {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		buf,err=json.Marshal(respon)
		if err!=nil{
			fmt.Println(err)
			return
		}

		writer.Write(buf)
	})
	http.ListenAndServe(":8103",nil)

	c := make(chan os.Signal, 0)
	signal.Notify(c,syscall.SIGINT, syscall.SIGTERM)
	<-c
	close(tm.ExitChan)
	tm.WG.Wait()

	fmt.Println("close")

}

//curl -X DELETE localhost:4567 -d ' {"id":"print-time"}'
//curl -X POST localhost:4567 -d '{"id":"print-time","cmd":"date","args":["-R"],"interval":5000}'