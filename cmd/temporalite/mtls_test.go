// MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2021 Datadog, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"text/template"
	"time"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
)

func TestMTLSConfig(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, thisFile, _, _ := runtime.Caller(0)
	mtlsDir := filepath.Join(thisFile, "../../../internal/examples/mtls")

	// Create temp config dir
	confDir, err := os.MkdirTemp("", "temporalite-conf-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(confDir)

	// Run templated config and put in temp dir
	var buf bytes.Buffer
	tmpl, err := template.New("temporalite.yaml.template").
		Funcs(template.FuncMap{"qualified": func(s string) string { return strconv.Quote(filepath.Join(mtlsDir, s)) }}).
		ParseFiles(filepath.Join(mtlsDir, "temporalite.yaml.template"))
	if err != nil {
		t.Fatal(err)
	} else if err = tmpl.Execute(&buf, nil); err != nil {
		t.Fatal(err)
	} else if err = os.WriteFile(filepath.Join(confDir, "temporalite.yaml"), buf.Bytes(), 0644); err != nil {
		t.Fatal(err)
	}

	// Run ephemerally using temp config
	args := []string{
		"temporalite",
		"start",
		"--ephemeral",
		"--config", confDir,
		"--namespace", "default",
		"--log-format", "pretty",
		"--port", "10233",
	}
	go func() {
		if err := buildCLI().RunContext(ctx, args); err != nil {
			t.Logf("CLI failed: %v", err)
		}
	}()

	// Load client cert/key for auth
	clientCert, err := tls.LoadX509KeyPair(
		filepath.Join(mtlsDir, "client-cert.pem"),
		filepath.Join(mtlsDir, "client-key.pem"),
	)
	if err != nil {
		t.Fatal(err)
	}
	// Load server cert for CA check
	serverCAPEM, err := os.ReadFile(filepath.Join(mtlsDir, "server-ca-cert.pem"))
	if err != nil {
		t.Fatal(err)
	}
	serverCAPool := x509.NewCertPool()
	serverCAPool.AppendCertsFromPEM(serverCAPEM)

	// Build client options and try to connect client every 100ms for 5s
	options := client.Options{
		HostPort: "localhost:10233",
		ConnectionOptions: client.ConnectionOptions{
			TLS: &tls.Config{
				Certificates: []tls.Certificate{clientCert},
				RootCAs:      serverCAPool,
			},
		},
	}
	var c client.Client
	for i := 0; i < 50; i++ {
		if c, err = client.Dial(options); err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	if err != nil {
		t.Fatal(err)
	}

	// Make a call
	resp, err := c.WorkflowService().DescribeNamespace(ctx, &workflowservice.DescribeNamespaceRequest{
		Namespace: "default",
	})
	if err != nil {
		t.Fatal(err)
	} else if resp.NamespaceInfo.State != enums.NAMESPACE_STATE_REGISTERED {
		t.Fatalf("Bad state: %v", resp.NamespaceInfo.State)
	}
}