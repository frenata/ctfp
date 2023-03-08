import typing
import math

optional = typing.Tuple[any, bool]

def safe_root(n: int) -> optional:
    if n >= 0:
        return math.sqrt(n), True

    return None, False

def safe_reciprocal(n: int) -> optional:
    if n != 0:
        return 1/n, True
    else:
        return None, False

def id(v: any) -> optional:
    return v, True

def compose(f, g):
    def both(v):
        x, ok = g(v)
        if ok:
            return f(x)
        else:
            return x,ok

    return both
