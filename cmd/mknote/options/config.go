package options

import (
	"flag"
	"os"
	"path/filepath"
)

var Instance *Config

type Config struct {
	HostName      string
	HttpPort      string
	TlsPort       string
	HomeDir       string
	LogLevel      int
	ArticlesDir   string
	DownloadDir   string
	StaticDir     string
	HtmlDir       string
	IsTls         bool
	TlsCertFile   string
	TlsKeyFile    string
	ArticleAuthor string
	Version       bool
	DebugPort     string
}

func NewDefaultConfig() *Config {
	// setting up default work location
	self, _ := filepath.Abs(os.Args[0])
	binDir := filepath.Dir(self)
	workDir := filepath.Dir(binDir)
	Instance = &Config{
		"",
		"80",
		"443",
		workDir,
		1,
		workDir + "/articles",
		workDir + "/f",
		workDir + "/static",
		workDir + "/static/template",
		false,
		workDir + "/conf/fullchain.pem",
		workDir + "/conf/privkey.pem",
		"sycki",
		false,
		"8000",
	}

	return Instance
}

// AddFlags add options for command
func (c *Config) AddFlags(cmd *flag.FlagSet) {
	cmd.StringVar(&c.HostName, "hostname", c.HostName, "binding hostname, used to redirect the http port to https")
	cmd.StringVar(&c.HttpPort, "port-http", c.HttpPort, "http server listen port")
	cmd.StringVar(&c.TlsPort, "port-tls", c.TlsPort, "https server listen port")
	cmd.IntVar(&c.LogLevel, "log-level", c.LogLevel, "set log output level, 0...4")
	cmd.StringVar(&c.ArticlesDir, "articles-dir", c.ArticlesDir, "markdown files dir")
	cmd.StringVar(&c.DownloadDir, "uploads-dir", c.DownloadDir, "refence images files dir of all articles")
	cmd.StringVar(&c.StaticDir, "static-dir", c.StaticDir, "html css js etc.")
	cmd.StringVar(&c.HtmlDir, "html-dir", c.HtmlDir, "html template dir")
	cmd.BoolVar(&c.IsTls, "tls", c.IsTls, "specify https model instead of http")
	cmd.StringVar(&c.TlsCertFile, "tls-cert", c.TlsCertFile, "server cert file for https")
	cmd.StringVar(&c.TlsKeyFile, "tls-key", c.TlsKeyFile, "server key file for https")
	cmd.StringVar(&c.ArticleAuthor, "author", c.ArticleAuthor, "generate author name for article metadata")
	cmd.BoolVar(&c.Version, "version", c.Version, "print version and exit")
	cmd.StringVar(&c.DebugPort, "debug-port", c.DebugPort, "used pprofile server port")
}
