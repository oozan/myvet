package context

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/guregu/null.v4"

	"myvet-v2-api/mmqtt"
	"myvet-v2-api/mws"
	"myvet-v2-api/repo"
	"myvet-v2-api/structs"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

//Base application context with various connections.
type Base struct {
	WS        *mws.Hub
	Mmqtt     *mmqtt.Mmqtt
	Config    *Config
	Repo      *repo.Repo
	Router    *mux.Router
	Version   string
	BuildDate string
	Branch    string
}

func (b *Base) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//normally take source from here
		//ses, err := store.Get(r, "session-name")
		//check(err)
		/*check(err)
		log.Println("Is new?", ses.IsNew)
		source := ses.Values["user"]

		log.Println(source)*/
		source := null.NewString("current user", true)
		//source := sql.NullString{String: "current user", Valid: true}
		log.Println(source)
		//do the logging

		//example: /update-payment
		//source = current user
		//target = payment
		//logType = UPDATE

		//really ugly way to do this for starters
		//log.Println(r.URL.Path)
		pathParts := strings.Split(r.URL.Path, "/")
		//log.Println("pathParts: ")
		//log.Println(pathParts)
		endpoint := pathParts[2]
		//log.Println("endpoint: ")
		//log.Println(endpoint)

		logType := "OTHER"
		target := null.NewString("", false)
		//target := sql.NullString{String: "", Valid: false}

		//assume update-X/delete-X/create-X endpoint
		//no error checking yet
		if strings.ContainsAny(endpoint, "-") {
			log.Println("endpointissa '-'")
			parts := strings.Split(endpoint, "-")
			logType = strings.ToUpper(parts[0])
			log.Println("iffissä logType: " + logType)
			target.String = strings.ToUpper(parts[1])
			target.Valid = true
			log.Println(logType)
			log.Println(target)
		} else {
			target.String = endpoint
			if len(endpoint) > 0 {
				target.Valid = true
			}
		}

		body, err := ioutil.ReadAll(r.Body)
		log.Println(err)
		logData := null.NewString("", false)
		//logData := sql.NullString{String: "", Valid: false}
		logData.String = string(body)

		if len(logData.String) > 0 {
			logData.Valid = true
		}

		//log.Println("payload: " + string(payload))

		logEvent := structs.LogEvent{
			Target:  target,
			Source:  source,
			LogType: logType,
			LogData: logData}
		log.Println("ennen login tallentamista")
		num, err := b.Repo.CreateLogEvent(logEvent)
		log.Println(logEvent)

		log.Println(num)

		log.Println(err)

		//DEVELOP FURTHER BASED ON THIS: https://dev.to/ale_ukr/go-middleware-example-how-to-alter-a-handler-result-499c ?
		//https://upgear.io/blog/golang-tip-wrapping-http-response-writer-for-middleware/
		//https://gist.github.com/fiorix/372801082efb50cb7fc2

		next.ServeHTTP(w, r)
		return
	})
}

//DBInit initializes database connection for klinikkaohjelma_kehitys database.
func (b *Base) DBInit() error {
	ConnectionStr := fmt.Sprintf(b.Config.DBConnString, b.Config.DbUser, b.Config.DbPass, b.Config.DBName)

	var err error
	var db *sqlx.DB
	db, err = sqlx.Connect("mysql", ConnectionStr)
	if err != nil {
		return errors.New("Database could not be connected to:\n" + err.Error())
	}
	log.Println("Database connected.")
	// https://gitlab.com/adaptek/myvet-backend/issues/1
	// db.SetMaxIdleConns(0)
	// db.SetMaxOpenConns(500)
	b.Repo = repo.InitDB(db)

	return nil
}

// ReadConf fills Base.Config from configuration file.
func (b *Base) ReadConf() error {
	var conf Config
	configPaths := []string{"myvet-v2-api.conf", "/etc/myvet-v2-api/myvet-v2-api.conf", "/etc/myvet-v2-api.conf", "myvet-v2-api.toml", "/etc/myvet-v2-api/myvet-v2-api.toml", "/etc/myvet-v2-api.toml"}
	var err error
	strConfig := []byte("")
	for _, location := range configPaths {
		log.Println("Searching for the configuration location " + location)
		strConfig, err = ioutil.ReadFile(location)
		if err == nil {
			log.Println("Found at: " + location)
			break
		}
	}
	if len(strConfig) == 0 {
		log.Fatal("Configuration file is either missing or could not be found.")
	}
	if _, err = toml.Decode(string(strConfig), &conf); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", conf)
	b.Config = &conf
	return nil
}

// Run the application.
func (b *Base) Run() {
	log.Println("Server is started")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methodsOn := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"https://test1.avec.vet/v2", "http://localhost:3000", "http://localhost:3001"})
	log.Fatal(http.ListenAndServe(b.Config.APIUrl, handlers.CORS(headersOk, methodsOn, origins, handlers.AllowCredentials())(b.Router)))
}

//Config - collection of string parameters to create both database and API connection.
type Config struct {
	WebsocketRoute string
	APIUrl         string
	APIPrefix      string
	/*
		MQTTURL        string
		MQTTClientID   string
		MQTTUsername   string
		MQTTPassword   string
	*/
	DBName       string
	DbUser       string
	DbPass       string
	DBConnString string
	NoAuth       bool
}
