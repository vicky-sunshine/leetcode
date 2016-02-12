# 1024ms, beats 4.02%
from collections import OrderedDict


class LRUCache(object):
    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.max_capacity = capacity
        self.cache = OrderedDict()

    def get(self, key):
        """
        :rtype: int
        """
        if key in self.cache:
            # move Recently Used element to end of the OrderedDict
            value = self.cache.get(key)
            self.cache.pop(key)
            self.cache[key] = value
            return value
        else:
            return -1

    def set(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: nothing
        """
        if key in self.cache:
            # move Recently Used element to end of the OrderedDict
            self.cache.pop(key)

        self.cache[key] = value
        if len(self.cache) == self.max_capacity:
            # delete unused for longest time element
            self.cache.pop(self.cache.keys()[0])
