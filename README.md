# Bedrock
Distributed, reference counted, content addressable storage.

## What
Bedrock stores blobs, keyed by the sha256 of their bytes.

## How
www.serf.io is used for system membership.

The data space is split into S shards (360 by default) labeled 0 .. S-1.
S is fixed at system create.
The system is configured for R replicas or copies of each shard.
R is fixed at system create.
Each node in the system gets assigned a number evenly distributed between 0 and S-1.
Each node handles shards greater to its number and less than the next Rth node number
Each node is a primary for shards greater than its number and less than or equal to 

### Shards
Shards are ranges in the key space.
Shards are stored on disk, one index file, and one data file per shard.
Shard index files are sorted by key.

Shard Index File
-------------------------------------------
| key-32 | refcount-8 | offset-8 | size-8 |
-------------------------------------------

If refcount == 0, this record is a tombstone.

The index file is named S.idx
The data file is named S.dat

Shard files are compacted during compaction phase.
During compaction tombstoned records are removed.

Shard Data File
--------
| data |
--------

The data file is just record after record and cannot be used without the index.

### Logs
A log is an append only structure.
There is a single log for membership, and a log for each shard.

Shard Log Record
-----------------------------
| key-32 | size-4 | bytes-N |
-----------------------------

If size > 0 insert op.
If size == 0 delete op.


## Why
Unix philosophy, do one thing, and do it well.

Bedrock is meant as a building block for distributed storage engines.
Bedrock reliably stores the objects. 
Higher level services handle query/index mechanims.
