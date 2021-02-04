package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Results struct {
	Totals struct {
		TotalPass int `json:"total_pass"`
		TotalFail int `json:"total_fail"`
		TotalWarn int `json:"total_warn"`
		TotalInfo int `json:"total_info"`
	} `json:"Totals"`
}

type repo struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			Wgpolicyk8SIoEngine string `yaml:"wgpolicyk8s.io/engine"`
		} `yaml:"labels"`
		Annotations struct {
			Name     string `yaml:"name"`
			Category string `yaml:"category"`
			Version  string `yaml:"version"`
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Summary struct {
		Pass  int `yaml:"pass"`
		Fail  int `yaml:"fail"`
		Warn  int `yaml:"warn"`
		Info  int `yaml:"info"`
		Error int `yaml:"error"`
		Skip  int `yaml:"skip"`
	} `yaml:"summary"`
}

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.SetConfigName("sample-resource")
		viper.AddConfigPath(currentdir())
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			log.Fatal(err)
		}

		C := repo{}

		err = viper.Unmarshal(&C)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

		fmt.Println(C)

		fmt.Println("JSON starts here: ")
		config := LoadResults("./logs.json")
		fmt.Println(config)

		// Change value in map and marshal back into yaml
		C.Summary.Pass = config.Totals.TotalPass
		C.Summary.Fail = config.Totals.TotalFail
		C.Summary.Warn = config.Totals.TotalWarn
		C.Summary.Info = config.Totals.TotalInfo

		fmt.Println(C)

		d, err := yaml.Marshal(&C)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Println(string(d))

		// write to file
		f, err := os.Create("/tmp/dat2")
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile("actual-resource.yaml", d, 0644)
		if err != nil {
			log.Fatal(err)
		}

		f.Close()

	},
}

func currentdir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return cwd
}

func LoadResults(filename string) Results {
	var results Results
	resultsFile, _ := os.Open(filename)
	defer resultsFile.Close()
	jsonParser := json.NewDecoder(resultsFile)
	jsonParser.Decode(&results)
	return results
}

func init() {
	rootCmd.AddCommand(produceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// produceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// produceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
