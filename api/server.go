package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var pidFile string = "/tmp/nts.pid"

var apiBase string = "/api"
var apiVersion string = "v1"

func RunServers() {
	pid := os.Getpid()
	os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)

	// Create separate ServeMux and http.Server for each server
	jsMux := http.NewServeMux()
	jsMux.Handle("/", http.FileServer(http.Dir("./pkg/ntscli/web")))
	jsServer := &http.Server{Addr: ":8000", Handler: jsMux}

	apiMux := http.NewServeMux()
	// NTP Server routes
	/*
		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.Root,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.Status),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.Status,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.InstanceNumber),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.InstanceNumber,
				true))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.IpMode),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.IpMode,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.IpAddress),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.IpAddress,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.MacAddress),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.MacAddress,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.VlanStatus),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.VlanStatus,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.VlanAddress),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.VlanAddress,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.UnicastMode),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.UnicastMode,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.MulticastMode),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.MulticastMode,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.BroadcastMode),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.BroadcastMode,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.PrecisionValue),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.PrecisionValue,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.PollIntervalValue),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.PollIntervalValue,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.StratumValue),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.StratumValue,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.ReferenceId),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.ReferenceId,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.SmearingStatus),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.SmearingStatus,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.Leap61Progress),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.Leap61Progress,
				true))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.Leap59Progress),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.Leap59Progress,
				true))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.Leap61Status),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.Leap61Status,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.Leap59Status),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.Leap59Status,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.UtcOffsetStatus),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.UtcOffsetStatus,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.UtcOffsetValue),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.UtcOffsetValue,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.RequestsValue),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.RequestsValue,
				true))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.ResponsesValue),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.ResponsesValue,
				true))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.RequestsDroppedValue),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.RequestsDroppedValue,
				true))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.BroadcastsValue),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.BroadcastsValue,
				true))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.ClearCountersStatus),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.ClearCountersStatus,
				false))

		apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.NtpServer.Root, axi.NtpServer.Version),
			Handler(axi.ReadNtpServer,
				axi.WriteNtpServer,
				axi.NtpServer.Version,
				true))
	*/
	// PTP OC routes
	//apiMux.HandleFunc(path.Join(apiBase, apiVersion, axi.PtpOc.Root, axi.PtpOc.Version),
	//	Handler(axi.ReadPtpOc,
	//		axi.WritePtpOc,
	//		axi.PtpOc.Version,
	//		false))
	//apiMux.HandleFunc("/ptp-oc/version", PtpVersionHandler)
	//apiMux.HandleFunc("/ptp-oc/instance", PtpInstanceHandler)
	//apiMux.HandleFunc("/ptp-oc/mac", PtpMacHandler)
	//apiMux.HandleFunc("/ptp-oc/vlan", PtpVlanHandler)

	//apiMux.HandleFunc("/posts", handler)
	apiServer := &http.Server{Addr: ":8080", Handler: corsMiddleware(apiMux)}

	go func() {
		fmt.Println("Serving static files on :8000")
		if err := jsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("8000 error:", err)
		}
	}()

	go func() {
		fmt.Println("Serving API on :8080")
		if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("8080 error:", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jsServer.Shutdown(ctx)
	apiServer.Shutdown(ctx)
	fmt.Println("Servers stopped gracefully.")
	// Call Shutdown on your servers
	// e.g., apiServer.Shutdown(ctx), staticServer.Shutdown(ctx)
	fmt.Println("Servers stopped gracefully.")
}

func StopServers() {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		fmt.Println("Could not read PID file:", err)
		return
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Println("Invalid PID:", err)
		return
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Could not find process:", err)
		return
	}
	// Send SIGTERM (or SIGINT)
	proc.Signal(syscall.SIGTERM)
	fmt.Println("Sent stop signal to server process")
}
