Custom timeouts can be configured for HAProxy via annotations. Supports all valid timeout option
for defaults section of HAProxy. [Read More](https://cbonte.github.io/haproxy-dconv/1.7/configuration.html#4.2-timeout%20check)

`ingress.appscode.com/default-timeout` expects a JSON encoded map of timeouts values.

Ingress Example:
```yaml
apiVersion: voyager.appscode.com/v1beta1
kind: Ingress
metadata:
  name: test-ingress
  namespace: default
  annotations:
    ingress.appscode.com/default-timeout: '{"connect": "5s", "server": "10s"}'
spec:
  backend:
    serviceName: test-server
    servicePort: '80'
  rules:
  - host: appscode.example.com
    http:
      paths:
      - backend:
          serviceName: test-service
          servicePort: '80'
```

This ingress will generate a HAProxy template with provided timeouts. like
```console
defaults
	log global

	option http-server-close
	option dontlognull

	# Timeout values
	# timeout {{ key }}  {{ value }}
	timeout  connect         5s
	timeout  server          10s
	timeout  client          50000
	timeout  client-fin      50000
	timeout  tunnel          50000

```


If any required timeouts is not provided timeouts will be populated with the following values.
```
	timeout  connect         50000
	timeout  client          50000
	timeout  client-fin      50000
	timeout  server          50000
	timeout  tunnel          50000
```

### Time Format
Some parameters involve values representing time, such as timeouts. These
values are generally expressed in milliseconds (unless explicitly stated
otherwise) but may be expressed in any other unit by suffixing the unit to the
numeric value. It is important to consider this because it will not be repeated
for every keyword. Supported units are :

  - us : microseconds. 1 microsecond = 1/1000000 second
  - ms : milliseconds. 1 millisecond = 1/1000 second. This is the default.
  - s  : seconds. 1s = 1000ms
  - m  : minutes. 1m = 60s = 60000ms
  - h  : hours.   1h = 60m = 3600s = 3600000ms
  - d  : days.    1d = 24h = 1440m = 86400s = 86400000ms

### Examples Annotations
```
ingress.appscode.com/default-timeout: '{"client": "5s"}'
```
