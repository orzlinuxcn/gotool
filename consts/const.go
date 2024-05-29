package consts

const (
	LimiterTypeSingle   = "single"
	DefaultLimiterRate  = 10
	DefaultLimiterBurst = 10
)

const (
	RedisIpPort   = "81.70.205.15:6400"
	RedisPassWord = "orzlinuxcn"

	RedisLimiterDefaultName = "orzlinuxcn_limiter"
)

const LuaRedisLimiter = `
local key = KEYS[1]
local limiterMap = redis.call('HGETALL',key)
local burst = tonumber(ARGV[1])
local curTime = tonumber(ARGV[2])
local timePerToken = tonumber(ARGV[3])
local wait = ARGV[4]
local tokens
local lastUpdateTime
if table.getn(limiterMap) == 0 then
	tokens = burst
	lastUpdateTime = curTime
else
	tokens = tonumber(redis.call('HGET',key,'tokens'))
	lastUpdateTime = tonumber(redis.call('HGET',key,'lastUpdateTime'))
end
local periodTimePass = curTime - lastUpdateTime
local addToken = periodTimePass/timePerToken
local oldTokenNum = tokens

tokens = tokens + addToken
if addToken > 0 then
	redis.call('HSET',key,'lastUpdateTime',curTime)
end
if tokens > burst then
	tokens = burst
end

if wait == 'wait' then
	redis.call('HSET',key,'tokens',tokens)
	return tokens
end

tokens = tokens - 1
redis.call('HSET',key,'tokens',tokens)
if tokens >= 0 then
	return 0
else
	return tokens
end
`
