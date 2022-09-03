## [0.2.10](https://github.com/uptrace/go-clickhouse/compare/v0.2.9...v0.2.10) (2022-09-03)



## [0.2.9](https://github.com/uptrace/go-clickhouse/compare/v0.2.8...v0.2.9) (2022-08-30)


### Bug Fixes

* **migrate:** upping was applying zero migrations ([f1d380c](https://github.com/uptrace/go-clickhouse/commit/f1d380c16590cc2055274c2dc9418792682a8378))


### Features

* add bfloat16 support ([75cc666](https://github.com/uptrace/go-clickhouse/commit/75cc6664576884120b629f38a473135cbe5214bd))
* add Raw ([07c1f88](https://github.com/uptrace/go-clickhouse/commit/07c1f88173bb056e476b56d8a35dc3e5cf00c596))
* add WithAutoCreateDatabase option ([74e949e](https://github.com/uptrace/go-clickhouse/commit/74e949e01d00e10718d375b43c6f72269165a19d))
* **chmigrate:** add WithReplicated option ([76433f0](https://github.com/uptrace/go-clickhouse/commit/76433f0158277aaa93fec681bbfca7af623baf8a))
* **migrate:** added option to only mark migration up/down as applied on success ([0b4f7bf](https://github.com/uptrace/go-clickhouse/commit/0b4f7bf56588c1060375f094406fe530b7086dcf))



## [0.2.8](https://github.com/uptrace/go-clickhouse/compare/v0.2.7...v0.2.8) (2022-05-29)


### Features

* enable opentelemetry support in protocol ([fb79ac4](https://github.com/uptrace/go-clickhouse/commit/fb79ac4b753bbf6ea794acb1d86fd8d116cf539c))



## [0.2.7](https://github.com/uptrace/go-clickhouse/compare/v0.2.6...v0.2.7) (2022-05-02)


### Features

* allow disabling compression for benchmarks ([a0d867b](https://github.com/uptrace/go-clickhouse/commit/a0d867b5f4478ac4879e73e1c8bb7cf0a8565142))



## [0.2.6](https://github.com/uptrace/go-clickhouse/compare/v0.2.5...v0.2.6) (2022-04-30)


### Features

* add proper Rows implementation and some optimizations ([aca5cfe](https://github.com/uptrace/go-clickhouse/commit/aca5cfeb91514cf6dccb4ebc261755940b290449))
* add support for DateTime64 ([1281505](https://github.com/uptrace/go-clickhouse/commit/1281505a77f39e0ff3203eddd969fded776e72f0))



#### 0.2.6 (2022-03-29)

# [](https://github.com/uptrace/go-clickhouse/compare/v0.2.4...v) (2022-03-29)


### Bug Fixes

* change rollback to always record migrations ([d6e6e55](https://github.com/uptrace/go-clickhouse/commit/d6e6e55142d6cb369d838357a0700dd1becd50a8))
* continue working with non UTC timezone ([d003d44](https://github.com/uptrace/go-clickhouse/commit/d003d44e55049b612610d48607809fe3fff5f151))



# [](https://github.com/uptrace/go-clickhouse/compare/v0.2.3...v) (2022-03-23)



# [](https://github.com/uptrace/go-clickhouse/compare/v0.2.2...v) (2022-03-23)


### Features

* close idle connections after 30 minutes ([844981b](https://github.com/uptrace/go-clickhouse/commit/844981bf1a831ab476e8854d413d2ea31c087d42))



# [](https://github.com/uptrace/go-clickhouse/compare/v0.2.1...v) (2022-03-21)



# [](https://github.com/uptrace/go-clickhouse/compare/v0.2.0...v) (2022-03-21)



#  (2022-03-21)


### Features

* initial commit ([092a2db](https://github.com/uptrace/go-clickhouse/commit/092a2dbf28ca070bd6d6cc3426ecbc1d9bc02c6e))



# [](https://github.com/uptrace/go-clickhouse/compare/v0.1.0...v) (2022-03-17)



# (2022-03-09)

### Bug Fixes

- parse query settings from DSN
  ([6dd2a1a](https://github.com/uptrace/go-clickhouse/commit/6dd2a1adde7a6992d25bf319ce447556fd21aa39))

### Features

- add CreateTableQuery.Order
  ([50192cd](https://github.com/uptrace/go-clickhouse/commit/50192cd8fb1bb6aa65f50daee5e7b11435627255))
- add migrations example
  ([98ecef3](https://github.com/uptrace/go-clickhouse/commit/98ecef3fdb7b10dc947fccb31d641a4ebce2f650))
- initial commit
  ([2f20600](https://github.com/uptrace/go-clickhouse/commit/2f20600f5e4fc9a20e12f1f027e65e0c2bd4f046))
