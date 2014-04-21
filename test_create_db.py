#coding: utf-8

import os
from pony.orm import Database

CURRENT_DIR = os.path.dirname(os.path.abspath(__file__)) + '/'
SQL_NAME = 'test_sql.sql'
SQL_TYPE = 'sqlite'


def is_file():
    '''判断数据库文件是否存在'''

    return os.path.isfile(CURRENT_DIR+SQL_NAME)

def create_db():
    '''创建数据库 create_db=True 参数没有则创建'''

    return Database(SQL_TYPE, SQL_NAME, create_db=True)


if __name__ == '__main__':

    if not is_file():
        # 不存在则创建数据库文件
        db = create_db()
        if is_file():
            print 'Create success.'
    else:
        print '{} is exist.'.format(SQL_NAME) 
