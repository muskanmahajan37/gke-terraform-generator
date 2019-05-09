/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"flag"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/klog"
	"os"
)

const header = `
+-------------------------------------------------------------------+
|    __.--/)  .-~~   ~~>>>>>>>>   .-.    gke-tf                     |
|   (._\~  \ (        ~~>>>>>>>>.~.-'                               |
|     -~}   \_~-,    )~~>>>>>>>' /                                  |
|       {     ~/    /~~~~~~. _.-~                                   |
|        ~.(   '--~~/      /~ ~.                                    |
|   .--~~~~_\  \--~(   -.-~~-.  \                                   |
|   '''-'~~ /  /    ~-.  \ .--~ /                                   |
|        (((_.'    (((__.' '''-'                                    |
+-------------------------------------------------------------------+
`

var (
	Version = "undefined"
	RootCMD = &cobra.Command{
		Use:   "gke-tf [command]",
		Short: "CLI Interface for creating terraform for GKE Cluster",
	}
)

func init() {
	// initialize klog
	klog.InitFlags(flag.CommandLine)
	// Sync the klog flags with ours
	flag.CommandLine.VisitAll(func(f1 *flag.Flag) {
		RootCMD.Flags().AddFlag(pflag.PFlagFromGoFlag(f1))
	})
}

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		exitWithError(err)
	}
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	klog.Errorf("Error: %v", err)
	os.Exit(1)
}
