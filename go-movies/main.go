package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
) 

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Tittle string `json:"title"`
	Director *Director `json:"director"`

}

type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname`
}