from collections.abc import Sequence

from employee import (
    Employee,
    accountant, programmer, frontender, backender
)

empoyees_seq: Sequence[Employee] = [
    accountant,
    programmer,
    frontender,
    backender,
    Employee('smth')
]