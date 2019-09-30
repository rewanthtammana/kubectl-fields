/*
Copyright © 2019 Rewanth Cool

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "fmt"
  "os"
  "os/exec"
  "github.com/spf13/cobra"

  "github.com/rewanth1997/kubectl-fields/pkg/fields"
  "github.com/rewanth1997/kubectl-fields/pkg/stdin"
)


var (
  cfgFile string
  ignoreCaseFlag bool
  stdinFlag bool
  rootCmdDescriptionShort = "Kubectl resources hierarchy parsing plugin"
  rootCmdDescriptionLong = `Kubectl resources hierarchy parser.
  
More info: https://github.com/rewanth1997/kubectl-fields`

  rootCmdExamples = `$ kubectl fields po.spec capa
containers.securityContext.capabilities
initContainers.securityContext.capabilities

$ kubectl fields -i svc ip
spec.clusterIP
spec.externalIPs
spec.loadBalancerIP
spec.sessionAffinityConfig.clientIP
status.loadBalancer.ingress.ip

Additional kubectl-fields example (the hard way: not recommended). Developed to run tests on pipeline

$ kubectl explain --recursive po.spec | ./kubectl-fields --stdin ver
dnsConfig.nameservers
volumes.csi.driver
volumes.flexVolume.driver
volumes.iscsi.chapAuthDiscovery
volumes.nfs.server`
)


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "kubectl-fields",
  Short: rootCmdDescriptionShort,
  Long: rootCmdDescriptionLong,
  Example: rootCmdExamples,
  Run: func(cmd *cobra.Command, args []string) {

    if stdinFlag {
      input := stdin.GetStdInput()
      fields.Parse(input, os.Args[1:], ignoreCaseFlag)
      return
    }

    output, err := exec.Command("kubectl", "explain" , "--recursive", args[0]).Output()
    if err != nil {
      fmt.Println(err)
      return
    }

    fields.Parse(string(output), args[1:], ignoreCaseFlag)
  },
}

// Initiates ignore-case and stdin flags
func init() {
  rootCmd.Flags().BoolVarP(&ignoreCaseFlag, "ignore-case", "i", false, "Ignore case distinction")
  rootCmd.Flags().BoolVarP(&stdinFlag, "stdin", "", false, "Expects input via pipes")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
