package main

import "fmt"

func Set(mysql *Database, redis *URedis, short, long string) error {
	if redis.Enable {
		if err := redis.Set(short, long); err != nil {
			return err
		}
	}

	if mysql.Enable {
		if err := mysql.Set(short, long); err != nil {
			return err
		}
	}

	return nil
}

func GetShort(mysql *Database, redis *URedis, long string) (string, error) {
	return Get(mysql, redis, "short", long)
}
func GetLong(mysql *Database, redis *URedis, short string) (string, error) {
	return Get(mysql, redis, "long", short)
}

func Get(mysql *Database, redis *URedis, key, value string) (string, error) {
	var answer string

	if redis.Enable && key == "long" {
		// 获取原始链接
		answer, err := redis.Get(value)
		if err == nil {
			return answer, err
		}
	}

	if mysql.Enable {

		if key == "long" {
			// 获取原始链接
			shorts, err := mysql.Get("short", value)
			if err == nil {
				_ = redis.Set(shorts.Short, shorts.Long)
			}
			return shorts.Long, err
		} else {
			shorts, err := mysql.Get("long", value)
			return shorts.Short, err
		}
	}

	return answer, nil
}

func Del(mysql *Database, redis *URedis, short string) {
	if redis.Enable {
		if err := redis.Del(short); err != nil {
			fmt.Println(err.Error())
		}
	}

	if mysql.Enable {
		if err := mysql.Del(short); err != nil {
			fmt.Println(err.Error())
		}
	}
}
