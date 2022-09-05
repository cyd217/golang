--- 获取旧值，设置新值
local key1,value1,value2=KEYS[1],ARGV[1],ARGV[2]
local current=redis.call('get',key1)
if current==value1
then redis.call('set',key1,value2)
    return true
end
    return false