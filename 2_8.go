
/*
Структурная схема нашей системы на текущий момент
const ()

type healthCheck struct{
	ServiceID string
	Status	  string
}



Хотя мы только начали знакомиться с языком Go, мы уже можем сделать что-то полезное в рамках нашего большого задания по разработке микросервисов.

Например нашему модулю Checker потребуются структуры для хранения результатов проверок внешних отслеживаемых сервисов.
Давайте создадим такую структуру, назовем ее HealthCheck. Она должна содержать два поля: ServiceID (здесь будет хранить идентификатор отслеживаемого сервиса) и Status (результат его проверки) - обе типа string. Объявим также константы для состояний Status


const ( PassStatus = "pass" FailStatus = "fail" )

Так же для вас разработан вспомогательноый сервис GoMetr, вам необходимо сделать git clone к себе на компьютер и запустить его. После запуска открыть в браузере ссылки http://localhost:8000/metrics и http://localhost:8000/health
8000 порт указан для текущего конфига, если он у вас занят изменить значение в конфиг файле и повторите запуск.
Подсказка: в репозитрии есть Makefile где прописана команда запуска.

Задачу необходимо залить в репозиторий по ранее описанному правилу.

Код по разрабатываемому сервису будем располагать в папке ./Service
*/