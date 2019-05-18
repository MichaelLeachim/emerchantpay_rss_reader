# This is a first part of the RSS reader task for the Emerchantpay

The goal is to:

```
Create an RSS Reader package, which can parse asynchronous multiple RSS
feeds.
```

## Package interface

The package exports the following struct

```golang

type RssItem struct{
  Title string
  Source string
  SourceURL string
  Link string
  PubishDate time.Time
  Description string
}

```

And the function:

```golang

Parse(urls []string) []RssItem

```

The `Parse`  function is asynchronous and is implemented
with `sync` package and goroutines. It also requires 
`github.com/sirupsen/logrus` external package for logginng. 


## How to install run and test

```bash

# Install
git clone https://github.com/MichaelLeachim/emerchantpay_rss_reader;
cd emerchantpay_rss_reader; 

# Run tests
go test; 

# Run benchmark
go test -bench; 

```

## Additional notes

* Parser is implemented using DI principle in order 
  to separate concerns and test the main method easily.
  


































