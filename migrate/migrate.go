package main

import (
	"github.com/Desmond51/go-CRUD/initializers"
	"github.com/Desmond51/go-CRUD/models"
)
 
func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func Migrate(){
	  // Migrate the schema
  initializers.DB.AutoMigrate(&models.Post{})
}