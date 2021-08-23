package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	"go-oracle/config"
	"go-oracle/src/database"
	"go-oracle/src/endpoints"
	serviceHttp "go-oracle/src/http"
	"go-oracle/src/service"
)

func main() {
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "caller", log.DefaultCaller)
		logger = log.With(logger, "time", log.DefaultTimestamp)
	}

	cfg, err := config.New()
	if err != nil {
		logger.Log("config", err)
		panic(err)
	}

	if cfg.Port == "" {
		cfg.Port = "8004"
	}

	port := fmt.Sprintf(":%s", cfg.Port)
	// logger.Log("StartWithPort %s", cfg.Port)

	var (
		httpAddr = flag.String("addr", port, "HTTP listen address")
	)

	// db, err := database.NewDB(cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	// if err != nil {
	// 	logger.Log("NewDB", err)
	// 	panic(err)
	// }

	db, err := database.OracleConnection("username/password@localhost:9090/DBNAME")
	if err != nil {
		fmt.Println("OracleConnection", err)
	}
	type RuleConfig struct {
		Code string
	}
	var rules = RuleConfig{}
	cc, err := db.Table("RULE_CONFIG").Count(&rules)
	if err != nil {
		fmt.Println("RuleConfig", err)
	}
	fmt.Println(cc)
	var s service.Service
	var h http.Handler

	s = service.Service{
		Logger: logger,
		Config: cfg,
		// DB:     repository.NewRepository(db, logger),
	}
	h = serviceHttp.NewHTTPHandler(
		s,
		endpoints.MakeServerEndpoints(&s),
		logger,
		false,
	)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
