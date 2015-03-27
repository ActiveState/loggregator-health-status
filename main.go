package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/ActiveState/loggregator_health_status/config"
	"github.com/ActiveState/loggregator_health_status/http_consumer"
)

const (
	APPGUID   = "APPGUID"
	INSTANCE  = "INSTANCE"
	ORGGUID   = "ORGGUID"
	SPACEGUID = "SPACEGUID"
)

var (
	defaultInstance = "0" // Always assuming the first instance
	appGuid         = flag.String("appGuid", "", "app guid")
	cfConfigPath    = fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".cf/config.json")
	csvPath         = "endpoints.csv"
)

func main() {
	flag.Parse()
	//appguid := "7102d156-bb28-4a12-b81a-5cee5b3e7116"

	cfConfig := &config.CfConfig{}

	err := config.ReadConfigInto(cfConfig, cfConfigPath)
	if err != nil {
		fmt.Print(err)

	}

	data, err := config.ReadCsv(csvPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Format in tab-separated columns with a tab stop of 8.
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "status\t\tquery")

	for _, each := range data {
		response, err := http_consumer.MakeHttpRequest(cfConfig.LoggregatorEndPoint, each[0], cfConfig.AccessToken, &tls.Config{InsecureSkipVerify: true}, nil)
		if err != nil {
			fmt.Println(err)

		} else{

			fmt.Fprintln(w, response.Status+"\t\t"+queryParser(cfConfig, each[0]))

		}
	}

	fmt.Fprintln(w)
	w.Flush()
}

func queryParser(cfConfig *config.CfConfig, query string) string {

	if strings.Contains(query, APPGUID) {

		query = strings.Replace(query, APPGUID, *appGuid, -1)
	}
	if strings.Contains(query, INSTANCE) {
		query = strings.Replace(query, INSTANCE, defaultInstance, -1)

	}
	if strings.Contains(query, ORGGUID) {
		query = strings.Replace(query, ORGGUID, cfConfig.OrganizationFields.Guid, -1)

	}
	if strings.Contains(query, SPACEGUID) {
		query = strings.Replace(query, SPACEGUID, cfConfig.SpaceFields.Guid, -1)

	}

	return query

}
