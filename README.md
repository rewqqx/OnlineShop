# Онлайн-магазин

## Деплой

Для дампа бд используйте:

`docker exec -i <container_name> /bin/bash -c "PGPASSWORD=<pgpassword> pg_dump --username <pgusername> <dbname>" > /desired/path/on/your/machine/dump.sql`

Для рестора бд используйте:

`docker exec -i <container_name> /bin/bash -c "PGPASSWORD=<pgpassword> psql --username <pgusername> <dbname>" < /desired/path/on/your/machine/dump.sql`

## [Описание прецедентов](https://github.com/rewqqx/OnlineShop/blob/main/uml/PrecedentDescr.pdf)

## [UML](https://github.com/rewqqx/OnlineShop/blob/main/uml)

### Диаграмма прецедентов

![Диаграмма прецедентов](https://github.com/rewqqx/OnlineShop/blob/main/uml/PrecedentDiagram.png)

### Диаграмма компонент

![Диаграмма компонент](https://github.com/rewqqx/OnlineShop/blob/main/uml/ComponentDiagram.png)

### Диаграмма концептуальных классов

![Диаграмма концептуальных классов](https://github.com/rewqqx/OnlineShop/blob/main/uml/ConceptDiagram.png)
