#coding: utf-8

from pony.orm import Required, db_session, commit, select

from test_create_db import create_db


def create_table(db):

    class Person(db.Entity):
        name = Required(unicode)
        age = Required(int)

    db.generate_mapping(create_tables=True)

    return Person


if __name__ == '__main__':

    db = create_db()

    Person = create_table(db)

    with db_session:
        p1 = Person(name='Bob', age=15)
        p2 = Person(name='Mike', age=14)
        commit()

        for i in select(j for j in Person):
            print i.id, i.name, i.age
