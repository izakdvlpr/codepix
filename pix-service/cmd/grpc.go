package cmd

import (
	"github.com/izakdvlpr/codepix/application/grpc"
	"github.com/izakdvlpr/codepix/infrastructure/database"

	"github.com/spf13/cobra"
)

var grpcServerPort int

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.ConnectDatabase()

		grpc.StartGrpcServer(db, grpcServerPort)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	grpcCmd.Flags().IntVarP(&grpcServerPort, "port", "p", 50051, "gRPC server port")
}
