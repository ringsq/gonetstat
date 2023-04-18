package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bastjan/netstat"
)

/* Get TCP information and show like netstat.
   Information like 'user' and 'name' of some processes will not show if you
   don't have root permissions */

const filename = "/var/lib/node_exporter/netstat_listen.%v"

func main() {

	log.Println("Determining listening ports")
	tmpfile := fmt.Sprintf(filename, os.Getpid())
	fd, err := os.Create(tmpfile)
	if err != nil {
		log.Fatalf("Could not open %s: %v", tmpfile, err)
	}
	defer fd.Close()

	fd.WriteString(`# HELP netstat_listening_port Information about listening ports
# TYPE netstat_listening_port gauge
`)

	rows := 0
	rows += formatConnections(fd, netstat.TCP)
	rows += formatConnections(fd, netstat.TCP6)
	rows += formatConnections(fd, netstat.UDP)
	rows += formatConnections(fd, netstat.UDP6)

	fd.Close()
	if rows > 0 {
		os.Rename(tmpfile, fmt.Sprintf(filename, "prom"))
	} else {
		log.Println("No rows written")
		os.Remove(tmpfile)
	}
}

// formatConnections writes the listening ports and returns the number of rows written
func formatConnections(fd io.StringWriter, loc *netstat.Protocol) int {
	rows := 0
	connections, _ := loc.Connections()
	for _, conn := range connections {
		if !isListening(conn) {
			continue
		}
		if _, err := fd.WriteString(fmt.Sprintf("netstat_listening_port{address=\"%v\",port=\"%v\",protocol=\"%v\",process=\"%v\"} 1\n",
			conn.IP, formatPort(conn.Port), conn.Protocol.Name, formatPidProgname(conn.Pid, conn.Exe))); err != nil {
			log.Printf("Error writing metric: %v", err)
		} else {
			rows++
		}
	}
	return rows
}

func isListening(conn *netstat.Connection) bool {
	tcpListen := strings.HasPrefix(conn.Protocol.Name, "tcp") && conn.State == netstat.TCPListen
	udpListen := strings.HasPrefix(conn.Protocol.Name, "udp") && conn.State == netstat.TCPClose
	return tcpListen || udpListen
}

func formatPort(port int) string {
	if port == 0 {
		return "*"
	}
	return strconv.Itoa(port)
}

func formatPidProgname(pid int, exe string) string {
	if pid == 0 {
		return "-"
	}
	_, binary := filepath.Split(exe)
	// return fmt.Sprintf("%d/%s", pid, binary)
	return binary
}
