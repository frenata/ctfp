class Either:
    def __init__(self, *, left=None, right=None):
        if left is not None and right is not None:
            raise ValueError("there can only be one")
        if left is None and right is None:
            raise ValueError("there must be one")
        self._left = left
        self._right = right

    @property
    def left(self):
        return self._left

    @property
    def right(self):
        return self._right

