package db

import (
	"planetas/external/github.com/boltdb/bolt"
	"fmt"
	"planetas/entity"
	"encoding/json"
	"strings"
	"strconv"
)

var BucketName = string ("galaxy")
const dbPath = "bolt.db"


func InitBolt () (*bolt.DB)  {
	db, err := bolt.Open(dbPath, 0600, nil)
	if (err != nil) {
		//este error cierra todo
		panic(err)
	}
	return db
}

//intento levantar una key de un bucket solo para ver si existe
func CheckExistsBucket() bool {
	dataBase := InitBolt()
	defer dataBase.Close()

	_, err:= Get(dataBase, []byte("anyKey"), BucketName)
	if (err != nil) {
		return  false
	}
	return  true
}

func Put (db *bolt.DB, key []byte, value []byte) error {
	return store(db, key, value, BucketName)
}

func store(db *bolt.DB, key []byte, value []byte, bucketName string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return bucket.Put(key, value)
	})
}

func Get(db *bolt.DB, key []byte, bucketName string) ([]byte, error) {
	var val []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("Ohhh el bucket %q no existe!", bucketName)
		}
		val = bucket.Get(key)
		return nil
	})
	return val, err
}

//No puedo devolver un bucket porque *tx se cierra y me da un bucket cerrado
func GetCantidadPeriodos (db *bolt.DB, climaBuscado string) (cantidadPeriodos int, diaMaximo int) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		var permimetroMaximo float64
		b.ForEach(func(k, v []byte) error {
			var clima entity.Clima
			err := json.Unmarshal(v, &clima)
			//por las dudas mando a Upper
			if ((strings.ToUpper((string(clima.TipoClima))) == strings.ToUpper(climaBuscado))) {
				cantidadPeriodos++
				if (permimetroMaximo < clima.PerimetroTriangulo) {
					permimetroMaximo = clima.PerimetroTriangulo
					diaMaximo,err = strconv.Atoi(string(k))
				}
			}
			return err
		})
		return nil
	})
	return cantidadPeriodos, diaMaximo
}


func GetClimaPorDia (base *bolt.DB, dia string) ([]byte, error) {
	return Get(base, [] byte(dia), BucketName)
}