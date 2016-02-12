# 452 ms, beats 26.01%
class LRUCache:

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.capacity = capacity
        self.cache = {}
        self.keys = []  # store

    def get(self, key):
        """
        :rtype: int
        """
        if key in self.cache:
            # move Recently Used element to end of the key list
            self.keys.remove(key)
            self.keys.append(key)
            return self.cache[key]
        else:
            return -1

    def set(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: nothing
        """
        if key in self.cache:
            # move Recently Used element to end of the key list
            self.cache[key] = value
            self.keys.remove(key)
            self.keys.append(key)
        else:
            if len(self.cache) == self.capacity:
                # delete unused for longest time element
                self.cache.pop(self.keys[0])
                self.keys.pop(0)
            self.cache[key] = value
            self.keys.append(key)
