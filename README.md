# Онлайн-магазин

## Деплой

Контейнер:

`docker run --name online_shop -p 5432:5432 -e POSTGRES_USER=pguser -e POSTGRES_PASSWORD=pgpass -e POSTGRES_DB=postgres -d postgres:latest`

Для дампа бд используйте:

`docker exec -i online_shop /bin/bash -c "PGPASSWORD=pgpass pg_dump --username pguser postgres" > dump.sql`

Для рестора бд используйте:

`docker exec -i online_shop /bin/bash -c "PGPASSWORD=pgpass psql --username pguser postgres" < dump.sql`

## [Описание прецедентов](https://github.com/rewqqx/OnlineShop/blob/main/uml/PrecedentDescr.pdf)

## [UML](https://github.com/rewqqx/OnlineShop/blob/main/uml)

### Диаграмма прецедентов

![Диаграмма прецедентов](https://github.com/rewqqx/OnlineShop/blob/main/uml/PrecedentDiagram.png)

### Диаграмма компонент

![Диаграмма компонент](https://github.com/rewqqx/OnlineShop/blob/main/uml/ComponentDiagram.png)

### Диаграмма концептуальных классов

![Диаграмма концептуальных классов](https://github.com/rewqqx/OnlineShop/blob/main/uml/ConceptDiagram.png)
