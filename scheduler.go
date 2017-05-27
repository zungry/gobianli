package main

import (
	"time"
	"fmt"
	"container/list"
	"os"
	"io/ioutil"
	"strings"
    "strconv"
    "github.com/Shopify/sarama"
	"log"
	"os/signal"
)

var MapCrawlerSpeed map[string]int
var AllTimeStamp [86400]*list.List
var ConfigFile string = "/home/kang/work/src/github.com/broadroad/gobianli/LinkedList/host.go"


type Scheduler struct {
    Consumer sarama.Consumer
}

func NewScheduler() *Scheduler {
    return &Scheduler{}
}
// init a list of 86400
func init()  {
    MapCrawlerSpeed = make(map[string]int)
    for i := 0; i < 86400; i++ {
        listi := list.New()
        listi.PushBack(i)
        AllTimeStamp[i] = listi
    }
    
    ReadConfig()
}

func ReadConfig(){
    // Read File
    // Update File
    f, err := os.Open(ConfigFile)
    if err != nil {
        //return nil, err
    }

    linestr, _ := ioutil.ReadAll(f)
    lines := strings.Split(string(linestr), "\n")
    for _,v := range lines {
        fmt.Println(v)
        kvs := strings.Split(v, " ")
        value, _ := strconv.Atoi(kvs[1])
        MapCrawlerSpeed[kvs[0]] = value
        fmt.Println(MapCrawlerSpeed[kvs[0]])
    }
}

func (s *Scheduler)ConsumerUrls() {

    urlConsumer, err := s.Consumer.ConsumePartition("url", 0, 0)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err := urlConsumer.Close(); err != nil {
            log.Fatalln(err)
        }
    }()

    //Trap SIGINT to trigger a shutdown
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)

    consumed := 0
    ConsumerLoop:
    for {
        select {
        case msg:= <-urlConsumer.Messages():
            log.Printf("Consumed message offset %d\n", msg.Offset)
            consumed++
        case <- signals:
            break ConsumerLoop
        }
    }
    log.Printf("Consumed: %d\n", consumed)
}

func (scheduler *Scheduler)Start()  {
    consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println("kafka client has been connected")

    scheduler.Consumer = consumer
    defer func ()  {
        if err = scheduler.Consumer.Close(); err != nil {
            log.Fatalln(err)
        }
    }()
    go scheduler.ConsumerUrls()
    timer := time.NewTicker(1 * time.Second)
    timerConfig := time.NewTicker(10 * time.Second)
    for {
        select {
        case <- timer.C:
            now0 := time.Now()
            time0 := time.Date(now0.Year(), now0.Month(), now0.Day(), 0,0,0,0,now0.Location())
            fmt.Println(now0.Unix() - time0.Unix())
            key := now0.Unix() - time0.Unix()
            fmt.Println(AllTimeStamp[key].Front().Value) 
        
        case <- timerConfig.C:
            fmt.Println("10s after")
            //go ReadConfig()
        }

    }
}   

func main() {
    s := NewScheduler()
    s.Start()
}