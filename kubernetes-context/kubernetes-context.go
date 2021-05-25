package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Context struct {
	Cluster   string
	Namespace string
	User      string
}

type ContextAndName struct {
	Context Context
	Name    string
}

type Config struct {
	Contexts       []ContextAndName
	CurrentContext string `yaml:"current-context"`
}

func main() {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig != "" {
		log.Fatal("$KUBECONFIG is not supported")
	}

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")))
	if err != nil {
		log.Fatalf("Cannot read config: %v", err)
	}

	config := Config{}
	if err := yaml.Unmarshal([]byte(data), &config); err != nil {
		log.Fatalf("Cannot parse config: %v", err)
	}

	if config.CurrentContext == "" {
		return
	}

	namespace := ""
	for _, context := range config.Contexts {
		if context.Name == config.CurrentContext {
			namespace = context.Context.Namespace
			break
		}
	}

	fmt.Printf("%v", config.CurrentContext)
	if namespace != "" {
		fmt.Printf(" (%v)", namespace)
	}
}
