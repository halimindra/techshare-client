package main

import (
	"flag"
	"log"
	"net/http"
	"path"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"orami.com/techshare-client/services"
	pb "orami.com/techshare/pkg"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	mode               = flag.String("mode", "grpc", "server type to get data from")
)

func main() {
	flag.Parse()
	start := time.Now()
	limit := int64(2000000)

	if *mode == "grpc" {
		var opts []grpc.DialOption

		if *tls {
			if *caFile == "" {
				*caFile = path.Dir("ca.pem")
			}
			creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
			if err != nil {
				log.Fatalf("Failed to create TLS credentials %v", err)
			}
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			opts = append(opts, grpc.WithInsecure())
		}

		conn, err := grpc.Dial(*serverAddr, opts...)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		client := pb.NewTechShareClient(conn)
		services.PrintPersonGRPC(client, 1)
		services.PrintPeopleGRPC(client, limit)
	} else {
		client := http.Client{}
		services.PrintPersonREST(&client, *serverAddr, 1)
		services.PrintPeopleREST(&client, *serverAddr, limit)
	}

	elapsed := time.Since(start)
	log.Printf("Process took %s", elapsed)
}
