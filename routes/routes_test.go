package routes

import (
    "testing"

    "github.com/gin-contrib/cors"
)

func TestBuildCORSConfigFromString_Exact(t *testing.T) {
    cfg := buildCORSConfigFromString("http://example.com, https://app.local")

    // When no regexps are present, AllowOrigins should be set
    if len(cfg.AllowOrigins) != 2 {
        t.Fatalf("expected 2 AllowOrigins, got %d", len(cfg.AllowOrigins))
    }
    expected := map[string]bool{"http://example.com": true, "https://app.local": true}
    for _, a := range cfg.AllowOrigins {
        if !expected[a] {
            t.Fatalf("unexpected origin %s", a)
        }
    }

    if cfg.AllowOriginFunc != nil {
        t.Fatalf("expected nil AllowOriginFunc when no wildcards present")
    }
}

func TestBuildCORSConfigFromString_Wildcards(t *testing.T) {
    cfg := buildCORSConfigFromString("http://localhost:*, https://*.example.org,https://exact.com")

    if cfg.AllowOriginFunc == nil {
        t.Fatalf("expected AllowOriginFunc to be set for wildcard patterns")
    }

    matches := []struct{
        origin string
        want bool
    }{
        {"http://localhost:8080", true},
        {"http://localhost", false},
        {"https://sub.example.org", true},
        {"https://deep.sub.example.org", true},
        {"https://example.org", false},
        {"https://exact.com", true},
    }

    for _, tt := range matches {
        got := cfg.AllowOriginFunc(tt.origin)
        if got != tt.want {
            t.Fatalf("origin %s: got %v want %v", tt.origin, got, tt.want)
        }
    }
}

// A tiny sanity test verifying the default behaviour when empty string provided
func TestBuildCORSConfigFromString_Defaults(t *testing.T) {
    cfg := buildCORSConfigFromString("")
    // defaults contain wildcard form so AllowOriginFunc should be set
    if cfg.AllowOriginFunc == nil {
        t.Fatalf("expected AllowOriginFunc for defaults")
    }
    if !cfg.AllowOriginFunc("http://localhost:3000") {
        t.Fatalf("expected default to allow http://localhost:3000")
    }
}

// small helper to ensure the returned type compiles with cors.Config
func ensureType(c cors.Config) {}
