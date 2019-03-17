import io

from setuptools import find_packages, setup

with io.open('README.rst', 'rt', encoding='utf8') as f:
    readme = f.read()

setup(
    name='first',
    version='1.0.0',
    url='http://localhost:9315/',
    license='BSD',
    maintainer='Win Chua',
    maintainer_email='winchua@foxmail.com',
    description='try',
    long_description=readme,
    packages=find_packages(),
    include_package_data=True,
    zip_safe=False,
    install_requires=[
        'flask',
    ],
    extras_require={
        'test': [
            'pytest',
            'coverage',
        ],
    },
)
