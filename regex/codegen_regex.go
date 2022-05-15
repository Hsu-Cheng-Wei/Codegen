package regex

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Regex(args *TemplateArgs) {
	if args.Configure {
		configure()
		return
	}

	if !checkConfigure() {
		fmt.Println("Please init config....")
		return
	}

	data, _ := os.ReadFile(getConfigurePath())

	var cfg TemplateConfigure
	_ = json.Unmarshal(data, &cfg)

	if len(args.Type) == 0 {
		fmt.Println("Please enter the type [query|command]....")
		return
	}

	(&TemplateRegx{
		Args: args,
		Cfg:  &cfg,
	}).Regex()
}

const configName string = "codegen.json"

func checkConfigure() bool {
	configPath := getConfigurePath()

	_, err := os.Stat(configPath)
	return err == nil
}

func getConfigurePath() string {
	path, _ := os.Getwd()
	return filepath.Join(path, configName)
}

func configure() {

	var cfg TemplateConfigure
	scanner := bufio.NewScanner(os.Stdin)

	path, _ := os.Getwd()

	dir := filepath.Base(path)

	cfg.Namespace = dir
	cfg.QueryPrefix = "Query"
	cfg.CommandPrefix = "Command"

	/*---------ApplicationPath-----------*/
	fmt.Print("ApplicationPath [default Applications]:")
	scanner.Scan()
	txt := scanner.Text()
	if len([]rune(txt)) == 0 {
		cfg.ApplicationPath = "Applications"
	} else {
		cfg.ApplicationPath = txt
	}
	cfg.Namespace = cfg.Namespace + "." + cfg.ApplicationPath
	configPath := getConfigurePath()

	if _, err := os.Stat(configPath); err == nil {
		_ = os.Remove(configPath)
	}

	fmt.Println("Generator json file " + configName + "....")
	f, _ := os.Create(configPath)

	data, _ := json.Marshal(cfg)
	_, _ = f.Write(data)

	_ = f.Sync()

	_ = f.Close()
}
