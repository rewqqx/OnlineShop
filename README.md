# Онлайн-магазин

## Деплой

Для запуска используйте

`docker compose up`

Для дампа бд используйте:

`docker exec -i online_shop /bin/bash -c "PGPASSWORD=pgpass pg_dump --username pguser postgres" > dump.sql`


## [Описание прецедентов](https://github.com/rewqqx/OnlineShop/blob/main/uml/PrecedentDescr.pdf)

## [UML](https://github.com/rewqqx/OnlineShop/blob/main/uml)

### Диаграмма прецедентов

![Диаграмма прецедентов](https://github.com/rewqqx/OnlineShop/blob/main/uml/PrecedentDiagram.png)

### Диаграмма компонент

![Диаграмма компонент](https://github.com/rewqqx/OnlineShop/blob/main/uml/ComponentDiagram.png)

### Диаграмма концептуальных классов

![Диаграмма концептуальных классов](https://github.com/rewqqx/OnlineShop/blob/main/uml/ConceptDiagram.png)
