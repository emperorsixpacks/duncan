package redis 

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/emperorsixpacks/duncan"
	"github.com/redis/go-redis/v9"
)

func mapToStruct(i interface{}, o *interface{}) error {
	newStrVal, err := json.Marshal(i)
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(newStrVal), o); err != nil {
		return err
	}
	return nil
}

var ctx = context.Background() // I do not know, should I put this in the struct

// TODO look into making this a singleton
type RedisClient struct {
	rdb *redis.Client
}

func (this *RedisClient) clearDB() error {
	if err := this.rdb.FlushAll(ctx).Err(); err != nil {
		return err
	}
	return nil
}

// TODO try to make this simpler
func (this RedisClient) GetJSON(k string, o interface{}) error {
	// NOTE this works
	val, err := this.getJSON(k)
	if err != nil {
		return err
	}
	strMapping, ok := val.([]interface{})
	if !ok {
		return errors.New("internal error")
	}
	err = mapToStruct(strMapping[0], &o)
	if err != nil {
		return err
	}
	return nil
}

// this is a low level method, from here, we can perform things like deleting a single key or updating a single key
func (this RedisClient) getJSON(k string, inner_key ...string) (interface{}, error) {
	if len(inner_key) == 0 {
		inner_key = []string{"$"}
	}
	val, err := this.rdb.JSONGet(ctx, k, inner_key[0]).Expanded()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// we can even expand this further to get the data in a nestad json
// let us go ahead now and create some hidden methods to handle this
func (this RedisClient) SetJSON(key string, value interface{}) error {
	if err := this.setJSON(key, value); err != nil {
		return err
	}

	return nil
}

func (this RedisClient) setJSON(key string, value interface{}, inner_key ...string) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// TODO put this into a function
	if len(inner_key) == 0 {
		inner_key = []string{"$"}
	}
	err = this.rdb.JSONSet(ctx, key, inner_key[0], val).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Keeing it like this for now, later if needed we may need to go into nested objets to delete specific keys, but that should be from the client
func (this RedisClient) DeleteJSON(key string, value interface{}) error {
	err := this.rdb.JSONDel(ctx, key, "$").Err()
	if err != nil {
		// log error here
		return err
	}
	// log here
	return nil
}
func (this RedisClient) UpdateJSON(key string, value interface{}) {}

// this should be private, and later, we should have only getconnection, var, should com from duncan config
// We can use an interface here something like duncan.cache, but that should be later
// Fix null pointer error
func New(conn duncan.RedisConnetion) (*RedisClient, error) {
	newClient := new(RedisClient)
	options := &redis.Options{
		Addr:     conn.Addr,
		Password: conn.Password,
		DB:       conn.DB,
	}
	client := redis.NewClient(options)
	err := client.Ping(ctx).Err()
	if err != nil {
		message := fmt.Sprintf("could not connect on %s \n%v", conn.Addr, err)
		fmt.Println(message)
		return nil, err
	}
	newClient.rdb = client
	return newClient, nil
}
