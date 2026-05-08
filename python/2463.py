from typing import List
from functools import lru_cache

class Solution:
    def minimumTotalDistance(self, robot: List[int], factory: List[List[int]]) -> int:
        robot.sort()
        factory.sort()  # sort by position

        n = len(robot)
        m = len(factory)
        INF = 10**18

        @lru_cache(None)
        def dp(i: int, j: int) -> int:
            """
            Minimum distance to repair robots[i:] using factories[j:].
            """
            # All robots assigned
            if i == n:
                return 0

            # No factories left but robots remain
            if j == m:
                return INF

            pos, cap = factory[j]

            # Option 1: skip this factory entirely
            ans = dp(i, j + 1)

            # Option 2: assign k robots (1..cap) to this factory, in order
            dist_sum = 0
            upper = min(cap, n - i)
            for k in range(1, upper + 1):
                dist_sum += abs(robot[i + k - 1] - pos)
                ans = min(ans, dist_sum + dp(i + k, j + 1))

            return ans

        return dp(0, 0)