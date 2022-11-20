# newrelic-zap-logs-in-context

This repository is for using [Logs in Context](https://docs.newrelic.com/docs/logs/logs-context/logs-in-context/) with New Relic using [zap](https://github.com/uber-go/zap).

## Steps 

1. Expose NEW_RELIC_LICENSE_KEY

```shell
export NEW_RELIC_LICENSE_KEY=<your license key>
```

2. Run docker compose

```shell
docker compose up --build
```

3. HTTP Reqeust

```shell
 curl localhost:8000/example/<parameter>
```

4. Check your New Relic UI

APM → Logs

<img width="1571" alt="Screen Shot 2022-11-20 at 18 21 03" src="https://user-images.githubusercontent.com/39491874/202894577-a6897347-77bf-4c31-b273-c3d064ff1b60.png">
