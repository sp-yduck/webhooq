# WebhooQ

Queue server with RESTful interface 

## How it works

1. Enqueue

Send `POST` request to WebhooQ server to enque items

```sh
+-----------------+                   +----------------+   +-queue-+
| webhook service | ==POST request==> | WebhooQ server | - | 1:foo | (Enque foo)
+-----------------+  {"foo": "bar"}   +----------------+   +-------+

+-----------------+                   +----------------+   +-----queue-----+
| webhook service | ==POST request==> | WebhooQ server | - | 1:foo, 2:hoge | (Enque hoge)
+-----------------+  {"hoge": "fuga"} +----------------+   +---------------+
```

2. Dequeue

Send `Get` request to WebhooQ server to dequeue items

```sh
    +----queue------+   +----------------+  <==GET request==  +--------+
    | 1:foo, 2:hoge | - | WebhooQ server |  ==return item==>  | client | (Deque foo)
    +---------------+   +----------------+   {"foo": "bar"}   +--------+

    +-queue--+          +----------------+  <==GET request==  +--------+
    | 1:hoge | -------- | WebhooQ server |  ==return item==>  | client | (Deque hoge)
    +--------+          +----------------+  {"hoge": "fuga"}  +--------+

    +--queue--+         +----------------+  <==GET request==  +--------+
    | <empty> | ------- | WebhooQ server |  ==return 204 ==>  | client | (cannot Deque)
    +---------+         +----------------+   204 No Content   +--------+
```