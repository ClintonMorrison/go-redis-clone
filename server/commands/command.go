package commands

/*
TODO: commands to implement

> Strings
append,bitcount,decr,decrby,get,getset,incr,incrby,mget,mset,msetnx,setnx

> Lists
lindex,llen,lpop,lpush,lrange,lrem,lset,ltrim,rpoplpush,rpop,rpush

> Sets
sadd,scard,smembers,sismember,sdiff,sinter,sunion,sdiffstore,sinterstore,sunionstore,spop,srandmember,srem,smove

> Connection
echo,select

> Server
auth,bgrewriteaof,bgsave,config,dbsize,debug,flushdb,info,lastsave,monitor,save,shutdown

> Keys
expireat,expire,keys,move,randomkey,renamenx,sort,ttl,type

*/

type Command interface {
	Apply(store Store) string
	String() string
}

type ExistsCommand struct {
	Key string
}
