# Задание 
Решить кубическое уравнение: $a x^3 + b x^2 + c x + d = 0.$
= Требования 
- CLI
- Вход: коэффициенты
- Выход: действительные корни или сообщение об ошибке
- Обработка ошибок
- Документация
- Тесты

Проект будет содержать содержать два файла:
- `main.go` --- основная часть программы, которая будет:
	- содержать решатель кубического уравнения;
	- парсить входные данные от пользователя;
	- обрабатывать ошибки ввода и решателя;
	- осуществлять вывод результатов.
- `main_test.go` --- тесты работы решателя.

# Содержимое проекта

Основная часть программы содержит:
- Функцию `cubicRoots`, рассматривающую все возможные случаи для действительных корней кубического уравнения, в том числе частные случаи, когда оно выраждается до квадратного или линейного уравнений. Функции должны получать на вход коэфициены уравнения (a, b, c и d) и возвращать массив `float`'ов и количество найденных корней, при наличии бесконечного количества решений договоримся возвращать -1;
- Функцию `main`, в которой реализуется CLI с помощью пакета [`Cobra`](https://cobra.dev/), а также происходит обработка ошибок и всех возможных случаев поведения программы. Для обработки ошибок флаги коэффициентов будем счиать обязательными, для удобства пользователя создадим также флаг `--h`, который будет объяснять как использовать программу и что она делает.
