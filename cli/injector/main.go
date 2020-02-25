// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2020 Datadog, Inc.

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "chaos-fi",
	Short: "Datadog chaos failures injection application",
	Run:   nil,
}

var log *zap.SugaredLogger

func init() {
	rootCmd.AddCommand(networkFailureCmd)
	rootCmd.AddCommand(nodeFailureCmd)
	rootCmd.AddCommand(networkLatencyCmd)
	rootCmd.PersistentFlags().String("uid", "", "UID of the failure resource")
	_ = cobra.MarkFlagRequired(rootCmd.PersistentFlags(), "uid")
}

func main() {
	// prepare logger
	zapInstance, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("error while creating logger: %v", err)
		os.Exit(2)
	}
	log = zapInstance.Sugar()

	// execute command
	if err = rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}