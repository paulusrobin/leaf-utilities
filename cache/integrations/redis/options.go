package leafRedis

import (
	"github.com/go-redis/redis"
	"time"
)

type (
	Option interface {
		Apply(o *options)
	}
	options struct {
		// Credentials options
		address  []string
		password string
		dB       int

		// Common options
		maxRetries         int
		minRetryBackoff    time.Duration
		maxRetryBackoff    time.Duration
		dialTimeout        time.Duration
		readTimeout        time.Duration
		writeTimeout       time.Duration
		poolSize           int
		minIdleConns       int
		maxConnAge         time.Duration
		poolTimeout        time.Duration
		idleTimeout        time.Duration
		idleCheckFrequency time.Duration

		// Cluster options
		maxRedirects   int
		readOnly       bool
		routeByLatency bool
		routeRandomly  bool

		// The sentinel master name, fail over clients.
		sentinelMasterName string
	}
)

func defaultOption() options {
	readTimeout := 3 * time.Second
	idleTimeout := 5 * time.Minute
	return options{
		dB:                 0,
		maxRetries:         0,
		minRetryBackoff:    8 * time.Millisecond,
		maxRetryBackoff:    512 * time.Millisecond,
		dialTimeout:        5 * time.Second,
		readTimeout:        readTimeout,
		writeTimeout:       readTimeout,
		poolSize:           10,
		minIdleConns:       15,
		maxConnAge:         time.Minute,
		poolTimeout:        readTimeout + time.Second,
		idleTimeout:        idleTimeout,
		idleCheckFrequency: idleTimeout,
		maxRedirects:       8,
		readOnly:           true,
		routeByLatency:     true,
		routeRandomly:      true,
		sentinelMasterName: "",
	}
}

func (o options) universalOption() *redis.UniversalOptions {
	return &redis.UniversalOptions{
		Addrs:              o.address,
		DB:                 o.dB,
		Password:           o.password,
		MaxRetries:         o.maxRetries,
		MinRetryBackoff:    o.minRetryBackoff,
		MaxRetryBackoff:    o.maxRetryBackoff,
		DialTimeout:        o.dialTimeout,
		ReadTimeout:        o.readTimeout,
		WriteTimeout:       o.writeTimeout,
		PoolSize:           o.poolSize,
		MinIdleConns:       o.minIdleConns,
		MaxConnAge:         o.maxConnAge,
		PoolTimeout:        o.poolTimeout,
		IdleTimeout:        o.idleTimeout,
		IdleCheckFrequency: o.idleCheckFrequency,
		MaxRedirects:       o.maxRedirects,
		ReadOnly:           o.readOnly,
		RouteByLatency:     o.routeByLatency,
		RouteRandomly:      o.routeRandomly,
		MasterName:         o.sentinelMasterName,
	}
}

type withAddress []string

func (w withAddress) Apply(o *options) {
	o.address = w
}
func WithAddress(address []string) Option {
	return withAddress(address)
}

type withPassword string

func (w withPassword) Apply(o *options) {
	o.password = string(w)
}
func WithPassword(password string) Option {
	return withPassword(password)
}

type withDB int

func (w withDB) Apply(o *options) {
	o.dB = int(w)
}
func WithDB(db int) Option {
	return withDB(db)
}

type withMaxRetries int

func (w withMaxRetries) Apply(o *options) {
	o.maxRetries = int(w)
}
func WithMaxRetries(maxRetries int) Option {
	return withMaxRetries(maxRetries)
}

type withMinRetryBackoff time.Duration

func (w withMinRetryBackoff) Apply(o *options) {
	o.minRetryBackoff = time.Duration(w)
}
func WithMinRetryBackoff(minRetryBackoff time.Duration) Option {
	return withMinRetryBackoff(minRetryBackoff)
}

type withMaxRetryBackoff time.Duration

func (w withMaxRetryBackoff) Apply(o *options) {
	o.maxRetryBackoff = time.Duration(w)
}
func WithMaxRetryBackoff(maxRetryBackoff time.Duration) Option {
	return withMaxRetryBackoff(maxRetryBackoff)
}

type withDialTimeout time.Duration

func (w withDialTimeout) Apply(o *options) {
	o.dialTimeout = time.Duration(w)
}
func WithDialTimeout(dialTimeout time.Duration) Option {
	return withDialTimeout(dialTimeout)
}

type withReadTimeout time.Duration

func (w withReadTimeout) Apply(o *options) {
	o.readTimeout = time.Duration(w)
}
func WithReadTimeout(readTimeout time.Duration) Option {
	return withReadTimeout(readTimeout)
}

type withWriteTimeout time.Duration

func (w withWriteTimeout) Apply(o *options) {
	o.writeTimeout = time.Duration(w)
}
func WithWriteTimeout(writeTimeout time.Duration) Option {
	return withWriteTimeout(writeTimeout)
}

type withPoolSize int

func (w withPoolSize) Apply(o *options) {
	o.poolSize = int(w)
}
func WithPoolSize(poolSize int) Option {
	return withPoolSize(poolSize)
}

type withMinIdleConns int

func (w withMinIdleConns) Apply(o *options) {
	o.minIdleConns = int(w)
}
func WithMinIdleConns(minIdleConns int) Option {
	return withMinIdleConns(minIdleConns)
}

type withMaxConnAge time.Duration

func (w withMaxConnAge) Apply(o *options) {
	o.maxConnAge = time.Duration(w)
}
func WithMaxConnAge(maxConnAge time.Duration) Option {
	return withMaxConnAge(maxConnAge)
}

type withPoolTimeout time.Duration

func (w withPoolTimeout) Apply(o *options) {
	o.poolTimeout = time.Duration(w)
}
func WithPoolTimeout(poolTimeout time.Duration) Option {
	return withPoolTimeout(poolTimeout)
}

type withIdleTimeout time.Duration

func (w withIdleTimeout) Apply(o *options) {
	o.idleTimeout = time.Duration(w)
}
func WithIdleTimeout(idleTimeout time.Duration) Option {
	return withIdleTimeout(idleTimeout)
}

type withIdleCheckFrequency time.Duration

func (w withIdleCheckFrequency) Apply(o *options) {
	o.idleCheckFrequency = time.Duration(w)
}
func WithIdleCheckFrequency(idleCheckFrequency time.Duration) Option {
	return withIdleCheckFrequency(idleCheckFrequency)
}

type withMaxRedirects int

func (w withMaxRedirects) Apply(o *options) {
	o.maxRedirects = int(w)
}
func WithMaxRedirects(maxRedirects time.Duration) Option {
	return withMaxRedirects(maxRedirects)
}

type withReadOnly bool

func (w withReadOnly) Apply(o *options) {
	o.readOnly = bool(w)
}
func WithReadOnly(readOnly bool) Option {
	return withReadOnly(readOnly)
}

type withRouteByLatency bool

func (w withRouteByLatency) Apply(o *options) {
	o.routeByLatency = bool(w)
}
func WithRouteByLatency(routeByLatency bool) Option {
	return withRouteByLatency(routeByLatency)
}

type withRouteRandomly bool

func (w withRouteRandomly) Apply(o *options) {
	o.routeRandomly = bool(w)
}
func WithRouteRandomly(routeRandomly bool) Option {
	return withRouteRandomly(routeRandomly)
}

type withSentinelMasterName string

func (w withSentinelMasterName) Apply(o *options) {
	o.sentinelMasterName = string(w)
}
func WithSentinelMasterName(sentinelMasterName string) Option {
	return withSentinelMasterName(sentinelMasterName)
}
