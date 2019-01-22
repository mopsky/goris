package cache

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type Config struct {
	host string
	port int
}

type Redis struct {
	conn redis.Conn
	dbid int
	config Config
	reply interface{}
	err error
}

// 校验连接
func CheckConn(r * Redis) bool {
	if r.conn == nil {
		r.err = errors.New("Redis未初始化")
		return false
	}

	return true
}

func (r *Redis) Open(host string, port int) error {
	if host == "" || port <= 0 {
		r.err = errors.New("非法的主机IP或端口")
	} else {
		r.conn, r.err = redis.Dial("tcp", host+":"+strconv.Itoa(port))
	}

	return r.err
}

func (r *Redis) Select(dbid int) (interface{}, error)  {
	if !CheckConn(r) {
		return nil, r.err
	}

	if dbid < 0 {
		r.err = errors.New("数据库ID不能小于0")
		return "FAIL", r.err
	}

	r.dbid = dbid
	r.reply, r.err = r.conn.Do("SELECT", r.dbid)
	return r.reply, r.err
}

func (r *Redis) Set(key, value interface{}, expire ...int) error {
	if !CheckConn(r) {
		return r.err
	}

	// 换成json存储
	var data string
	switch value.(type) {
	case string:
		data = value.(string)
	default:
		var jsonByte []byte
		jsonByte, r.err = json.Marshal(value)
		if r.err != nil {
			return r.err
		}
		data = string(jsonByte)
	}

	r.reply, r.err = r.conn.Do("SET", key, data)

	// 设置过期时间
	if r.err == nil && len(expire) > 0 && expire[0] > 0 {
		r.reply, r.err = r.conn.Do("EXPIRE", key, expire[0])
	}

	return r.err
}

func (r *Redis) Get(key string) (interface{}, error) {
	if !CheckConn(r) {
		return nil, r.err
	}

	r.reply, r.err = redis.String(r.conn.Do("GET", key))
	return r.reply, r.err
}

func (r *Redis) TTL(key string) (int64, error) {
	if !CheckConn(r) {
		return -2, r.err
	}

	r.reply, r.err = r.conn.Do("TTL", key)
	if r.err != nil {
		return -2, r.err
	}
	return r.reply.(int64), nil
}

func (r *Redis) Delete(keys string, all bool) error {
	if !CheckConn(r) {
		return r.err
	}

	if all {
		keys += "*"
	}

	r.reply, r.err = r.conn.Do("DEL", keys)
	return nil
}

func (r *Redis) Keys(keys string, all bool) ([]string, error) {
	if !CheckConn(r) {
		return nil, r.err
	}

	if all {
		keys += "*"
	}

	r.reply, r.err = r.conn.Do("KEYS", keys)
	if r.err != nil {
		return nil, r.err
	}

	var res []string
	for _, v := range r.reply.([]interface{}){
		res = append(res, string(v.([]byte)))
	}

	return res, r.err
}

func (r *Redis) Close() error {
	return r.conn.Close()
}

