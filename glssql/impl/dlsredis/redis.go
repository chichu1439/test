package dlsredis

import (
	"fmt"
	"gopkg.in/redis.v5"
	"strings"
	"sync"
	"time"
)

var (
	DlsConfig = DlsQueryConfig{}
)

type DlsQueryConfig struct {
	DlsFlag         int
	PrimaryType     string
	PrimaryTypeLen  int
	PrimaryClassLen int
	SessionName     string
	HostOrg         string
	HostDcn         string
	HostTopicId     string
	QueryExist      string
	RedisType       string
	RedisAddr       string
	RedisPoolNum    int
	RedisPassword   string
	RedisMultiFlag  bool
	RedisReadonly   bool
	TopicDcnTitle   string
}

type alias struct {
	Type          string
	Multi         bool
	Client        *redis.Client
	ClusterClient *redis.ClusterClient
}

type _redisCache struct {
	mux   sync.RWMutex
	cache map[string]*alias
}

type DlsRedisClient interface {
	GetRedis(aliasNames ...string) (*DlsRedisClient, error)
	SetValue(key string) error
	SetValueByTimeout(key string, timeout time.Duration) error
	GetValue(key string) (string, error)
}

type DlsRedisElement struct {
	Key string
	Field string
	Value string
}

var (
	redisClientCache = &_redisCache{cache: make(map[string]*alias)}
	registFlag       = false
)

func (ac *_redisCache) get(name string) (al *alias, ok bool) {
	ac.mux.RLock()
	defer ac.mux.RUnlock()
	al, ok = ac.cache[name]
	return
}

func (ac *_redisCache) add(name string, al *alias) (added bool) {
	ac.mux.Lock()
	defer ac.mux.Unlock()
	if _, ok := ac.cache[name]; !ok {
		ac.cache[name] = al
		added = true
	}
	return
}

func getRedisAlias(name string) *alias {
	if al, ok := redisClientCache.get(name); ok {
		return al
	}
	panic(fmt.Errorf("unknown Redis alias name %s", name))
}

func IsRedisRegister() bool {
	return registFlag
}

func RegisterRedis(aliasNames ...string) (err error) {
	var name string
	var multi bool = false
	// 不考虑支持事务
	multi = false
	if len(aliasNames) > 0 {
		name = aliasNames[0]
	} else {
		name = "default"
	}
	var tp = strings.Trim(strings.ToLower(DlsConfig.RedisType), " ")
	var al = alias{}
	al.Type = tp
	if tp == "cluster" {
		cli := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    strings.Split(DlsConfig.RedisAddr, ","),
			Password: DlsConfig.RedisPassword,
			ReadOnly: true,
			PoolSize: DlsConfig.RedisPoolNum,
		})
		if multi {
			_, err := cli.ExistsMulti("multi").Result()
			if err != nil {
				multi = false
			}
			_, err = cli.ExistsMulti("discard").Result()
			if err != nil {
				multi = false
			}
		}
		al.Multi = multi
		al.ClusterClient = cli
		if _, err = al.ClusterClient.Ping().Result(); err != nil {
			return fmt.Errorf("Register Cluster redis[%v] failed, err=[%v]", DlsConfig, err)
		}
	} else {
		cli := redis.NewClient(&redis.Options{
			Addr:     DlsConfig.RedisAddr,
			Password: DlsConfig.RedisPassword,
			PoolSize: DlsConfig.RedisPoolNum,
		})
		if multi {
			_, err := cli.ExistsMulti("multi").Result()
			if err != nil {
				multi = false
			}
			_, err = cli.ExistsMulti("discard").Result()
			if err != nil {
				multi = false
			}
		}
		al.Multi = multi
		al.Client = cli
		if _, err = al.Client.Ping().Result(); err != nil {
			return fmt.Errorf("Register Single redis[%v] failed, err=[%v]", DlsConfig, err)
		}
	}

	if !redisClientCache.add(name, &al) {
		return fmt.Errorf("RedisCache alias name `%s` already registered, cannot reuse", name)
	}

	registFlag = true

	return nil
}

func InitRedisChache() error {
	if IsRedisRegister() {
		return nil
	}
	return RegisterRedis()
}

func HSetValue(key string, field string, value string) (str string, err error) {
	al, err := getRedis()
	if nil != err {
		return str, err
	}
	fv := map[string]string{field: value}
	if al.Type == "cluster" {
		v, err := al.ClusterClient.HMSet(key, fv).Result()
		if nil != err {
			return "", err
		}
		return v, nil
	} else {
		v, err := al.Client.HMSet(key, fv).Result()
		if nil != err {
			return "", err
		}
		return v, nil
	}
}

