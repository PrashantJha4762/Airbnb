package db

type Storage struct { //inside this struct we will have all the repositories that we will use in our application
	UserRepository UserRepository
}

//This fnction will create a new instance of the Storage struct and return it. It will also create a new instance of the UserRepositoryImpl and all other repositories and assign them to their respective fields of the Storage struct.
func NewStorage() *Storage {
	return &Storage{
		UserRepository:&UserRepositoryImpl{}, //we are creating a new instance of UserRepositoryImpl and assigning it to the UserRepository field of the Storage struct
	}
}