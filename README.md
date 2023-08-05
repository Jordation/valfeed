# What it do

I'm using this data to mock a "feed" of data being output and handled by a series of services to 

- Evaluate
- Transform
- Display the state of the game / results on a per-update basis. 

The output is about 100mb if jsonl, 6500ish events.
I have a few goals in mind but this is mostly for fun and to experiment with these design patterns

### targets

Handle this processing in a modular fashion, such that multiple concurrent feeds can be handled, will just require some basic linking and pub/sub style containers for my logic.
The order the events are received should not matter either (to mock some fault tolerance in receiving messages) the system will process the events in their true order (its in the metadata)rather than fifo

Do some cool outputs with the data. I started playing around with it but it was very finicky getting the positioning of the dots right. More testing required but the scaffolding is there for it

Build the system modularly using sub routers to allow external communication with the different components  