func getRedis(aliasNames ...string) (*alias, error) {
	if !registFlag {
		return nil, fmt.Errorf("RedisCache of alias should be Regist First!")
	}
	var name string
	if len(aliasNames) > 0 {
		name = aliasNames[0]
	} else {
		name = "default"
	}
	al, ok := redisClientCache.get(name)
	if ok {
		return al, nil
	}
	return nil, fmt.Errorf("RedisCache of alias name `%s` not found", name)
}

func ReKey(old string, new string) (err error) {
	al, err := getRedis()
	if nil != err {
		return fmt.Errorf("RedisCache not found")
	}
	if al.Type == "cluster" {
		al.ClusterClient.Rename(old, new)
	} else {
		al.Client.Rename(old, new)
	}

	return nil
}

func DelKey(key string) (err error) {
	al, err := getRedis()
	if nil != err {
		return fmt.Errorf("RedisCache not found")
	}
	if al.Type == "cluster" {
		_, err = al.ClusterClient.Del(key).Result()
		err = err
		if err != nil {
			return err
		}
		/*_, err = al.ClusterClient.Get(key).Result()
		err = err
		if err == nil {
			log.Error(err)
			return fmt.Errorf("Delete Redis Key[%s] Faild", key)
		}*/
	} else {
		al.Client.Del(key)
		err = err
		if err != nil {
			return err
		}
		/*_, err = al.Client.Get(key).Result()
		err = err
		if err == nil {
			log.Error(err)
			return fmt.Errorf("Delete Redis Key[%s] Faild", key)
		}*/
	}

	return nil
}

func HDelKey(key string, field string) (err error) {
	al, err := getRedis()
	if nil != err {
		return fmt.Errorf("RedisCache not found")
	}
	if al.Type == "cluster" {
		_, err = al.ClusterClient.HDel(key, field).Result()
		err = err
		if err != nil {
			return err
		}
		/*_, err = al.ClusterClient.Get(key).Result()
		err = err
		if err == nil {
			log.Error(err)
			return fmt.Errorf("Delete Redis Key[%s] Faild", key)
		}*/
	} else {
		al.Client.HDel(key, field)
		err = err
		if err != nil {
			return err
		}
		/*_, err = al.Client.Get(key).Result()
		err = err
		if err == nil {
			log.Error(err)
			return fmt.Errorf("Delete Redis Key[%s] Faild", key)
		}*/
	}

	return nil
}

func SetValue(key string, value string) (str string, err error) {
	al, err := getRedis()
	if nil != err {
		return str, err
	}
	if al.Type == "cluster" {
		v, err := al.ClusterClient.Set(key, value, 0).Result()
		if nil != err {
			return "", fmt.Errorf("Redis SetValueTimeout[%s][%s] Failed! ", key, value)
		}
		return v, nil
	} else {
		v, err := al.Client.Set(key, value, 0).Result()
		if nil != err {
			return "", fmt.Errorf("Redis SetValue[%s][%s] Failed! ", key, value)
		}
		return v, nil
	}
}

