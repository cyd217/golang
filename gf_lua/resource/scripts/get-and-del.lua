--- 获取旧值，设置新值
local key1,value1=KEYS[1],ARGV[1]
local current=redis.call('get',key1)
if current==value1
then redis.call('del',key1)
    return true
end
    return false