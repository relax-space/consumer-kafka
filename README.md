# consumer-kafka
this a simple consumer sample

## prepare

install kafka
```
https://www.cnblogs.com/mignet/p/window_kafka.html
https://www.apache.org/dyn/closer.cgi?path=/kafka/2.3.0/kafka_2.12-2.3.0.tgz
```

## environment
 - KAFKA_HOST=127.0.0.1
 - PORT=9092
 - TOPIC=fruit

## start
```bash
$ go run .
```

## send message

```bash
./kafka-console-producer.bat --broker-list localhost:9092 --topic fruit
> {"authToken":"","createdAt":"2019-07-30T23:51:30.5899707Z","payload":{"name":"apple"},"requestId":"","status":"FruitCreated"}
```

## accept message

you will see the new message in your console
