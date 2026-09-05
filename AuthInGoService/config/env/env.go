package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// This fn is used to load the env values from the env files
func load() {
	err:=godotenv.Load() 
	if(err!=nil){
		fmt.Println("Not able to load env files")
	}
}




//This Getstring fn is used to get the string value of an environment variable. in typescript
func GetString(key string,fallback string) string{
	load()
	value,ok:=os.LookupEnv(key)
	if(!ok){
		return fallback
	}
	return value;
}

//This Getint fn is used to get the int value of an environment variable. in typescript
func GeInt(key string,fallback int) int{
	//load()
	value,ok:=os.LookupEnv(key)	
	if(!ok){
		return fallback
	}	
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return valueInt
}