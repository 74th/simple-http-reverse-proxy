package simplehttpreverseproxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	addr        string
	publicDir   string
	HostsConfig *ConfigRoot
}

func NewServer(addr string, hostsConfig *ConfigRoot, publicDir string) *Server {
	return &Server{
		addr:        addr,
		publicDir:   publicDir,
		HostsConfig: hostsConfig,
	}
}

func (s *Server) Run() {
	log.Printf("Starting HTTP server on %s", s.addr)
	mux := http.NewServeMux()

	for _, c := range s.HostsConfig.Hosts {
		targetURL, err := url.Parse(c.Host)
		if err != nil {
			log.Fatalf("Failed to parse host URL: %v", err)
		}

		proxy := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				orgURL := req.URL.String()
				req.URL.Scheme = targetURL.Scheme
				req.URL.Host = targetURL.Host
				if len(c.Path) > 0 {
					req.URL.Path = targetURL.Path + req.URL.Path[len(c.Path)+1:]
				}
				for k, v := range c.Headers {
					req.Header.Set(k, v)
				}

				dstURL := req.URL.String()
				log.Printf("Proxying request %s -> %s", orgURL, dstURL)
			},
		}
		mux.Handle(fmt.Sprintf("/%s/", c.Path), proxy)
	}

	if len(s.publicDir) > 0 {
		mux.Handle("/", http.FileServer(http.Dir(s.publicDir)))
	}

	err := http.ListenAndServe(s.addr, mux)
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
