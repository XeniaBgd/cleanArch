package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/XeniaBgd/CleanArch/internal/application"
	"github.com/XeniaBgd/CleanArch/internal/transport/http/server"
)

type (
	Configs struct {
		Srv server.Conf
		App application.Conf
	}

	appConf struct {
		InterruptTimeout string `json:"interruptTimeout"`
	}

	serverConf struct {
		Addr         string `json:"addr"`
		ReadTimeout  string `json:"readTimeout"`
		WriteTimeout string `json:"writeTimeout"`
		IdleTimeout  string `json:"idleTimeout"`
	}

	tmpConfigs struct {
		appConf
		serverConf
	}
)

func MustGetConfData(path string) Configs {
	return createConf(unmarshallInTmp(readFile(path)))
}

func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func unmarshallInTmp(bytes []byte) tmpConfigs {
	var defaultConf = tmpConfigs{
		appConf: appConf{
			InterruptTimeout: "20s",
		},
		serverConf: serverConf{
			Addr:         ":8080",
			ReadTimeout:  "10s",
			WriteTimeout: "10s",
			IdleTimeout:  "1s",
		},
	}
	if err := json.Unmarshal(bytes, &defaultConf); err != nil {
		log.Fatal(err)
	}
	return defaultConf
}

func createConf(tmp tmpConfigs) Configs {
	var partitionErr error
	var err error

	var conf Configs
	conf.Srv.Addr = tmp.Addr

	conf.Srv.ReadTimeout, partitionErr = time.ParseDuration(tmp.ReadTimeout)
	err = errors.Join(err, partitionErr)

	conf.Srv.WriteTimeout, partitionErr = time.ParseDuration(tmp.WriteTimeout)
	err = errors.Join(err, partitionErr)

	conf.Srv.IdleTimeout, partitionErr = time.ParseDuration(tmp.IdleTimeout)
	err = errors.Join(err, partitionErr)

	conf.App.InterruptTimeout, partitionErr = time.ParseDuration(tmp.InterruptTimeout)
	err = errors.Join(err, partitionErr)

	if err != nil {
		log.Fatal(err)
	}

	return conf
}
