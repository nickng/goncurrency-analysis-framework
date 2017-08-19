# Demo tool for Go concurrency analysis framework

The basic tool consists of a runnable code editor (for use with [playground][playground] package).
Designed for 2-stage transformation of Go code to an intermediate format, then
analysis with an external tool.

[playground]: http://golang.org/x/tools/playground

## Tools support

- CFSM approach
    - [CFSM Synthesis](https://github.com/nickng/dingo-hunter)
    - Paper: [Static Deadlock Detection for Concurrent Go by Global Session Graph Synthesis](http://mrg.doc.ic.ac.uk/publications/static-deadlock-detection-for-concurrent-go-by-global-session-graph-synthesis/)
- Gong approach
    - [Gong safety and liveness checker](https://github.com/nickng/gong)
    - Paper: [Fencing off Go: Liveness and Safety for Channel-based Programming](http://mrg.doc.ic.ac.uk/publications/fencing-off-go-liveness-and-safety-for-channel-based-programming/)
- Godel Checker
    - [migoinfer](https://github.com/nickng/gospal)
    - Godel checker frontend

## Configuration and usage

### Server side: HTTP handler

A webservice handler is defined in `webservice/webservice.go`, defined as name
and an init function. `InitFunc` sets up the HTTP handler path (using the
default `http.HandleFunc`).

    type Handler struct {
        Name     string
        InitFunc func()
    }

HTTP handlers can be loaded selectively by flags in the executable.
By default, only Godel is loaded.

### Client side: Javascript

Event handlers are set up in `static/script.js` as ajax calls, responses should
be encoded as JSON object.

- MiGo (v1 and v2)
    - Endpoint: '/migo.v1' and '/migo.v2'
    - Response: `{ 'MiGo': migo_output, 'time': execution_time, 'Error': error }`
- CFSM synthesis
    - Endpoint: '/cfsm'
    - Response: `{ 'CFSM': cfsm_output, 'time': execution_time, 'Error': error }`
- Gong
    - Endpoint: '/gong'
    - Response: `{ 'Gong': gong_output, 'time': execution_time, 'Error': error }`
- Godel checker
    - Endpoint: '/godel'
    - Response: `{ 'Godel': godel_output, 'time': execution_time, 'Error': error }`
