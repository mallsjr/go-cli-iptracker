/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP address",
	Long: `Trace the IP address.`,
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) == 0) {
			fmt.Println("Please provide an IP address")
			return
		}
		for _, ip := range args {
			fmt.Println(ip)
			showData(ip)
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// traceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// traceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type IPInfo struct {
	IP string `json:"ip"`
	City string `json:"city"`
	Region string `json:"region"`
	Country string `json:"country"`
	Loc string `json:"loc"`
	Postal string `json:"postal"`
	Timezone string `json:"timezone"`
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	return responseBytes
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseBytes := getData(url)

	data := IPInfo{}

	err := json.Unmarshal(responseBytes, &data)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(data)
	//TODO: Look for JSON Pretty Print
	//TOdo: Look into colorizing the output <faith/color>
}