package params

import (
	"encoding/json"
	"errors"
	"log"
	"time"
)

type Conf struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type tmpConf struct {
	Addr         string `json:"addr"`
	ReadTimeout  string `json:"readTimeout"`
	WriteTimeout string `json:"writeTimeout"`
	IdleTimeout  string `json:"idleTimeout"`
}

func mustUnmarshallInTmp(bytes []byte) tmpConf {
	var tmp = tmpConf{
		Addr:         ":8080",
		ReadTimeout:  "10s",
		WriteTimeout: "10s",
		IdleTimeout:  "1s",
	}
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		log.Fatal(err)
	}
	return tmp
}

func mustCreateConf(tmp tmpConf) Conf {
	var partitionErr error
	var err error

	var conf Conf
	conf.Addr = tmp.Addr

	conf.ReadTimeout, partitionErr = time.ParseDuration(tmp.ReadTimeout)
	err = errors.Join(err, partitionErr)

	conf.WriteTimeout, partitionErr = time.ParseDuration(tmp.WriteTimeout)
	err = errors.Join(err, partitionErr)

	conf.IdleTimeout, partitionErr = time.ParseDuration(tmp.IdleTimeout)
	err = errors.Join(err, partitionErr)

	if err != nil {
		log.Fatal(err)
	}

	return conf
}

func MustGetConfData(bytes []byte) Conf {
	return mustCreateConf(mustUnmarshallInTmp(bytes))
}
