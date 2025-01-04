class Employee:
    def __init__(self, name: str):
        self.name = name

    def pay_salary(self, salary: int) -> None:
        print(f'Paid salary: {salary} for {self.name}')

    def __repr__(self) -> str:
        return f'{type(self).__name__}(name={self.name})'

class Accounttant(Employee):
    def calc_salary_for_employee(self, employee: Employee) -> int:
        print(f'{self} calculated salary for employee {employee.name}')
        return 1_000_000 * len(employee.name)

class Programmer(Employee):
    def fix_bug_in_code(self, bug: int) -> None:
        print(f'{self} fixed bug in code: {bug}')

class Frontender(Programmer):
    def release_markup(self, issue: int) -> None:
        print(f'{self} released markup {issue}')

class Backender(Programmer):
    def release_api(self, issue: int) -> None:
        print(f'{self} released api {issue}')

accountant = Accounttant('accountant')
programmer = Programmer('programmer')
frontender = Frontender('frontender')
backender = Backender('backend')
