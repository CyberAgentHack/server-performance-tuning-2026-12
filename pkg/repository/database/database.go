package database

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("github.com/CyberAgentHack/server-performance-tuning-2026-12/pkg/repository/database")
