# Web Client Heimdall

Web Client wrapper based on github.com/gojek/heimdall

## Usage (without Circuit Breaker)
- #### Initialize Web Client Factory
```go
factory := leafHeimdall.NewClientFactory()
```

- #### Initialize Web Client Instance
```go
webClient := factory.Create(time.Second)
webClientRetry := factory.CreateWithRetry(time.Second, 2)
```

- #### GET Request Example
  - ##### Response Struct
```go
type Data struct {
	Array   []int       `json:"array"`
	Boolean bool        `json:"boolean"`
	Color   string      `json:"color"`
	Null    interface{} `json:"null"`
	Number  int         `json:"number"`
	Object  struct {
		A string `json:"a"`
		C string `json:"c"`
	} `json:"object"`
	String string `json:"string"`
}
```
  - ##### API Call
```go
    headers := http.Header{"x-data": []string{"my-data"}}
    headers.Set("Content-Type", "application/json")
    url := "https://run.mocky.io/v3/a815080c-f86b-4bad-8463-c538897a3405"
    
    response, err := webClient.Get(ctx, url, headers)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    
    jsonBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(jsonBody))
    
    var data Data
    if err := json.Unmarshal(jsonBody, &data); err != nil {
        panic(err)
    }
    fmt.Println(data)
  ```

## Usage (with Circuit Breaker)
- #### Initialize Web Client Factory
```go
factory := leafHystrix.NewClientFactory()
```

- #### Initialize Web Client Instance
```go
webClient := factory.Create(
	leafWebClient.WithCommandName("command-name"))
```

- #### GET Request Example
    - ##### Response Struct
  ```go
    type Data struct {
        Array   []int       `json:"array"`
        Boolean bool        `json:"boolean"`
        Color   string      `json:"color"`
        Null    interface{} `json:"null"`
        Number  int         `json:"number"`
        Object  struct {
            A string `json:"a"`
            C string `json:"c"`
        } `json:"object"`
        String string `json:"string"`
    }
  ```
    - ##### API Call
  ```go
    headers := http.Header{"x-data": []string{"my-data"}}
    headers.Set("Content-Type", "application/json")
    url := "https://run.mocky.io/v3/a815080c-f86b-4bad-8463-c538897a3405"
    
    response, err := webClient.Get(ctx, url, headers)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    
    jsonBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(jsonBody))
    
    var data Data
    if err := json.Unmarshal(jsonBody, &data); err != nil {
        panic(err)
    }
    fmt.Println(data)
  ```