package main
import (
  "github.com/influxdata/influxdb/client/v2"
  "fmt"
  "flag"
  "time"
  "os"
  "log"
  "strconv"
   )
   
func influxDBClient() client.Client {
  c, err := client.NewHTTPClient(client.HTTPConfig{
    Addr: "https://<<host>>:8086",
    Username: "XXX",
    Password: "YYYYY",
    InsecureSkipVerify: true,
  })
  if err!= nil {
    fmt.Println("Error: ",err)
  }
  return c
}

func main() {
  log.SetOutput(os.Stdout)
  t := time.Now()

  serverPtr := flag.String("s","","Server Name")
  ipPtr := flag.String("i","","IP address")
  vpcPtr := flag.String("v","","VPC")
  clusterPtr := flag.String("c","","Cluster")
  poolPtr := flag.String("p","","Pool")
  cpuPtr := flag.String("cpu","","CPU")
  memPtr := flag.String("mem","","mem")
  osPtr := flag.String("os","","OS")
  domainPtr := flag.String("domain","","Domain")
  requestorPtr := flag.String("r","","Requestor")
  flag.Parse()
  
  c := influxDBClient()
   bp, err := client.NewBatchPoints(client.BatchPointsConfig{
    Database: "provisioning",
    Precision: "s",
  })
  if err != nil {
    log.Fatalln("Error: ", err)
  }
  
  tags := map[string]string{
    "cluster": *clusterPtr,
    "vpc": *vpcPtr,
    "pool": *poolPtr,
    "os": *osPtr,
    }
    
    proc, err:= strconv.Atoi(*cpuPtr)
    if err!= nil {
      fmt.Println("Error: ",err)
    }
    ram, err:= strconv.Atoi(*memPtr)
     if err!= nil {
      fmt.Println("Error: ",err)
    }
    
  fields := map[string]interface{}{
    "server": *serverPtr,
    "cpu": proc,
    "mem": ram,
    "ip": *ipPtr,
    "domain": *domainPtr,
    "requestor": *requestorPtr,
    }
  
  point, err := client.NewPoint(
        fmt.Sprintf("vm_provisioning1",),
        tags,
        fields,
        t.AddDate(0, 0, -1).UTC(),
        )
    if err != nil {
                log.Fatalln("Error: ", err)
            }

    bp.AddPoint(point)
    
  err = c.Write(bp)
  if err != nil {
    log.Fatal(err)
  }
}
