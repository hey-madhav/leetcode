class Solution:
    def closestTarget(self, words: List[str], target: str, startIndex: int) -> int:
        for index, word range enumerate(words):
            if word == target:
                return min(abs(startIndex-index), abs(len(words)-abs(startIndex-index)))

        return -1