package main
import (
    "fmt"
    "sync"
    )

//function to send value into the buffered channel
func buff(str string, buff_chan chan string, wg1 *sync.WaitGroup){
    defer wg1.Done()
    buff_chan<-str
    fmt.Println("Value sent into buffered channel")
}

//function to send values into the unbuffered channel
func unbuff(unbuff_chan chan string, wg2 *sync.WaitGroup){
    defer wg2.Done()
    for i:=0; i<4; i++{
        unbuff_chan<-fmt.Sprintf("Value %d",i)
        fmt.Println("Value ",i," sent successfully")
    }
}

func mainG() {
  // illustrating a buffered channel and unbuffered channel
  
  //buffered channel
  buff_chan:=make(chan string,4)
  //unbuffered channel
  unbuff_chan:=make(chan string)
  
  //creating a wg for buffered channel
  wg1:=&sync.WaitGroup{}
  wg2:=&sync.WaitGroup{}
  
  //goroutine to send values into the channel
  for i:=0; i<4; i++{
      wg1.Add(1)
      fmt.Println("Sending value " ,i," into the buffered channel")
      go buff(fmt.Sprintf("Hola Amigo %d",i), buff_chan, wg1)
  }
  wg1.Wait()
  close(buff_chan)
  
  //Reading values from the buffered channel
  for s:= range buff_chan{
      fmt.Println(s)
      fmt.Println("Value recieved successfully")
  }
  
  //goroutine to send values into the unbuffered channel
  wg2.Add(1)
  go unbuff(unbuff_chan,wg2)
  
  //Ranging over the unbuffered channel to read values
  for i:=0; i<4; i++{
      fmt.Println(<-unbuff_chan)
      fmt.Println("Value "+fmt.Sprintf("%d",i)+" recieved successfully")
  }
  
  //Waiting for
  wg2.Wait()
}