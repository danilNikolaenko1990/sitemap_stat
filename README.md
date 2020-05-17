**sitemap_stat**

Простая консольная утилита для вывода статистики по страницам сайта 
на основе sitemap.xml.
Программа берет sitemap.xml, который вы указываете в sitemap-url, и обходит 
все страницы, указанные в этом xml, замеряя время ответа страницы. Данная операция происходит параллельно в нескольких воркерах, а 
результат пишется в результирующий канал. Далее программа вычитывает результирующий канал и генерирует отчет в виде csv файла. 
Я надеюсь добавить также web-интерфейс и расширить статистику.

**Как запускать в linux**
* Склонируйте данный репозиторий.
* Выполните `go build`
* Выполните `./sitemap_stat stat -u=https://www.google.com/gmail/sitemap.xml`
По завершению программы в папке с программой вы найдете файл с именем report.csv.
Дополнительные детали для работы с программой можно найти, выполнив `./sitemap_stat --help`

The simple console util to output site pages stat based on sitemap.xml.
The program gets sitemap.xml you pointed in sitemap-url and crawls 
all the pages pointed in this xml and measure response time of these pages then.
This operation performs in parallel way. Result writes into result channel in "fan-out way". 
Next step the program reads result channel and generates csv report.
In the future I am planning to add web interface and expand data.  

**How to launch in linux**
* clone this repo
* run `go build`
* run `./sitemap_stat stat -u=https://www.google.com/gmail/sitemap.xml`
After end of execution of the program you will find file named report.csv
Run `./sitemap_stat --help` to find additional details.

**Used packages**
* https://github.com/spf13/cobra for cli 
* https://github.com/sirupsen/logrus to do logging
* https://github.com/cheggaaa/pb to show nice progress bar