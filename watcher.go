package main

import (  
    "errors"
    "fmt"
    "gopkg.in/olivere/elastic.v5"
    "reflect"
    "time"
    "net/http"
    "encoding/json"
    "context"
    "log"
)

const (  
    indexName    = "meminfo"
    docType      = "log"
    name      = "ethosdistro"
    indexMapping = `{
                        "rigs" : {
                            "ID" : {
                                    "rack_loc" : { "type" : "string", "index" : "not_analyzed" },
                                    "miner_hashes" : { "type" : "string", "index" : "not_analyzed" },
                                    "current_time" : { "type" : "date" }
                                }
                        }
                    }`
)

type Log struct {  
    Rack_loc     string    `json:"rack_loc"`
    Miner_hashes  string    `json:"miner_hashes"`
    Time          time.Time `json:"current_time"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
    r, err := myClient.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func getData(client *elastic.Client) {
    var data []Log
    getJson("http://jaylin.ethosdistro.com/?json=yes", &data)
    ctx := context.Background()
    for _,item := range data{
        _,err := client.Index().Index(indexName).Type(Log).BodyJson(Log).Do(ctx)   
        if err !=nil{
            log.Fatal(err)
        }
    }  
    return (data.Time)
}







func main() {  
    client, err := elastic.NewClient(elastic.SetURL("http://35.202.37.2/elasticsearch/"))
    if err != nil {
        panic(err)
    }

    err = createIndexWithLogsIfDoesNotExist(client)
    if err != nil {
        panic(err)
    }

    err = getData(client)
    if err != nil {
        panic(err)
    }
}

func createIndexWithLogsIfDoesNotExist(client *elastic.Client) error {  
    exists, err := client.IndexExists(indexName).Do()
    if err != nil {
        return err
    }

    if exists {
        return nil
    }

    res, err := client.CreateIndex(indexName).
        Body(indexMapping).
        Do()

    if err != nil {
        return err
    }
    if !res.Acknowledged {
        return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
    }

    return addLogsToIndex(client)
}

func addLogsToIndex(client *elastic.Client) error {  
    for i := 0; i < 10; i++ {
        l := Log{
            Rack_loc:      "Name",
            Miner_hashes:   fmt.Sprintf("message %d", i),
            Time:           time.Now(),
        }

        _, err := client.Index().
            Index(indexName).
            Type(docType).
            BodyJson(l).
            Do()

        if err != nil {
            return err
        }
    }

    return nil
}

func findAndPrintAppLogs(client *elastic.Client) error {  
    termQuery := elastic.NewTermQuery("rack_loc", rack_loc)

    res, err := client.Search(indexName).
        Index(indexName).
        Query(termQuery).
        Sort("time", true).
        Do()

    if err != nil {
        return err
    }

    fmt.Println("Logs found:")
    var l Log
    for _, item := range res.Each(reflect.TypeOf(l)) {
        l := item.(Log)
        fmt.Printf("time: %s miner_hashes: %s\n", l.Time, l.Miner_hashes)
    }

    return nil
}
