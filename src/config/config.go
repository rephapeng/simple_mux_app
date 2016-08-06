package config
import(
    "encoding/json"
    "os"
    "log"
)

var globalConfiguration *configuration = new(configuration)

type configuration struct {
    Port string
}
func init(){
    loadConfiguration()
}

func Generate() (config *configuration) {
    //loadConfiguration()
    return globalConfiguration
}
//configuration in json format
func loadConfiguration() {
    configFilePath := "src/config/main.conf"

    log.Printf("starting %s load\n", configFilePath)
    configFile, err := os.Open(configFilePath)
    if err != nil {
        log.Println("[ERROR] ", err)
        log.Println("For your happiness an example config file is provided in the 'conf' directory in the repository.")
        os.Exit(1)
    }

    configDecoder := json.NewDecoder(configFile)
    err = configDecoder.Decode(&globalConfiguration)
    if err != nil {
        log.Println("[ERROR] ", err)
        log.Println("Please ensure that your config file is in valid JSON format.")
        os.Exit(1)
    }

    log.Printf("finished %s load\n", configFilePath)
}
