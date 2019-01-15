 Crawler-baba  
==============  

crawl videos from Vimeo API... learning-by-doing Golang!    

## LEARN - MAKE - SHOW - SHARE   

### Prerequisites

In order to build and run this app you need to have a couple of things installed:  

- The [Go SDK](https://golang.org)   
- An IDE for the development, like [Atom](https://atom.io) or IntelliJ/Goland       

### Building the App  

#### Clone this repo     

```bash
  # clone this repo  
$ git clone https://github.com/capgemininorway/crawlerbaba  
$ cd crawlerbaba  

```   

#### Prepare env-vars  

**WARNING** This APP requires _a valid token_ for connecting to **Vimeo API**!!!          
You can get it from [Vimeo Developer portal](https://developer.vimeo.com/api/start)    

By default, app checks for env-var to find _vimeoToken_, if not found then it asks user for input.            
```bash
# set vimeoToken as env-variable    
$ export VIMEO_TOKEN=<your-vimeo-token>     

# OR just enter it when app asks your for it     

```

#### Build & Run the App        

```bash

  # run the App from the code    
$ go run crawler.go

# build & run it as executable 
$ go build crawler.go
$ ./crawler  

# by default, 'go build' generates executable per env it runs.  
# see https://golang.org/pkg/go/build/    

# if you want to specify per OS compatibility, see each below      
## - linux   
$ GOOS=linux GOARCH=amd64 go build crawler.go     

## - mac   
$ GOOS=darwin GOARCH=amd64 go build crawler.go  

## - windows   
$ GOOS=windows GOARCH=amd64 go build crawler.go    

```