func DeferPipe(pipe *redis.Pipeline) {
	_, err := pipe.Exec()
	if err != nil {
		fmt.Println(err)
	}
	err = pipe.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func GetPipe() *redis.Pipeline {
	al, err := getRedis()
	if nil != err {
		return nil
	}
	var pipe *redis.Pipeline
	if al.Type == "cluster" {
		pipe = al.ClusterClient.Pipeline()
	} else {
		pipe = al.Client.Pipeline()
	}
	return pipe
}

func SetPipeValue(pipe *redis.Pipeline, k, f, v string) error {
	if f == "" {
		_, err := pipe.Set(k, v, 0).Result()
		if nil != err {
			return fmt.Errorf("Redis SetValue[%s][%s][%s] Failed! ", k, f, v)
		}
	} else {
		kv := map[string]string{f: v}
		_, err := pipe.HMSet(k, kv).Result()
		if nil != err {
			return fmt.Errorf("Redis SetValue[%v] Failed! ", v)
		}
	}
	return nil
}

func SetPipeValues(kv []DlsRedisElement) error {
	al, err := getRedis()
	if nil != err {
		return err
	}
	var pipe *redis.Pipeline
	if al.Type == "cluster" {
		pipe = al.ClusterClient.Pipeline()
	} else {
		pipe = al.Client.Pipeline()
	}
	defer DeferPipe(pipe)
	for _, v := range kv {
		if v.Field == "" {
			_, err := pipe.Set(v.Key, v.Value, 0).Result()
			if nil != err {
				return fmt.Errorf("Redis SetValue[%v] Failed! ", v)
			}
		} else {
			kv := map[string]string{v.Field: v.Value}
			_, err := pipe.HMSet(v.Key, kv).Result()
			if nil != err {
				return fmt.Errorf("Redis SetValue[%v] Failed! ", v)
			}
		}
	}
	return nil
}

func SetValueTimeout(key string, value string, duration time.Duration) (str string, err error) {
	al, err := getRedis()
	if nil != err {
		return str, err
	}
	if al.Type == "cluster" {
		v, err := al.ClusterClient.Set(key, value, duration).Result()
		if nil != err {
			return str, fmt.Errorf("Redis SetValueTimeout[%s][%s][%v] Failed! ", key, value, duration)
		}
		return v, nil
	} else {
		v, err := al.Client.Set(key, value, duration).Result()
		if nil != err {
			return "", fmt.Errorf("Redis SetValueTimeout[%s][%s][%v] Failed! ", key, value, duration)
		}
		return v, nil
	}
}

func GetValue(key string) (str string, err error) {
	al, err := getRedis()
	if nil != err {
		return "", err
	}
	if al.Type == "cluster" {
		v, err := al.ClusterClient.Get(key).Result()
		if nil != err {
			return "", fmt.Errorf("Redis GetValue[%s] Failed! ", key)
		}
		return v, nil
	} else {
		v, err := al.Client.Get(key).Result()
		if nil != err {
			return "", fmt.Errorf("Redis GetValue[%s] Failed! ", key)
		}
		return v, nil
	}
}

func HGetValue(key string, field string) (str string, err error) {
	// 时间结算
	al, err := getRedis()
	if nil != err {
		return "", err
	}
	if al.Type == "cluster" {
		v, err := al.ClusterClient.HGet(key, field).Result()
		if nil != err {
			return "", err
		}
		return v, nil
	} else {
		v, err := al.Client.HGet(key, field).Result()
		if nil != err {
			return "", err
		}
		return v, nil
	}
}

func HGetAll(key string) (str map[string]string, err error) {
	al, err := getRedis()
	if nil != err {
		return nil, err
	}
	if al.Type == "cluster" {
		v, err := al.ClusterClient.HGetAll(key).Result()
		if nil != err {
			return nil, err
		}
		return v, nil
	} else {
		v, err := al.Client.HGetAll(key).Result()
		if nil != err {
			return nil, err
		}
		return v, nil
	}
}

func MultiProc(cmds ...string) (err error) {
	al, err := getRedis()
	if nil != err {
		return err
	}
	if al.Type == "cluster" {
		_, err = al.ClusterClient.ExistsMulti(cmds...).Result()
		if nil != err {
			return err
		}
		return nil
	} else {
		_, err = al.Client.ExistsMulti(cmds...).Result()
		if nil != err {
			return err
		}
		return nil
	}
	return nil
}

func Multi() error {
	al, err := getRedis()
	if nil != err {
		return err
	}
	if !al.Multi {
		return nil
	}
	if al.Type == "cluster" {
		err = al.ClusterClient.Process(redis.NewCmd("multi"))
		if nil != err {
			return err
		}
		return nil
	} else {
		err = al.Client.Process(redis.NewCmd("multi"))
		if nil != err {
			return err
		}
		return nil
	}
	return nil
}

func Exec() error {
	al, err := getRedis()
	if nil != err {
		return err
	}
	if !al.Multi {
		return nil
	}
	if al.Type == "cluster" {
		err = al.ClusterClient.Process(redis.NewCmd("exec"))
		if nil != err {
			return err
		}
		return nil
	} else {
		err = al.Client.Process(redis.NewCmd("exec"))
		if nil != err {
			return err
		}
		return nil
	}
	return nil
}

func Discard() error {
	al, err := getRedis()
	if nil != err {
		return err
	}
	if !al.Multi {
		return nil
	}
	if al.Type == "cluster" {
		err = al.ClusterClient.Process(redis.NewCmd("discard"))
		if nil != err {
			return err
		}
		return nil
	} else {
		err = al.Client.Process(redis.NewCmd("discard"))
		if nil != err {
			return err
		}
		return nil
	}
	return nil
}
