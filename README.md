# tcp-go

This repository contains a simple TCP client-server implementation in Go. The client connects to a TCP server, sends messages, and receives responses. This example demonstrates basic TCP communication, including handling multi-word messages and concurrent connections.

## TCP/IP
TCP stands for Transmission Control Protocol. It is a transport layer protocol that facilitates the transmission of packets from source to destination. It is a connection-oriented protocol that means it establishes the connection prior to the communication that occurs between the computing devices in a network. This protocol is used with an IP protocol, so together, they are referred to as a TCP/IP.

**Features of TCP**
- Connection Oriented
- Reliable
- Provides Guaranteed Delivery of Data Packets
- Maintains the order of the Data Packets

## Features

- **TCP Server**: Listens for incoming client connections and echoes back received messages.
- **TCP Client**: Connects to the server, sends messages, and displays server responses.
- **Multi-Word Message Support**: Handles messages with spaces (e.g., "Hello World").

## Getting Started

### Prerequisites

- Go 1.21+ installed on your machine.
- A terminal to run the commands.

### Running the Server

To run the TCP server, execute the following command in your terminal:

```bash
cd tcp-server
go run main.go
```

### Running Client

To run the TCP Client, execute the following command in your terminal:

```bash
cd tcp-client
go run main.go
```

## Dependencies
Standard Go libraries: net, bufio, fmt, log, os, strings, time