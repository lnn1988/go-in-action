package main

import (
	"log"
	"os"
	"~/search"
	"matchers"
)

func init()  {
	log.SetOutPut(os.Stdout)

}

func main()  {
	search.Run("president")
	
}