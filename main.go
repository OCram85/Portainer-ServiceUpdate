package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	con "ServiceUpdater/cconsole"

	"github.com/spf13/viper"
)

var (
	verboseMode  bool
	portainerURI string
	webhookToken string
	updateURI    string
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("PLUGIN")

	if viper.IsSet("VERBOSE") {
		verboseMode = viper.GetBool("VERBOSE")
	}

	if viper.IsSet("URI") {
		portainerURI = viper.GetString("uri")
	} else {
		log.Fatal(con.Warn("Missing URI variable!"))
	}

	if viper.IsSet("TOKEN") {
		webhookToken = viper.GetString("token")
	} else {
		log.Fatal(con.Warn("Missing TOKEN variable!"))
	}

	updateURI = fmt.Sprintf("%v/api/webhooks/%v", portainerURI, webhookToken)

	_, err := url.ParseRequestURI(updateURI)
	if err != nil {
		log.Fatal(con.Fata("Could not parse final post URL!\n", err))
	}
}

func main() {
	if verboseMode {
		log.Println(con.Magenta("==== Verbose Mode ===="))
		log.Println(con.Info("PortainerURI : ", portainerURI))
		log.Println(con.Info("WebhookToken : ", webhookToken))
		log.Println(con.Info("updateURI    : ", updateURI))
	}

	resp, err := http.Post(updateURI, "application/json", nil)
	if err != nil {
		log.Fatalln(con.Fata("Portainer service update error: ", err))
	}

	if resp.StatusCode != 204 {
		log.Fatal(con.Fata("Update Failed!\n", resp))
	} else {
		log.Println(con.Warn("Service updated!"))
	}
}
