select 
    sts.student_id, sts.student_name, sts.subject_name, count(e.student_id) as attended_exams 
from 
    (select * from Students join Subjects) sts  
    left join 
        Examinations e 
    on 
        sts.student_id = e.student_id and sts.subject_name = e.subject_name 

group by sts.student_id, sts.subject_name 
order by sts.student_id, sts.subject_name