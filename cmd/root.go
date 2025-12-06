package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nghSia/Mini-CRM/internal/app"
	"github.com/nghSia/Mini-CRM/internal/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	store   storage.Storer
	reader  *bufio.Reader
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "gomincrm",
		Short: "A CLI tool to manage contacts in Mini-CRM",
		Long: `gomincrm is a command-line interface application for managing contacts.
It provides commands to add, list, view, update, and delete contacts.
You can use it in interactive mode or with flags for quick operations.

The storage backend (memory, json, or gorm) can be configured via config.yaml.

Examples:
  gomincrm add              # Interactive mode to add a contact
  gomincrm add -n "John" -e "john@example.com"  # Add with flags
  gomincrm list             # List all contacts
  gomincrm get 1            # Get contact with ID 1
  gomincrm update 1 -n "Jane"  # Update contact name
  gomincrm delete 1         # Delete contact with ID 1`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			store = initStore()
		},
		Run: func(cmd *cobra.Command, args []string) {
			app.Run(store)
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("⚠️  Warning: Could not read config file, using default (gorm): %v", err)
		viper.SetDefault("storage.type", "gorm")
	} else {
		log.Printf("📄 Using config file: %s", viper.ConfigFileUsed())
	}
}

func initStore() storage.Storer {
	storageType := viper.GetString("storage.type")
	storageType = strings.ToLower(strings.TrimSpace(storageType))

	log.Printf("🔧 Initializing storage backend: %s", storageType)

	switch storageType {
	case "memory":
		log.Println("💾 Using MemoryStore (non-persistent)")
		return storage.NewMemoryStore()
	case "json":
		log.Println("📝 Using JSONStore (contacts.json)")
		return storage.NewJsonStore()
	case "gorm":
		log.Println("🗄️  Using GORMStore (contacts.db)")
		return storage.NewGORMStore()
	default:
		log.Printf("⚠️  Unknown storage type '%s', defaulting to gorm", storageType)
		return storage.NewGORMStore()
	}
}

func Execute() {
	reader = bufio.NewReader(os.Stdin)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func GetStore() storage.Storer { return store }
func GetReader() *bufio.Reader { return reader }

func intit() {}
