from collections.abc import Callable

from employee import (
    Employee, Programmer, Frontender,
    programmer,
)

def give_task_for_programmer(
    task: Callable[[Programmer], None],
    programmer: Programmer,
) -> None:
    task(programmer)

def task_for_programmer(programmer: Programmer) -> None:
    programmer.fix_bug_in_code(1)

def task_for_employee(employee: Employee) -> None:
    employee.pay_salary(123)

def task_for_frontend(frontend: Frontender) -> None:
    frontend.release_markup(123)

give_task_for_programmer(task_for_programmer, programmer)
give_task_for_programmer(task_for_employee, programmer)

give_task_for_programmer(task_for_frontend, programmer)
