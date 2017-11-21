package DataSource

import (
	"sync"
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
	"strings"
	"crypto/tls"
	"net"
)

type singleton struct {
	dbContext *mgo.Database
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		var uri = "mongodb://denizgokce:testpassword@,cluster0-shard-00-02-ias3l.mongodb.net:27017/admin?authSource=admin&replicaSet=Cluster0-shard-0"
		if uri == "" {
			fmt.Println("No connection string provided - set MONGODB_URL")
			os.Exit(1)
		}
		uri = strings.TrimSuffix(uri, "?ssl=true")

		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true

		dialInfo, err := mgo.ParseURL(uri)

		if err != nil {
			fmt.Println("Failed to parse URI: ", err)
			os.Exit(1)
		}

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}

		session, err := mgo.DialWithInfo(dialInfo)
		//session, err := mgo.Dial("")
		if err != nil {
			panic(err)
		}
		//defer session.Close()
		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)
		instance = &singleton{session.DB("default")}
	})
	return instance
}
