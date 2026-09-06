package db

type UserRepository interface{
	Create() error
}

type UserRepositoryImpl struct { // Hme UserRepository interface ko implement krna tha uske lie ek
								// ek struct chahiye thi so hmne ye banaya

}
func (u * UserRepositoryImpl) Create() error{ //jaise hi hmne ye method banaya ye UserRepository interface ko implement krne lag gya
	return nil
}