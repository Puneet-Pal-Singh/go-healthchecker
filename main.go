package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func formatURL(rawURL string) (string, error) {
	if _, err := net.LookupHost(rawURL); err != nil {
		return "", err
	}

	return rawURL, nil
}

func main() {
	var timeout int

	var rootCmd = &cobra.Command{
		Use:   "website-health-check [url]",
		Short: "Check the health of a website",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rawURL := args[0]

			formattedURL, err := formatURL(rawURL)
			if err != nil {
				fmt.Printf("Invalid URL: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Checking health for website: %s with timeout: %d seconds\n", formattedURL, timeout)
			isHealthy, err := checkWebsiteHealth(formattedURL, time.Duration(timeout)*time.Second)
			if err != nil {
				fmt.Printf("Error checking website health: %v\n", err)
				os.Exit(1)
			}

			if isHealthy {
				fmt.Println("Website is up!")
			} else {
				fmt.Println("Website is down!")
			}
		},
	}

	rootCmd.Flags().IntVarP(&timeout, "timeout", "t", 5, "Timeout in seconds")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
