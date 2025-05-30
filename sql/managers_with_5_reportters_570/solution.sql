select 
    e1.name 
from 
    Employee e1 left join Employee e2 on e1.id = e2.managerID 
where 
    e2.id is not null 
group by 
    e2.managerId 
having 
    count(*) > 4

-- Лучше использовать иннер так как не придется писать where

SELECT e.name
FROM Employee AS e 
INNER JOIN Employee AS m ON e.id=m.managerId 
GROUP BY m.managerId 
HAVING COUNT(m.managerId) >= 5

SELECT e1.name
FROM employee e1
LEFT JOIN employee e2 ON e1.id=e2.managerId
GROUP BY e1.id
HAVING COUNT(e2.name) >= 5;