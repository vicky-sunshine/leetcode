# 160 ms
class LRUCache
    # Initialize your data structure here
    # @param {Integer} capacity
    def initialize(capacity)
        @cache = Hash.new
        @capacity = capacity
    end

    # @param {Integer} key
    # @return {Integer}
    def get(key)
        val = @cache[key]
        if val.nil?
            -1
        else
            # let the pair {k,v} in the tail of hash
            @cache.delete(key)
            @cache[key] = val
        end
    end

    # @param {Integer} key
    # @param {Integer} valuea
    # @return {Void}
    def set(key, value)
        # let the {k,v} in the tail of hash
        @cache.delete(key)
        @cache[key] = value

        if @cache.size > @capacity
            # delete the head of hash
            @cache.delete(@cache.first[0])
        end
    end
end
