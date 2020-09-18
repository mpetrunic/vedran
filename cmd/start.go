package cmd

import (
	"errors"
	"fmt"
	"github.com/NodeFactoryIo/vedran/internal/loadbalancer"
	"github.com/NodeFactoryIo/vedran/pkg/util/random"
	"github.com/spf13/cobra"
)

var (
	authSecret string
	name       string
	capacity   int64
	whitelist  []string
	fee        float32
	selection  string
	port       int32
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts vedran load balancer",
	Run:   startCommand,
	Args: func(cmd *cobra.Command, args []string) error {
		// valid values are round-robin and random
		if selection != "round-robin" && selection != "random" {
			return errors.New("invalid selection option selected")
		}
		// all positive integers are valid, and -1 representing unlimited capacity
		if capacity < -1 {
			return errors.New("invalid capacity value")
		}
		// valid value is between 0-1
		if fee < 0 || fee > 1 {
			return errors.New("invalid fee value")
		}
		// well known ports and registered ports
		if port > 0 && port <= 49151 {
			return errors.New("invalid port number")
		}
		return nil
	},
}

func init() {
	startCmd.Flags().StringVar(
		&authSecret,
		"auth-secret",
		"",
		"[REQUIRED] authentication secret used for generating tokens")

	startCmd.Flags().StringVar(
		&name,
		"name",
		fmt.Sprintf("load-balancer-%s", random.String(12, random.Alphabetic)),
		"[OPTIONAL] public name for load balancer, autogenerated name used if omitted")

	startCmd.Flags().Int64Var(
		&capacity,
		"capacity",
		-1,
		"[OPTIONAL] maximum number of nodes allowed to connect, where -1 represents no upper limit")

	startCmd.Flags().StringArrayVar(
		&whitelist,
		"whitelist",
		nil,
		"[OPTIONAL] comma separated list of node id-s, if provided only these nodes will be allowed to connect")

	startCmd.Flags().Float32Var(
		&fee,
		"fee",
		0.1,
		"[OPTIONAL] value between 0-1 representing fee percentage")

	startCmd.Flags().StringVar(
		&selection,
		"selection",
		"round-robin",
		"[OPTIONAL] type of selection used for choosing nodes")

	startCmd.Flags().Int32Var(
		&port,
		"port",
		4000,
		"[OPTIONAL] port on which load balancer will be started")

	RootCmd.AddCommand(startCmd)
}

func startCommand(_ *cobra.Command, _ []string) {
	loadbalancer.StartLoadBalancerServer(loadbalancer.Properties{
		AuthSecret: authSecret,
		Name:       name,
		Capacity:   capacity,
		Whitelist:  whitelist,
		Fee:        fee,
		Selection:  selection,
		Port:       port,
	})
}
