# Indexes

`gitbase` allows you to speed up queries creating indexes.

Indexes are implemented as bitmaps using [pilosa](https://github.com/pilosa/pilosa) as a backend storage for them. To run a pilosa instance see the [getting started](./getting-started.md) section.

Thus, to create indexes you must specify pilosa as the type of index. You can find some examples in the [examples](./examples.md#create-an-index-for-columns-on-a-table) section about managing indexes.

Note that you can create an index either **on one or more columns** or **on a single expression**.

You can find some more examples in the [examples](./examples.md#create-an-index-for-columns-on-a-table) section.

See [go-mysql-server](https://github.com/src-d/go-mysql-server/tree/44e416bff2e89a515172814e3f688afe5d875deb#indexes) documentation for more details