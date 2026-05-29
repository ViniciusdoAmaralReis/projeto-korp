package main

import (
    "encoding/json"
    "net/http"
    "time"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Resposta struct {
    Nome    string `json:"nome"`
    Horario string `json:"horario"`
}

var (
    requestsTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint"},
    )
    serviceUp = promauto.NewGauge(
        prometheus.GaugeOpts{
            Name: "service_up",
            Help: "Service availability (1 = up, 0 = down)",
        },
    )
)

func handler(w http.ResponseWriter, r *http.Request) {
    requestsTotal.WithLabelValues(r.Method, "/projeto-korp").Inc()

    resp := Resposta{
        Nome:    "Projeto Korp",
        Horario: time.Now().UTC().Format(time.RFC3339),
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func main() {
    serviceUp.Set(1)

    http.HandleFunc("/projeto-korp", handler)
    http.Handle("/metrics", promhttp.Handler())

    http.ListenAndServe(":8080", nil)
}
